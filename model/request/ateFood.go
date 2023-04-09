package request

type AteFoodStore struct {
	Type  uint8  `binding:"required,min=1,max=4"`
	Name  string `binding:"required"`
	Pic   string
	AteAt string `json:"ate_at" binding:"required,datetime=2006-01-02"`
}
