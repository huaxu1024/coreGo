package test

import (
	"coreGo/nosql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	key, value := "goland:redis:set:01", "rookie"
	nosql.Set(key, value)
	assert.Equal(t, value, nosql.Get(key))
}

func TestSetEx(t *testing.T) {
	key, value := "goland:redis:set:ex:01", "rookie_ex"
	nosql.SetEx(key, value, "2")
	assert.Equal(t, value, nosql.Get(key))
	time.Sleep(2 * time.Second)
	assert.Equal(t, "", nosql.Get(key))
}

func TestDel(t *testing.T) {
	key, value := "goland:redis:set:del:01", "rookie_ex"
	nosql.Set(key, value)
	assert.Equal(t, value, nosql.Get(key))
	nosql.Del(key)
	assert.Equal(t, "", nosql.Get(key))
}

