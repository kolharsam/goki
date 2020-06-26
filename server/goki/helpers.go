package goki

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/kolharsam/goki/server/logger"
)

// Some helper methods

func makeNewStoreValue(value string) (StoredValue, error) {
	if num, err := checkIfInt(value); err == nil {
		return StoredValue{
			Value:     num,
			TimeAlive: -1,
			Type:      Int,
		}, nil
	}

	if num, err := checkIfFloat(value); err == nil {
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
	if num, err := checkIfInt(value); err == nil {
		return StoredValue{
			Value:     num,
			TimeAlive: time,
			Type:      Int,
		}, nil
	}

	if num, err := checkIfFloat(value); err == nil {
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

func checkIfInt(value string) (int32, error) {
	num, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return -1, errors.New("not an int")
	}
	return int32(num), nil
}

func checkIfFloat(value string) (float32, error) {
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
