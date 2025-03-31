package domain

type ActivityResp struct {
	ActivityID  int64  `json:"activity_id"`
	UserID      int64  `json:"user_id"`
	RelationID  int64  `json:"relation_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
type ActivityListResp struct {
	Total       int32           `json:"total"`
	HasMore     bool            `json:"has_more"`
	Activities  []*ActivityResp `json:"activities"`
	ChatCount   int32           `json:"chat_count"`
	MemoirCount int32           `json:"memoir_count"`
	UseDay      int32           `json:"use_day"`
}
