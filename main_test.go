package main

import (
	"database/sql"
	"testing"
)

func TestIsValidURLTrue(t *testing.T) {
	if isValidURL("https://tutorialedge.net/golang/intro-testing-in-go/") == false {
		t.Error("isValidURL was supposed to return true but returned false")
	}
}

func TestIsValidURLFalse(t *testing.T) {
	if isValidURL("http://hjsb") == true {
		t.Error("isValidURL was supposed to return false but returned true")
	}
}

func TestGetshortURL(t *testing.T) {
	// delete from table new_table where surl= "2JNLKjUjlpZG"
	url := "https://golang.org/pkg/database/sql/#DB.Prepare "
	if getshortURL(url) != "AwuX3dqBAQ==" {
		t.Error("Your shortening Service is not working")
	}
}

func TestLengthenURLTrue(t *testing.T) {
	surl := "AwuX3dqBAQ=="
	url := "https://golang.org/pkg/database/sql/#DB.Prepare"
	if getLongURL(surl) == url {
		t.Error("your short-URL doesn't return the right long-URL")
	}
}

func TestLengthenURLFalse(t *testing.T) {
	surl := "AwuX3dqBAQ==rsgstgrgdtgdr"
	url := "Invalid Shortened URL"
	if getLongURL(surl) != url {
		t.Error("your short-URL doesn't return the right long-URL")
	}
}

func TestAddtoDB(t *testing.T) {

	surl := "AwuX3dqBAQ=="
	url := "https://golang.org/pkg/database/sql/#DB.Prepare "

	db, err := sql.Open("mysql", "shivi:papa&DAD2016@tcp(mydb.cmxpivutp5cx.us-east-2.rds.amazonaws.com:3306)/urlify")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Query("Delete from new_table where surl=?", surl)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	var lurldb string

	err = db.QueryRow("Select lurl from new_table where surl= ?", surl).Scan(&lurldb)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	if compareURL(lurldb, url) {
		t.Error("Add to DB is not adding to the database")
	}
}
