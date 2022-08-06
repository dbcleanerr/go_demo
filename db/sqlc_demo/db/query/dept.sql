-- name: CreateDept :one
INSERT INTO dept (dname, loc)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteGetDept :one
DELETE FROM dept
 WHERE deptno = $1
RETURNING *;

-- name: UpdateGetDept :one
UPDATE dept
   SET dname = $1,
       loc = $2
 WHERE deptno = $3
RETURNING *;

-- name: GetDept :one
SELECT *
  FROM dept
 WHERE deptno = $1 LIMIT 1;
