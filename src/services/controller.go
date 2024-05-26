package services

type PageInfo struct {
	CurrentPage int `json:"page"`
	NextPage    any `json:"nextPage"`
	PrevPage    any `json:"prevPage"`
	Limit       int `json:"limit"`
	TotalPage   int `json:"totalPage"`
	TotalData   int `json:"totalData"`
}

type Info struct {
	Data  interface{}
	Count int
}
type ResponseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
