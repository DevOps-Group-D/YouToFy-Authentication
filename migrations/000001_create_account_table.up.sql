CREATE TABLE account (
    username VARCHAR(15) PRIMARY KEY,
    password VARCHAR(60) NOT NULL,
    session_token CHAR(44),
    csrf_token CHAR(44)
);

CREATE TABLE spotify (
    account_username VARCHAR(15) PRIMARY KEY,
    access_token CHAR(44),
    FOREIGN KEY (account_username) REFERENCES account(username)
);
