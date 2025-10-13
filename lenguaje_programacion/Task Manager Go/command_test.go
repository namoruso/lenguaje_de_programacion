package main

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func setupTestEnv(t *testing.T) func() {
	originalHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	path, _ := tasksFilePath()
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}

	cmdAdd.Flags().Set("title", "")
	cmdAdd.Flags().Set("desc", "")
	cmdList.Flags().Set("state", "all")
	cmdEdit.Flags().Set("title", "")
	cmdEdit.Flags().Set("desc", "")

	return func() {
		os.Setenv("HOME", originalHome)
	}
}

func TestCmdAddSuccess(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	tempCmd := &cobra.Command{
		Use:   "add",
		Short: "Agregar una nueva tarea",
		Run:   cmdAdd.Run,
	}
	tempCmd.Flags().StringP("title", "t", "", "Título de la tarea")
	tempCmd.Flags().StringP("desc", "d", "", "Descripción")

	tempCmd.SetArgs([]string{"--title", "Nueva Tarea", "--desc", "Descripción de prueba"})
	tempCmd.Execute()

	tasks, err := loadTasks()
	if err != nil {
		t.Fatalf("Error cargando tareas: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Se esperaba 1 tarea, obtenidas %d", len(tasks))
	}

	if tasks[0].Title != "Nueva Tarea" {
		t.Errorf("Título = %v, esperado 'Nueva Tarea'", tasks[0].Title)
	}

	if tasks[0].Description != "Descripción de prueba" {
		t.Errorf("Descripción = %v, esperado 'Descripción de prueba'", tasks[0].Description)
	}

	if tasks[0].Status != TODO {
		t.Errorf("Status = %v, esperado TODO", tasks[0].Status)
	}
}

func TestCmdAddWithoutTitle(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	output := captureOutput(func() {
		tempCmd := &cobra.Command{
			Use:   "add",
			Short: "Agregar una nueva tarea",
			Run:   cmdAdd.Run,
		}
		tempCmd.Flags().StringP("title", "t", "", "Título de la tarea")
		tempCmd.Flags().StringP("desc", "d", "", "Descripción")

		tempCmd.SetArgs([]string{"--desc", "Solo descripción"})
		tempCmd.Execute()
	})

	if output == "" {
		t.Error("Se esperaba mensaje de error")
	}

	tasks, _ := loadTasks()
	if len(tasks) != 0 {
		t.Errorf("No se debería haber creado ninguna tarea, pero hay %d", len(tasks))
	}
}

func TestCmdListEmpty(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	output := captureOutput(func() {
		tempCmd := &cobra.Command{
			Use:   "list",
			Short: "Listar tareas",
			Run:   cmdList.Run,
		}
		tempCmd.Flags().StringP("state", "s", "all", "Filtrar por estado")

		tempCmd.SetArgs([]string{})
		tempCmd.Execute()
	})

	if output != "No hay tareas.\n" {
		t.Errorf("Output inesperado: %v", output)
	}
}

func TestCmdListWithTasks(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	tasks := []Task{
		NewTask(1, "Tarea TODO", ""),
		NewTask(2, "Tarea en progreso", ""),
		NewTask(3, "Tarea completada", ""),
	}
	tasks[1].Status = INPROGRESS
	tasks[2].Status = DONE
	saveTasks(tasks)

	cmdList.Flags().Set("state", "all")
	output := captureOutput(func() {
		cmdList.Run(cmdList, []string{})
	})

	if output == "" {
		t.Error("Output no debe estar vacío")
	}

	cmdList.Flags().Set("state", "todo")
	output = captureOutput(func() {
		cmdList.Run(cmdList, []string{})
	})

	if output == "" {
		t.Error("Debe mostrar tareas TODO")
	}
}

func TestCmdViewExistingTask(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	task := NewTask(1, "Tarea de prueba", "Esta es la descripción")
	saveTasks([]Task{task})

	output := captureOutput(func() {
		cmdView.Run(cmdView, []string{"1"})
	})

	if output == "" {
		t.Error("Output no debe estar vacío")
	}
}

func TestCmdViewNonExistingTask(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	output := captureOutput(func() {
		cmdView.Run(cmdView, []string{"999"})
	})

	if output != "Tarea 999 no encontrada\n" {
		t.Errorf("Output inesperado: %v", output)
	}
}

