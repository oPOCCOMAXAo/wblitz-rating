package data

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestSqliteStorage(t *testing.T) {
	dbFile := t.TempDir() + "/wblitz.test.db"

	stor, err := NewSqliteStorage(dbFile)
	require.NoError(t, err)

	suite.Run(t, &StorageSuite{
		Storage: stor,
	})
}
