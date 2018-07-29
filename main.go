package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbconnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "shivi:papa&DAD2016@tcp(mydb.cmxpivutp5cx.us-east-2.rds.amazonaws.com:3306)/urlify")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func updateCount(count int) {
	db := dbconnect()
	time := time.Now()
	stmtIns, err := db.Query("UPDATE new_table SET visit_count = ? , LastVisit =?", count+1, time)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
}

func addURLtoDB(surl string, lurl string) {

	db := dbconnect()
	var lurldb string
	err := db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
	if err != nil {
		addURL(surl, lurl)
		err = db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
		if err != nil {
			log.Fatal(err)
		}
	}

	var visitCount int
	err = db.QueryRow("SELECT lurl, visit_count FROM new_table WHERE surl =?", surl).Scan(&lurldb, &visitCount)
	if err != nil {
		log.Fatal(err)
	}

	if compareURL(lurl, lurldb) == true {
		updateCount(visitCount)
		fmt.Println("ShortURL was Successfully created/updated")
	} else {
		fmt.Println("Oops there is a collision try refreshing the db")
	}
}

func addURL(surl string, lurl string) {

	db := dbconnect()
	time := time.Now()
	stmtIns, err := db.Query("INSERT INTO new_table VALUES( ?, ?, ?,?)", surl, lurl, 0, time)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
}

func compareURL(lurl1 string, lurl2 string) bool {
	return (lurl1 == lurl2)
}

func getLongURL(surl string) string {

	db := dbconnect()
	var lurldb string
	var visitCount int

	err := db.QueryRow("SELECT lurl, visit_count FROM new_table WHERE surl =?", surl).Scan(&lurldb, &visitCount)
	if err != nil {
		return ("Invalid Shortened URL")
	}
	updateCount(visitCount)
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
	fmt.Println("Enter c: to continue")
	fmt.Println("Enter e: To EXIT")
	fmt.Println("NOTE: The choices are case sensitive")
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
func input() {
	fmt.Println("Would you like to continue ? ")
	var text string
	fmt.Scanln(&text)
	text = strings.TrimSpace(text)
	choices(text)
}

func choices(choice string) {
	switch choice {
	case "s":
		shortenURL()
		input()

	case "l":
		lengthenURL()
		input()

	case "c":
		input()
	default:
		break
	}
}

func main() {
	EntryText()
	var text string
	fmt.Scanln(&text)
	text = strings.TrimSpace(text)
	choices(text)
}