func TestCmdStart(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	task := NewTask(1, "Tarea", "Desc")
	saveTasks([]Task{task})

	cmdStart.Run(cmdStart, []string{"1"})

	tasks, _ := loadTasks()
	if tasks[0].Status != INPROGRESS {
		t.Errorf("Status = %v, esperado IN_PROGRESS", tasks[0].Status)
	}

	if tasks[0].UpdatedAt.Equal(tasks[0].CreatedAt) {
		t.Error("UpdatedAt debería haber cambiado")
	}
}

func TestCmdDone(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	task := NewTask(1, "Tarea", "Desc")
	task.Status = INPROGRESS
	saveTasks([]Task{task})

	cmdDone.Run(cmdDone, []string{"1"})

	tasks, _ := loadTasks()
	if tasks[0].Status != DONE {
		t.Errorf("Status = %v, esperado DONE", tasks[0].Status)
	}
}

func TestCmdEditTitle(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	task := NewTask(1, "Título Original", "Descripción")
	saveTasks([]Task{task})

	captureOutput(func() {
		tempCmd := &cobra.Command{
			Use:   "edit <id>",
			Short: "Editar título y/o descripción de una tarea",
			Args:  cobra.ExactArgs(1),
			Run:   cmdEdit.Run,
		}
		tempCmd.Flags().StringP("title", "t", "", "Nuevo título")
		tempCmd.Flags().StringP("desc", "d", "", "Nueva descripción")

		tempCmd.SetArgs([]string{"1", "--title", "Título Nuevo"})
		tempCmd.Execute()
	})

	tasks, _ := loadTasks()
	if len(tasks) == 0 {
		t.Fatal("No hay tareas después de editar")
	}

	if tasks[0].Title != "Título Nuevo" {
		t.Errorf("Title = %v, esperado 'Título Nuevo'", tasks[0].Title)
	}

	if tasks[0].Description != "Descripción" {
		t.Errorf("La descripción = %v, esperada 'Descripción'", tasks[0].Description)
	}
}

func TestCmdEditDescription(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	task := NewTask(1, "Título", "Descripción Original")
	saveTasks([]Task{task})

	captureOutput(func() {
		tempCmd := &cobra.Command{
			Use:   "edit <id>",
			Short: "Editar título y/o descripción de una tarea",
			Args:  cobra.ExactArgs(1),
			Run:   cmdEdit.Run,
		}
		tempCmd.Flags().StringP("title", "t", "", "Nuevo título")
		tempCmd.Flags().StringP("desc", "d", "", "Nueva descripción")

		tempCmd.SetArgs([]string{"1", "--desc", "Nueva Descripción"})
		tempCmd.Execute()
	})

	tasks, _ := loadTasks()
	if len(tasks) == 0 {
		t.Fatal("No hay tareas después de editar")
	}

	if tasks[0].Description != "Nueva Descripción" {
		t.Errorf("Description = %v, esperado 'Nueva Descripción'", tasks[0].Description)
	}

	if tasks[0].Title != "Título" {
		t.Errorf("El título = %v, esperado 'Título'", tasks[0].Title)
	}
}

func TestCmdEditWithoutFlags(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	task := NewTask(1, "Título", "Descripción")
	saveTasks([]Task{task})

	cmdEdit.Flags().Set("title", "")
	cmdEdit.Flags().Set("desc", "")

	output := captureOutput(func() {
		cmdEdit.Run(cmdEdit, []string{"1"})
	})

	if output == "" {
		t.Error("Debería mostrar mensaje de error")
	}
}

func TestCmdRemove(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	tasks := []Task{
		NewTask(1, "Tarea 1", ""),
		NewTask(2, "Tarea 2", ""),
		NewTask(3, "Tarea 3", ""),
	}
	saveTasks(tasks)

	cmdRemove.Run(cmdRemove, []string{"1"})

	loadedTasks, _ := loadTasks()
	if len(loadedTasks) != 2 {
		t.Errorf("Se esperaban 2 tareas, obtenidas %d", len(loadedTasks))
	}

	for _, task := range loadedTasks {
		if task.ID == 1 {
			t.Error("La tarea 1 no debería existir")
		}
	}
}

func TestCmdVersion(t *testing.T) {
	output := captureOutput(func() {
		cmdVersion.Run(cmdVersion, []string{})
	})

	if output == "" {
		t.Error("Version output no debe estar vacío")
	}
}

func TestTimeNow(t *testing.T) {
	before := time.Now()
	result := timeNow()
	after := time.Now()

	if result.Before(before) || result.After(after) {
		t.Error("timeNow() debería retornar tiempo actual")
	}
}
