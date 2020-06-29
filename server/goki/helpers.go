package goki

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/kolharsam/goki/common"
	"github.com/kolharsam/goki/server/logger"
)

// Some helper methods

func makeNewStoreValue(value string) (StoredValue, error) {
	if num, err := getIntValue(value); err == nil {
		return StoredValue{
			Value:     num,
			TimeAlive: -1,
			Type:      Int,
		}, nil
	}

	if num, err := getFloatValue(value); err == nil {
		return StoredValue{
			Value:     num,
			TimeAlive: -1,
			Type:      Float,
		}, nil
	}

	return StoredValue{
		Value:     value,
		TimeAlive: -1,
		Type:      String,
	}, nil
}

func makeTimedValue(value string, time int) (StoredValue, error) {
	if num, err := getIntValue(value); err == nil {
		return StoredValue{
			Value:     num,
			TimeAlive: time,
			Type:      Int,
		}, nil
	}

	if num, err := getFloatValue(value); err == nil {
		return StoredValue{
			Value:     num,
			TimeAlive: time,
			Type:      Float,
		}, nil
	}

	return StoredValue{
		Value:     value,
		TimeAlive: time,
		Type:      String,
	}, nil
}

func getIntValue(value string) (int, error) {
	num, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return -1, errors.New("not an int")
	}
	return int(num), nil
}

func getFloatValue(value string) (float32, error) {
	num, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return -1, errors.New("not an float")
	}
	return float32(num), nil
}

func encodeStructToBytes(v StoredValue) ([]byte, error) {
	return json.Marshal(v)
}

func decodeStructToBytes(bytesArr []byte) (StoredValue, error) {
	var unmarshalledVal StoredValue
	err := json.Unmarshal(bytesArr, &unmarshalledVal)
	return unmarshalledVal, err
}

func updateTTL(t int, key string) {
	for {
		// get current object
		currentValue, err := decodeStructToBytes(store.gokiStore[key])
		if err != nil {
			logger.LogWarning(time.Now().String(), "Unable to read key from memory")
			return
		}
		valueString := fmt.Sprintf("%v", currentValue.Value)

		// decrement time count

		// this case covers the point if the key has been edited while the timer is running
		if currentValue.TimeAlive <= 0 {
			return
		}
		t--

		newValue, err := makeTimedValue(valueString, t)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// encode the new object
		encodedNewVal, err := encodeStructToBytes(newValue)
		if err != nil {
			logger.LogWarning(time.Now().String(), "Unable to commit key to memory")
			return
		}

		store.gokiStore[key] = encodedNewVal
		if t <= 0 {
			delete(store.gokiStore, key)
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func performIntFunction(key string, updateFunc func(int) int) common.GokiResult {
	if currentValue, ok := store.gokiStore[key]; ok != false {
		strValue, err := decodeStructToBytes(currentValue)
		if err != nil {
			return common.GokiResult{
				ToFormat: false,
				Value:    err.Error(),
			}
		}
		if strValue.Type != Int {
			return common.GokiResult{
				ToFormat: false,
				Value:    "ERR: value for given key is not integer",
			}
		}
		currentVal := fmt.Sprintf("%v", strValue.Value)
		currentIntVal, err := getIntValue(currentVal)
		if err != nil {
			return common.GokiResult{
				ToFormat: false,
				Value:    err.Error(),
			}
		}
		newIntValue := updateFunc(currentIntVal)
		var newStr StoredValue
		newStr.TimeAlive = -1
		newStr.Type = Int
		newStr.Value = newIntValue
		newStructVal, err := encodeStructToBytes(newStr)
		if err != nil {
			return common.GokiResult{
				ToFormat: false,
				Value:    "ERR: could not commit new value to memory!",
			}
		}
		store.gokiStore[key] = newStructVal

		return common.GokiResult{
			ToFormat: false,
			Value:    "(integer) " + strconv.Itoa(int(newIntValue)),
		}
	}

	return common.GokiResult{
		Value:    "ERR: Key not found",
		ToFormat: false,
	}
}
