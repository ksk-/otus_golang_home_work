package storage

import (
	"context"
	"fmt"
	"math"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	testEventsCount   = 3
	testEventDuration = time.Minute
	testEventGap      = 10 * time.Second
	testNotifyIn      = 2 * time.Second
)

type MemoryStorageTestSuite struct {
	suite.Suite
	events  []Event
	storage Storage
	begin   time.Time
}

func (s *MemoryStorageTestSuite) SetupTest() {
	s.begin = time.Now()
	s.events = make([]Event, testEventsCount)
	t := s.begin
	for i := 0; i < testEventsCount; i++ {
		s.events[i] = makeTestEvent(fmt.Sprintf("Event_%d", i), t)
		t = t.Add(testEventDuration).Add(testEventGap)
	}
	s.storage = &memoryStorage{
		events: makeEventMap(s.events),
	}
}

func (s *MemoryStorageTestSuite) TestListEvents() {
	ctx := context.Background()
	s.Run("simple case", func() {
		events, err := s.storage.ListEvents(ctx, 2, 0)
		s.NoError(err)
		s.Equal(s.events[:2], events)

		events, err = s.storage.ListEvents(ctx, 2, 1)
		s.NoError(err)
		s.Equal(s.events[1:3], events)

		events, err = s.storage.ListEvents(ctx, 2, 2)
		s.NoError(err)
		s.Equal(s.events[2:3], events)

		events, err = s.storage.ListEvents(ctx, 2, 3)
		s.NoError(err)
		s.Empty(events)
	})
	s.Run("get all", func() {
		events, err := s.storage.ListEvents(ctx, math.MaxUint64, 0)
		s.NoError(err)
		s.Equal(s.events, events)
	})
	s.Run("empty storage", func() {
		storage := NewMemoryStorage()
		events, err := storage.ListEvents(ctx, 100, 0)
		s.NoError(err)
		s.Empty(events)
	})
}

func (s *MemoryStorageTestSuite) TestGetEventsForPeriod() {
	ctx := context.Background()
	s.Run("simple case", func() {
		s.Run("all events", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin, s.begin.Add(150*time.Second))
			s.NoError(err)
			s.Equal(s.events, events)
		})
		s.Run("1st and 2nd events", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(time.Minute), s.begin.Add(100*time.Second))
			s.NoError(err)
			s.Equal(s.events[:2], events)
		})
		s.Run("2st and 3rd events", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(100*time.Second), s.begin.Add(150*time.Second))
			s.NoError(err)
			s.Equal(s.events[1:], events)
		})
		s.Run("2st event", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(100*time.Second), s.begin.Add(70*time.Second))
			s.NoError(err)
			s.Equal(s.events[1:1], events)
		})
	})
	s.Run("no events", func() {
		s.Run("before first", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(-time.Second), s.begin)
			s.NoError(err)
			s.Empty(events)
		})
		s.Run("after last", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(211*time.Second), s.begin.Add(time.Hour))
			s.NoError(err)
			s.Empty(events)
		})
		s.Run("in a gap", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(65*time.Second), s.begin.Add(70*time.Second))
			s.NoError(err)
			s.Empty(events)
		})
	})
}

func (s *MemoryStorageTestSuite) TestGetGetEventsToNotify() {
	ctx := context.Background()
	s.Run("simple case", func() {
		s.Run("all events", func() {
			events, err := s.storage.GetEventsToNotify(ctx, s.begin.Add(-time.Minute), s.begin.Add(150*time.Second))
			s.NoError(err)
			s.Equal(s.events, events)
		})
		s.Run("1st and 2nd events", func() {
			events, err := s.storage.GetEventsToNotify(ctx, s.begin.Add(-10*time.Second), s.begin.Add(70*time.Second))
			s.NoError(err)
			s.Equal(s.events[:2], events)
		})
		s.Run("2st and 3rd events", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(65*time.Second), s.begin.Add(time.Hour))
			s.NoError(err)
			s.Equal(s.events[1:], events)
		})
		s.Run("2st event", func() {
			events, err := s.storage.GetEventsForPeriod(ctx, s.begin.Add(100*time.Second), s.begin.Add(70*time.Second))
			s.NoError(err)
			s.Equal(s.events[1:1], events)
		})
	})
	s.Run("no events", func() {
		s.Run("before first", func() {
			events, err := s.storage.GetEventsToNotify(ctx, s.begin.Add(-time.Second), s.begin)
			s.NoError(err)
			s.Empty(events)
		})
		s.Run("after last", func() {
			events, err := s.storage.GetEventsToNotify(ctx, s.begin.Add(211*time.Second), s.begin.Add(time.Hour))
			s.NoError(err)
			s.Empty(events)
		})
		s.Run("in a gap", func() {
			events, err := s.storage.GetEventsToNotify(ctx, s.begin.Add(time.Minute), s.begin.Add(65*time.Second))
			s.NoError(err)
			s.Empty(events)
		})
	})
}

