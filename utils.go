package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

func ShouldAddDayStats(dailyStat DailyStatisticsTotals, start_date, end_date string) bool {
	competitionStart, _ := time.Parse("2006-01-02", start_date)
	competitionEnd, _ := time.Parse("2006-01-02", end_date)
	date, err := time.Parse("2006-01-02", dailyStat.Date)
	if err != nil {
		log.WithFields(log.Fields{
			"err":     err,
			"date":    dailyStat.Date,
			"dateObj": date,
		}).Error("Could not parse time string.")
	}
	if date.Unix() < competitionStart.Unix() || date.Unix() > competitionEnd.Unix() {
		log.WithFields(log.Fields{
			"date": date,
		}).Debug("Skipping date outside range...")
		return false
	}
	return true
}
