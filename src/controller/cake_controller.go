package controller

import (
	"cake-store/src/constant"
	"cake-store/src/model"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type cakeController struct {
	cakeService model.CakeService
}

func NewCakeController(cakeService model.CakeService) model.CakeController {
	return &cakeController{
		cakeService: cakeService,
	}
}

func (cC *cakeController) HandleCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := cC.cakeService.Create(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, model.ResponseSuccess{
			Success: true,
			Data:    create,
		})
	}
}

func (cC *cakeController) HandleFindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		cakes, err := cC.cakeService.FindAll(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, model.ResponseSuccess{
			Success: true,
			Data:    cakes,
		})
	}
}

func (cC *cakeController) HandleFindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		cake, err := cC.cakeService.FindById(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, model.ResponseSuccess{
			Success: true,
			Data:    cake,
		})
	}
}

func (cC *cakeController) HandleUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		idStr := c.Param("id")
		log.Printf("Received id parameter: %s", idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		update, err := cC.cakeService.Update(c.Request().Context(), req, id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, model.ResponseSuccess{
			Success: true,
			Data:    update,
		})
	}
}

func (cC *cakeController) HandleDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		delete, err := cC.cakeService.Delete(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, model.ResponseSuccess{
			Success: true,
			Data:    delete,
		})
	}
}
