package handler

import (
	"context"

	"github.com/xzlinux/category/domain/model"
	"github.com/xzlinux/category/domain/service"
	"github.com/xzlinux/category/proto/category"
)

type Category struct {
	CategoryDataService service.ICartDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	category := &model.Category{}
}
func (c *Category) UpdateCategory(context.Context, *CategoryRequest, *UpdateCategoryResponse) error {

}
func (c *Category) DeleteCategory(context.Context, *DeleteCategoryRequest, *DeleteCategoryResponse) error {

}
func (c *Category) FindCategoryByName(context.Context, *FindByNameRequest, *CategoryResponse) error {

}
func (c *Category) FindCategoryByID(context.Context, *FindByIdRequest, *CategoryResponse) error {

}
func (c *Category) FindCategoryByLevel(context.Context, *FindByLevelRequest, *FindAllResponse) error {

}
func (c *Category) FindCategoryByParent(context.Context, *FindByParentRequest, *FindAllResponse) error {

}
func (c *Category) FindAllCategory(context.Context, *FindAllRequest, *FindAllResponse) error {

}
