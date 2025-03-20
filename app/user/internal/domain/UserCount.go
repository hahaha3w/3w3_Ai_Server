package domain

import "time"

type UserCount struct {
	ChatCount   int64     `json:"chat_count"`
	MemoirCount int64     `json:"memoir_count"`
	UseDays     int64     `json:"use_days"`
	LastUpdated time.Time `json:"last_updated"`
}

type UpdateUserCount struct {
	UserId          int64     `json:"user_id"`
	ChatCountIncr   int64     `json:"chat_count_incr"`
	MemoirCountIncr int64     `json:"memoir_count_incr"`
	UseDaysIncr     int64     `json:"use_days_incr"`
	LastUpdated     time.Time `json:"last_updated"`
}
