package models

import "time"

type Service struct {
	ID            int
	Name          string
	URL           string
	CheckInterval int
	Timeout       int
	CreatedAt     time.Time
}
