package main

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

func GetTeams(c *gin.Context) {
	client := GetClient()
	defer client.Close()

	res := client.HKeys("teams")
	keys, err := res.Result()

	if err != nil {
		c.JSON(500, gin.H{"message": "Something went wrong."})
		return
	}
	teams := make([]string, len(keys))

	for i := range keys {
		teams[i] = keys[i]
	}

	if err != nil {
		log.WithField("error", err).Error("Marshaling JSON failed")
		c.JSON(500, gin.H{"message": "Something went wrong."})
	} else {
		c.JSON(200, gin.H{"teams": teams})
	}
}

func GetTeamDates(c *gin.Context) {
	client := GetClient()
	defer client.Close()

	res := client.HKeys(c.Param("team_id"))
	dates, err := res.Result()
	if err != nil {
		c.JSON(500, gin.H{"message": "Something went wrong."})
		return
	}

	log.Infof("dates %+v", dates)
	c.JSON(200, gin.H{"dates": dates})
}

func GetTeamStats(c *gin.Context) {
	client := GetClient()
	defer client.Close()

	res := client.HVals(c.Param("team_id"))

	vals, err := res.Result()

	if err != nil {
		log.Error("Error getting result")
		return
	}

	stats := make([]DailyStatisticsTotals, len(vals))

	for i := range vals {
		err = json.Unmarshal([]byte(vals[i]), &stats[i])
	}

	// log.Infof("stats %+v", stats)

	// Add last updated
	// Add last checked
	c.JSON(200, gin.H{
		"count":   len(vals),
		"results": stats,
	})
}
