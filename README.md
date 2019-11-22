# Squeat Cassandra REST API

- We are using a Revel server to handle our HTTP request
- We are commuicating with our Cassandra database using GoCQL package

## Requirements 

#### In order to implement this API (if you have to)

- You need to go get Revel and GoCQL.
   - go get github.com/revel/revel
   - go get github.com/gocql/gocql
   
- You need copy this repository in your Go workspace at
   ->Go/src
 
 - Then if all the packages are installed and this repository in the right place you can move on to the Revel part to start the API

# Revel

### Start the API:

   revel run -a cassandraRest

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here


## Help with Revel

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).
