package goki

import "sync"

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
	value := key + " was not found"
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

// Some helper methods

func encodeToBytes(str string) []byte {
	return []byte(str)
}

func decodeToString(bs []byte) string {
	return string(bs)
}
