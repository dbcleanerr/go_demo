package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDept(t *testing.T) {
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
}

func TestDeleteGetDept(t *testing.T) {
	testQueries.DeleteGetDept(context.Background(), 2)
}
