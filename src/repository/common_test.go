package repository

import (
	"cake-store/src/database"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type repoTestKit struct {
	dbmock    sqlmock.Sqlmock
	db        *sql.DB
	ctrl      *gomock.Controller
	miniredis *miniredis.Miniredis
	redis     *redis.Client
}

func initializeRepoTestKit(t *testing.T) (kit *repoTestKit, close func()) {
	dbconn, dbmock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}

	mr, _ := miniredis.Run()
	r := database.NewRedisConn(mr.Addr())

	ctrl := gomock.NewController(t)

	tk := &repoTestKit{
		ctrl:      ctrl,
		dbmock:    dbmock,
		db:        dbconn,
		miniredis: mr,
		redis:     r,
	}

	close = func() {
		if conn := tk.db; conn != nil {
			_ = conn.Close()
		}
		tk.miniredis.Close()
	}

	return tk, close
}
