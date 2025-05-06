-- name: GetTasksByCode :one
SELECT * FROM TASKS
WHERE CODE = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM TASKS
ORDER BY ID;

-- name: CreateTask :one
INSERT INTO TASKS (
  CODE, NAME, FREQUENCY_DATE, FREQUENCY_TIME, MAX_RETRIES
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateTasks :exec
UPDATE TASKS
  SET NAME = $2,
  STATUS = $3
WHERE ID = $1;

-- name: DeleteTasks :exec
DELETE FROM TASKS
WHERE ID = $1;

-- name: GetEnabledTask :many
SELECT * FROM TASKS WHERE IS_ENABLED = TRUE;