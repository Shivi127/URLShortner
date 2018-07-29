

# URL-Shortner build in GoLang

- Shorten a URL to a smallURL (10 characters long)

- Given the shortURL get the longURL

- Visitor Counting



### Prerequisites

Go should be installed on your machine. You can find the instructions on https://golang.org/dl/. 

Install mysql driver( used to query the database): run the following command

```
go get github.com/go-sql-driver/mysql
```

### Installing

To run the application first clone the application via
``` git clone https://github.com/Shivi127/URLShortner.git``` 

Then cd into the folder 
```
cd URLShortner
```

To run the go application straight from your terminal

```
go run main.go
```



## Running the tests

These are the various tests I ahve implemented to make sure that the code is running as I expect it to. 
The code to the test can be found in main_test.go

  - TestIsValidURLTrue
  
  - TestIsValidURLFalse

  - TestGetshortURL
  
  - TestLengthenURLTrue
  
  - TestLengthenURLFalse
  
  - TestAddtoDB
  
  - TestIncrementCount
  
  You can check the functionality of the code by running 
  ```
    go test
  ```
  
### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system





## Authors

* **Shivangi Singh** 


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Keep Learning
