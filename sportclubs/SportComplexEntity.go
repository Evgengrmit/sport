package sportclubs

import "time"

type SportComplex struct {
	Title       string `json:"title"`
	ScheduledAt string `json:"scheduledAt"`
	Description string `json:"description"`
}

type ComplexJSON struct {
	Duration    time.Duration `json:"duration"`
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	ScheduledAt time.Time     `json:"scheduledAt"`
	Trainer     Trainer       `json:"trainer"`
}
type Trainer struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}
