CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE posts
(
    id serial not null unique,
    title varchar(255) not null,
    content varchar(255)
);

CREATE TABLE comments
(
    id serial not null unique,
    content varchar(255)
);

CREATE TABLE users_posts
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    post_id int references posts (id) on delete cascade not null
);

CREATE TABLE users_comments
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    comment_id int references comments (id) on delete cascade not null
);

CREATE TABLE posts_comments
(
    id serial not null unique,
    comment_id int references comments (id) on delete cascade not null,
    post_id int references posts (id) on delete cascade not null
);