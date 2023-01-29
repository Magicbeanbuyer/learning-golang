package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
)

func TestQueries_CreateAccount(t *testing.T) { // question: what is testing.T?
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQuery.CreateAccount(context.Background(), args) // question: what is context background?
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

/*
go caches test results, running `make test` in terminal twice, only the first time will write a row to the database.
to rerun the test, force clean the go test cache.
*/