func (s *MemoryStorageTestSuite) TestGetEvent() {
	ctx := context.Background()
	s.Run("simple case", func() {
		for i := 0; i < testEventsCount; i++ {
			eventID := s.events[i].ID
			event, err := s.storage.GetEvent(ctx, eventID)
			s.NoError(err)
			s.Equal(&s.events[i], event)
		}
	})
	s.Run("non-existent event", func() {
		event, err := s.storage.GetEvent(ctx, uuid.New())
		s.ErrorIs(err, ErrEventNotFound)
		s.Nil(event)
	})
}

func (s *MemoryStorageTestSuite) TestInsertEvent() {
	ctx := context.Background()
	s.Run("simple case", func() {
		testEvent := makeTestEvent("test_event", time.Now())
		err := s.storage.InsertEvent(ctx, &testEvent)
		s.NoError(err)

		event, err := s.storage.GetEvent(ctx, testEvent.ID)
		s.NoError(err)
		s.Equal(&testEvent, event)
	})
	s.Run("already existent event", func() {
		for i := 0; i < testEventsCount; i++ {
			eventID := s.events[i].ID
			changedEvent := makeTestEvent("changed_event", time.Now())
			changedEvent.ID = eventID
			err := s.storage.InsertEvent(ctx, &changedEvent)
			s.ErrorIs(err, ErrEventAlreadyExists)
		}
	})
}

func (s *MemoryStorageTestSuite) TestUpdateEvent() {
	ctx := context.Background()
	s.Run("simple case", func() {
		eventID := s.events[0].ID
		changedEvent := makeTestEvent("changed_event", time.Now())
		changedEvent.ID = eventID
		err := s.storage.UpdateEvent(ctx, &changedEvent)
		s.NoError(err)

		event, err := s.storage.GetEvent(ctx, eventID)
		s.NoError(err)
		s.Equal(&changedEvent, event)
	})
	s.Run("non-existent event", func() {
		testEvent := makeTestEvent("test_event", time.Now())
		err := s.storage.UpdateEvent(ctx, &testEvent)
		s.ErrorIs(err, ErrEventNotFound)
	})
}

func (s *MemoryStorageTestSuite) TestDeleteEvent() {
	ctx := context.Background()
	s.Run("simple case", func() {
		for i := 0; i < testEventsCount; i++ {
			eventID := s.events[i].ID
			err := s.storage.DeleteEvent(ctx, eventID)
			s.NoError(err)

			event, err := s.storage.GetEvent(ctx, eventID)
			s.ErrorIs(err, ErrEventNotFound)
			s.Nil(event)
		}
	})
	s.Run("non-existent event", func() {
		err := s.storage.DeleteEvent(ctx, uuid.New())
		s.ErrorIs(err, ErrEventNotFound)
	})
}

func (s *MemoryStorageTestSuite) TestDeletePastEvents() {
	ctx := context.Background()
	s.Run("simple case", func() {
		deleted, err := s.storage.DeletePastEvents(ctx, s.begin.Add(3*time.Minute))
		s.NoError(err)
		s.Equal(int64(2), deleted)
	})
	s.Run("all events are past", func() {
		deleted, err := s.storage.DeletePastEvents(ctx, s.begin.Add(time.Hour))
		s.NoError(err)
		s.Equal(int64(3), deleted)
	})
	s.Run("no past events", func() {
		deleted, err := s.storage.DeletePastEvents(ctx, s.begin.Add(-time.Hour))
		s.NoError(err)
		s.Zero(deleted)
	})
}

func TestStorage(t *testing.T) {
	suite.Run(t, new(MemoryStorageTestSuite))
}

func TestStorage_Concurrency(t *testing.T) {
	ctx := context.Background()
	s := NewMemoryStorage()
	count := 100

	var wg sync.WaitGroup
	wg.Add(2 * count)
	for i := 0; i < count; i++ {
		eventID := uuid.New()

		// NOTE: insert
		go func(i int) {
			defer wg.Done()
			event := makeTestEvent(fmt.Sprintf("Event_%d", i), time.Now().Add(time.Duration(i)*time.Second))
			event.ID = eventID
			require.NoError(t, s.InsertEvent(ctx, &event))
		}(i)

		// NOTE: get
		go func() {
			defer wg.Done()
			require.Eventually(t, func() bool {
				event, err := s.GetEvent(ctx, eventID)
				return err == nil && event.ID == eventID
			}, time.Minute, time.Millisecond)
		}()
	}
}

func makeTestEvent(title string, beginTime time.Time) Event {
	return Event{
		ID:               uuid.New(),
		Title:            title,
		BeginTime:        beginTime,
		EndTime:          beginTime.Add(testEventDuration),
		Description:      "description...",
		UserID:           uuid.New(),
		NotificationTime: beginTime.Add(-testNotifyIn),
	}
}

func makeEventMap(events []Event) map[uuid.UUID]Event {
	m := make(map[uuid.UUID]Event, len(events))
	for _, event := range events {
		m[event.ID] = event
	}
	return m
}
