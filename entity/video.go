package entity

type Person struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=18,lte=100"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=100"`
	Description string `json:"description" binding:"max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

type ResponseVideo struct {
	Data    Video  `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
