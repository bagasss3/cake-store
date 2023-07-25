package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateUpdateRequest struct {
	Title       string  `json:"title" validate:"required,min=3,max=60"`
	Description string  `json:"description" validate:"min=3"`
	Rating      float32 `json:"rating" validate:"gt=0,lte=10"`
	Image       string  `json:"image"`
}

func (c *CreateUpdateRequest) Validate() error {
	return validate.Struct(c)
}

type Cake struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      float32    `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type CakeRepository interface {
	Save(ctx context.Context, cake *Cake) error
	Update(ctx context.Context, cake *Cake) error
	Delete(ctx context.Context, cake *Cake) error
	FindAll(ctx context.Context) ([]*Cake, error)
	FindById(ctx context.Context, id int) (*Cake, error)
}

type CakeService interface {
	Create(ctx context.Context, req CreateUpdateRequest) (*Cake, error)
	Update(ctx context.Context, req CreateUpdateRequest, cakeId int) (*Cake, error)
	Delete(ctx context.Context, cakeId int) (*Cake, error)
	FindById(ctx context.Context, cakeId int) (*Cake, error)
	FindAll(ctx context.Context) ([]*Cake, error)
}

type CakeController interface {
	HandleCreate() echo.HandlerFunc
	HandleUpdate() echo.HandlerFunc
	HandleDelete() echo.HandlerFunc
	HandleFindById() echo.HandlerFunc
	HandleFindAll() echo.HandlerFunc
}
