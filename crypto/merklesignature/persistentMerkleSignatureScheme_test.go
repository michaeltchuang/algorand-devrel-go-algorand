// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package merklesignature

import (
	"context"
	"database/sql"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/algorand/go-algorand/test/partitiontest"
	"github.com/algorand/go-algorand/util/db"
)

func TestFetchRestoreAllSecrets(t *testing.T) {
	partitiontest.PartitionTest(t)
	a := require.New(t)
	store := createTestDB(a)
	defer store.Close()

	firstValid := uint64(1)
	LastValid := uint64(5000)

	interval := uint64(256)
	mss, err := New(firstValid, LastValid, interval)
	a.NoError(err)
	a.NoError(mss.Persist(*store))

	newMss := Secrets{}
	newMss.SignerContext = mss.SignerContext
	err = newMss.RestoreAllSecrets(*store)
	a.NoError(err)

	for i := uint64(1); i < LastValid; i++ {
		key1 := mss.GetKey(i)
		key2 := newMss.GetKey(i)
		if i%interval == 0 {
			a.NotNil(key1)
			a.NotNil(key2)
			a.Equal(*key1, *key2)
			continue
		}
		a.Nil(key1)
		a.Nil(key2)
	}

}

func createTestDB(a *require.Assertions) *db.Accessor {
	tmpname := uuid.NewV4().String() // could this just be a constant string instead? does it even matter?
	store, err := db.MakeAccessor(tmpname, false, true)
	a.NoError(err)
	a.NotNil(store)

	err = store.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err = tx.Exec(`CREATE TABLE schema (
			tablename TEXT PRIMARY KEY,
			version INTEGER
		);`)
		return err
	})
	a.NoError(err)

	return &store
}
