package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestDept(t *testing.T) Dept {
	arg := CreateDeptParams{
		Dname: "sale",
		Loc:   "shanghai",
	}

	dept, err := testQueries.CreateDept(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, dept)
	require.Equal(t, arg.Dname, dept.Dname)
	require.Equal(t, arg.Loc, dept.Loc)
	// default 值测试
	require.NotZero(t, dept.Deptno)

	return dept
}

func TestQueries_CreateDept(t *testing.T) {
	createTestDept(t)
}

func TestQueries_GetDept(t *testing.T) {
	dept1 := createTestDept(t)
	dept2, err := testQueries.GetDept(context.Background(), dept1.Deptno)

	require.NoError(t, err)
	require.NotEmpty(t, dept2)

	require.Equal(t, dept1.Dname, dept2.Dname)
	require.Equal(t, dept1.Deptno, dept2.Deptno)
	require.Equal(t, dept1.Loc, dept2.Loc)
}

func TestQueries_DeleteGetDept(t *testing.T) {
	dept1 := createTestDept(t)
	dept2, err := testQueries.DeleteGetDept(context.Background(), dept1.Deptno)

	require.NoError(t, err)
	require.NotEmpty(t, dept2)

	require.Equal(t, dept1.Dname, dept2.Dname)
	require.Equal(t, dept1.Deptno, dept2.Deptno)
	require.Equal(t, dept1.Loc, dept2.Loc)

	dept3, err := testQueries.GetDept(context.Background(), dept1.Deptno)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, dept3)
}

func TestQueries_UpdateGetDept(t *testing.T) {
	dept1 := createTestDept(t)

	args := UpdateGetDeptParams{
		Dname:  "jiuge dev",
		Deptno: dept1.Deptno,
		Loc:    dept1.Loc,
	}

	dept2, err := testQueries.UpdateGetDept(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, dept2)

	require.Equal(t, dept2.Dname, args.Dname)
	require.Equal(t, dept2.Deptno, dept1.Deptno)
	require.Equal(t, dept2.Loc, dept1.Loc)

	//require.WithinDuration(t, )
}
