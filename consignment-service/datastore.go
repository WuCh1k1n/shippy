package main

import (
	"gopkg.in/mgo.v2"
)

// 创建与 MongoDB 交互的主会话
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}