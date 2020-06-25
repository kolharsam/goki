package goki

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/kolharsam/goki/common"
	"github.com/kolharsam/goki/server/logger"
)

// TODO: A lot of the functions have more or less the similar
// function body so far. Maybe an abstraction can be made in this regard.

// all the basic types that are currently supported
const (
	String = "string"
	Int    = "int"
	Float  = "float"
)

// NOTE: Redis doesn't support floats out of the box. This does. I guess it's okay

// StoredValue is the Struct used to store the values within the store
type StoredValue struct {
	Value     interface{}
	TimeAlive int
	Type      string
}

// this is trying to make the usage of this map concurrent
var store = struct {
	sync.RWMutex
	gokiStore map[string][]byte
}{gokiStore: make(map[string][]byte)}

// SetToStorage sets values to the map
func SetToStorage(key string, value string) common.GokiResult {
	store.Lock()
	byteVal, err := makeNewStoreValue(value)
	if err != nil {
		store.Unlock()
		return common.GokiResult{
			Value:    err.Error(),
			ToFormat: false,
		}
	}
	newEncodedVal, err := encodeStructToBytes(byteVal)
	if err != nil {
		store.Unlock()
		return common.GokiResult{
			Value:    err.Error(),
			ToFormat: false,
		}
	}
	store.gokiStore[key] = newEncodedVal
	store.Unlock()

	// making it redis-cli like
	return common.GokiResult{
		Value:    "OK",
		ToFormat: false,
	}
}

// GetValue returns the value of the
func GetValue(key string) common.GokiResult {
	value := ""
	toFormat := false
	store.RLock()
	if _, ok := store.gokiStore[key]; ok != false {
		strVal, err := decodeStructToBytes(store.gokiStore[key])
		if err != nil {
			store.RUnlock()
			return common.GokiResult{
				Value:    "(nil)",
				ToFormat: false,
			}
		}
		if strVal.Type == "int" {
			value += "(integer) "
		}
		if strVal.Type == "float" {
			value += "(float) "
		}
		value += fmt.Sprintf("%v", strVal.Value)
		if strVal.Type == "string" {
			toFormat = true
		}

		return common.GokiResult{
			Value:    value,
			ToFormat: toFormat,
		}
	}
	store.RUnlock()

	return common.GokiResult{
		Value:    "Err: Key not found",
		ToFormat: toFormat,
	}
}

// CheckIfExists returns "YES" if a key is present in the db
func CheckIfExists(key string) common.GokiResult {
	value := "NO"
	store.RLock()
	if _, ok := store.gokiStore[key]; ok != false {
		value = "YES"
	}
	store.RUnlock()

	return common.GokiResult{
		Value:    value,
		ToFormat: false,
	}
}

// DeleteKey deletes a key if it is present
func DeleteKey(key string) common.GokiResult {
	value := "(nil)"
	store.RLock()
	if _, ok := store.gokiStore[key]; ok != false {
		delete(store.gokiStore, key)
		value = "OK"
	}
	store.RUnlock()

	return common.GokiResult{
		Value:    value,
		ToFormat: false,
	}
}

// ExpireKey deletes a particular key upon a certain given expiry time
// the time needs to be in seconds (for now, TODO: to add more precision soon)
func ExpireKey(key string, expireTime string) {
	timeNumber, err := strconv.Atoi(expireTime)
	if err != nil {
		logger.LogWarning(time.Now().String(), "Input for key expiry wasn't correct.")
		return
	}

	if currentStr, ok := store.gokiStore[key]; ok != false {
		// There are 2 issues to be fixed with this:
		// 1. There's no global way of stopping the timer, if the value is updated from the user
		//    This might work out because setting a value sets the TimeAlive to -1 and so essentially
		//    the changes should come on updateTTL where we can verify the current objects time and
		//    the time left, that way we can easily stop the timer
		// 2. Updates do not use the current mutex present on store, the timer stops working for some
		//    reason. Should look for more references regarding this issue

		go updateTTL(timeNumber, currentStr, key)
	}
}

// FetchTTL returns the Time To Live value for a particular key. Default is -1.
func FetchTTL(key string) common.GokiResult {
	store.RLock()
	if currentValue, ok := store.gokiStore[key]; ok != false {
		decodedValue, err := decodeStructToBytes(currentValue)
		if err != nil {
			store.RUnlock()
			return common.GokiResult{
				Value:    "ERR: in fetching value from memory",
				ToFormat: false,
			}
		}
		return common.GokiResult{
			Value:    strconv.Itoa(decodedValue.TimeAlive),
			ToFormat: false,
		}
	}
	store.RUnlock()

	return common.GokiResult{
		Value:    "ERR: Key not found",
		ToFormat: false,
	}
}
