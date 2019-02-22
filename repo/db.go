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
	updatePasswordStmt        	*sqlx.Stmt
	getGenre					*sqlx.Stmt
	getInstrument				*sqlx.Stmt
	userProfileStmt				*sqlx.Stmt
	updateUserGenreStmt			*sqlx.Stmt
	updateUserInstrumentStmt	*sqlx.Stmt
	deletePostStmt				*sqlx.Stmt
	insertNewUserStmt         	*sqlx.NamedStmt
	insertUserGenreStmt			*sqlx.NamedStmt
	insertUserInstrumentStmt	*sqlx.NamedStmt
	insertNewPostStmt			*sqlx.NamedStmt
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
	r.updateUserGenreStmt = r.MustPrepareStmt("UPDATE musiciandb.user_genre SET genre_id=? WHERE user_id=?")
	r.updateUserInstrumentStmt = r.MustPrepareStmt("UPDATE musiciandb.user_instrument SET instrument_id=? WHERE user_id=?")
	r.getGenre = r.MustPrepareStmt("SELECT * FROM musiciandb.genre_list")
	r.getInstrument = r.MustPrepareStmt("SELECT * FROM musiciandb.instrument_list")
	r.deletePostStmt = r.MustPrepareStmt("DELETE FROM musiciandb.user_post WHERE post_id=?")
	r.userProfileStmt = r.MustPrepareStmt("UPDATE musiciandb.user_detail SET name=?, gender=?, birthdate=?, bio=?, avatar_url=? WHERE id=?")
	r.insertNewUserStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_detail (id, email, password) VALUES (:id, :email, :password)")
	r.insertUserGenreStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_genre (user_id, genre_id) VALUES (:user_id, :genre_id)")
	r.insertUserInstrumentStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_instrument (user_id, instrument_id) VALUES (:user_id, :instrument_id)")
	r.insertNewPostStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_post (post_id, user_id, post_type, file_url, caption) VALUES (:post_id, :user_id, :post_type, :file_url, :caption)")
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
	success = false

	_, err = db.insertNewUserStmt.Exec(newUser)
	if err != nil {
		log.Println("Error inserting new user:  ", err)
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

func (db *userRepository) InsertProfile(name string, gender string, birthdate string, bio string, avatarurl string, id string, genre UserGenre, instrument UserInstrument) (success bool, err error){
	success = false

	_, err = db.userProfileStmt.Exec(name, gender, birthdate, bio, avatarurl, id)
	if err != nil {
		log.Println("Failed making profile: ", err)
		return
	}

	_, err = db.insertUserGenreStmt.Exec(genre)
	if err != nil {
		log.Println("Failed inserting user genre:  ", err)		
		return
	}

	_, err = db.insertUserInstrumentStmt.Exec(instrument)
	if err != nil {
		log.Println("Failed inserting user instrument:  ", err)
		return
	}

	success = true;
	return
}

func (db *userRepository) UpdateProfile(name string, gender string, birthdate string, bio string, avatarurl string, id string, genre string, instrument string) (success bool, err error){
	success = false

	_, err = db.userProfileStmt.Exec(name, gender, birthdate, bio, avatarurl, id)
	if err != nil {
		log.Println("Failed updating profile:", err)
		return
	}

	_, err = db.updateUserGenreStmt.Exec(genre, id)
	if err != nil {
		log.Println("Failed updating user's genre:", err)
		return
	}

	_, err = db.updateUserInstrumentStmt.Exec(instrument, id)
	if err != nil {
		log.Println("Failed updating user's instrument:", err)
		return
	}

	success = true;
	return
}

func (db *userRepository) InsertNewPost(newPost UserPost) (success bool, err error){
	success = false
	_, err = db.insertNewPostStmt.Exec(newPost)
	if err != nil {
		log.Println("Failed uploading new post:", err)
		return
	}
	success = true
	return
}

func (db *userRepository) DeletePost(postId string) (success bool, err error){
	success = false
	_, err = db.deletePostStmt.Exec(postId)
	if err != nil {
		log.Println("Failed to delete post:,", err)
		return
	}
	success = true
	return
}