package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn                      *sqlx.DB
	findIDStmt                *sqlx.Stmt
	findEmailStmt             *sqlx.Stmt
	findUsernameStmt          *sqlx.Stmt
	insertNewUserStmt         *sqlx.NamedStmt
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
	r.findIDStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.user_detail WHERE id=?")
	r.findEmailStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.user_detail WHERE email=?")
	r.findUsernameStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.user_detail WHERE username=?")
	r.insertNewUserStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_detail (id, email, username, password, name, gender, birthdate, bio, role) VALUES (:id, :email, :username, :password, :name. :gender, :birthdate, :bio, :role)")
	return &r
}

func (db *userRepository) FindByID(id string) (usr UserDetail, err error) {
	err = db.findIDStmt.Get(&usr, id)
	if err != nil {
		log.Printf("ID: %v , doesn't exist", id)
		log.Println("Error at finding id:  ", err)
	}
	return
}

func (db *userRepository) FindByEmail(email string) (usr UserDetail, err error) {
	err = db.findEmailStmt.Get(&usr, email)
	if err != nil {
		log.Printf("Email: %v, doesn't exist", email)
		log.Println("Error at finding email:  ", err)
	}
	return
}

func (db *userRepository) FindByUsername(username string) (usr UserDetail, err error) {
	err = db.findUsernameStmt.Get(&usr, username)
	if err != nil {
		log.Printf("Username: %v doesn't exist", username)
		log.Println("Error at finding username:  ", err)
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