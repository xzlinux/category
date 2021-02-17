package service

import (
	"github.com/xzlinux/category/domain/model"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(uint64) ([]model.Category, error)
}

type CategoryDataService struct {
	CategoryRepository respository.ICategoryRepository
}
