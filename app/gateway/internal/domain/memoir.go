package domain

type GenerateMemoirResp struct {
	Success bool   `json:"success"`
	Memoir  Memoir `json:"memoir"`
}
type Memoir struct {
	Id        int32  `json:"id"`
	UserId    int32  `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	Style     string `json:"style"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	CreatedAt string `json:"created_at"`
}
type ListMemoirResp struct {
	Memoirs []Memoir `json:"memoirs"`
	Total   int32    `json:"total"`
}