package tokens

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

func TestNewToken(t *testing.T) {
	tests := []struct {
		name         string
		body         *bytes.Buffer
		expectedCode int
	}{
		{
			name: "testSucceed",
			body: bytes.NewBufferString(`{
				"refreshToken": "6f388eca-e72e-4461-a478-172f50a7cbaf"
			}`),
			expectedCode: http.StatusOK,
		},
		{
			name: "testErrorScan",
			body: bytes.NewBufferString(`{
				"refreshToken": "6f388eca-e72e-4461-a478-172f50a7cbaf"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testPrepareError",
			body: bytes.NewBufferString(`{
				"refreshToken": "6f388eca-e72e-4461-a478-172f50a7cbaf"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testCreateRowError",
			body: bytes.NewBufferString(`{
				"refreshToken": "6f388eca-e72e-4461-a478-172f50a7cbaf"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testBadRequest",
			body: bytes.NewBufferString(`{
				"refreshToken:
			}`),
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "testInternalServerError",
			body: bytes.NewBufferString(`{
				"refreshToken": "6f388eca-e72e-4461-a478-172f50a7cbaf"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "testExecError",
			body: bytes.NewBufferString(`{
				"refreshToken": "6f388eca-e72e-4461-a478-172f50a7cbaf"
			}`),
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec, c := setupTestServer(http.MethodPost, "/newtoken", tt.body)
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			refreshToken := "6f388eca-e22e-4461-a478-172f50a7cbaf"
			tokenRow := sqlmock.NewRows([]string{"refreshToken", "account_id", "exp", "isValid"}).
				AddRow("6f388eca-e72e-4461-a478-172f50a7cbaf", 1, time.Now(), true)
			newReToen := sqlmock.NewRows([]string{"refreshToken"}).AddRow(refreshToken)
			var accountRow *sqlmock.Rows
			if tt.name != "testErrorScan" {
				accountRow = sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
					"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id", "company_id"}).
					AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
						"Patchara", "Kleebbua", "", "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1, 1)
			} else {
				accountRow = sqlmock.NewRows([]string{"account_id", "email", "password_hash", "name",
					"surname", "avatar_url", "apply_date", "is_active", "hp", "level", "created", "department_id", "posirion_id"}).
					AddRow(1, "user1", "$2a$11$nn97NWhYg8ALx5FclSJlVOSIhrNL4eR3hyX7exa/ZRVmsEnPDIkkm",
						"Patchara", "Kleebbua", "", "2019-01-01 00:00:00", true, 100, 1, "2019-01-01 00:00:00", 1, 1)
			}

			if tt.name != "testInternalServerError" {
				mock.ExpectQuery("SELECT * FROM refresh_token WHERE refreshToken = ? && isValid = ? && exp >= ?").WithArgs("6f388eca-e72e-4461-a478-172f50a7cbaf", true, AnyTime{}).WillReturnRows(tokenRow)
				mock.ExpectQuery("SELECT * FROM accounts WHERE account_id=?").WithArgs(1).WillReturnRows(accountRow)
				if tt.name != "testPrepareError" {
					expectPrepare := mock.ExpectPrepare("UPDATE refresh_token SET isValid = ? WHERE refreshToken = ?")
					if tt.name != "testExecError" {
						expectPrepare.ExpectExec().WithArgs(false, "6f388eca-e72e-4461-a478-172f50a7cbaf").WillReturnResult(sqlmock.NewResult(0, 0))
						if tt.name != "testCreateRowError" {
							mock.ExpectQuery("INSERT INTO refresh_token (refreshToken, account_id, exp, isValid) VALUES (?, ?, ?, ?) RETURNING refreshToken;").WithArgs(AnyToken{}, 1, AnyTime{}, true).WillReturnRows(newReToen)
						}
					}
				}

			}
			h := Handler{db}
			err = h.NewTokenHandler(c)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedCode, rec.Code)
				t.Log(rec.Body)
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
	expected := "*tokens.Handler"

	// Act
	n := NewApplication(db)
	actual := fmt.Sprintf("%T", n)

	// Assert
	assert.Equal(t, expected, actual)
}
