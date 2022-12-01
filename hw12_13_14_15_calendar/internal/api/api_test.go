package api_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/api"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
	pb "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const bufferSize = 1024 * 1024

type apiTestSuite struct {
	suite.Suite
	storage  storage.Storage
	srv      *grpc.Server
	listener *bufconn.Listener

	conn   *grpc.ClientConn
	client pb.CalendarApiClient
}

func (s *apiTestSuite) SetupSuite() {
	s.storage = storage.NewMemoryStorage()
	s.srv = grpc.NewServer()
	s.listener = bufconn.Listen(bufferSize)

	pb.RegisterCalendarApiServer(s.srv, api.NewAPI(app.New(s.storage, logger.Global())))
	go func() {
		s.NoError(s.srv.Serve(s.listener))
	}()
}

func (s *apiTestSuite) SetupTest() {
	conn, err := grpc.DialContext(
		context.Background(), "api_test",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	s.NoError(err)
	s.conn = conn
	s.client = pb.NewCalendarApiClient(conn)
}

func (s *apiTestSuite) TearDownTest() {
	s.NoError(s.conn.Close())
}

func (s *apiTestSuite) TearDownSuite() {
	s.srv.Stop()
}

func (s *apiTestSuite) TestCreateEvent() {
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

		event, err := s.storage.GetEvent(ctx, eventID)
		s.NoError(err)
		s.Equal(eventID, event.ID)
		s.Equal("event_1", event.Title)
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

func (s *apiTestSuite) TestUpdateEvent() {
	ctx := context.Background()
	eventID := s.addTestEvent()
	s.Run("simple case", func() {
		_, err := s.client.UpdateEventV1(ctx, &pb.UpdateEventV1Request{
			Event: &pb.Event{
				Id:        eventID.String(),
				Title:     "changed_event",
				BeginTime: timestamppb.New(time.Now()),
				EndTime:   timestamppb.New(time.Now().Add(time.Second)),
				UserId:    uuid.New().String(),
				NotifyIn:  durationpb.New(time.Second),
			},
		})
		s.NoError(err)

		event, err := s.storage.GetEvent(ctx, eventID)
		s.NoError(err)
		s.Equal("changed_event", event.Title)
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

func (s *apiTestSuite) TestDeleteEvent() {
	ctx := context.Background()
	eventID := s.addTestEvent()
	s.Run("simple case", func() {
		_, err := s.client.DeleteEventV1(ctx, &pb.DeleteEventV1Request{EventId: eventID.String()})
		s.NoError(err)

		_, err = s.storage.GetEvent(ctx, eventID)
		s.ErrorContains(err, "event not found")
	})
	s.Run("non existent event", func() {
		_, err := s.client.DeleteEventV1(ctx, &pb.DeleteEventV1Request{EventId: uuid.New().String()})
		s.ErrorContains(err, "event not found")
	})
}

func (s *apiTestSuite) TestGetEventsOfDay() {
	today := time.Now().Truncate(24 * time.Hour)
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)
	s.fillEventStorage(yesterday, 4*time.Hour, 18)

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

func (s *apiTestSuite) TestGetEventsOfWeek() {
	s.fillEventStorage(time.Now(), 24*time.Hour, 10)

	ctx := context.Background()
	s.Run("in period", func() {
		for _, date := range []time.Time{time.Now(), time.Now().AddDate(0, 0, 7)} {
			res, err := s.client.GetEventsOfWeekV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.NotEmpty(res.Events)
		}
	})
	s.Run("no in period", func() {
		for _, date := range []time.Time{time.Now().AddDate(0, 0, -8), time.Now().AddDate(0, 0, 15)} {
			res, err := s.client.GetEventsOfWeekV1(ctx, &pb.GetEventsV1Request{Since: timestamppb.New(date)})
			s.NoError(err)
			s.Empty(res.Events)
		}
	})
}

func (s *apiTestSuite) GetEventsOfMonth() {
	s.fillEventStorage(time.Now(), 24*time.Hour, 35)

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

func TestAPI(t *testing.T) {
	suite.Run(t, new(apiTestSuite))
}

func (s *apiTestSuite) addTestEvent() uuid.UUID {
	eventID := uuid.New()
	event := storage.Event{
		ID:               eventID,
		BeginTime:        time.Now(),
		EndTime:          time.Now().Add(time.Second),
		UserID:           uuid.New(),
		NotificationTime: time.Now().Add(-time.Second),
	}
	s.NoError(s.storage.InsertEvent(context.Background(), &event))
	return eventID
}

func (s *apiTestSuite) fillEventStorage(since time.Time, d time.Duration, count int) {
	t := since
	for i := 0; i < count; i++ {
		event := storage.Event{
			ID:               uuid.New(),
			BeginTime:        t,
			EndTime:          t.Add(d),
			UserID:           uuid.New(),
			NotificationTime: t,
		}
		s.NoError(s.storage.InsertEvent(context.Background(), &event))
		t = t.Add(d)
	}
}
