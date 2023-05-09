/* Users*/
CREATE TABLE IF NOT EXISTS users ( id UUID PRIMARY KEY default UUID(), first_name varchar(255) NOT NULL, last_name varchar(255) NOT NULL, email varchar(255) NOT NULL, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL);

* Orders */
create table if not exists orders ( id uuid primary key default uuid(), task varchar(50) not null, details varchar(255) null, user_id UUID not null , created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL, foreign key (user_id) references users(id));

/* Image Groups */
create table if not exists image_groups ( id uuid primary key default uuid(), name varchar(50) NOT NULl, user_id UUID not null, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL, foreign key (user_id) references users(id));

/* Images */
create table if not exists images ( id uuid primary key default uuid(), name varchar(50) NOT NULL, url varchar(50) NOT NULL, small_url varchar(50), image_group_id UUID not null, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL, deleted_at DATETIME  NULL, foreign key (image_group_id) references image_groups(id));
