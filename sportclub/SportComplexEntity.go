package sportclub

import "time"

type Complex struct {
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