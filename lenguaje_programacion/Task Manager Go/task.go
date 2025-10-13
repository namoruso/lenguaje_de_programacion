package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Status int

const (
	TODO Status = iota
	INPROGRESS
	DONE
)

func (s Status) String() string {
	switch s {
	case TODO:
		return "TODO"
	case INPROGRESS:
		return "IN_PROGRESS"
	case DONE:
		return "DONE"
	default:
		return "UNKNOWN"
	}
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(b []byte) error {
	var st string
	if err := json.Unmarshal(b, &st); err != nil {
		return err
	}
	switch st {
	case "TODO":
		*s = TODO
	case "IN_PROGRESS":
		*s = INPROGRESS
	case "DONE":
		*s = DONE
	default:
		return fmt.Errorf("estado desconocido: %s", st)
	}
	return nil
}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(id int, title, desc string) Task {
	now := time.Now()
	return Task{
		ID:          id,
		Title:       title,
		Description: desc,
		Status:      TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
