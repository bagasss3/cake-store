package service

import (
	"cake-store/src/constant"
	"cake-store/src/model"
	"cake-store/src/model/mock"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCakeService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockCakeRepo := mock.NewMockCakeRepository(ctrl)

	cakeService := &cakeService{
		cakeRepository: mockCakeRepo,
	}

	cake := &model.Cake{
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}

		mockCakeRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Times(1).Return(nil)
		res, err := cakeService.Create(ctx, cakeReq)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("validate error", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       "a",
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}
		mockCakeRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Times(0).Return(nil)
		res, err := cakeService.Create(ctx, cakeReq)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error from repo", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}
		mockCakeRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("err db"))
		res, err := cakeService.Create(ctx, cakeReq)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestCakeService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockCakeRepo := mock.NewMockCakeRepository(ctrl)

	cakeService := &cakeService{
		cakeRepository: mockCakeRepo,
	}
	id := 1
	cake := &model.Cake{
		Id:          id,
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}

		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(cake, nil)
		mockCakeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(nil)

		res, err := cakeService.Update(ctx, cakeReq, cake.Id)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("validate error", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       "a",
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}
		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(cake, nil)
		mockCakeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(0).Return(nil)
		res, err := cakeService.Update(ctx, cakeReq, cake.Id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("id not found", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}
		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(nil, nil)
		mockCakeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(0).Return(nil)
		res, err := cakeService.Update(ctx, cakeReq, cake.Id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error from repo", func(t *testing.T) {
		cakeReq := model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}
		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(cake, nil)
		mockCakeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("err db"))
		res, err := cakeService.Update(ctx, cakeReq, cake.Id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestCakeService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockCakeRepo := mock.NewMockCakeRepository(ctrl)

	cakeService := &cakeService{
		cakeRepository: mockCakeRepo,
	}

	var cakes []*model.Cake
	cake := &model.Cake{
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	cake2 := &model.Cake{
		Title:       "Kue Test 2",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	cakes = append(cakes, cake, cake2)

	t.Run("ok", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindAll(gomock.Any()).Times(1).Return(cakes, nil)
		res, err := cakeService.FindAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("data empty", func(t *testing.T) {
		cakes := make([]*model.Cake, 0)

		mockCakeRepo.EXPECT().FindAll(gomock.Any()).Times(1).Return(cakes, nil)
		res, err := cakeService.FindAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("error from repo", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindAll(gomock.Any()).Times(1).Return(nil, errors.New("err db"))
		res, err := cakeService.FindAll(ctx)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestCakeService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockCakeRepo := mock.NewMockCakeRepository(ctrl)

	cakeService := &cakeService{
		cakeRepository: mockCakeRepo,
	}

	var id int = 1

	cake := &model.Cake{
		Id:          id,
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindById(gomock.Any(), id).Times(1).Return(cake, nil)
		res, err := cakeService.FindById(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("data empty", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindById(gomock.Any(), id).Times(1).Return(nil, constant.ErrNotFound)
		res, err := cakeService.FindById(ctx, id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error from repo", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindById(gomock.Any(), id).Times(1).Return(nil, errors.New("err db"))
		res, err := cakeService.FindById(ctx, id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestCakeService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockCakeRepo := mock.NewMockCakeRepository(ctrl)

	cakeService := &cakeService{
		cakeRepository: mockCakeRepo,
	}
	id := 1
	cake := &model.Cake{
		Id:          id,
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(cake, nil)
		mockCakeRepo.EXPECT().Delete(gomock.Any(), cake).Times(1).Return(nil)

		res, err := cakeService.Delete(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("not found", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(nil, constant.ErrNotFound)
		mockCakeRepo.EXPECT().Delete(gomock.Any(), cake).Times(0).Return(nil)
		res, err := cakeService.Delete(ctx, cake.Id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("already deleted", func(t *testing.T) {
		cake := &model.Cake{
			Id:          id,
			Title:       "Kue Test",
			Description: "Desc test",
			Rating:      5.5,
			Image:       "test image",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}
		cake.DeletedAt = new(time.Time)
		*cake.DeletedAt = time.Now()

		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(cake, nil)
		mockCakeRepo.EXPECT().Delete(gomock.Any(), cake).Times(0).Return(nil)
		res, err := cakeService.Delete(ctx, cake.Id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error from repo", func(t *testing.T) {
		mockCakeRepo.EXPECT().FindById(gomock.Any(), cake.Id).Times(1).Return(cake, nil)
		mockCakeRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("err db"))
		res, err := cakeService.Delete(ctx, cake.Id)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
