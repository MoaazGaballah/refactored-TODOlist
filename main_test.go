package main_test

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	_ "encoding/json"
	"github.com/MoaazGaballah/refactored-TODOlist"
	"log"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"os"
	"strconv"
	_ "strconv"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize("todouser", "1", "todos") // database

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM todos")
	a.DB.Exec("ALTER SEQUENCE todos_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS todos
(
    id SERIAL,
    Name TEXT NOT NULL,
    CONSTRAINT todos_pkey PRIMARY KEY (id)
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/todo", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateTodo(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"name":"test todo"}`)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test todo" {
		t.Errorf("Expected todo to be 'test todo'. Got '%v'", m["name"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected todo ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetTodos(t *testing.T) {
	clearTable()
	addTodo(1)

	req, _ := http.NewRequest("GET", "/todo", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addTodo(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO todos(name) VALUES($1)", "Todo "+strconv.Itoa(i), (i+1.0)*10)
	}
}
