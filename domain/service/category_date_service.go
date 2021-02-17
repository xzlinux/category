package service

import (
	"github.com/xzlinux/category/domain/model"
	"github.com/xzlinux/category/domain/respository"
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

func NewCategoryDataService(categoryRepository respository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{categoryRepository}
}

func (u *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return u.CategoryRepository.CreateCategory(category)
}
func (u *CategoryDataService) DeleteCategory(categoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(categoryID)
}
func (u *CategoryDataService) UpdateCategory(category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(category)
}
func (u *CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return &u.CategoryRepository.FindCategoryByID(categoryID)
}
func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

func (u *CategoryDataService) FindCategoryByName(categoryName string) (*model.Category, error) {
	return &u.CategoryRepository.FindCategoryByName(categoryName)
}

func (u *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}
func (u *CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}
