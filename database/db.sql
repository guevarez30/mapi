/* Users*/
CREATE TABLE IF NOT EXISTS users ( id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, UUID varchar(50) not null, first_name varchar(255) NOT NULL, last_name varchar(255) NOT NULL, email varchar(255) NOT NULL, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL);
        

/* Orders */
create table if not exists orders ( id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, UUID varchar(50) not null, task varchar(50) not null, details varchar(255) null, user_uuid varchar(50) not null , created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL);

/* Image Groups */
create table if not exists image_groups ( id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, UUID varchar(50) not null,  name varchar(50) NOT NULl, user_uuid varchar(50) not null, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL);

/* Images */
create table if not exists images ( id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, UUID varchar(50) not null, name varchar(50) NOT NULL, url varchar(50) NOT NULL, small_url varchar(50), image_group_uuid varchar(50) not null,user_uuid varchar(50), created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL);

INSERT INTO users (uuid, first_name, last_name, email, created_at, updated_at) VALUES ("b5c6379a-ebf9-4845-841b-e187ece03d4d", "first_name", "last_name", "guevarezfamily30@gmail.com", "2022-07-27 17:57:26", "2022-07-27 17:57:26");
