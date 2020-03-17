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
	"gopkg.in/mgo.v2/bson"
)

var (
	collectionBlock       *mgo.Collection
	collectionTransaction *mgo.Collection
	collectionSyncInfo    *mgo.Collection
)

func getOrInitCollection(table string, collection **mgo.Collection, indexKey string) *mgo.Collection {
	if *collection == nil {
		*collection = database.C(table)
		if indexKey != "" {
			(*collection).EnsureIndexKey(indexKey)
		}
	}
	return *collection
}

func getCollection(table string) *mgo.Collection {
	switch table {
	case tbBlocks:
		return getOrInitCollection(table, &collectionBlock, "number")
	case tbTransactions:
		return getOrInitCollection(table, &collectionTransaction, "blockNumber")
	case tbSyncInfo:
		return getOrInitCollection(table, &collectionSyncInfo, "")
	}
	panic("unknown talbe " + table)
}

// --------------- add ---------------------------------

func AddBlock(mb *MgoBlock, overwrite bool) error {
	if overwrite {
		_, err := getCollection(tbBlocks).UpsertId(mb.Key, mb)
		return err
	}
	return getCollection(tbBlocks).Insert(mb)
}

func AddTransaction(mt *MgoTransaction, overwrite bool) error {
	if overwrite {
		_, err := getCollection(tbTransactions).UpsertId(mt.Key, mt)
		return err
	}
	return getCollection(tbTransactions).Insert(mt)
}

func AddSyncInfo(msi *MgoSyncInfo) error {
	return getCollection(tbSyncInfo).Insert(msi)
}

func InitLatestSyncInfo() error {
	_, err := FindLatestSyncInfo()
	if err == nil {
		return nil
	}
	msi := &MgoSyncInfo{
		Key: KeyOfLatestSyncInfo,
	}
	return AddSyncInfo(msi)
}

// --------------- update ---------------------------------

func UpdateSyncInfo(num, timestamp, td uint64) error {
	return getCollection(tbSyncInfo).UpdateId(KeyOfLatestSyncInfo,
		bson.M{"$set": bson.M{
			"number":          num,
			"timestamp":       timestamp,
			"totalDifficulty": td,
		}})
}

func UpdateBlockInfo(key string, blocktime, td uint64) error {
	return getCollection(tbBlocks).UpdateId(key,
		bson.M{"$set": bson.M{
			"blockTime":       blocktime,
			"totalDifficulty": td,
		}})
}

// --------------- delete ---------------------------------

func DeleteBlock(hash string) error {
	return getCollection(tbBlocks).Remove(bson.M{"hash": hash})
}

func DeleteTransaction(hash string) error {
	return getCollection(tbTransactions).Remove(bson.M{"hash": hash})
}

func DeleteLatestSyncInfo() error {
	return getCollection(tbSyncInfo).RemoveId(KeyOfLatestSyncInfo)
}

// --------------- find ---------------------------------

func FindBlockByNumber(num uint64) (*MgoBlock, error) {
	var block MgoBlock
	err := getCollection(tbBlocks).Find(bson.M{"number": num}).One(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func FindBlock(hash string) (*MgoBlock, error) {
	var block MgoBlock
	err := getCollection(tbBlocks).Find(bson.M{"hash": hash}).One(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func FindBlocksInRange(start, end uint64) ([]*MgoBlock, error) {
	count := int(end - start + 1)
	blocks := make([]*MgoBlock, count)
	iter := getCollection(tbBlocks).Find(bson.M{"number": bson.M{"$gte": start, "$lte": end}}).Limit(count).Iter()
	err := iter.All(&blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func FindTransaction(hash string) (*MgoTransaction, error) {
	var tx MgoTransaction
	err := getCollection(tbTransactions).Find(bson.M{"hash": hash}).One(&tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func FindLatestSyncInfo() (*MgoSyncInfo, error) {
	var info MgoSyncInfo
	err := getCollection(tbSyncInfo).FindId(KeyOfLatestSyncInfo).One(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
