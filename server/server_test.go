package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/henrikkorsgaard/crud-perf/database"
	"github.com/matryer/is"
)

var testdb = "studiepraktik_test.db"

func TestGetUser(t *testing.T) {
	defer cleanup()
	is := is.New(t)

	db := database.New(testdb)
	u := createTestUser("test@example.com", "Test User")

	_, err := db.CreateUser(u)
	is.NoErr(err)

	ts := httptest.NewServer(addRoutes(db))
	defer ts.Close()

	client := ts.Client()
	r, err := client.Get(fmt.Sprintf("%v/users/%s", ts.URL, u.Email))
	is.NoErr(err)
	is.Equal(r.StatusCode, http.StatusOK)

	var user database.User
	json.NewDecoder(r.Body).Decode(&user)
	is.Equal(u.Name, user.Name)
}

func TestPostUser(t *testing.T) {
	defer cleanup()
	is := is.New(t)

	db := database.New(testdb)
	ts := httptest.NewServer(addRoutes(db))
	defer ts.Close()

	client := ts.Client()

	u := createTestUser("test@example.com", "Test User")
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(u)
	is.NoErr(err)
	r, err := client.Post(fmt.Sprintf("%v/users", ts.URL), "application/json", b)
	is.NoErr(err)
	is.Equal(r.StatusCode, http.StatusCreated)
	user, err := db.GetUserByEmail(u.Email)
	is.NoErr(err)
	is.Equal(user.Name, u.Name)
}

func TestGetUsers(t *testing.T) {
	defer cleanup()
	is := is.New(t)

	db := database.New(testdb)

	u1 := createTestUser("test_1@example.com", "Test User 1")
	u2 := createTestUser("test_2@example.com", "Test User 2")

	_, err := db.CreateUser(u1)
	is.NoErr(err)
	_, err = db.CreateUser(u2)
	is.NoErr(err)

	ts := httptest.NewServer(addRoutes(db))
	defer ts.Close()

	client := ts.Client()
	r, err := client.Get(fmt.Sprintf("%v/users", ts.URL))
	is.NoErr(err)
	is.Equal(r.StatusCode, http.StatusOK)

	var users []database.User
	json.NewDecoder(r.Body).Decode(&users)
	is.Equal(u1.Name, users[0].Name)

}

func cleanup() {
	err := os.Remove(testdb)
	if err != nil {
		panic(err)
	}
}

func createTestUser(email, name string) database.User {
	return database.User{
		Email:                      email,
		Name:                       name,
		Password:                   "Test_User_1!",
		PhoneNumber:                "+4512345678",
		PostCode:                   8000,
		City:                       "Aarhus",
		DoesNotLiveInDenmark:       false,
		SchoolId:                   "79710b5a-6f7e-4554-ad57-9abd310fa146",
		FieldOfStudyId:             "9b272d25-f7aa-4333-80d6-9ec5eabf6b10",
		NoLongerAttendingTheCourse: false,
		TermsAccepted:              true,
	}
}
