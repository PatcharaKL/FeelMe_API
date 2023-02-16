package users

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{}
type AnyToken struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func (a AnyToken) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
func TestUserLogin(t *testing.T) {
	tests := []struct {
		name         string
		body         *bytes.Buffer
		expectedCode int
	}{
		{
			name: "testSucceed",
			body: bytes.NewBufferString(`{
				"email": "user1",
				"password": "123"
				}`),
			expectedCode: http.StatusOK,
		},
		{
			name: "testUnauthorized",
			body: bytes.NewBufferString(`{
				"email": "user1",
				"password": "1234"
			}`),
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "testScanError",
			body: bytes.NewBufferString(`{
				"email": "user1",
				"password": "123"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testCreateRowError",
			body: bytes.NewBufferString(`{
				"email": "user1",
				"password": "123"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testPrepareError",
			body: bytes.NewBufferString(`{
				"emails": "user1",
				"password": "123"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testInternalServerError",
			body: bytes.NewBufferString(`{
				"emails": "user1",
				"password": "123"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testBadRequest",
			body: bytes.NewBufferString(`{
				"emails": "user1",
				: "124"
			}`),
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "testExecError",
			body: bytes.NewBufferString(`{
				"email": "user1",
				"password": "123"
				}`),
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec, c := setupTestServer(http.MethodPost, "/login", tt.body)
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			var tokenRows *sqlmock.Rows
			accountRow := sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
				"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id", "company_id"}).
				AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
					"Patchara", "Kleebbua", "", "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1, 1)

			tokenRows = sqlmock.NewRows([]string{"refreshToken"}).AddRow("1234568632")
			newtokenRows := sqlmock.NewRows([]string{"refreshToken"}).AddRow("1234")

			if tt.name != "testInternalServerError" {
				if tt.name != "testScanError" {
					mock.ExpectQuery("SELECT \\* FROM accounts WHERE email=\\?").WithArgs("user1").WillReturnRows(accountRow)
					mock.ExpectQuery("SELECT (.+) FROM refresh_token WHERE (.+)").WithArgs(1, true).WillReturnRows(tokenRows)
				} else {
					mock.ExpectQuery("SELECT \\* FROM accounts WHERE email=\\?").WithArgs("user1").WillReturnRows(accountRow)
					mock.ExpectQuery("SELECT refreshToken FROM refresh_token WHERE account_id=? && isValid=?").WithArgs(1).WillReturnRows(tokenRows)
				}
				if tt.name != "testPrepareError" {
					expectPrepare := mock.ExpectPrepare("UPDATE refresh_token SET (.+) WHERE (.+)")
					if tt.name != "testExecError" {
						expectPrepare.ExpectExec().WithArgs(false, "1234568632").WillReturnResult(sqlmock.NewResult(0, 0))
					}
				}

				if tt.name != "testCreateRowError" {
					mock.ExpectQuery("INSERT INTO (.+)").WithArgs(AnyToken{}, 1, AnyTime{}, true).WillReturnRows(newtokenRows)
				}

			}
			h := Handler{db}
			err = h.UserLoginHandler(c)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedCode, rec.Code)
			}
		})
	}
}

func TestUserLogOut(t *testing.T) {
	tests := []struct {
		name         string
		body         *bytes.Buffer
		expectedCode int
	}{
		{
			name: "testSucceed",
			body: bytes.NewBufferString(`{
				"refreshToken":"d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9"
				}`),
			expectedCode: http.StatusOK,
		},
		{
			name: "testBadRequest",
			body: bytes.NewBufferString(`{
				"refreshTokensssss":"4dbb-9c28-963117f4dec9",12313
				}`),
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "testInternalServerError",
			body: bytes.NewBufferString(`{
				"refreshToken":"d840e7da-bd5b-490a-8f2a-83648c114d6c"
				}`),
			expectedCode: http.StatusInternalServerError,
		},

		{
			name: "testExecError",
			body: bytes.NewBufferString(`{
				"refreshToken":"d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9"
				}`),
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec, c := setupTestServer(http.MethodPost, "/users/userlogout", tt.body)

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			sqlmock.NewRows([]string{"refreshToken", "account_id", "exp", "isValid"}).
				AddRow("d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9",
					1, "2023-02-04 07:52:09", true).
				AddRow("d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9",
					1, "2023-02-04 07:52:09", true)
			if tt.name != "testInternalServerError" {
				expectPrepare := mock.ExpectPrepare("UPDATE refresh_token SET (.+) WHERE (.+)")
				if tt.name != "testExecError" {
					expectPrepare.ExpectExec().WithArgs(false, "d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9").WillReturnResult(sqlmock.NewResult(0, 0))
				}
			}
			h := Handler{db}

			err = h.UserLogOutHandler(c)

			if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedCode, rec.Code)
			}
		})
	}
}
func setupTestServer(method, uri string, body *bytes.Buffer) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(method, uri, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return rec, c
}
func TestNewApplicationInit(t *testing.T) {
	// Arrange
	db, _, _ := sqlmock.New()
	expected := "*users.Handler"

	// Act
	n := NewApplication(db)
	actual := fmt.Sprintf("%T", n)

	// Assert
	assert.Equal(t, expected, actual)
}
