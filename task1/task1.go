package task1

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type Task struct {
	Name      string    `json:"taskname"`
	Status    status    `json:"status"`
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// конструктор для задачи
func CreateTask(name string, id int) (*Task, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	for k := range taskStorage {
		if k == id {
			return nil, errors.New("invalid id")
		}
	}

	return &Task{Name: name, ID: id, Status: Todo, CreatedAt: time.Now()}, nil

}

type status string

var (
	Todo      status = "to do"
	InProgres status = "in progress"
	Done      status = "done"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrEmptyName     = errors.New("task name cannot be empty")
	ErrInvalidStatus = errors.New("invalid status")
	ErrNoData        = errors.New("no data for request")
)

var taskStorage = make(map[int]*Task)

// записывает в файл в формате JSON из TaskStorage
func WriteInFile(fileName string) error {
	var file *os.File
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)

	for _, v := range taskStorage {
		err = encoder.Encode(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// считывает из файла и переводит из JSON в структуры Task, храним в TaskStorage
func ReadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	var t Task
	for {
		if err := decoder.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		taskStorage[t.ID] = &t
	}
	return nil
}

// добавляем task в TaskStorage
func AddTask(t *Task) error {
	if t.Name == "" {
		return ErrEmptyName
	}
	taskStorage[t.ID] = t
	return nil
}

// обновляем имя задачи в TaskStorage
func UpdateTaskName(id int, newName string) error {
	if _, ok := taskStorage[id]; !ok {
		return ErrTaskNotFound
	}
	if newName == "" {
		return ErrEmptyName
	}

	taskStorage[id].Name = newName
	taskStorage[id].UpdatedAt = time.Now()

	return nil
}

// обновляем статус задачи в TaskStorage
func UpdateTaskStatus(id int, newStatus status) error {
	if _, ok := taskStorage[id]; !ok {
		return ErrTaskNotFound
	}
	if newStatus == taskStorage[id].Status {
		return errors.New("status must change")
	}

	taskStorage[id].Status = newStatus
	taskStorage[id].UpdatedAt = time.Now()
	return nil
}

// удаляем задачу из TaskStorage
func DeleteTask(id int) error {
	if _, ok := taskStorage[id]; !ok {
		return ErrTaskNotFound
	}

	delete(taskStorage, id)
	return nil
}

// выводим весь список задач
func ShowAllTasks() error {
	if len(taskStorage) == 0 {
		return ErrNoData
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range taskStorage {
		str := fmt.Sprintf("ID: %v, Task name: %v, Status: %v, Creating time: %v, Last update time: %v \n", v.ID, v.Name, v.Status, v.CreatedAt.Format("15:04:05 02/01/2006"), v.UpdatedAt.Format("15:04:05 02/01/2006"))
		_, err := writer.WriteString(str)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

// вывод всех задач со статусом todo
func ShowTodoTasks() error {
	if len(taskStorage) == 0 {
		return ErrNoData
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range taskStorage {
		if v.Status != Todo {
			continue
		}
		str := fmt.Sprintf("ID: %v, Task name: %v, Status: %v, Creating time: %v, Last update time: %v \n", v.ID, v.Name, v.Status, v.CreatedAt.Format("14:59:59 22/01/2006"), v.UpdatedAt.Format("14:59:59 22/01/2006"))
		_, err := writer.WriteString(str)
		if err != nil {
			return err
		}
	}
	return nil
}

// выводит все задачи со статусом in progress
func ShowInprogressTasks() error {
	if len(taskStorage) == 0 {
		return ErrNoData
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range taskStorage {
		if v.Status != InProgres {
			continue
		}
		str := fmt.Sprintf("ID: %v, Task name: %v, Status: %v, Creating time: %v, Last update time: %v \n", v.ID, v.Name, v.Status, v.CreatedAt.Format("14:59:59 22/01/2006"), v.UpdatedAt.Format("14:59:59 22/01/2006"))
		_, err := writer.WriteString(str)
		if err != nil {
			return err
		}
	}
	return nil
}

// выводит все задачи со статусом done
func ShowDoneTasks() error {
	if len(taskStorage) == 0 {
		return ErrNoData
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range taskStorage {
		if v.Status != Done {
			continue
		}
		str := fmt.Sprintf("ID: %v, Task name: %v, Status: %v, Creating time: %v, Last update time: %v \n", v.ID, v.Name, v.Status, v.CreatedAt.Format("14:59:59 22/01/2006"), v.UpdatedAt.Format("14:59:59 22/01/2006"))
		_, err := writer.WriteString(str)
		if err != nil {
			return err
		}
	}
	return nil
}
