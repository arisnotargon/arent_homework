package request

type BodyInfoStore struct {
	Weight  string `binding:"required,numeric"`
	FatRate string `json:"fat_rate" binding:"required,numeric"`
}
