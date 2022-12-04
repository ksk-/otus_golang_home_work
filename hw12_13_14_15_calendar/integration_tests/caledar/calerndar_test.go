//go:build integration

package caledar_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/notify"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/rmq"
	pb "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb"
	"github.com/pressly/goose"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	serviceAddr   string
	eventsDSN     string
	migrationsDir string
	rmqURI        string
	rmqQueue      string
)

func TestMain(m *testing.M) {
	flag.StringVar(&serviceAddr, "service", "localhost:6703", "calendar service address")
	flag.StringVar(&eventsDSN, "events-db", "postgres://user:password123@localhost:5432/events", "events DB DSN")
	flag.StringVar(&migrationsDir, "migrations-dir", "migrations", "events DB migrations directory")
	flag.StringVar(&rmqURI, "rmq-uri", "amqp://user:password123@localhost:5672", "RMQ service URI")
	flag.StringVar(&rmqQueue, "queue", "event_notifications_sent", "RMQ queue name (to read sent notifications)")
	os.Exit(m.Run())
}

func TestCalendar(t *testing.T) {
	suite.Run(t, new(calendarTestSuite))
}

type calendarTestSuite struct {
	suite.Suite
	db     *sql.DB
	client pb.CalendarApiClient
}

func (s *calendarTestSuite) SetupSuite() {
	db, err := sql.Open("pgx", eventsDSN)
	s.NoError(err)
	s.NoError(goose.Up(db, migrationsDir))
	s.db = db
}

func (s *calendarTestSuite) SetupTest() {
	_, err := s.db.ExecContext(context.Background(), `TRUNCATE TABLE events`)
	s.NoError(err)

	conn, err := grpc.Dial(
		serviceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	s.NoError(err)
	s.client = pb.NewCalendarApiClient(conn)
}

func (s *calendarTestSuite) TearDownSuite() {
	s.NoError(s.db.Close())
}

func (s *calendarTestSuite) TestCreateEvent() {
	ctx := context.Background()
	s.Run("simple case", func() {
		res, err := s.client.CreateEventV1(ctx, &pb.CreateEventV1Request{
			Title:     "event_1",
			BeginTime: timestamppb.New(time.Now()),
			EndTime:   timestamppb.New(time.Now().Add(time.Second)),
			UserId:    uuid.New().String(),
		})
		s.NoError(err)

		eventID, err := uuid.Parse(res.EventId)
		s.NoError(err)

		created, err := s.client.GetEvent(ctx, &pb.GetEventV1Request{EventId: eventID.String()})
		s.NoError(err)
		s.NotNil(created.Event)
		s.Equal(res.EventId, created.Event.Id)
		s.Equal("event_1", created.Event.Title)
	})
	s.Run("create unique events id", func() {
		res1, err := s.client.CreateEventV1(ctx, &pb.CreateEventV1Request{
			Title:     "event_1",
			BeginTime: timestamppb.New(time.Now()),
			EndTime:   timestamppb.New(time.Now().Add(time.Second)),
			UserId:    uuid.New().String(),
		})
		s.NoError(err)
		res2, err := s.client.CreateEventV1(ctx, &pb.CreateEventV1Request{
			Title:     "event_2",
			BeginTime: timestamppb.New(time.Now()),
			EndTime:   timestamppb.New(time.Now().Add(time.Second)),
			UserId:    uuid.New().String(),
		})
		s.NoError(err)
		s.NotEqual(res1.EventId, res2.EventId)
	})
	s.Run("invalid request", func() {
		_, err := s.client.CreateEventV1(ctx, &pb.CreateEventV1Request{Title: "event_1"})
		s.ErrorContains(err, "invalid request")
	})
}

func (s *calendarTestSuite) TestUpdateEvent() {
	ctx := context.Background()
	eventID := s.addTestEvent()

	s.Run("simple case", func() {
		_, err := s.client.UpdateEventV1(ctx, &pb.UpdateEventV1Request{
			Event: &pb.Event{
				Id:        eventID,
				Title:     "changed_event",
				BeginTime: timestamppb.New(time.Now()),
				EndTime:   timestamppb.New(time.Now().Add(time.Second)),
				UserId:    uuid.New().String(),
				NotifyIn:  durationpb.New(time.Second),
			},
		})
		s.NoError(err)

		updated, err := s.client.GetEvent(ctx, &pb.GetEventV1Request{EventId: eventID})
		s.NoError(err)
		s.NotNil(updated.Event)
		s.Equal("changed_event", updated.Event.Title)
	})
	s.Run("non existent event", func() {
		_, err := s.client.UpdateEventV1(ctx, &pb.UpdateEventV1Request{
			Event: &pb.Event{
				Id:        uuid.New().String(),
				Title:     "changed_event",
				BeginTime: timestamppb.New(time.Now()),
				EndTime:   timestamppb.New(time.Now().Add(time.Second)),
				UserId:    uuid.New().String(),
				NotifyIn:  durationpb.New(time.Second),
			},
		})
		s.ErrorContains(err, "event not found")
	})
	s.Run("invalid request", func() {
		_, err := s.client.UpdateEventV1(ctx, &pb.UpdateEventV1Request{
			Event: &pb.Event{Title: "event_1"},
		})
		s.ErrorContains(err, "invalid request")
	})
}

func (s *calendarTestSuite) TestDeleteEvent() {
	ctx := context.Background()
	eventID := s.addTestEvent()
	s.Run("simple case", func() {
		_, err := s.client.DeleteEventV1(ctx, &pb.DeleteEventV1Request{EventId: eventID})
		s.NoError(err)

		_, err = s.client.GetEvent(ctx, &pb.GetEventV1Request{EventId: eventID})
		s.ErrorContains(err, "event not found")
	})
	s.Run("non existent event", func() {
		_, err := s.client.DeleteEventV1(ctx, &pb.DeleteEventV1Request{EventId: uuid.New().String()})
		s.ErrorContains(err, "event not found")
	})
}

func (s *calendarTestSuite) TestGetEventsOfDay() {
	today := time.Now().Truncate(24 * time.Hour)
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)
	s.addTestEvents(yesterday, 4*time.Hour, 18)

	ctx := context.Background()
	s.Run("in period", func() {
		for _, date := range []time.Time{yesterday, today, tomorrow} {
			res, err := s.client.GetEventsOfDayV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.NotEmpty(res.Events)
		}
	})
	s.Run("no in period", func() {
		for _, date := range []time.Time{yesterday.AddDate(0, 0, -5), tomorrow.AddDate(1, 0, 5)} {
			res, err := s.client.GetEventsOfDayV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.Empty(res.Events)
		}
	})
}

