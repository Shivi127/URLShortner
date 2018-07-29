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
		// Means that the url doesnt exist
		// log.Fatal(err, " Error from here")
		addURL(surl, lurl)
		err = db.QueryRow("SELECT lurl FROM new_table WHERE surl =?", surl).Scan(&lurldb)
		if err != nil {
			// Means that the url doesnt exist
			log.Fatal(err)
		}
	}

	// Means that the URL is already there so we need to match the incoming url
	if compareURL(lurl, lurldb) == true {
		fmt.Println("No collision so just return from here")
	} else {
		fmt.Println("Need to rehash- very small probability though")
	}
	fmt.Println("Am I connected Yet", lurldb)
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

func returnLongURL(surl string) string {
	fmt.Println("I am open to retreive ")
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
	fmt.Println("Yo this is the short URl", surl, " Long URL", lurldb)
	return (lurldb)

}

// Function to validate that the user is entering a valid URL
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}

// Function that produces a hash using SHA-2
func getSHA3Hash(text string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(text))
	x := sha_512.Sum(nil)[0:9]
	fmt.Println("x[0] before ", x[0])
	x[0] = x[0] % 16
	fmt.Println("Bytes Array", x)

	return base64.URLEncoding.EncodeToString(x)

}

func main() {

	fmt.Println("Hello, How can I help you today?")
	fmt.Println("Enter 1: To Shorten URL")
	fmt.Println("Enter 2: To Lengthen the shortURL")
	fmt.Println("Enter 3 to EXIT")

	var text string
	fmt.Scanln(&text)
	text = strings.TrimSpace(text)
	for text != "3" {
		if text == "1" {
			// Shorten
			fmt.Println("Enter your URL")
			var url string
			fmt.Scanln(&url)
			url = strings.TrimSpace(url)
			valid := isValidURL(url)
			if valid == true {
				sURL := getSHA3Hash(url)
				addURLtoDB(sURL, url)
				fmt.Println("Your new ShortURL is: ", sURL)
			} else {
				// Give an error message
				fmt.Println("Please enter a valid URL")
			}
		}
		if text == "2" {
			// lengthen
			var lurl string
			fmt.Println("Enter your Shortened URL")
			fmt.Scanln(&lurl)
			lurl = strings.TrimSpace(lurl)
			fmt.Println(returnLongURL(lurl))
		}
		fmt.Println("Try Again? Enter your choice")
		fmt.Scanln(&text)

	}
}
