package repo

import (
	"time"
)

type UserDetail struct{
	ID			string		`json:"id" db:"id"`
	Email		string		`json:"email" db:"email"`
	Username	string		`json:"username" db:"username"`
	Password	string		`json:"password" db:"password"`
	Name		string		`json:"name" db:"name"`
	Gender		string		`json:"gender" db:"gender"`
	Birthdate	string		`json:"birthdate" db:"birthdate"`
	Bio			string		`json:"bio" db:"bio"`
	Role		string		`json:"role" db:"role"`
	CreatedAt 	time.Time 	`json:"created_at" db:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at" db:"updated_at"`
}

type GenreList struct{
	ID		string	`json:"genre_id" db:"genre_id"`
	Genre	string	`json:"genre" db"genre"`
}

type InstrumentList struct{
	ID			string	`json:"instrument_id" db:"instrument_id"`
	Instrument	string	`json:"instrument"	db:"instrument"`
}

type UserInstrument struct{
	IDUInstrument	string	`json:"uinstrument_id" db:"uinstrument_id"`
	UserId			string	`json:"user_id" db:"user_id"`
	InstrumentId	string	`json:"instrument_id" db:"instrument_id"`
}

type UserGenre struct{
	IDUGenre	string	`json:"ugenre_id" db:"ugenre_id"`
	UserId		string	`json:"user_id" db:"user_id"`
	GenreId		string	`json:"genre_id" db:"genre_id"`
}