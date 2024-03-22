package main

import (
	"finalassignment/database"
	_ "finalassignment/docs"
	"finalassignment/route"
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
)

// var PORT = os.Getenv("PGPORT")
var PORT = "8080"

// @title Final Assignment - M.Irvan Muhandis
// @version 1.0
// @description Final Assignment Class 005
// @license.name Golang Hacktiv8
// @host localhost:8080
// @BasePath /
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				To access the app
func main() {
	// Add your own struct validation tags
	govalidator.TagMap["age8"] = govalidator.Validator(func(val string) bool {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num > 8
	})
	fmt.Println("Swagger running on : http://localhost:8080/swagger/index.html")

	database.StartDB()
	fmt.Println("Server starting at :", PORT)

	route.StartApp().Run(":" + PORT)
}