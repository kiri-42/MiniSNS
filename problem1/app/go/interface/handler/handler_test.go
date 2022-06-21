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

	testResponseCode(t, rec.Code, http.StatusOK)

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

	path := "/get_user/"
	tCases := map[string]struct {
		inUserID string
		wantCode int
	}{
		"OK_user_id:1":       {"1", 200},
		"NG_user_idがアルファベット": {"a", 500},
		"NG_user_idが空":       {"", 500},
		"NG_存在しないuser_id1":   {"100", 500},
		"NG_存在しないuser_id2":   {"0", 500},
		"NG_存在しないuser_id3":   {"-1", 500},
	}

	for name, tc := range tCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			url := path + tc.inUserID
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			testResponseCode(t, rec.Code, tc.wantCode)
		})
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

	testResponseCode(t, rec.Code, http.StatusOK)
}

func TestGetUserListPaging(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	path := "/get_user_list_paging/"
	tCases := map[string]struct {
		inLimit  string
		inPage   string
		wantCode int
	}{
		"OK_limit:1_page:1":   {"1", "1", 200},
		"OK_limit:3_page:4":   {"3", "4", 200},
		"OK_limit:3_page:100": {"3", "100", 200},
		"NG_limitがアルファベット":    {"a", "1", 500},
		"NG_pageがアルファベット":     {"3", "a", 500},
		"NG_存在しないlimit1":      {"0", "4", 500},
		"NG_存在しないlimit2":      {"-1", "4", 500},
		"NG_存在しないpage1":       {"3", "0", 500},
		"NG_存在しないpage2":       {"3", "-1", 500},
		"NG_limitが空":          {"", "a", 500},
		"NG_pageが空":           {"3", "", 500},
	}

	for name, tc := range tCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			url := path + tc.inLimit + "/" + tc.inPage
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			testResponseCode(t, rec.Code, tc.wantCode)
		})
	}
}

func TestGetFriendList(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	path := "/get_friend_list/"
	tCases := map[string]struct {
		in       string
		wantCode int
	}{
		"OK_user_id:1":       {"1", 200},
		"NG_存在しないuser_id1":   {"0", 500},
		"NG_存在しないuser_id2":   {"-1", 500},
		"NG_存在しないuser_id3":   {"100", 500},
		"NG_user_idがアルファベット": {"a", 500},
		"NG_user_idが空":       {"", 500},
	}

	for name, tc := range tCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			url := path + tc.in
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			testResponseCode(t, rec.Code, tc.wantCode)
		})
	}
}

func TestGetFriendOfFriendList(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	path := "/get_friend_of_friend_list/"
	tCases := map[string]struct {
		in       string
		wantCode int
	}{
		"OK_user_id:1":       {"1", 200},
		"NG_存在しないuser_id1":   {"0", 500},
		"NG_存在しないuser_id2":   {"-1", 500},
		"NG_存在しないuser_id3":   {"100", 500},
		"NG_user_idがアルファベット": {"a", 500},
		"NG_user_idが空":       {"", 500},
	}

	for name, tc := range tCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			url := path + tc.in
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			testResponseCode(t, rec.Code, tc.wantCode)
		})
	}
}

func TestGetFriendOfFriendListPaging(t *testing.T) {
	db, err := configs.GetDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)

	path := "/get_friend_of_friend_list_paging/"
	tCases := map[string]struct {
		inUserID string
		inLimit  string
		inPage   string
		wantCode int
	}{
		"OK_user_id:1":     {"1", "1", "1", 200},
		"OK_user_id:2":     {"3", "3", "3", 200},
		"NG_存在しないuser_id1": {"0", "3", "3", 500},
		"NG_存在しないuser_id2": {"-1", "3", "3", 500},
		"NG_存在しないuser_id3": {"100", "3", "3", 500},
		"NG_存在しないlimit1":   {"3", "0", "3", 500},
		"NG_存在しないlimit2":   {"3", "-1", "3", 500},
		"NG_存在しないpage1":    {"3", "3", "0", 500},
		"NG_存在しないpage2":    {"3", "3", "-1", 500},
		"NG_user_idが空":     {"", "3", "3", 500},
		"NG_limitが空":       {"3", "", "3", 500},
		"NG_pageが空":        {"3", "3", "", 500},
	}

	for name, tc := range tCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			url := path + tc.inUserID + "/" + tc.inLimit + "/" + tc.inPage
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			testResponseCode(t, rec.Code, tc.wantCode)
		})
	}
}

func testResponseCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got: %d want: %d\n", got, want)
	}
}
