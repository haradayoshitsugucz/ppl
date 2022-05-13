package output

type pager struct {
	Total  int64       `json:"total"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Result interface{} `json:"result"`
}

func ToPager(total int64, offset int, limit int, result interface{}) *pager {
	return &pager{Total: total, Offset: offset, Limit: limit, Result: result}
}