func (s *calendarTestSuite) TestGetEventsOfWeek() {
	s.addTestEvents(time.Now(), 24*time.Hour, 10)

	ctx := context.Background()
	s.Run("in period", func() {
		for _, date := range []time.Time{time.Now(), time.Now().AddDate(0, 0, 7)} {
			res, err := s.client.GetEventsOfWeekV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.NotEmpty(res.Events)
		}
	})
	s.Run("no in period", func() {
		for _, date := range []time.Time{time.Now().AddDate(0, 0, -8), time.Now().AddDate(0, 0, 14)} {
			res, err := s.client.GetEventsOfWeekV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.Empty(res.Events)
		}
	})
}

func (s *calendarTestSuite) TestGetEventsOfMonth() {
	s.addTestEvents(time.Now(), 24*time.Hour, 35)

	ctx := context.Background()
	s.Run("in period", func() {
		for _, date := range []time.Time{time.Now(), time.Now().AddDate(0, 1, 0)} {
			res, err := s.client.GetEventsOfMonthV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.NotEmpty(res.Events)
		}
	})
	s.Run("no in period", func() {
		for _, date := range []time.Time{time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, 2, 0)} {
			res, err := s.client.GetEventsOfMonthV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.Empty(res.Events)
		}
	})
}

func (s *calendarTestSuite) TestSendNotifications() {
	queue, err := rmq.NewQueue(rmqURI, rmqQueue)
	s.NoError(err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := queue.ConsumeChannel(ctx, "TestSendNotifications")
	s.NoError(err)

	count := 5
	since := time.Now().Add(3 * time.Second)
	events := s.addTestEvents(since, time.Second, count)
	notifications := make([]notify.Notification, 0, count)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer cancel()
		defer wg.Done()
		s.Eventually(func() bool {
			return len(notifications) == count
		}, time.Minute, time.Millisecond)
	}()
	go func() {
		defer wg.Done()
		for msg := range ch {
			var notification notify.Notification
			s.NoError(json.Unmarshal(msg, &notification))
			notifications = append(notifications, notification)
		}
	}()
	<-ctx.Done()
	wg.Wait()

	s.Equal(len(events), len(notifications))
	for i := 0; i < len(events); i++ {
		s.Equal(events[i], notifications[i].EventID.String())
	}
}

func (s *calendarTestSuite) addTestEvent() string {
	res, err := s.client.CreateEventV1(context.Background(), &pb.CreateEventV1Request{
		BeginTime: timestamppb.New(time.Now()),
		EndTime:   timestamppb.New(time.Now().Add(time.Second)),
		UserId:    uuid.New().String(),
		NotifyIn:  durationpb.New(time.Second),
	})
	s.NoError(err)
	return res.EventId
}

func (s *calendarTestSuite) addTestEvents(since time.Time, d time.Duration, count int) []string {
	added := make([]string, 0)

	t := since
	for i := 0; i < count; i++ {
		res, err := s.client.CreateEventV1(context.Background(), &pb.CreateEventV1Request{
			BeginTime: timestamppb.New(t),
			EndTime:   timestamppb.New(t.Add(d)),
			UserId:    uuid.New().String(),
			NotifyIn:  durationpb.New(0),
		})
		s.NoError(err)
		t = t.Add(d)
		added = append(added, res.EventId)
	}

	return added
}
