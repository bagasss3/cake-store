package repository

import (
	"cake-store/src/model"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCakeRepository_Create(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := cakeRepository{
		db: kit.db,
	}

	ctx := context.TODO()

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
		mock.ExpectExec("INSERT INTO cakes").
			WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := repo.Save(ctx, cake)
		require.NoError(t, err)
	})

	t.Run("failed to save cake", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO cakes").
			WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt).
			WillReturnError(errors.New("db error"))
		err := repo.Save(ctx, cake)
		require.Error(t, err)
	})
}

func TestCakeRepository_Update(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := cakeRepository{
		db: kit.db,
	}

	ctx := context.TODO()

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
		mock.ExpectExec("UPDATE cakes").
			WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.UpdatedAt, cake.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := repo.Update(ctx, cake)
		require.NoError(t, err)
	})

	t.Run("failed to update cake", func(t *testing.T) {
		mock.ExpectExec("UPDATE cakes").
			WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt).
			WillReturnError(errors.New("db error"))
		err := repo.Update(ctx, cake)
		require.Error(t, err)
	})
}

func TestCakeRepository_Delete(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := cakeRepository{
		db: kit.db,
	}

	ctx := context.TODO()

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
	cake.DeletedAt = new(time.Time)
	*cake.DeletedAt = time.Now()

	t.Run("ok", func(t *testing.T) {
		mock.ExpectExec("UPDATE cakes").
			WithArgs(cake.DeletedAt, cake.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := repo.Delete(ctx, cake)
		require.NoError(t, err)
	})

	t.Run("failed to delete cake", func(t *testing.T) {
		mock.ExpectExec("UPDATE cakes").
			WithArgs(cake.DeletedAt, cake.Id).
			WillReturnError(errors.New("db error"))
		err := repo.Delete(ctx, cake)
		require.Error(t, err)
	})
}

func TestCakeRepository_FindAll(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := cakeRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	t.Run("ok - found", func(t *testing.T) {
		resRows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"}).
			AddRow(1, "Kue Test", "Desc test", 5.5, "test image", time.Now(), time.Now(), nil).
			AddRow(2, "Kue Test 2", "Desc test", 6, "test image", time.Now(), time.Now(), nil)

		mock.ExpectQuery("SELECT \\* FROM cakes WHERE deleted_at IS null ORDER BY rating DESC, title ASC").
			WillReturnRows(resRows)

		res, err := repo.FindAll(ctx)
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, 2, len(res))
	})

	t.Run("ok - not found", func(t *testing.T) {
		resRows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"})

		mock.ExpectQuery("SELECT \\* FROM cakes WHERE deleted_at IS null ORDER BY rating DESC, title ASC").
			WillReturnRows(resRows)

		res, err := repo.FindAll(ctx)
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, 0, len(res))
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("SELECT \\* FROM cakes WHERE deleted_at IS null ORDER BY rating DESC, title ASC").
			WillReturnError(errors.New("invalid db"))

		res, err := repo.FindAll(ctx)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestCakeRepository_FindByID(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := cakeRepository{
		db: kit.db,
	}

	ctx := context.TODO()

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

	t.Run("ok - retrieve from db", func(t *testing.T) {
		resRows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"}).
			AddRow(cake.Id, cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt)
		mock.ExpectQuery("SELECT \\* FROM cakes").
			WithArgs(cake.Id).
			WillReturnRows(resRows)

		res, err := repo.FindById(ctx, cake.Id)
		require.NoError(t, err)
		require.NotNil(t, res)
	})

	t.Run("not found", func(t *testing.T) {
		resRows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"})
		mock.ExpectQuery("SELECT \\* FROM cakes").
			WithArgs(cake.Id).
			WillReturnRows(resRows)

		res, err := repo.FindById(ctx, cake.Id)
		require.NoError(t, err)
		require.Nil(t, res)
	})

	t.Run("err db", func(t *testing.T) {
		mock.ExpectQuery("SELECT \\* FROM cakes").
			WithArgs(cake.Id).
			WillReturnError(errors.New("err db"))

		res, err := repo.FindById(ctx, cake.Id)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
