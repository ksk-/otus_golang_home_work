package api

import (
	"context"
	"fmt"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/mapper"
	pb "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb"
)

func NewAPI(app *app.App) pb.CalendarApiServer {
	return &api{app: app}
}

type api struct {
	pb.UnimplementedCalendarApiServer
	app *app.App
}

func (a *api) CreateEventV1(ctx context.Context, req *pb.CreateEventV1Request) (*pb.CreateEventV1Response, error) {
	cmd, err := mapper.CreateEventCommand(req)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	eventID, err := a.app.CreateEvent(ctx, cmd)
	if err != nil {
		return nil, fmt.Errorf("create event: %w", err)
	}
	return &pb.CreateEventV1Response{EventId: eventID.String()}, nil
}

func (a *api) UpdateEventV1(ctx context.Context, req *pb.UpdateEventV1Request) (*pb.UpdateEventV1Response, error) {
	event, err := mapper.Event(req)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	if err = a.app.UpdateEvent(ctx, event); err != nil {
		return nil, fmt.Errorf("update event: %w", err)
	}
	return &pb.UpdateEventV1Response{}, nil
}

func (a *api) DeleteEventV1(ctx context.Context, req *pb.DeleteEventV1Request) (*pb.DeleteEventV1Response, error) {
	eventID, err := mapper.EventID(req.EventId)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	if err = a.app.DeleteEvent(ctx, eventID); err != nil {
		return nil, fmt.Errorf("delete event: %w", err)
	}
	return &pb.DeleteEventV1Response{}, nil
}

func (a *api) GetEvent(ctx context.Context, req *pb.GetEventV1Request) (*pb.GetEventV1Response, error) {
	eventID, err := mapper.EventID(req.EventId)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	event, err := a.app.GetEvent(ctx, eventID)
	if err != nil {
		return nil, fmt.Errorf("get event: %w", err)
	}
	return mapper.GetEventV1Response(event), nil
}

func (a *api) GetEventsOfDayV1(ctx context.Context, req *pb.GetEventsV1Request) (*pb.GetEventsV1Response, error) {
	events, err := a.app.GetEventsOfDay(ctx, mapper.BeginOfDay(req))
	if err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return mapper.GetEventsV1Response(events), nil
}

func (a *api) GetEventsOfWeekV1(ctx context.Context, req *pb.GetEventsV1Request) (*pb.GetEventsV1Response, error) {
	events, err := a.app.GetEventsOfWeek(ctx, mapper.BeginOfDay(req))
	if err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return mapper.GetEventsV1Response(events), nil
}

func (a *api) GetEventsOfMonthV1(ctx context.Context, req *pb.GetEventsV1Request) (*pb.GetEventsV1Response, error) {
	events, err := a.app.GetEventsOfMonth(ctx, mapper.BeginOfDay(req))
	if err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return mapper.GetEventsV1Response(events), nil
}
