package goki

import (
	"strconv"
	"sync"
	"time"

	"github.com/kolharsam/goki/server/logger"
)

// TODO: A lot of the functions have more or less the similar
// function body so far. Maybe an abstraction can be made in this regard.

// this is trying to make the usage of this map concurrent
var store = struct {
	sync.RWMutex
	gokiStore map[string][]byte
}{gokiStore: make(map[string][]byte)}

// SetToStorage sets values to the map
func SetToStorage(key string, value string) string {
	store.Lock()
	store.gokiStore[key] = encodeToBytes(value)
	store.Unlock()

	// making it redis-cli like
	return "OK"
}

// GetValue returns the value of the
func GetValue(key string) string {
	value := "(nil)"
	store.RLock()
	if _, ok := store.gokiStore[key]; ok != false {
		value = decodeToString(store.gokiStore[key])
	}
	store.RUnlock()

	return value
}

// CheckIfExists returns "YES" if a key is present in the db
func CheckIfExists(key string) string {
	value := "NO"
	store.RLock()
	if _, ok := store.gokiStore[key]; ok != false {
		value = "YES"
	}
	store.RUnlock()

	return value
}

// DeleteKey deletes a key if it is present
func DeleteKey(key string) string {
	value := "(nil)"
	store.RLock()
	if _, ok := store.gokiStore[key]; ok != false {
		delete(store.gokiStore, key)
		value = "OK"
	}
	store.RUnlock()

	return value
}

// ExpireKey deletes a particular key upon a certain given expiry time
// the time needs to be in seconds (for now, TODO: will add more precision soon)
func ExpireKey(key string, expireTime string) {
	timeNumber, err := strconv.Atoi(expireTime)
	if err != nil {
		logger.LogWarning(time.Now().String(), "Input for key expiry wasn't right.")
		return
	}

	time.AfterFunc(time.Duration(timeNumber)*time.Second, func() {
		store.RLock()
		if _, ok := store.gokiStore[key]; ok != false {
			delete(store.gokiStore, key)
		}
		store.RUnlock()
	})
}

// TODO: Now that expiry is here. we can add ttl - this might have me make some changes to how data is stored
// Some helper methods

func encodeToBytes(str string) []byte {
	return []byte(str)
}

func decodeToString(bs []byte) string {
	return string(bs)
}
