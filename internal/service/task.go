package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/to404hanga/TaskTracker/internal/domain"
	"github.com/to404hanga/TaskTracker/internal/model"
)

type TaskService struct {
	AutoIncrement int
}

func NewTaskService() *TaskService {
	var autoIncrement int
	if file, err := os.Open("../json/auto-increment.txt"); err == nil {
		defer file.Close()
		fmt.Fscan(file, &autoIncrement)
	}
	return &TaskService{
		AutoIncrement: autoIncrement,
	}
}

func (s *TaskService) ModelSliceToDomainSlice(tasks []model.Task) []domain.Task {
	ret := make([]domain.Task, len(tasks))
	for i, task := range tasks {
		ret[i] = domain.FromModel(task)
	}
	return ret
}

func (s *TaskService) GetTasks(mode model.StatusCode) ([]domain.Task, error) {
	res, err := s.getTasksFromJson(mode)
	if err != nil {
		return nil, err
	}
	return s.ModelSliceToDomainSlice(res), nil
}

func (s *TaskService) GetAllTasks() ([]domain.Task, error) {
	ret := make([]model.Task, 0)
	statusList := []model.StatusCode{model.Todo, model.InProgress, model.Done}
	for _, status := range statusList {
		tasks, err := s.getTasksFromJson(status)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tasks...)
	}
	return s.ModelSliceToDomainSlice(ret), nil
}

func (s *TaskService) AddTask(description string) error {
	todoTasks, err := s.getTasksFromJson(model.Todo)
	if err != nil {
		return err
	}

	s.AutoIncrement++
	now := time.Now().Unix()
	task := model.Task{
		Id:          s.AutoIncrement,
		Description: description,
		Status:      model.Todo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	todoTasks = append(todoTasks, task)

	if err = s.saveTasksToJson(model.Todo, todoTasks); err != nil {
		return err
	}
	return os.WriteFile("../json/auto-increment.txt", []byte(fmt.Sprintf("%d", s.AutoIncrement)), 0644)
}

func (s *TaskService) UpdateTask(id int, description string) error {
	// statusList := []model.StatusCode{model.Todo, model.InProgress, model.Done}
	// for _, status := range statusList {
	// 	tasks, err := s.getTasksFromJson(status)
	// 	if err != nil {
	// 		return fmt.Errorf("error getting tasks: %v", err)
	// 	}
	// 	for i, task := range tasks {
	// 		if task.Id == id {
	// 			tasks[i].Description = description
	// 			tasks[i].UpdatedAt = time.Now().Unix()
	// 			s.saveTasksToJson(model.Todo, tasks)
	// 			return nil
	// 		}
	// 	}
	// }
	// return fmt.Errorf("error getting tasks: cannot find the task id")
	return s.doSomethingById(id, func(idx int, tasks []model.Task) []model.Task {
		tasks[idx].Description = description
		tasks[idx].UpdatedAt = time.Now().Unix()
		return tasks
	})
}

func (s *TaskService) RemoveTask(id int) error {
	// statusList := []model.StatusCode{model.Todo, model.InProgress, model.Done}
	// for _, status := range statusList {
	// 	tasks, err := s.getTasksFromJson(status)
	// 	if err != nil {
	// 		return fmt.Errorf("error getting tasks: %v", err)
	// 	}
	// 	for i, task := range tasks {
	// 		if task.Id == id {
	// 			tasks = append(tasks[:i], tasks[i+1:]...)
	// 			s.saveTasksToJson(model.Todo, tasks)
	// 			return nil
	// 		}
	// 	}
	// }
	// return fmt.Errorf("error getting tasks: cannot find the task id")
	return s.doSomethingById(id, func(idx int, tasks []model.Task) []model.Task {
		return append(tasks[:idx], tasks[idx+1:]...)
	})
}

func (s *TaskService) MarkInProgress(id int) error {
	var task model.Task

	err := s.doSomethingById(id, func(idx int, tasks []model.Task) []model.Task {
		if tasks[idx].Status == model.InProgress {
			return tasks
		}
		task = tasks[idx]
		return append(tasks[:idx], tasks[idx+1:]...)
	})
	if task.Id == 0 {
		return err
	}

	tasks, err := s.getTasksFromJson(model.InProgress)
	if err != nil {
		return err
	}
	task.Status = model.InProgress
	task.UpdatedAt = time.Now().Unix()
	tasks = append(tasks, task)
	return s.saveTasksToJson(model.InProgress, tasks)
}

func (s *TaskService) MarkDone(id int) error {
	var task model.Task

	err := s.doSomethingById(id, func(idx int, tasks []model.Task) []model.Task {
		if tasks[idx].Status == model.Done {
			return tasks
		}
		task = tasks[idx]
		return append(tasks[:idx], tasks[idx+1:]...)
	})
	if task.Id == 0 {
		return err
	}

	tasks, err := s.getTasksFromJson(model.Done)
	if err != nil {
		return err
	}
	task.Status = model.Done
	task.UpdatedAt = time.Now().Unix()
	tasks = append(tasks, task)
	return s.saveTasksToJson(model.Done, tasks)
}

func (s *TaskService) doSomethingById(id int, fn func(idx int, tasks []model.Task) []model.Task) error {
	statusList := []model.StatusCode{model.Todo, model.InProgress, model.Done}
	for _, status := range statusList {
		tasks, err := s.getTasksFromJson(status)
		if err != nil {
			return err
		}
		for i, task := range tasks {
			if task.Id == id {
				newTasks := fn(i, tasks)
				log.Println(newTasks)
				return s.saveTasksToJson(status, newTasks)
			}
		}
	}
	return fmt.Errorf("error getting tasks: cannot find the task id")
}

func (s *TaskService) getTasksFromJson(mode model.StatusCode) ([]model.Task, error) {
	var path string
	switch mode {
	case model.Todo:
		path = "../json/todo.json"
	case model.InProgress:
		path = "../json/in-progress.json"
	case model.Done:
		path = "../json/done.json"
	default:
		return nil, fmt.Errorf("invalid mode: %v", mode)
	}

	// read from json file
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.WriteFile(path, []byte("[]"), 0644)
			return make([]model.Task, 0), nil
		}
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var jsonTasks []model.Task
	if err = json.Unmarshal(data, &jsonTasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %v", err)
	}

	return jsonTasks, nil
}

func (s *TaskService) saveTasksToJson(mode model.StatusCode, tasks []model.Task) error {
	var path string
	switch mode {
	case model.Todo:
		path = "../json/todo.json"
	case model.InProgress:
		path = "../json/in-progress.json"
	case model.Done:
		path = "../json/done.json"
	default:
		return fmt.Errorf("invalid mode: %v", mode)
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error marshalling json: %v", err)
	}

	if err = os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
