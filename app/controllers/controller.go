package controllers

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

// CtrlErr Generic controller error interface
type CtrlErr map[string]interface{}

func parseUintOrDefault(intStr string, _default uint64) uint64 {
	value, err := strconv.ParseUint(intStr, 0, 64)
	if err != nil {
		return _default
	}
	return value
}

func parseIntOrDefault(intStr string, _default int64) int64 {
	value, err := strconv.ParseInt(intStr, 0, 64)
	if err != nil {
		return _default
	}
	return value
}

func convertToObjectIDHex(id string) (result bson.ObjectId, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Unable to convert %v to object id", id)
		}
	}()

	return bson.ObjectIdHex(id), err
}

func buildErrResponse(err error, errorCode string) CtrlErr {
	ctrlErr := CtrlErr{}
	ctrlErr["error_message"] = err.Error()
	ctrlErr["error_code"] = errorCode
	return ctrlErr
}
