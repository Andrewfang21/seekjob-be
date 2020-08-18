package requests

// JobRequest model
type JobRequest struct {
	Query    string `form:"q"`
	Category string `form:"category"`
	Country  string `form:"country"`
	Source   string `form:"source"`
	PageNo   int    `form:"page_no" binding:"required"`
	PerPage  int    `form:"per_page" binding:"required"`
}
