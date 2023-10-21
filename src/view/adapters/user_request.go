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

type TUserUpdateRequest struct {
	Name     string `json:"name,omitempty" binding:"omitempty,min=2,max=100"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6,max=64,containsany=!@#$%&_-^~/?"`
	Age      int8   `json:"age,omitempty" binding:"omitempty,min=1,max=150"`
}

type TUserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&_-^~/?"`
}
