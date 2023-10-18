package adapters

type TUserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int8   `json:"age"`
	Active bool   `json:"active"`
}
