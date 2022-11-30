CREATE TABLE posts (
    id int primary key,
    user_id int not null,
    title text not null,
    body text not null
);