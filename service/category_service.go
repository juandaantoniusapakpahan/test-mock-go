package service

import (
	"errors"

	"github.com/juandaantoniusapakpahan/test-mock-go/entity"
	"github.com/juandaantoniusapakpahan/test-mock-go/repository"
)

type CategoryService struct {
	Repository repository.RepositoryCategory
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	result := service.Repository.FindById(id)

	if result == nil {
		return nil, errors.New("No found category")
	} else {
		return result, nil
	}
}
