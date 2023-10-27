package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string, dv int) int {
	if len(key) == 0 {
		return dv
	}

	v, err := strconv.Atoi(GetString(key))
	mustParseKey(err, key)
	return v
}

func mustParseKey(err error, key string) {
	if err != nil {
		log.Warning("Could not parse key :%s,error:%s", key, err)
	}
}

func GetIntt(key string) int {
	v, err := strconv.Atoi(GetString(key))
	mustParseKey(err, key)
	return v
}
