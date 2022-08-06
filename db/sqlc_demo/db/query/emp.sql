-- name: CreateEmp :one
INSERT INTO emp (ename, job, mgr, hiredate, sal, comm, deptno)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteEmp :exec
DELETE FROM emp
 WHERE empno = $1;

-- name: DeleteGetEmp :one
DELETE FROM emp
 WHERE empno = $1
RETURNING *;

-- name: UpdateGetEmp :one
UPDATE emp
  SET ename = $1,
      job = $2,
      mgr = $3,
      hiredate = $4,
      sal = $5,
      comm = $6,
      deptno = $7
WHERE empno = $8
RETURNING *;

-- name: GetEmp :one
SELECT *
  FROM emp
 WHERE empno = $1 LIMIT 1;

-- name: ListEmpByEmpnoArray :many
SELECT *
  FROM emp
 WHERE empno = ANY($1::int[]);

-- name: ListEmps :many
SELECT *
  FROM emp
 LIMIT $1
OFFSET $2;





