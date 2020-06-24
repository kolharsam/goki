package goki

import (
	"encoding/json"
	"errors"
	"strconv"
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
