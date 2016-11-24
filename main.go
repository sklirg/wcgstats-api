package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Static("/client", "./public")
	router.Static("/static", "./public/static")

	v1 := router.Group("api/v1")
	team_routes := v1.Group("teams")
	{
		team_routes.GET("/", GetTeams)
		team_routes.GET("/:team_id", GetTeamDates)
	}

	stats_routes := v1.Group("stats")
	{
		//stats_routes.GET("/", nil)
		stats_routes.GET("/:team_id", GetTeamStats)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
