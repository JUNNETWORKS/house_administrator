package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JUNNETWORKS/house_administrator/data"
	"github.com/julienschmidt/httprouter"
)

func TestGetRooms(t *testing.T) {
	setup()

	// テスト用のデータの挿入
	room1 := data.NewRoom("TestRoom1", "Test1")
	room1.Create()
	room2 := data.NewRoom("TestRoom2", "Test2")
	room2.Create()

	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// マルチプレクサを作成して,ハンドラーを登録
	mux := httprouter.New()
	mux.GET("/rooms", GetRooms)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	mux.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 日付はテスト出来ないので部屋の名前と説明文だけ
	var rooms []data.Room
	if err := json.Unmarshal(rr.Body.Bytes(), &rooms); err != nil {
		log.Fatal(err)
	}
	// Check the response body is what we expect.
	expected := 2
	if actual := len(rooms); actual != expected {
		t.Errorf("想定されている部屋数と違います! (expected: %d, actual: %d", expected, actual)
	}
}
