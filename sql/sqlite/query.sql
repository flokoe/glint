-- name: GetProviders :many
SELECT * FROM providers ORDER BY id;

-- name: AddProvider :one
INSERT INTO providers (type, url) VALUES (?, ?) RETURNING *;

-- name: GetLLMs :many
SELECT * FROM llms ORDER BY id;

-- name: AddLLM :one
INSERT INTO llms (string, name, provider_id, context_size) VALUES (?, ?, ?, ?) RETURNING *;
