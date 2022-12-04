package mapper

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
	pb "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateEventCommand(req *pb.CreateEventV1Request) (*app.CreateEventCommand, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UserID: %w", err)
	}

	cmd := app.CreateEventCommand{
		Title:       req.Title,
		BeginTime:   req.BeginTime.AsTime(),
		EndTime:     req.EndTime.AsTime(),
		Description: req.Description,
		UserID:      userID,
	}

	if req.NotifyIn == nil {
		cmd.NotificationTime = cmd.BeginTime
	} else {
		cmd.NotificationTime = cmd.BeginTime.Add(-req.NotifyIn.AsDuration())
	}

	return &cmd, nil
}

func Event(req *pb.UpdateEventV1Request) (*storage.Event, error) {
	if req.Event == nil {
		return nil, errors.New("event field is empty")
	}

	eventID, err := uuid.Parse(req.Event.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid EventID: %w", err)
	}

	userID, err := uuid.Parse(req.Event.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UserID: %w", err)
	}

	return &storage.Event{
		ID:               eventID,
		Title:            req.Event.Title,
		BeginTime:        req.Event.BeginTime.AsTime(),
		EndTime:          req.Event.EndTime.AsTime(),
		Description:      req.Event.Description,
		UserID:           userID,
		NotificationTime: req.Event.BeginTime.AsTime().Add(-req.Event.NotifyIn.AsDuration()),
	}, nil
}

func EventID(str string) (uuid.UUID, error) {
	eventID, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid eventID: %w", err)
	}
	return eventID, nil
}

func BeginOfDay(req *pb.GetEventsV1Request) time.Time {
	return req.Since.AsTime().Truncate(24 * time.Hour)
}

func GetEventV1Response(event *storage.Event) *pb.GetEventV1Response {
	return &pb.GetEventV1Response{
		Event: &pb.Event{
			Id:          event.ID.String(),
			Title:       event.Title,
			BeginTime:   timestamppb.New(event.BeginTime),
			EndTime:     timestamppb.New(event.EndTime),
			Description: event.Description,
			UserId:      event.UserID.String(),
			NotifyIn:    durationpb.New(event.BeginTime.Sub(event.NotificationTime)),
		},
	}
}

func GetEventsV1Response(events []storage.Event) *pb.GetEventsV1Response {
	mapped := make([]*pb.Event, len(events))
	for i, event := range events {
		mapped[i] = &pb.Event{
			Id:          event.ID.String(),
			Title:       event.Title,
			BeginTime:   timestamppb.New(event.BeginTime),
			EndTime:     timestamppb.New(event.EndTime),
			Description: event.Description,
			UserId:      event.UserID.String(),
			NotifyIn:    durationpb.New(event.BeginTime.Sub(event.NotificationTime)),
		}
	}
	return &pb.GetEventsV1Response{Events: mapped}
}
