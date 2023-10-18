package adapters

type TUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64,containsany=!@#$%&_-^~/?"`
	Age      int8   `json:"age" binding:"required,min=1,max=150"`
}

type TUserEmailRequest struct {
	Email string `json:"email"`
}
