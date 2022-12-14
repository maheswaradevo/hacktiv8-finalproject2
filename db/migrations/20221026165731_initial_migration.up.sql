CREATE TABLE IF NOT EXISTS `user`(
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR (255) NOT NULL,
    email VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    age INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (id),
    UNIQUE KEY unique_username(username)
);

CREATE TABLE IF NOT EXISTS `photo`(
	id INT NOT NULL AUTO_INCREMENT,
	title VARCHAR (255) NOT NULL,
    caption VARCHAR (255) NOT NULL,
    photo_url VARCHAR (255) NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY(user_id) REFERENCES `user`(id)
);

CREATE TABLE IF NOT EXISTS `social_media`(
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR (255) NOT NULL,
    social_media_url VARCHAR (255) NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY(user_id) REFERENCES `user`(id)
);

CREATE TABLE IF NOT EXISTS `comment`(
	id INT NOT NULL AUTO_INCREMENT,
	message VARCHAR (255) NOT NULL,
    photo_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY(user_id) REFERENCES `user`(id),
    FOREIGN KEY(photo_id) REFERENCES `photo`(id)
);