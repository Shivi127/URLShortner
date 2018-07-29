package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func addURLtoDB(surl string, lurl string) {

	db, err := sql.Open("mysql", "shivi:papa&DAD2016@tcp(mydb.cmxpivutp5cx.us-east-2.rds.amazonaws.com:3306)/urlify")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var lurldb string

	err = db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
	if err != nil {
		addURL(surl, lurl)
		err = db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
	if err != nil {
		log.Fatal(err)
	}

	if compareURL(lurl, lurldb) == true {
		fmt.Println("No collision so just return from here")
	} else {
		fmt.Println("Need to rehash- very small probability though")
	}
}

func compareURL(lurl1 string, lurl2 string) bool {
	return (lurl1 == lurl2)
}

func addURL(surl string, lurl string) {
	db, err := sql.Open("mysql", "shivi:papa&DAD2016@tcp(mydb.cmxpivutp5cx.us-east-2.rds.amazonaws.com:3306)/urlify")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtIns, err := db.Query("INSERT INTO new_table VALUES( ?, ? )", surl, lurl)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
}

func getLongURL(surl string) string {
	db, err := sql.Open("mysql", "shivi:papa&DAD2016@tcp(mydb.cmxpivutp5cx.us-east-2.rds.amazonaws.com:3306)/urlify")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var lurldb string

	err = db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
	if err != nil {
		return ("Invalid Shortened URL")
	}
	fmt.Println("This is the short URl", surl, " for ", lurldb)
	return (lurldb)

}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}

func getshortURL(text string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(text))
	x := sha_512.Sum(nil)[0:9]
	return base64.URLEncoding.EncodeToString(x)

}

func EntryText() {
	fmt.Println("Hello, How can I help you today?")
	fmt.Println("Enter s: To Shorten URL")
	fmt.Println("Enter l: To Lengthen the shortURL")
	fmt.Println("Enter q to EXIT")
}

func shortenURL() {
	fmt.Println("Enter your URL")
	var url string
	fmt.Scanln(&url)
	url = strings.TrimSpace(url)
	valid := isValidURL(url)
	if valid == true {
		sURL := getshortURL(url)
		addURLtoDB(sURL, url)
		fmt.Println("Your new ShortURL is: ", sURL)
	} else {
		// Give an error message
		fmt.Println("Please enter a valid URL")
	}
}

func lengthenURL() {
	var surl string
	fmt.Println("Enter your Shortened URL")
	fmt.Scanln(&surl)
	surl = strings.TrimSpace(surl)
	fmt.Println(getLongURL(surl))
}

func main() {
	EntryText()
	var text string
	fmt.Scanln(&text)
	text = strings.TrimSpace(text)
	for text != "e" {
		if text == "s" {
			shortenURL()

		}
		if text == "l" {
			lengthenURL()
		}
		fmt.Println("")
		fmt.Println("Try Again? Enter your choice")
		EntryText()
		fmt.Scanln(&text)

	}
}
