package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type cacheHandler struct {
	conn *redis.Client
}

func (ch *cacheHandler) get(key string, to interface{}) error {
	doc, err := ch.conn.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(doc, to)
}

func (ch *cacheHandler) set(key string, doc interface{}) error {
	docBytes, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	return ch.conn.Set(context.Background(), key, docBytes, 0).Err()
}

func (ch *cacheHandler) remove(key string) error {
	return ch.conn.Del(context.Background(), key).Err()
}

func setCacheHandler(host string) (*cacheHandler, error) {
	opts, err := redis.ParseURL(fmt.Sprintf(
		"redis://%s/0?protocol=3",
		host,
	))
	if err != nil {
		return nil, err
	}

	conn := redis.NewClient(opts)
	return &cacheHandler{
		conn,
	}, nil
}
