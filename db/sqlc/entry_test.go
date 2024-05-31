package db

import (
	"context"
	"testing"

	"github.com/jesseinvent/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {

	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: util.RandomSqlInt(account.ID),
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}
