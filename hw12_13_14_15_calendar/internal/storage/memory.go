package storage

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
)

func NewMemoryStorage() Storage {
	return &memoryStorage{
		events: make(map[uuid.UUID]Event),
	}
}

type memoryStorage struct {
	events map[uuid.UUID]Event
	mtx    sync.RWMutex
}

func (m *memoryStorage) ListEvents(_ context.Context, limit, offset uint64) ([]Event, error) {
	events := make([]Event, 0, len(m.events))

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, event := range m.events {
		events = append(events, event)
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].BeginTime.Before(events[j].BeginTime)
	})

	return events[offset:min(offset+limit, uint64(len(events)))], nil
}

func (m *memoryStorage) GetEvent(_ context.Context, id uuid.UUID) (*Event, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if event, ok := m.events[id]; ok {
		return &event, nil
	}
	return nil, ErrEventNotFound
}

func (m *memoryStorage) GetEventsForPeriod(_ context.Context, from, to time.Time) ([]Event, error) {
	events := make([]Event, 0, len(m.events))

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, event := range m.events {
		if event.inPeriod(from, to) {
			events = append(events, event)
		}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].BeginTime.Before(events[j].BeginTime)
	})

	return events, nil
}

func (m *memoryStorage) InsertEvent(_ context.Context, event *Event) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.contains(event.ID) {
		return ErrEventAlreadyExists
	}
	m.events[event.ID] = *event
	return nil
}

func (m *memoryStorage) UpdateEvent(_ context.Context, event *Event) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if !m.contains(event.ID) {
		return ErrEventNotFound
	}
	m.events[event.ID] = *event
	return nil
}

func (m *memoryStorage) DeleteEvent(_ context.Context, id uuid.UUID) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if !m.contains(id) {
		return ErrEventNotFound
	}
	delete(m.events, id)
	return nil
}

func (m *memoryStorage) contains(id uuid.UUID) bool {
	_, ok := m.events[id]
	return ok
}

func min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}
