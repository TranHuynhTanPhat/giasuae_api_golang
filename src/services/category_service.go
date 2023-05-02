package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type CategoryService interface {
	InsertCategory(ctg *entities.Category) error
	UpdateCategory(ctg *entities.Category) error
	DeleteCategory(id int) error
	FindAllCategory() []entities.Category
	FindByID(id int) entities.Category
	FilterCategory(value ...interface{}) []entities.Category
}
type categoryService struct {
	CategoryRepository repositories.CategoryRepository
}

// DeleteCategory implements CategoryService
func (svc *categoryService) DeleteCategory(id int) error {
	return svc.CategoryRepository.DeleteCategory(id)
}

// FilterCategory implements CategoryService
func (svc *categoryService) FilterCategory(value ...interface{}) []entities.Category {
	return svc.CategoryRepository.FilterCategory(value[0])
}

// FindAllCategory implements CategoryService
func (svc *categoryService) FindAllCategory() []entities.Category {
	return svc.CategoryRepository.FindAllCategory()
}

// FindByID implements CategoryService
func (svc *categoryService) FindByID(id int) entities.Category {
	return svc.CategoryRepository.FindByID(id)
}

// InsertCategory implements CategoryService
func (svc *categoryService) InsertCategory(ctg *entities.Category) error {
	return svc.CategoryRepository.InsertCategory(ctg)
}

// UpdateCategory implements CategoryService
func (svc *categoryService) UpdateCategory(ctg *entities.Category) error {
	return svc.CategoryRepository.UpdateCategory(ctg)
}

func NewCategoryService(CategoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryRepository: CategoryRepo,
	}
}
