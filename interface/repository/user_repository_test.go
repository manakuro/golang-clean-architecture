package repository_test

import (
	"regexp"
	"testing"

	"github.com/manakuro/golang-clean-architecture/testutil"

	"github.com/manakuro/golang-clean-architecture/interface/repository"

	"github.com/stretchr/testify/assert"

	"github.com/manakuro/golang-clean-architecture/domain/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func setup(t *testing.T) (r repository.UserRepository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, _ := testutil.NewDBMock(t)

	r = repository.NewUserRepository(db)

	return r, mock, func() {
		db.Close()
	}
}

func TestUserRepository_FindAll(t *testing.T) {
	t.Helper()
	r, mock, teardown := setup(t)
	defer teardown()

	cases := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, u []*model.User, err error)
	}{
		"正常なSQLのとき": {
			arrange: func(t *testing.T) {
				u := model.User{ID: 1, Name: "manato", Age: "20"}
				c := model.CreditCard{ID: 1, UserID: 1, Number: "111111"}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(u.ID, u.Name, u.Age))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "number"}).AddRow(c.ID, c.UserID, c.Number))
			},
			assert: func(t *testing.T, u []*model.User, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 1, len(u))
				assert.NotEmpty(t, u[0].CreditCards)
			},
		},
		"エラーのSQLのとき": {
			arrange: func(t *testing.T) {
				u := model.User{ID: 1, Name: "manato", Age: "20"}
				c := model.CreditCard{ID: 1, UserID: 1, Number: "111111"}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(u.ID, u.Name, u.Age))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "number"}).AddRow(c.ID, c.UserID, c.Number))
			},
			assert: func(t *testing.T, u []*model.User, err error) {
				//assert.NotEmpty(t, err)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			tt.arrange(t)

			u, err := r.FindAll([]*model.User{})

			tt.assert(t, u, err)
		})
	}
}
