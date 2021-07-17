package entity

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type ResponseVideo struct {
	Data    Video  `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
