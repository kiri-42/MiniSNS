package handler_test

import (
	_ "database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"

	"problem1/configs"
	"problem1/interface/handler"
)

func TestRoot(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got: %d want: %d\n", rec.Code, http.StatusOK)
	}

	if rec.Body.String() != "mini sns" {
		t.Errorf("got: %s want: %sn", rec.Body.String(), "mini sns")
	}
}
