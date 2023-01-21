package users

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

//	claims := &jwtCustomClaims{
//		"user1", "Patchara", "Kleebbua", 1, 1, 1, 1, jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 120)),
//		},
//	}
func TestUser(t *testing.T) {
	Unauthorized(t)
	InternalServerError(t)
	Succeed(t)
	BadRequest(t)

}

//	func initTestDatabase() *sql.DB {
//		db, err := sql.Open("mysql", "feelme_admin:zj439$Z2p@tcp(119.59.96.90)/feelme_db")
//		if err != nil {
//			log.Fatal("Connect to database error", err)
//		}
//		return db
//	}
func Succeed(t *testing.T) {
	body := bytes.NewBufferString(`{
		"email": "user1",
		"password": "123"
	  }`)
	rec, c := setupTestServer(http.MethodPost, "/userlogin", body)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
		"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id", "company_id"}).
		AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
			"Patchara", "Kleebbua", nil, "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1, 1)

	mock.ExpectQuery("SELECT \\* FROM accounts WHERE email=\\?").WithArgs("user1").WillReturnRows(accountRow).WillReturnError(err)

	expectedRow := sqlmock.NewRows([]string{"refreshToken", "account_id", "exp", "isValid"}).
		AddRow("d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9", 1, "2019-01-01 00:00:00", true).
		AddRow("e840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9", 1, "2019-01-01 00:00:00", true)

	mock.ExpectQuery("SELECT \\refreshToken FROM refresh_token WHERE account_id=\\? && isValid=\\?").WillReturnRows(expectedRow)
	if err != nil {
		return
	}
	h := Handler{db}
	err = h.UserLoginHandler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
func Unauthorized(t *testing.T) {
	body := bytes.NewBufferString(`{
		"email": "user1",
		"password": "1234"
	  }`)
	rec, c := setupTestServer(http.MethodPost, "/userlogin", body)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
		"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id", "company_id"}).
		AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
			"Patchara", "Kleebbua", nil, "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1, 1)
	mock.ExpectQuery("SELECT \\* FROM accounts WHERE email=\\?").WithArgs("user1").WillReturnRows(accountRow).WillReturnError(err)
	if err != nil {
		return
	}
	h := Handler{db}
	err = h.UserLoginHandler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
func InternalServerError(t *testing.T) {
	body := bytes.NewBufferString(`{
		"emails": "user1",
		"password": "123"
	  }`)
	rec, c := setupTestServer(http.MethodPost, "/userlogin", body)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
		"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id", "company_id"}).
		AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
			"Patchara", "Kleebbua", "", "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1, 1)

	mock.ExpectQuery("SELECT \\* FROM accounts WHERE email=\\?").WithArgs("user1").WillReturnRows(accountRow).WillReturnError(err)
	refreshTokemRows := sqlmock.NewRows([]string{"refreshToken"}).
		AddRow("d840e7da-bd5b-490a-8f2a-83648c114d6c-f0a82297-0ec7-4d74-baf2-1db6b4a03c14-43a2d592-09e3-4520-abf3-531ee2b64ea0-7dde1858-e8e5-4dbb-9c28-963117f4dec9")
	mock.ExpectQuery("SELECT refreshToken FROM refresh_token WHERE account_id=?&&isValid=?").WithArgs(1, true).WillReturnRows(refreshTokemRows).WillReturnError(err)
	h := Handler{db}
	err = h.UserLoginHandler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}
func BadRequest(t *testing.T) {

	body := bytes.NewBufferString(`{
		"email": "user1",
		"password": 
	  }`)
	rec, c := setupTestServer(http.MethodPost, "/userlogin", body)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
		"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id", "company_id"}).
		AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
			"Patchara", "Kleebbua", "", "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1, 1)
	mock.ExpectQuery("SELECT \\* FROM accounts WHERE email=\\?").WithArgs("user1").WillReturnRows(accountRow)

	if err != nil {
		return
	}
	h := Handler{db}
	err = h.UserLoginHandler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
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
