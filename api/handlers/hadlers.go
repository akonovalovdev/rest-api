package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"

	"rest-api/entity"
)

var tasks = map[string]entity.Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// GetTasks - бработчик для получения всех задач
// (ИЗ ЗАДАНИЯ) При ошибке сервер должен вернуть статус 500 Internal Server Error.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при сериализации данных, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if len(tasks) == 0 {
		http.Error(w, "Ошибка при получении данных. На данный момент у вас нет задач", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при записи данных, %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

// GetTask - oбработчик для получения задачи по ID
// (ИЗ ЗАДАНИЯ) В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
func GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, fmt.Sprintf("Ошибка при получении данных. Задача c id = %s, не найдена", id), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при сериализации данных, %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при записи данных, %s", err.Error()), http.StatusBadRequest)
		return
	}
}

// AddTask - обработчик для отправки задачи на сервер
// (ИЗ ЗАДАНИЯ) При ошибке сервер должен вернуть статус 400 Bad Request
func AddTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при чтении тела запроса: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при разборе JSON: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// ID и Description - обязательные поля для заполнения при создании новой задачи
	if task.ID == "" || task.Description == "" {
		http.Error(w, "Ошибка при добавлении данных. Поля ID и Description обязательны", http.StatusBadRequest)
		return
	}

	// Проверяем, что задачи с таким ID нет в мапе
	_, ok := tasks[task.ID]
	if ok {
		http.Error(w, fmt.Sprintf("Ошибка при добавлении данных. Задача c id = %s, уже существует", task.ID), http.StatusBadRequest)
		return
	}

	tasks[task.ID] = task
	w.WriteHeader(http.StatusCreated)
}

// UpdateTask - обработчик для обновления задачи по ID
// Дополнительно сделал самостоятельно, т.к. в задании не было, а эндпоинт POST не может обновлять данные
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при чтении тела запроса: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при разборе JSON: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// ID и Description - обязательные поля для заполнения при создании новой задачи
	if task.ID == "" || task.Description == "" {
		http.Error(w, "Ошибка при обновлении данных. Поля ID и Description обязательны", http.StatusBadRequest)
		return
	}

	// Проверяем, что задачи с таким ID нет в мапе
	_, ok := tasks[task.ID]
	if !ok {
		http.Error(w, fmt.Sprintf("Ошибка при обновлении данных. Задача c id = %s, не найдена", task.ID), http.StatusBadRequest)
		return
	}

	tasks[task.ID] = task
	w.WriteHeader(http.StatusOK)
}

// DeleteTask - бработчик удаления задачи по ID
// (ИЗ ЗАДАНИЯ) В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, ok := tasks[id]
	if !ok {
		http.Error(w, fmt.Sprintf("Ошибка при удалении данных. Задача c id = %s, не найдена", id), http.StatusBadRequest)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusOK)
}
