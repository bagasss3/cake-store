package controller

import (
	"cake-store/src/constant"
	"cake-store/src/model"
	"cake-store/src/model/mock"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestHTTP_handleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCakeService := mock.NewMockCakeService(ctrl)
	cakeController := &cakeController{
		cakeService: mockCakeService,
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
		ec := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/cakes", strings.NewReader(`
		{
            "title":"Kue Test",
            "description" :"Desc test",
            "rating": 5.5,
            "image":"test image"
		}`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ctx := context.TODO()

		mockCakeService.EXPECT().Create(ctx, model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}).Times(1).Return(cake, nil)

		err := cakeController.HandleCreate()(ectx)
		require.NoError(t, err)

		resBody := map[string]interface{}{}
		err = json.NewDecoder(rec.Result().Body).Decode(&resBody)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("handle error - validate", func(t *testing.T) {
		ec := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/cakes", strings.NewReader(`
		{
            "title":"K",
            "description" :"Desc test",
            "rating": 5.5,
            "image":"test image"
		}`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ctx := context.TODO()
		cake := &model.Cake{
			Title:       "K",
			Description: "Desc test",
			Rating:      5.5,
			Image:       "test image",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}

		mockCakeService.EXPECT().Create(ctx, model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}).Times(1).Return(nil, constant.ErrInvalidArgument)

		err := cakeController.HandleCreate()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("handle error - internal", func(t *testing.T) {
		ec := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/cakes", strings.NewReader(`
		{
            "title":"Kaaaa",
            "description" :"Desc test",
            "rating": 5.5,
            "image":"test image"
		}`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ctx := context.TODO()
		cake := &model.Cake{
			Title:       "Kaaaa",
			Description: "Desc test",
			Rating:      5.5,
			Image:       "test image",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}

		mockCakeService.EXPECT().Create(ctx, model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}).Times(1).Return(nil, constant.ErrInternal)

		err := cakeController.HandleCreate()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}

func TestHTTP_handleUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCakeService := mock.NewMockCakeService(ctrl)
	cakeController := &cakeController{
		cakeService: mockCakeService,
	}

	cake := &model.Cake{
		Id:          1,
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodPut, "/cakes", strings.NewReader(`
		{
            "title":"Kue Test",
            "description" :"Desc test",
            "rating": 5.5,
            "image":"test image"
		}`,
		))
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))
		mockCakeService.EXPECT().Update(ctx, model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}, cake.Id).Times(1).Return(cake, nil)

		err := cakeController.HandleUpdate()(ectx)
		require.NoError(t, err)

		resBody := map[string]interface{}{}
		err = json.NewDecoder(rec.Result().Body).Decode(&resBody)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("handle error - validate", func(t *testing.T) {
		ec := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/cakes", strings.NewReader(`
		{
            "title":"K",
            "description" :"Desc test",
            "rating": 5.5,
            "image":"test image"
		}`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))
		ctx := context.TODO()
		cake := &model.Cake{
			Id:          1,
			Title:       "K",
			Description: "Desc test",
			Rating:      5.5,
			Image:       "test image",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}

		mockCakeService.EXPECT().Update(ctx, model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}, cake.Id).Times(1).Return(nil, constant.ErrInvalidArgument)

		err := cakeController.HandleUpdate()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("handle error - internal", func(t *testing.T) {
		ec := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/cakes", strings.NewReader(`
		{
            "title":"Kaaaa",
            "description" :"Desc test",
            "rating": 5.5,
            "image":"test image"
		}`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))
		ctx := context.TODO()
		cake := &model.Cake{
			Id:          1,
			Title:       "Kaaaa",
			Description: "Desc test",
			Rating:      5.5,
			Image:       "test image",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}

		mockCakeService.EXPECT().Update(ctx, model.CreateUpdateRequest{
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			Image:       cake.Image,
		}, cake.Id).Times(1).Return(nil, constant.ErrInternal)

		err := cakeController.HandleUpdate()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}

func TestHTTP_handleDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCakeService := mock.NewMockCakeService(ctrl)
	cakeController := &cakeController{
		cakeService: mockCakeService,
	}

	cake := &model.Cake{
		Id:          1,
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodDelete, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))
		mockCakeService.EXPECT().Delete(ctx, cake.Id).Times(1).Return(cake, nil)

		err := cakeController.HandleDelete()(ectx)
		require.NoError(t, err)

		resBody := map[string]interface{}{}
		err = json.NewDecoder(rec.Result().Body).Decode(&resBody)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("handle error - not found", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodDelete, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))

		mockCakeService.EXPECT().Delete(ctx, cake.Id).Times(1).Return(nil, constant.ErrNotFound)

		err := cakeController.HandleDelete()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusNotFound, rec.Result().StatusCode)
	})

	t.Run("handle error - internal", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodDelete, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))

		mockCakeService.EXPECT().Delete(ctx, cake.Id).Times(1).Return(nil, constant.ErrInternal)

		err := cakeController.HandleDelete()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}

func TestHTTP_handleFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCakeService := mock.NewMockCakeService(ctrl)
	cakeController := &cakeController{
		cakeService: mockCakeService,
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
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodGet, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)

		mockCakeService.EXPECT().FindAll(ctx).Times(1).Return(cakes, nil)

		err := cakeController.HandleFindAll()(ectx)
		require.NoError(t, err)

		resBody := map[string]interface{}{}
		err = json.NewDecoder(rec.Result().Body).Decode(&resBody)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("handle not found", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		cakes := make([]*model.Cake, 0)
		req := httptest.NewRequest(http.MethodGet, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)

		mockCakeService.EXPECT().FindAll(ctx).Times(1).Return(cakes, nil)

		err := cakeController.HandleFindAll()(ectx)
		require.NoError(t, err)

		resBody := map[string]interface{}{}
		err = json.NewDecoder(rec.Result().Body).Decode(&resBody)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("handle error - internal", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodGet, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)

		mockCakeService.EXPECT().FindAll(ctx).Times(1).Return(nil, constant.ErrInternal)

		err := cakeController.HandleFindAll()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}

func TestHTTP_handleFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCakeService := mock.NewMockCakeService(ctrl)
	cakeController := &cakeController{
		cakeService: mockCakeService,
	}

	cake := &model.Cake{
		Id:          1,
		Title:       "Kue Test",
		Description: "Desc test",
		Rating:      5.5,
		Image:       "test image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	t.Run("ok", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodGet, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))

		mockCakeService.EXPECT().FindById(ctx, cake.Id).Times(1).Return(cake, nil)

		err := cakeController.HandleFindById()(ectx)
		require.NoError(t, err)

		resBody := map[string]interface{}{}
		err = json.NewDecoder(rec.Result().Body).Decode(&resBody)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("handle not found", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodGet, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))

		mockCakeService.EXPECT().FindById(ctx, cake.Id).Times(1).Return(nil, constant.ErrNotFound)

		err := cakeController.HandleFindById()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusNotFound, rec.Result().StatusCode)
	})

	t.Run("handle error - internal", func(t *testing.T) {
		ec := echo.New()
		rec := httptest.NewRecorder()
		ctx := context.TODO()
		req := httptest.NewRequest(http.MethodGet, "/cakes", nil)
		req.Header.Set("Content-Type", "application/json")
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues(strconv.Itoa(cake.Id))

		mockCakeService.EXPECT().FindById(ctx, cake.Id).Times(1).Return(nil, constant.ErrInternal)

		err := cakeController.HandleFindById()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.EqualValues(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}
