package request

type Response struct {
	Message string `json:"message"`
}

type RegisterRequest struct{
	ID string `json:"user_id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Bio string `json:"bio"`
	Role string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}