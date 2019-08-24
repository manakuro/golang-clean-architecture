package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func NewDBMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}

	gdb.LogMode(true)

	return gdb, mock, nil
}
