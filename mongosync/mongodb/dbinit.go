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
	"fmt"
	"time"

	"github.com/fsn-dev/fsn-go-sdk/efsn/log"
	"gopkg.in/mgo.v2"
)

var (
	database *mgo.Database
	session  *mgo.Session

	dialInfo *mgo.DialInfo
)

// MongoServerInit int mongodb server session
func MongoServerInit(addrs []string, dbname, user, pass string) {
	initDialInfo(addrs, dbname, user, pass)
	mongoConnect()
	initCollections()
	go checkMongoSession()
}

func initDialInfo(addrs []string, db, user, pass string) {
	dialInfo = &mgo.DialInfo{
		Addrs:    addrs,
		Database: db,
		Username: user,
		Password: pass,
	}
}

func mongoConnect() {
	if session != nil { // when reconnect
		session.Close()
	}
	log.Info("[mongodb] connect database start.", "addrs", dialInfo.Addrs, "dbName", dialInfo.Database)
	var err error
	for {
		session, err = mgo.DialWithInfo(dialInfo)
		if err == nil {
			break
		}
		log.Warn("[mongodb] dial error", "err", err)
		time.Sleep(1 * time.Second)
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{FSync: true})
	database = session.DB(dialInfo.Database)
	deinintCollections()
	log.Info("[mongodb] connect database finished.", "dbName", dialInfo.Database)
}

// fix 'read tcp 127.0.0.1:43502->127.0.0.1:27917: i/o timeout'
func checkMongoSession() {
	for {
		time.Sleep(60 * time.Second)
		if err := ensureMongoConnected(); err != nil {
			log.Info("[mongodb] check session error", "err", err)
			log.Info("[mongodb] reconnect database", "dbName", dialInfo.Database)
			mongoConnect()
		}
	}
}

func sessionPing() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recover from error %v", r)
		}
	}()
	for i := 0; i < 6; i++ {
		err = session.Ping()
		if err == nil {
			break
		}
		time.Sleep(10 * time.Second)
	}
	return err
}

func ensureMongoConnected() (err error) {
	err = sessionPing()
	if err != nil {
		log.Error("[mongodb] session ping error", "err", err)
		log.Info("[mongodb] refresh session.", "dbName", dialInfo.Database)
		session.Refresh()
		database = session.DB(dialInfo.Database)
		deinintCollections()
		err = sessionPing()
	}
	return err
}

// Fsync flush memory to db
func Fsync(async bool) error {
	return session.Fsync(async)
}
