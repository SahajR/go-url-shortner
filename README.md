# Golang URL shortener
My `hello world` project for Golang.

## Getting started
Install the following dependencies:
* [Mux](https://github.com/gorilla/mux)
* [Mgo.V2](https://godoc.org/gopkg.in/mgo.v2)
* [Viper](https://github.com/spf13/viper)
* [Goblin](https://github.com/franela/goblin), for tests

This project uses MongoDB. You can get a free hosted one at [mLab](https://mlab.com/). Get the connection string and database name and put them in `example.db.json`. Rename it to `db.json`
. The default port the server runs is 3000. You can change it in the config if you wish.

Then you can start the server by running:
```
go run main.go
 
// OR
 
go build main.go
./main
```