package bo

type PageInfo struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}
