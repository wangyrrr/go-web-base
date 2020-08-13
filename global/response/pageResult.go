package response

type PageResult struct {
	Total   int         `json:"total"`
	Size    int         `json:"size"`
	Pages   int         `json:"pages"`
	Current int         `json:"current"`
	Records interface{} `json:"records"`
}
