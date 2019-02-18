package request

type Response struct {
	Message string `json:"message"`
}

type RegisterRequest struct{
	ID 			string `json:"user_id"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

type GenreResponse struct {
	ID 		string `json:"genre_id"`
	Genre 	string `json:"genre"`
}

type InstrumentResponse struct {
	ID			string `json:"instrument_id"`
	Instrument 	string `json:"instrument"`
}