package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestTasksFilePath(t *testing.T) {
	originalHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	path, err := tasksFilePath()
	if err != nil {
		t.Fatalf("Error obteniendo path: %v", err)
	}

	if path == "" {
		t.Error("Path no debe estar vacío")
	}

	if !filepath.IsAbs(path) {
		t.Error("Path debe ser absoluto")
	}

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Error("El directorio .taskcli no se creó")
	}
}

func TestALoadTasksEmptyFile(t *testing.T) {
	originalHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	path, _ := tasksFilePath()
	os.Remove(path)

	tasks, err := loadTasks()
	if err != nil {
		t.Fatalf("Error cargando tareas: %v", err)
	}

	if len(tasks) != 0 {
		t.Logf("Directorio temporal: %s", tmpDir)
		t.Logf("Path de tareas: %s", path)
		for i, task := range tasks {
			t.Logf("Tarea %d: ID=%d, Title=%s", i, task.ID, task.Title)
		}
		t.Errorf("Se esperaba slice vacío, obtenido %d tareas", len(tasks))
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	originalHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	testTasks := []Task{
		NewTask(1, "Tarea 1", "Descripción 1"),
		NewTask(2, "Tarea 2", "Descripción 2"),
	}
	testTasks[1].Status = DONE

	err := saveTasks(testTasks)
	if err != nil {
		t.Fatalf("Error guardando tareas: %v", err)
	}

	loadedTasks, err := loadTasks()
	if err != nil {
		t.Fatalf("Error cargando tareas: %v", err)
	}

	if len(loadedTasks) != len(testTasks) {
		t.Errorf("Se esperaban %d tareas, obtenidas %d", len(testTasks), len(loadedTasks))
	}

	for i, task := range testTasks {
		if loadedTasks[i].ID != task.ID {
			t.Errorf("ID[%d] = %v, esperado %v", i, loadedTasks[i].ID, task.ID)
		}
		if loadedTasks[i].Title != task.Title {
			t.Errorf("Title[%d] = %v, esperado %v", i, loadedTasks[i].Title, task.Title)
		}
		if loadedTasks[i].Status != task.Status {
			t.Errorf("Status[%d] = %v, esperado %v", i, loadedTasks[i].Status, task.Status)
		}
	}
}

func TestNextID(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []Task
		expected int
	}{
		{
			name:     "lista vacía",
			tasks:    []Task{},
			expected: 1,
		},
		{
			name: "una tarea",
			tasks: []Task{
				{ID: 5},
			},
			expected: 6,
		},
		{
			name: "múltiples tareas",
			tasks: []Task{
				{ID: 3},
				{ID: 1},
				{ID: 7},
				{ID: 2},
			},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nextID(tt.tasks)
			if result != tt.expected {
				t.Errorf("nextID() = %v, esperado %v", result, tt.expected)
			}
		})
	}
}

func TestFindTaskIndexByID(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "Task 1"},
		{ID: 5, Title: "Task 5"},
		{ID: 3, Title: "Task 3"},
	}

	tests := []struct {
		name        string
		id          int
		expectedIdx int
		shouldError bool
	}{
		{"encontrar primera tarea", 1, 0, false},
		{"encontrar tarea intermedia", 5, 1, false},
		{"encontrar última tarea", 3, 2, false},
		{"tarea no existe", 99, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, err := findTaskIndexByID(tasks, tt.id)
			if tt.shouldError {
				if err == nil {
					t.Error("Se esperaba error pero no hubo")
				}
				if !os.IsNotExist(err) {
					t.Errorf("Error esperado: os.ErrNotExist, obtenido: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Error inesperado: %v", err)
				}
				if idx != tt.expectedIdx {
					t.Errorf("Índice = %v, esperado %v", idx, tt.expectedIdx)
				}
			}
		})
	}
}

func TestSaveTasksJSONFormat(t *testing.T) {
	originalHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	tasks := []Task{
		NewTask(1, "Test", "Desc"),
	}

	err := saveTasks(tasks)
	if err != nil {
		t.Fatalf("Error guardando: %v", err)
	}

	path, _ := tasksFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Error leyendo archivo: %v", err)
	}

	var decoded []Task
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("JSON inválido: %v", err)
	}

	if len(decoded) != 1 {
		t.Errorf("Se esperaba 1 tarea, obtenidas %d", len(decoded))
	}
}
