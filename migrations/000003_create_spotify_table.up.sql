CREATE TABLE spotify (
    account_username VARCHAR(15) PRIMARY KEY,
    access_token CHAR(310),
    FOREIGN KEY (account_username) REFERENCES account(username)
);
