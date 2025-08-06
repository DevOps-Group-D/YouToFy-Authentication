CREATE TABLE account (
    username VARCHAR(15) primary key,
    password VARCHAR(60) NOT NULL,
    session_token CHAR(44),
    csrf_token CHAR(44)
);
