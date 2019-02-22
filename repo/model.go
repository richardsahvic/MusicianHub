package repo

import (
	"time"
)

type UserDetail struct{
	ID			string		`json:"id" db:"id"`
	Email		string		`json:"email" db:"email"`
	Password	string		`json:"password" db:"password"`
	Name		string		`json:"name" db:"name"`
	Gender		string		`json:"gender" db:"gender"`
	Birthdate	string		`json:"birthdate" db:"birthdate"`
	Bio			string		`json:"bio" db:"bio"`
	CreatedAt 	time.Time 	`json:"created_at" db:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at" db:"updated_at"`
	AvatarUrl	string		`json:"avatar_url" db:"avatar_url"`
}

type GenreList struct{
	ID		string	`json:"id" db:"id"`
	Genre	string	`json:"genre" db"genre"`
}

type InstrumentList struct{
	ID			string	`json:"id" db:"id"`
	Instrument	string	`json:"instrument" db:"instrument"`
}

type UserInstrument struct{
	IDUserInstrument	string	`json:"uinstrument_id" db:"id"`
	UserId				string	`json:"user_id" db:"user_id"`
	InstrumentId		string	`json:"instrument_id" db:"instrument_id"`
}

type UserGenre struct{
	IDUserGenre	string	`json:"ugenre_id" db:"id"`
	UserId		string	`json:"user_id" db:"user_id"`
	GenreId		string	`json:"genre_id" db:"genre_id"`
}

type UserPost struct{
	PostId	string `json:"post_id" db:"post_id"`
	UserId string `json:"user_id" db:"user_id"`
	PostType string `json:"post_type" db:"post_type"`
	FileUrl string `json:"file_url" db:"file_url"`
	Caption string `json:"caption" db:"caption"`
	CreatedAt string `json:"created_at" db:"created_at"`
}