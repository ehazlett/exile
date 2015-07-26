package api

import (
	log "github.com/Sirupsen/logrus"
)

func getTestAPI() (*API, error) {
	log.SetLevel(log.ErrorLevel)
	return NewAPI("", nil)
}
