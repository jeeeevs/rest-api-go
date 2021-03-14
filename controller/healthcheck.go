package controller

import (
	"encoding/json"

	"github.com/jeeeevs/timeutil"
)

var upTimeNow = timeutil.UpTime()

type Health struct {
	status string
	upTime int64
}

func Healthcheck() []byte {
	upTime := upTimeNow()
	m := Health{"available", upTime}
	b, _ := json.Marshal(m)
	return b
}
