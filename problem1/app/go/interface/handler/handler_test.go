package handler_test

import (
	_ "database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"

	"problem1/configs"
	"problem1/interface/handler"
)

func TestRoot(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got: %d want: %d\n", rec.Code, http.StatusOK)
	}

	if rec.Body.String() != "mini sns" {
		t.Errorf("got: %s want: %sn", rec.Body.String(), "mini sns")
	}
}

func TestGetUser(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	req := httptest.NewRequest(http.MethodGet, "/get_user/1", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got: %d want: %d\n", rec.Code, http.StatusOK)
	}
}

func TestGetUserList(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	req := httptest.NewRequest(http.MethodGet, "/get_user_list", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got: %d want: %d\n", rec.Code, http.StatusOK)
	}
}
