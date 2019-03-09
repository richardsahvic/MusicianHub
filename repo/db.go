package repo

import (
	"fmt"
	"log"
	"os"

	sqlx "github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn                      	*sqlx.DB
	findIDStmt                	*sqlx.Stmt
	findEmailStmt             	*sqlx.Stmt
	updatePasswordStmt        	*sqlx.Stmt
	getGenreStmt				*sqlx.Stmt
	getInstrumentStmt			*sqlx.Stmt
	getFollowerStmt				*sqlx.Stmt
	getFollowingStmt			*sqlx.Stmt
	getFollowedIdStmt			*sqlx.Stmt
	getRelatedPostStmt			*sqlx.Stmt
	getUserPostStmt				*sqlx.Stmt
	getUserProfileStmt			*sqlx.Stmt
	getFollowsDataStmt			*sqlx.Stmt
	userProfileStmt				*sqlx.Stmt
	updateUserGenreStmt			*sqlx.Stmt
	updateUserInstrumentStmt	*sqlx.Stmt
	deletePostStmt				*sqlx.Stmt
	insertNewUserStmt         	*sqlx.NamedStmt
	insertUserGenreStmt			*sqlx.NamedStmt
	insertUserInstrumentStmt	*sqlx.NamedStmt
	insertNewPostStmt			*sqlx.NamedStmt
	insertNewFollowStmt			*sqlx.NamedStmt
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
	r.getGenreStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.genre_list")
	r.getInstrumentStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.instrument_list")
	r.getFollowedIdStmt = r.MustPrepareStmt("SELECT followed_id FROM musiciandb.user_follow WHERE user_id=?")
	r.getRelatedPostStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.user_post WHERE user_id IN (?) ORDER BY created_at DESC")
	r.getFollowingStmt = r.MustPrepareStmt("SELECT d.id, d.name, d.avatar_url FROM user_follow f INNER JOIN user_detail d ON f.followed_id=d.id WHERE f.user_id=?")
	r.getFollowerStmt = r.MustPrepareStmt("SELECT d.id, d.name, d.avatar_url FROM user_follow f INNER JOIN user_detail d ON f.user_id=d.id WHERE f.followed_id=?")
	r.getUserPostStmt = r.MustPrepareStmt("SELECT * FROM musiciandb.user_post WHERE user_id=?")
	r.getUserProfileStmt = r.MustPrepareStmt("SELECT name, email, gender, birthdate, bio, avatar_url FROM musiciandb.user_detail WHERE id=?")
	r.deletePostStmt = r.MustPrepareStmt("DELETE FROM musiciandb.user_post WHERE post_id=?")
	r.userProfileStmt = r.MustPrepareStmt("UPDATE musiciandb.user_detail SET name=?, gender=?, birthdate=?, bio=?, avatar_url=? WHERE id=?")
	r.insertNewUserStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_detail (id, email, password) VALUES (:id, :email, :password)")
	r.insertUserGenreStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_genre (user_id, genre_id) VALUES (:user_id, :genre_id)")
	r.insertUserInstrumentStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_instrument (user_id, instrument_id) VALUES (:user_id, :instrument_id)")
	r.insertNewPostStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_post (post_id, user_id, post_type, file_url, caption) VALUES (:post_id, :user_id, :post_type, :file_url, :caption)")
	r.insertNewFollowStmt = r.MustPrepareNamedStmt("INSERT INTO musiciandb.user_follow (user_id, followed_id) VALUES (:user_id, :followed_id)")
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
	err = db.getGenreStmt.Select(&genres)
	if err != nil{
		log.Println("Failed to get genres: ", err)
	}
	return
}

func (db *userRepository) GetInstruments() (instruments []InstrumentList, err error) {
	err = db.getInstrumentStmt.Select(&instruments)
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
		log.Println("Failed to delete post:", err)
		return
	}
	success = true
	return
}

func (db *userRepository) InsertFollow(userFollow UserFollow) (success bool, err error){
	success = false
	_, err = db.insertNewFollowStmt.Exec(userFollow)
	if err != nil {
		log.Println("Failed to follow user:", err)
		return
	}
	success = true
	return
}

func (db *userRepository) GetFollower(id string) (follower []UserDetail, err error){
	err = db.getFollowerStmt.Select(&follower, id)
	if err != nil {
		log.Println("Failed to get follower:", err)
	}
	return
}

func (db *userRepository) GetFollowing(id string) (following []UserDetail, err error){
	err = db.getFollowingStmt.Select(&following, id)
	if err != nil{
		log.Println("Failed to get following:", err)
	}
	return
}

func (db *userRepository) GetUserPost(id string) (posts []UserPost, err error){
	err = db.getUserPostStmt.Select(&posts, id)
	if err != nil {
		log.Println("Failed to get posts:", err)
	}
	return
}

func (db *userRepository) GetUserProfile(id string) (profile UserDetail, err error){
	err = db.getUserProfileStmt.Get(&profile, id)
	if err != nil {
		log.Println("Failed to get profile:", err)
	}
	return
}

func (db *userRepository) GetFollowedId(id string) (followedId []string, err error){
	var tempFollowedIds []string

	err = db.getFollowedIdStmt.Select(&tempFollowedIds, id)
	if err != nil {
		log.Println("Failed to get followed id:", err)
		return
	}

	tempidsLength := len(tempFollowedIds)
	idsLength := tempidsLength + 1

	followedId = make([]string, idsLength, idsLength)

	for i := range followedId {
		if i + 1 != idsLength {
			followedId[i] = tempFollowedIds[i]
		}else if i + 1 == idsLength {
			followedId[i] = id
		}
		// log.Println(followedId[i], i)
	}

	return
}

func (db *userRepository) GetRelatedPost(id []string) (posts []UserPost, err error){
	query, args, err := sqlx.In("SELECT * FROM musiciandb.user_post WHERE user_id IN (?) ORDER BY created_at DESC;", id)
	if err != nil {
		log.Println("Failed to get related post:", err)
		return
	}else {
		query = db.conn.Rebind(query)
		rows, err2 := db.conn.Queryx(query, args)
		if err2 != nil {
			log.Fatalln(err2)
			return
		}
		for rows.Next() {
			err = rows.StructScan(&posts)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	
	return
}