package request

type TargetWeightStore struct {
	TargetWeight string `json:"target_weight" binding:"required,numeric"`
	NowWeight    string `json:"now_weight" binding:"required,numeric"`
}
