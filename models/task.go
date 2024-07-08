package models

type Tasks struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      User   `json:"user_id" gorm:"references:ID"`
	Priority    bool   `json:"priority" `
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
	CompletedAt string `json:"completed_at"`
	CreatedAt   string `json:"created_at"`
}
