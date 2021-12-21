CREATE TABLE IF NOT EXISTS users
(
    Id uuid PRIMARY KEY,
    Name TEXT,
    Password_hash TEXT
);