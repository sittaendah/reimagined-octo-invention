package organization

type Organization struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}
