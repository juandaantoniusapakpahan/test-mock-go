package service

import (
	"testing"

	"github.com/juandaantoniusapakpahan/test-mock-go/entity"
	"github.com/juandaantoniusapakpahan/test-mock-go/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	dataTest := entity.Category{
		Id:   "2",
		Name: "Richard Oldo",
	}
	categoryRepository.Mock.On("FindById", "2").Return(dataTest)

	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dataTest.Id, result.Id)
	assert.Equal(t, dataTest.Name, result.Name)
}

/**
 */

func BenchmarkCategoryFound(b *testing.B) {
	dataTest := entity.Category{
		Id:   "2",
		Name: "Richard Oldo",
	}
	categoryRepository.Mock.On("FindById", "2").Return(dataTest)

	for i := 0; i < b.N; i++ {
		categoryService.Get("2")
	}
}

func BenchmarkCategoryNotFound(b *testing.B) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	for i := 0; i < b.N; i++ {
		categoryService.Get("1")
	}
}
