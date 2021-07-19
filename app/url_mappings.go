package app

import (
	"github.com/alfonsovgs/bookstore_users-api/controllers/ping"
	"github.com/alfonsovgs/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// Mapping users controller
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
