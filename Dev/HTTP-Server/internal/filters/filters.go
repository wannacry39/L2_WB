package filters

import (
	"httpserver/internal/models"
	"time"
)

type Filter interface {
	FilterByDay(data map[string]models.Event, filterDate time.Time) map[string]models.Event
	FilterByWeek(data map[string]models.Event) map[string]models.Event
	FilterByMonth(data map[string]models.Event, monthTime time.Time) (map[string]models.Event, error)
}
type CurrentFilter struct{}

func NewFilter() Filter {
	return &CurrentFilter{}
}

func (cp *CurrentFilter) FilterByDay(data map[string]models.Event, filterDate time.Time) map[string]models.Event {
	response := make(map[string]models.Event)
	for uuid, event := range data {
		if event.Date == filterDate {
			response[uuid] = event
		}
	}
	return response
}

func (cp *CurrentFilter) FilterByWeek(data map[string]models.Event) map[string]models.Event {
	response := make(map[string]models.Event)
	now := time.Now()
	start := now.AddDate(0, 0, -int(now.Weekday()))
	end := start.AddDate(0, 0, 7)
	for uuid, event := range data {
		if event.Date.After(start) && event.Date.Before(end) {
			response[uuid] = event
		}
	}
	return response
}

func (cp *CurrentFilter) FilterByMonth(data map[string]models.Event, monthTime time.Time) (map[string]models.Event, error) {
	response := make(map[string]models.Event)

	startOfMonth := time.Date(time.Now().Year(), monthTime.Month(), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	for key, event := range data {
		if event.Date.After(startOfMonth) && event.Date.Before(endOfMonth) {
			response[key] = event
		}
	}

	return response, nil
}
