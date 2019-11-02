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

	"gopkg.in/mgo.v2"
)

var (
	database *mgo.Database
	session  *mgo.Session
)

func MongoServerInit(mongoURL, dbName string) {
	var err error
	url := fmt.Sprintf("mongodb://%v", mongoURL)
	fmt.Printf("mongodb url %v\n", url)
	for {
		session, err = mgo.Dial(url)
		if err == nil {
			break
		}
		fmt.Printf("mgo.Dial url=%v, err=%v\n", url, err)
		time.Sleep(time.Duration(1) * time.Second)
	}
	session.SetMode(mgo.Monotonic, true)
	database = session.DB(dbName)
	fmt.Printf("mongodb mongoServerInit finished.\n")
}

func Fsync(async bool) error {
	return session.Fsync(async)
}
