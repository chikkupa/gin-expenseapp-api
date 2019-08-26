package config

var Host = "http://localhost"
var Port = "8080"
var Host_location = Host + ":" + Port + "/"

/* Database config */
var Db_name = "expenseapp"
var Db_user = "root"
var Db_password = "123"
var Mysql = "mysql"

var Dbconnection = Db_user + ":" + Db_password + "@tcp(127.0.0.1:3306)/" + Db_name
var BaseUrl = Host + ":" + Port

// Authentication Key
var Auth_secret = "expense@golang"