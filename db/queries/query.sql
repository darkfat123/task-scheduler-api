-- name: GetTasksByCode :one
SELECT * FROM tasks
WHERE code = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY id;

-- name: CreateTask :one
INSERT INTO tasks (
  code, name, frequency_date, frequency_time, max_retries
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateTasks :exec
UPDATE tasks
  set name = $2,
  status = $3
WHERE id = $1;

-- name: DeleteTasks :exec
DELETE FROM tasks
WHERE id = $1;