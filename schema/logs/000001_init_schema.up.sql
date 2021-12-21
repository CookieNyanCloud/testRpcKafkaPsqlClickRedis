CREATE TABLE IF NOT EXISTS user_logs
(
    Id uuid PRIMARY KEY,
    Name TEXT,
    Key uuid,
    Created_at timestamp   NOT NULL DEFAULT (now())
);