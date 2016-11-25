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
		log.WithField("err", err).Error("Could not connect.")
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
		log.WithField("err", err).Error("Could not connect.")
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
	start_date := c.DefaultQuery("start_date", "2016-11-24")
	end_date := c.DefaultQuery("end_date", "2016-12-24")

	vals, err := res.Result()

	if err != nil {
		log.WithField("err", err).Error("Error getting result")
		c.JSON(500, gin.H{"message": "Something went wrong."})
		return
	}

	stats := make([]DailyStatisticsTotals, 0)

	for i := range vals {
		dailyStat := DailyStatisticsTotals{}
		err = json.Unmarshal([]byte(vals[i]), &dailyStat)
		if err != nil {
			log.WithField("err", err).Error("Error getting result")
			c.JSON(500, gin.H{"message": "Something went wrong."})
			return
		}
		if ShouldAddDayStats(dailyStat, start_date, end_date) {
			stats = append(stats, dailyStat)
		}
	}

	// Add last updated
	// Add last checked
	c.JSON(200, gin.H{
		"count":   len(stats),
		"results": stats,
	})
}
