package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

type repoTestKit struct {
	dbmock sqlmock.Sqlmock
	db     *sql.DB
	ctrl   *gomock.Controller
}

func initializeRepoTestKit(t *testing.T) (kit *repoTestKit, close func()) {
	dbconn, dbmock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}

	ctrl := gomock.NewController(t)

	tk := &repoTestKit{
		ctrl:   ctrl,
		dbmock: dbmock,
		db:     dbconn,
	}

	close = func() {
		if conn := tk.db; conn != nil {
			_ = conn.Close()
		}
	}

	return tk, close
}
