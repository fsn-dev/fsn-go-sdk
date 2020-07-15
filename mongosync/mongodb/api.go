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
	"github.com/fsn-dev/fsn-go-sdk/efsn/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// --------------- add ---------------------------------

func AddBlock(mb *MgoBlock, overwrite bool) (err error) {
	if overwrite {
		_, err = collectionBlock.UpsertId(mb.Key, mb)
	} else {
		err = collectionBlock.Insert(mb)
	}
	if err == nil {
		log.Info("[mongodb] AddBlock success", "number", mb.Number, "hash", mb.Hash)
	} else if !mgo.IsDup(err) {
		log.Warn("[mongodb] AddBlock failed", "number", mb.Number, "hash", mb.Hash, "err", err)
	}
	return err
}

func AddTransaction(mt *MgoTransaction, overwrite bool) error {
	if overwrite {
		_, err := collectionTransaction.UpsertId(mt.Key, mt)
		return err
	}
	return collectionTransaction.Insert(mt)
}

func AddSyncInfo(msi *MgoSyncInfo) error {
	return collectionSyncInfo.Insert(msi)
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

func AddContract(mc *MgoContract) error {
	return collectionContracts.Insert(mc)
}

// --------------- update ---------------------------------

func UpdateSyncInfo(num, timestamp, td uint64) error {
	return collectionSyncInfo.UpdateId(KeyOfLatestSyncInfo,
		bson.M{"$set": bson.M{
			"number":          num,
			"timestamp":       timestamp,
			"totalDifficulty": td,
		}})
}

func UpdateBlockInfo(key string, blocktime, td uint64) error {
	return collectionBlock.UpdateId(key,
		bson.M{"$set": bson.M{
			"blockTime":       blocktime,
			"totalDifficulty": td,
		}})
}

// --------------- delete ---------------------------------

func DeleteBlock(hash string) error {
	return collectionBlock.Remove(bson.M{"hash": hash})
}

func DeleteTransaction(hash string) error {
	return collectionTransaction.Remove(bson.M{"hash": hash})
}

func DeleteLatestSyncInfo() error {
	return collectionSyncInfo.RemoveId(KeyOfLatestSyncInfo)
}

// --------------- find ---------------------------------

func FindBlockByNumber(num uint64) (*MgoBlock, error) {
	var block MgoBlock
	err := collectionBlock.Find(bson.M{"number": num}).One(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func FindBlock(hash string) (*MgoBlock, error) {
	var block MgoBlock
	err := collectionBlock.Find(bson.M{"hash": hash}).One(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func FindBlocksInRange(start, end uint64) ([]*MgoBlock, error) {
	count := int(end - start + 1)
	blocks := make([]*MgoBlock, count)
	iter := collectionBlock.Find(bson.M{"number": bson.M{"$gte": start, "$lte": end}}).Limit(count).Iter()
	err := iter.All(&blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func FindTransaction(hash string) (*MgoTransaction, error) {
	var tx MgoTransaction
	err := collectionTransaction.Find(bson.M{"hash": hash}).One(&tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func FindLatestSyncInfo() (*MgoSyncInfo, error) {
	var info MgoSyncInfo
	err := collectionSyncInfo.FindId(KeyOfLatestSyncInfo).One(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func FindContract(address string) (*MgoContract, error) {
	var contract MgoContract
	err := collectionContracts.FindId(address).One(&contract)
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
