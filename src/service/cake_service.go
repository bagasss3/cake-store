package service

import (
	"cake-store/src/constant"
	"cake-store/src/model"
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type cakeService struct {
	cakeRepository model.CakeRepository
}

func NewCakeService(cakeRepository model.CakeRepository) model.CakeService {
	return &cakeService{
		cakeRepository: cakeRepository,
	}
}

func (c *cakeService) Create(ctx context.Context, req model.CreateUpdateRequest) (*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Create Cake Service",
		"req":     req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	cake := &model.Cake{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := c.cakeRepository.Save(ctx, cake)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return cake, err
}

func (c *cakeService) FindById(ctx context.Context, cakeId int) (*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Find By ID Cake Service",
		"cakeId":  cakeId,
	})

	if cakeId == 0 {
		log.Error(constant.ErrInvalidArgument)
		return nil, constant.ErrInvalidArgument
	}

	cake, err := c.cakeRepository.FindById(ctx, cakeId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if cake == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	return cake, nil
}

func (c *cakeService) Update(ctx context.Context, req model.CreateUpdateRequest, cakeId int) (*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Update Cake Service",
		"req":     req,
	})

	cake, err := c.FindById(ctx, cakeId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	cake.Title = req.Title
	cake.Description = req.Description
	cake.Rating = req.Rating
	cake.Image = req.Image
	cake.UpdatedAt = time.Now()

	if err = c.cakeRepository.Update(ctx, cake); err != nil {
		log.Error(err)
		return nil, err
	}

	return cake, err
}

func (c *cakeService) FindAll(ctx context.Context) ([]*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Find All Cake Service",
	})

	cakes, err := c.cakeRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return cakes, nil
}

func (c *cakeService) Delete(ctx context.Context, cakeId int) (*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Delete Cake Service",
		"cakeId":  cakeId,
	})

	cake, err := c.FindById(ctx, cakeId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if cake.DeletedAt != nil {
		log.Error(constant.ErrAlreadyDeleted)
		return nil, constant.ErrAlreadyDeleted
	}

	cake.DeletedAt = new(time.Time)
	*cake.DeletedAt = time.Now()

	if err = c.cakeRepository.Delete(ctx, cake); err != nil {
		log.Error(err)
		return nil, err
	}

	return cake, err
}
