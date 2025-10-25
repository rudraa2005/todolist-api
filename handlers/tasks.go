package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"todo-api/models"

	"github.com/go-chi/chi/v5"
)

func AddTasks(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadFile("data/tasks.json")
	if err != nil {
		http.Error(w, "Could not read file", 500)
		return
	}

	var tasks []models.Task
	json.Unmarshal(files, &tasks)

	var newTasks models.Task
	json.NewDecoder(r.Body).Decode(&newTasks)

	newTasks.ID = len(tasks) + 1
	newTasks.Completed = false

	tasks = append(tasks, newTasks)

	updatedData, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile("data/tasks.json", updatedData, 0644)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTasks)
}

func DeleteCompleted(w http.ResponseWriter, r *http.Request) {

	files, err := os.ReadFile("data/tasks.json")
	if err != nil {
		http.Error(w, "Could nor read file", 500)
		return
	}
	var tasks []models.Task
	json.Unmarshal(files, &tasks)

	newTasks := []models.Task{}
	for _, t := range tasks {
		if !t.Completed {
			newTasks = append(newTasks, t)
		}
	}
	updatedData, _ := json.MarshalIndent(newTasks, "", "  ")
	os.WriteFile("data/tasks.json", updatedData, 0644)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTasks)
}

func ToggleCompleted(w http.ResponseWriter, r *http.Request) {

	IdS := chi.URLParam(r, "id")
	Id, err := strconv.Atoi(IdS)

	if err != nil {
		http.Error(w, "Enter Valid Id", http.StatusBadRequest)
		return
	}

	files, err := os.ReadFile("data/tasks.json")
	if err != nil {
		http.Error(w, "Could not read file", 500)
		return
	}

	var tasks []models.Task
	json.Unmarshal(files, &tasks)

	found := false
	for i, t := range tasks {
		if t.ID == Id {
			tasks[i].Completed = !t.Completed
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "The ID does not exist", http.StatusNotFound)
		return
	}

	updatedData, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile("data/tasks.json", updatedData, 0644)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func ViewTasks(w http.ResponseWriter, r *http.Request) {

	files, err := os.ReadFile("data/tasks.json")
	if err != nil {
		http.Error(w, "could not read file", http.StatusInternalServerError)
	}

	var tasks []models.Task
	json.Unmarshal(files, &tasks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
