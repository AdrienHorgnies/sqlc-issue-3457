-- name: InsertUsers :exec
INSERT INTO users (name)
VALUES ("foo");

-- name: UpdateUsers :exec
UPDATE users
SET NAME = sqlc.arg(name)
WHERE sqlc.arg(name) IN ("bar");
