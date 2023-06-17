package usecase

import (
	"errors"
	"final-project3/pkg/task/dto"
	"final-project3/pkg/task/model"
	"final-project3/pkg/task/repository"
)

type UsecaseInterfaceTask interface {
	CreateNewTask(req dto.TaskRequest) (model.Task, error)
	GetAllTask() (*[]dto.TaskResponse, error)
	UpdateTaskById(taskId int, input dto.EditTask) (model.Task, error)
	UpdateStatusByTaskId(taskId int, input dto.EditTask) (model.Task, error)
	UpdateCategoryByTaskId(taskId int, input dto.EditTask) (model.Task, error)
	DeleteTaskById(taskId int) error
}

type usecaseTask struct {
	repository repository.RepositoryInterfaceTask
}

func InitUsecaseTask(repository repository.RepositoryInterfaceTask) UsecaseInterfaceTask {
	return &usecaseTask{
		repository: repository,
	}
}

// CreateNewTask implements UsecaseInterfaceTask
func (u *usecaseTask) CreateNewTask(req dto.TaskRequest) (model.Task, error) {
	var task model.Task
	isTaskExist, _ := u.repository.GetTaskById(int(task.Id))
	if isTaskExist.Id != 0 {
		return task, errors.New("tasks already exist")
	}

	payload := model.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserId:      req.UserId,
		CategoryId:  req.CategoryId,
	}
	newTask, err := u.repository.CreateNewTask(payload)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

// GetAllTask implements UsecaseInterfaceTask
func (u *usecaseTask) GetAllTask() (*[]dto.TaskResponse, error) {
	tasks, err := u.repository.GetAllTask()
	if err != nil {
		return nil, err
	}

	tasksPayload := []dto.TaskResponse{}
	for _, newTask := range tasks {
		task := dto.TaskResponse{
			Id:          newTask.Id,
			Title:       newTask.Title,
			Description: newTask.Description,
			Status:      newTask.Status,
			UserId:      newTask.UserId,
			CategoryId:  newTask.CategoryId,
			CreatedAt:   newTask.CreatedAt,
			UpdatedAt:   newTask.UpdatedAt,
			User: dto.UserData{
				Id:       newTask.User.Id,
				Email:    newTask.User.Email,
				Fullname: newTask.User.Fullname,
			},
		}

		tasksPayload = append(tasksPayload, task)
	}

	return &tasksPayload, nil
}

// UpdateTaskById implements UsecaseInterfaceTask
func (u *usecaseTask) UpdateTaskById(taskId int, input dto.EditTask) (model.Task, error) {
	task, err := u.repository.GetTaskById(taskId)
	if err != nil {
		return task, err
	}

	if input.Title != "" {
		task.Title = input.Title
	}

	if input.Description != "" {
		task.Description = input.Description
	}

	return u.repository.UpdateTaskById(task)
}

// UpdateStatusByTaskId implements UsecaseInterfaceTask
func (u *usecaseTask) UpdateStatusByTaskId(taskId int, input dto.EditTask) (model.Task, error) {
	task, err := u.repository.GetTaskById(taskId)
	if err != nil {
		return task, err
	}

	if input.Status {
		task.Status = input.Status
	}

	return u.repository.UpdateTaskById(task)
}

// UpdateCategoryByTaskId implements UsecaseInterfaceTask
func (u *usecaseTask) UpdateCategoryByTaskId(taskId int, input dto.EditTask) (model.Task, error) {
	task, err := u.repository.GetTaskById(taskId)
	if err != nil {
		return task, err
	}

	if input.CategoryId != 0 {
		task.CategoryId = input.CategoryId
	}

	return u.repository.UpdateTaskById(task)
}

// DeleteTaskById implements UsecaseInterfaceTask
func (u *usecaseTask) DeleteTaskById(taskId int) error {
	err := u.repository.DeleteTaskById(taskId)
	if err != nil {
		return err
	}

	return nil
}
