

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
  
  
  You can check the functionality of the code by running 
  ```
    go test
  ```
 

## Deployment Using Docker

To run this on Docker you should already have a Docker installed. 

### Installation

The installation instructions:

Mac: can be found at https://docs.docker.com/docker-for-mac/

Windows: https://docs.docker.com/docker-for-windows/

### Creating a docker image

Once you have Docker installed you can create an image by running (make sure that you are in URLShortner when you run this command

```
docker -t <imageName> .
````

### Running the docker image

```
docker run -i <imageName>
```

## Future Development

- Use the time stamp in the db to flush out URL's that have not been used for sometime.

- In case the there is a collision in the Hashing Function increase the length of the short-URL

- Build a Frontend (Web Application) 

- Implement a redirect

## Authors

* **Shivangi Singh** 


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Keep Learning
