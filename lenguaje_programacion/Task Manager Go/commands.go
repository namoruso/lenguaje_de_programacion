package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"os"
	"time"

	"github.com/spf13/cobra"
)

func timeNow() time.Time {
	return time.Now()
}

var version = "v1.0.0"

var cmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Agregar una nueva tarea",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		if strings.TrimSpace(title) == "" {
			fmt.Println("Error: --title es requerido")
			_ = cmd.Help()
			return
		}
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando tareas:", err)
			return
		}
		t := NewTask(nextID(tasks), title, desc)
		tasks = append(tasks, t)
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error guardando:", err)
			return
		}
		fmt.Printf("Tarea creada: ID=%d\n", t.ID)
	},
}

func init() {
	cmdAdd.Flags().StringP("title", "t", "", "Título de la tarea (requerido)")
	cmdAdd.Flags().StringP("desc", "d", "", "Descripción (opcional)")
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "Listar tareas",
	Run: func(cmd *cobra.Command, args []string) {
		state, _ := cmd.Flags().GetString("state")
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No hay tareas.")
			return
		}
		printed := 0
		for _, t := range tasks {
			show := false
			switch state {
			case "all":
				show = true
			case "todo":
				show = (t.Status == TODO)
			case "inprogress":
				show = (t.Status == INPROGRESS)
			case "done":
				show = (t.Status == DONE)
			default:
				fmt.Println("Estado inválido. Usa: all|todo|inprogress|done")
				return
			}
			if show {
				fmt.Printf("[%d] %s (%s)\n", t.ID, t.Title, t.Status)
				if t.Description != "" {
					fmt.Printf("    %s\n", t.Description)
				}
				printed++
			}
		}
		if printed == 0 {
			fmt.Println("No se encontraron tareas con ese filtro.")
		}
	},
}

func init() {
	cmdList.Flags().StringP("state", "s", "all", "Filtrar por estado: all|todo|inprogress|done")
}

var cmdView = &cobra.Command{
	Use:   "view <id>",
	Short: "Ver detalles de una tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID inválido:", args[0])
			return
		}
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando:", err)
			return
		}
		i, err := findTaskIndexByID(tasks, id)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Tarea %d no encontrada\n", id)
			return
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		t := tasks[i]
		fmt.Printf("ID: %d\nTítulo: %s\nDescripción: %s\nEstado: %s\nCreado: %s\nActualizado: %s\n",
			t.ID, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"), t.UpdatedAt.Format("2006-01-02 15:04"))
	},
}

var cmdStart = &cobra.Command{
	Use:   "start <id>",
	Short: "Marcar tarea como IN_PROGRESS",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID inválido:", args[0])
			return
		}
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando:", err)
			return
		}
		i, err := findTaskIndexByID(tasks, id)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Tarea %d no encontrada\n", id)
			return
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		tasks[i].Status = INPROGRESS
		tasks[i].UpdatedAt = timeNow()
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error guardando:", err)
			return
		}
		fmt.Printf("Tarea %d marcada como IN_PROGRESS\n", id)
	},
}

var cmdDone = &cobra.Command{
	Use:   "done <id>",
	Short: "Marcar tarea como DONE",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID inválido:", args[0])
			return
		}
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando:", err)
			return
		}
		i, err := findTaskIndexByID(tasks, id)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Tarea %d no encontrada\n", id)
			return
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		tasks[i].Status = DONE
		tasks[i].UpdatedAt = timeNow()
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error guardando:", err)
			return
		}
		fmt.Printf("Tarea %d marcada como DONE\n", id)
	},
}

var cmdEdit = &cobra.Command{
	Use:   "edit <id>",
	Short: "Editar título y/o descripción de una tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID inválido:", args[0])
			return
		}

		titleChanged := cmd.Flags().Changed("title")
		descChanged := cmd.Flags().Changed("desc")

		if !titleChanged && !descChanged {
			fmt.Println("Debe especificar --title o --desc para editar")
			_ = cmd.Help()
			return
		}

		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando:", err)
			return
		}
		i, err := findTaskIndexByID(tasks, id)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Tarea %d no encontrada\n", id)
			return
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if titleChanged {
			tasks[i].Title = title
		}
		if descChanged {
			tasks[i].Description = desc
		}
		tasks[i].UpdatedAt = timeNow()
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error guardando:", err)
			return
		}
		fmt.Printf("Tarea %d actualizada\n", id)
	},
}

func init() {
	cmdEdit.Flags().StringP("title", "t", "", "Nuevo título")
	cmdEdit.Flags().StringP("desc", "d", "", "Nueva descripción")
}

var cmdRemove = &cobra.Command{
	Use:   "rm <id>",
	Short: "Eliminar tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID inválido:", args[0])
			return
		}
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error cargando:", err)
			return
		}
		i, err := findTaskIndexByID(tasks, id)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Tarea %d no encontrada\n", id)
			return
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		tasks = append(tasks[:i], tasks[i+1:]...)
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error guardando:", err)
			return
		}
		fmt.Printf("Tarea %d eliminada\n", id)
	},
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Mostrar versión",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskcli", version)
	},
}
