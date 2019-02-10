USE musiciandb;

CREATE TABLE user_detail (
	id BIGINT NOT NULL,
	email VARCHAR(45) NOT NULL,
	username VARCHAR(20) NOT NULL,
	password VARCHAR(45) NOT NULL,
	name VARCHAR(45) NOT NULL,
	gender VARCHAR(10) NOT NULL,
	birthdate VARCHAR(10),
	bio VARCHAR(200) NOT NULL,
	role VARCHAR(80) NOT NULL, 
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
);

CREATE TABLE genre_list (
	id INT NOT NULL,
	genre VARCHAR(20) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE instrument_list(
	id INT NOT NULL,
	instrument VARCHAR(20) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE user_genre (
	id INT NOT NULL,
	user_id BIGINT NOT NULL,
	genre_id INT NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE user_instrument(
	id INT NOT NULL,
	user_id BIGINT NOT NULL,
	instrument_id INT NOT NULL,
	PRIMARY KEY(id)
);