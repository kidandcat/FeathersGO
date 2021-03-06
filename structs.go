package main

import (
	redis "gopkg.in/redis.v5"
)

type ResultString struct {
	Data  string
	Error error
}
type ResultStringArray struct {
	Data  []string
	Error error
}
type ResultJsnArray struct {
	Data  []jsn
	Error error
}
type ResultJsn struct {
	Data  []string
	Error error
}
type ResultBoolean struct {
	Data  bool
	Error error
}

type DataKey struct {
	Key  string `json:"key"`
	Data Anon   `json:"data"`
}

type Anon interface{}

type dbc interface {
	Find(string) ResultJsnArray
	Get(string) ResultString
	Set(string, string) ResultBoolean
	Del(string) ResultBoolean
	Exists(string) ResultBoolean
}

type Config struct {
	Port         string `json:"port"`
	StaticFolder string `json:"staticFolder"`
	ApiPrefix    string `json:"apiPrefix"`
}

type GlobalContext struct{}

type jsn map[string]interface{}

type memoryStorage struct {
	Memory map[string]string
}

type redisStorage struct {
	Memory *redis.Client
}
