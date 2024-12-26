package pg_rep

import (
	"testing"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		input   models.SignUpInput
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("testUsername", "testPassword").
					WillReturnRows(rows)
			},
			input: models.SignUpInput{
				Username: "testUsername",
				Password: "testPassword",
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("testUsername", "").
					WillReturnRows(rows)
			},
			input: models.SignUpInput{
				Username: "testUsername",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.CreateUser(testCase.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})

	}
}

func TestAuth_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	type args struct {
		username string
		password string
	}

	testTable := []struct {
		name    string
		mock    func()
		input   args
		want    models.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
					AddRow(1, "testUsername", "testPassword", false)
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("testUsername", "testPassword").WillReturnRows(rows)
			},
			input: args{
				username: "testUsername",
				password: "testPassword",
			},
			want: models.User{
				Id:      1,
				IsAdmin: false,
				SignUpInput: models.SignUpInput{
					Username: "testUsername",
					Password: "testPassword",
				},
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"})
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("testUsername", "testPassword").WillReturnRows(rows)
			},
			input: args{
				username: "testUsername",
				password: "testPassword",
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetUser(testCase.input.username, testCase.input.password)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())

		})
	}
}
