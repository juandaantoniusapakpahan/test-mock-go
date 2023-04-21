package repository

import (
	"github.com/juandaantoniusapakpahan/test-mock-go/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (m *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	category := args.Get(0).(entity.Category)
	return &category

}
