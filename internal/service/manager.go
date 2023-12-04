package service

import (
	"context"

	"github.com/lalizita/go-crud-boilerplate/internal/entity"
	"github.com/lalizita/go-crud-boilerplate/internal/repository"
)

type TaskService struct {
	ctx               context.Context
	managerRepository repository.ITaskRepository
}

type ITaskService interface {
	CreateConsent(context.Context, entity.TaskDTOInput) (entity.TaskDTOOutput, error)
}

func NewTaskService(ctx context.Context, r repository.ITaskRepository) ITaskService {
	return &TaskService{
		ctx:               ctx,
		managerRepository: r,
	}
}

func (s *TaskService) CreateConsent(ctx context.Context, input entity.TaskDTOInput) (entity.TaskDTOOutput, error) {
	id, err := s.managerRepository.Insert(ctx, input)
	if err != nil {
		return entity.TaskDTOOutput{}, err
	}

	output, err := s.managerRepository.FindOneByID(ctx, id)
	if err != nil {
		return entity.TaskDTOOutput{}, err
	}

	return output, nil
}
