package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const storeDirName = ".taskcli"
const storeFileName = "tasks.json"

func tasksFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, storeDirName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, storeFileName), nil
}

func loadTasks() ([]Task, error) {
	path, err := tasksFilePath()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	path, err := tasksFilePath()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func findTaskIndexByID(tasks []Task, id int) (int, error) {
	for i, t := range tasks {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, os.ErrNotExist
}
