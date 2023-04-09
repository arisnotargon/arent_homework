package response

import "github.com/arisnotargon/arent_homework/model/db"

type BannerInfo struct {
	NowDate      string        `json:"now_date"`
	CompleteRate string        `json:"complete_rate"`
	BodyHistory  []db.BodyInfo `json:"body_history"`
}
