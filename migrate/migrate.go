package main

import (
	"github.com/jcblastor/restapi-jwt/initializers"
)

func main() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}
