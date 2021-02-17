package handler

import (
	"context"

	"github.com/prometheus/common/log"
	"github.com/xzlinux/category/domain/model"
	"github.com/xzlinux/category/domain/service"
	"github.com/xzlinux/category/proto/category"
	"github.com/xzlinux/common"
)

type Category struct {
	CategoryDataService service.ICartDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId
	return nil

}
func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil

}
func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}
func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}
func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIdRequest, response *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(request.categoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}
func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}
func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)

	return nil
}
func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)

	return nil
}

func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &category.categoryToResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}
