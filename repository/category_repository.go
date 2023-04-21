package repository

import "github.com/juandaantoniusapakpahan/test-mock-go/entity"

type RepositoryCategory interface {
	FindById(id string) *entity.Category
}
