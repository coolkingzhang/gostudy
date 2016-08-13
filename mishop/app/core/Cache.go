package core

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var MemcacheClient = memcache.New("localhost:11211")
