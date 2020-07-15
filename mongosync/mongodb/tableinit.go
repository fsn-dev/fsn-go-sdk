// Copyright 2019 The fsn-go-sdk Authors
// This file is part of the fsn-go-sdk library.
//
// The fsn-go-sdk library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fsn-go-sdk library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fsn-go-sdk library. If not, see <http://www.gnu.org/licenses/>.

package mongodb

import (
	"gopkg.in/mgo.v2"
)

var (
	collectionBlock       *mgo.Collection
	collectionTransaction *mgo.Collection
	collectionSyncInfo    *mgo.Collection
	collectionContracts   *mgo.Collection
)

func deinintCollections() {
	collectionBlock = database.C(tbBlocks)
	collectionTransaction = database.C(tbTransactions)
	collectionSyncInfo = database.C(tbSyncInfo)
	collectionContracts = database.C(tbContracts)
}

// do this when reconnect to the database
func initCollections() {
	initCollection(tbBlocks, &collectionBlock, "number")
	initCollection(tbTransactions, &collectionTransaction, "blockNumber")
	initCollection(tbSyncInfo, &collectionSyncInfo, "")
	initCollection(tbContracts, &collectionContracts, "")

	InitLatestSyncInfo()
}

func initCollection(table string, collection **mgo.Collection, indexKey ...string) {
	*collection = database.C(table)
	if len(indexKey) != 0 && indexKey[0] != "" {
		_ = (*collection).EnsureIndexKey(indexKey...)
	}
}
