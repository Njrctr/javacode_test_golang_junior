package pg_rep

import (
	"log"
	"testing"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestWallet_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWalletPostgres(db)

	testTable := []struct {
		name         string
		mockBehavior func()
		userId       int
		want         uuid.UUID
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "uuid"}).AddRow(1, uuid.UUID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
				mock.ExpectQuery("INSERT INTO wallets").WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO users_wallets").
					WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			userId: 1,
			want:   uuid.UUID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		{
			name: "Empty Fields",
			mockBehavior: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "uuid"}).AddRow(1, uuid.UUID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
				mock.ExpectQuery("INSERT INTO wallets").WillReturnRows(rows)

				mock.ExpectRollback()
			},
			userId:  1,
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.Create(testCase.userId)
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

func TestWallet_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWalletPostgres(db)

	testTable := []struct {
		name         string
		mockBehavior func()
		input        int
		want         []models.Wallet
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func() {

				rows := sqlmock.NewRows([]string{"id", "uuid", "balance", "blocked"}).
					AddRow(1, uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, 10, false).
					AddRow(2, uuid.UUID{0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, 20, false).
					AddRow(3, uuid.UUID{0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, 30, false)

				mock.ExpectQuery("SELECT (.+) FROM wallets w INNER JOIN users_wallets uw on (.+) WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: []models.Wallet{
				{
					Id:      1,
					UUID:    uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
					Balance: 10,
					Blocked: false,
				},
				{
					Id:      2,
					UUID:    uuid.UUID{0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
					Balance: 20,
					Blocked: false,
				},
				{
					Id:      3,
					UUID:    uuid.UUID{0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
					Balance: 30,
					Blocked: false,
				},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.GetAll(testCase.input)
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

func TestWallet_GetByUUID(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWalletPostgres(db)

	testTable := []struct {
		name         string
		mockBehavior func()
		walletUUID   uuid.UUID
		want         models.Wallet
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "uuid", "balance", "blocked"}).
					AddRow(1, uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, 150, true)

				mock.ExpectQuery("SELECT (.+) FROM wallets WHERE (.+)").
					WithArgs(uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}).WillReturnRows(rows)
			},
			walletUUID: uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			want: models.Wallet{
				Id:      1,
				UUID:    uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				Balance: 150,
				Blocked: true,
			},
		},
		{
			name: "Not Found",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "uuid", "balance", "blocked"})

				mock.ExpectQuery("SELECT (.+) FROM wallets WHERE (.+)").
					WithArgs(uuid.UUID{0x1, 0x1, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}).WillReturnRows(rows)
			},
			walletUUID: uuid.UUID{0x1, 0x1, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			wantErr:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.GetByUUID(testCase.walletUUID)
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

func TestWallet_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWalletPostgres(db)

	type args struct {
		userId     int
		walletUUID uuid.UUID
	}

	testTable := []struct {
		name         string
		mockBehavior func()
		input        args
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func() {

				rows := sqlmock.NewRows([]string{"id", "uuid", "balance", "blocked"}).
					AddRow(1, uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, 0, false)
				mock.ExpectQuery("SELECT (.+) FROM wallets WHERE (.+)").
					WithArgs(uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}).WillReturnRows(rows)

				mock.ExpectExec("DELETE FROM wallets w USING users_wallets uw WHERE (.+) AND (.+) AND (.+)").
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			input: args{
				userId:     1,
				walletUUID: uuid.UUID{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			err := r.Delete(testCase.input.userId, testCase.input.walletUUID)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())

		})

	}
}

func TestWallet_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWalletPostgres(db)

	testTable := []struct {
		name         string
		mockBehavior func()
		input        models.WalletUpdate
		wantErr      bool
	}{
		{
			name: "Ok_WITHDRAW",
			mockBehavior: func() {
				mock.ExpectExec("UPDATE wallets SET (.+) WHERE (.+)").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: models.WalletUpdate{
				OperationType: "WITHDRAW",
				Amount:        150,
				WalletUUID:    uuid.UUID{0x1, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
		},
		{
			name: "Ok_DEPOSIT",
			mockBehavior: func() {
				mock.ExpectExec("UPDATE wallets SET (.+) WHERE (.+)").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: models.WalletUpdate{
				OperationType: "DEPOSIT",
				Amount:        150,
				WalletUUID:    uuid.UUID{0x1, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			err := r.Update(testCase.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())

		})
	}
}
