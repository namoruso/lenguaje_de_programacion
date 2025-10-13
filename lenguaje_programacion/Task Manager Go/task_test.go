package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStatusString(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{TODO, "TODO"},
		{INPROGRESS, "IN_PROGRESS"},
		{DONE, "DONE"},
		{Status(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		result := tt.status.String()
		if result != tt.expected {
			t.Errorf("Status.String() = %v, esperado %v", result, tt.expected)
		}
	}
}

func TestStatusMarshalJSON(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{TODO, `"TODO"`},
		{INPROGRESS, `"IN_PROGRESS"`},
		{DONE, `"DONE"`},
	}

	for _, tt := range tests {
		result, err := json.Marshal(tt.status)
		if err != nil {
			t.Errorf("Error al serializar: %v", err)
		}
		if string(result) != tt.expected {
			t.Errorf("MarshalJSON() = %v, esperado %v", string(result), tt.expected)
		}
	}
}

func TestStatusUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input       string
		expected    Status
		shouldError bool
	}{
		{`"TODO"`, TODO, false},
		{`"IN_PROGRESS"`, INPROGRESS, false},
		{`"DONE"`, DONE, false},
		{`"INVALID"`, Status(0), true},
	}

	for _, tt := range tests {
		var s Status
		err := json.Unmarshal([]byte(tt.input), &s)
		if tt.shouldError {
			if err == nil {
				t.Errorf("Se esperaba error para input %v", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("Error inesperado: %v", err)
			}
			if s != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, esperado %v", s, tt.expected)
			}
		}
	}
}

func TestNewTask(t *testing.T) {
	id := 1
	title := "Test Task"
	desc := "Test Description"

	task := NewTask(id, title, desc)

	if task.ID != id {
		t.Errorf("ID = %v, esperado %v", task.ID, id)
	}
	if task.Title != title {
		t.Errorf("Title = %v, esperado %v", task.Title, title)
	}
	if task.Description != desc {
		t.Errorf("Description = %v, esperado %v", task.Description, desc)
	}
	if task.Status != TODO {
		t.Errorf("Status = %v, esperado TODO", task.Status)
	}
	if task.CreatedAt.IsZero() {
		t.Error("CreatedAt no debe estar vacío")
	}
	if task.UpdatedAt.IsZero() {
		t.Error("UpdatedAt no debe estar vacío")
	}
	if !task.CreatedAt.Equal(task.UpdatedAt) {
		t.Error("CreatedAt y UpdatedAt deben ser iguales en una tarea nueva")
	}
}

func TestTaskJSONSerialization(t *testing.T) {
	now := time.Now()
	task := Task{
		ID:          1,
		Title:       "Test",
		Description: "Desc",
		Status:      INPROGRESS,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Marshal
	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Error al serializar: %v", err)
	}

	// Unmarshal
	var decoded Task
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Error al deserializar: %v", err)
	}

	if decoded.ID != task.ID {
		t.Errorf("ID = %v, esperado %v", decoded.ID, task.ID)
	}
	if decoded.Title != task.Title {
		t.Errorf("Title = %v, esperado %v", decoded.Title, task.Title)
	}
	if decoded.Status != task.Status {
		t.Errorf("Status = %v, esperado %v", decoded.Status, task.Status)
	}
}
