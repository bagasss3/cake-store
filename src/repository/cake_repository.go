package repository

import (
	"cake-store/src/model"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type cakeRepository struct {
	db *sql.DB
}

func NewCakeRepository(db *sql.DB) model.CakeRepository {
	return &cakeRepository{
		db: db,
	}
}

func (c *cakeRepository) Save(ctx context.Context, cake *model.Cake) error {
	log := logrus.WithFields(logrus.Fields{
		"message": "Save Cake Repository",
		"cake":    cake,
	})

	sql := "INSERT INTO cakes(title,description,rating,image,created_at,updated_at,deleted_at) VALUES (?,?,?,?,?,?,?)"
	res, err := c.db.ExecContext(ctx, sql, cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt)
	if err != nil {
		log.Error(err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Error(err)
		return err
	}

	cake.Id = int(id)
	return nil
}

func (c *cakeRepository) Update(ctx context.Context, cake *model.Cake) error {
	log := logrus.WithFields(logrus.Fields{
		"message": "Update Cake Repository",
		"cake":    cake,
	})

	query := `UPDATE cakes SET title = ?, description = ?, rating = ?, Image = ?, updated_at = ? WHERE id = ?;
    `
	_, err := c.db.ExecContext(ctx, query, cake.Title, cake.Description, cake.Rating, cake.Image, cake.UpdatedAt, cake.Id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *cakeRepository) Delete(ctx context.Context, cake *model.Cake) error {
	log := logrus.WithFields(logrus.Fields{
		"message": "Delete Cake Repository",
		"cake":    cake,
	})

	query := `UPDATE cakes SET deleted_at = ? where id = ?;
    `
	_, err := c.db.ExecContext(ctx, query, cake.DeletedAt, cake.Id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *cakeRepository) FindAll(ctx context.Context) ([]*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Find All Cake Repository",
	})

	sql := "SELECT * FROM cakes WHERE deleted_at IS null ORDER BY rating DESC, title ASC"
	rows, err := c.db.QueryContext(ctx, sql)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	cakes := make([]*model.Cake, 0)

	for rows.Next() {
		cake := &model.Cake{}
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		cakes = append(cakes, cake)
	}
	return cakes, nil
}

func (c *cakeRepository) FindById(ctx context.Context, id int) (*model.Cake, error) {
	log := logrus.WithFields(logrus.Fields{
		"message": "Find By ID Cake Repository",
		"id":      id,
	})

	sql := "SELECT * FROM cakes where id = ? AND deleted_at is null"
	rows, err := c.db.QueryContext(ctx, sql, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	cake := &model.Cake{}
	if rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return cake, nil
	}
	return nil, nil
}
