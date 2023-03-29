-- name: CreateUser :one 
INSERT INTO "user" (
    name
) VALUES ($1)
RETURNING *;

-- name: ListOfUsers :many
SELECT "user".*, "video".id AS video_id FROM "user","video"
WHERE video_id = video.video_id AND video.id = ANY($1::bigint[]);

-- name: GetAuthorByID :one
SELECT * FROM "user"
WHERE id=$1;
