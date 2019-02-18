package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn                      	*sqlx.DB
	findIDStmt                	*sqlx.Stmt
	findEmailStmt             	*sqlx.Stmt
	insertNewUserStmt         	*sqlx.NamedStmt
	updatePasswordStmt        	*sqlx.Stmt
	getGenre					*sqlx.Stmt
	getInstrument				*sqlx.Stmt
}

func (db *userRepository) MustPrepareStmt(query string) *sqlx.Stmt {
	stmt, err := db.conn.Preparex(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %s\n", err)
		os.Exit(1)
	}
	return stmt
}

func (db *userRepository) MustPrepareNamedStmt(query string) *sqlx.NamedStmt {
	stmt, err := db.conn.PrepareNamed(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %s\n", err)
		os.Exit(1)
	}
	return stmt
}

func NewRepository(db *sqlx.DB) AppRepository {
	r := userRepository{conn: db}
	r.findIDStmt = r.MustPrepareStmt("SELECT email, password FROM musiciandb.user_detail WHERE id=?")
	r.findEmailStmt = r.MustPrepareStmt("SELECT id, email, password FROM musiciandb.user_detail WHERE email=?")
	r.updatePasswordStmt = r.MustPrepareStmt("UPDATE musiciandb.user_detail SET password=? WHERE id=?")
	r.insertNewUserStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_detail (id, email, password) VALUES (:id, :email, :password)")
	r.getGenre = r.MustPrepareStmt("SELECT * FROM musiciandb.genre_list")
	r.getInstrument = r.MustPrepareStmt("SELECT * FROM musiciandb.instrument_list")
	return &r
}

func (db *userRepository) FindByID(id string) (usr UserDetail, err error) {
	err = db.findIDStmt.Get(&usr, id)
	if err != nil {
		log.Println("Error at finding id:  ", err)
	}
	return
}

func (db *userRepository) FindByEmail(email string) (usr UserDetail, err error) {
	err = db.findEmailStmt.Get(&usr, email)
	if err != nil {
		log.Println("Error at finding email:  ", err)
	}
	return
}

func (db *userRepository) InsertNewUser(newUser UserDetail) (success bool, err error) {
	_, err = db.insertNewUserStmt.Exec(newUser)
	if err != nil {
		log.Println("Error inserting new user:  ", err)
		success = false
		return
	}
	success = true
	return
}

func (db *userRepository) UpdatePassword(id string, newPassword string) (success bool, err error) {
	_, err = db.updatePasswordStmt.Exec(newPassword, id)
	if err != nil {
		log.Println("Failed to update password: ", err)
		success = false
	}
	success = true
	return
}

func (db *userRepository) GetGenres() (genres []GenreList, err error) {
	err = db.getGenre.Select(&genres)
	if err != nil{
		log.Println("Failed to get genres: ", err)
	}
	return
}

func (db *userRepository) GetInstruments() (instruments []InstrumentList, err error) {
	err = db.getInstrument.Select(&instruments)
	if err != nil{
		log.Println("Failed to get instruments: ", err)
	}
	return
}