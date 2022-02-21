package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var errInvalidEvent = errors.New("could not parse Type: Invalid event")

func ParseRequest(msg []byte, isAuthClient bool) (Request, error) {
	request, err := Parse(msg, isAuthClient)
	if err != nil {
		return request, err
	}

	return request, nil
}

func Parse(msg []byte, isAuthClient bool) (Request, error) {
	var v map[string]interface{}
	var parsed Request

	if err := json.Unmarshal(msg, &v); err != nil {
		return parsed, fmt.Errorf("could not parse message: %w", err)
	}

	switch v["event"] {
	case "subscribe":
		parsed.Method = "subscribe"
		streams, ok := v["streams"]
		if !ok {
			return parsed, fmt.Errorf("no streams provided")
		}
		switch reflect.TypeOf(streams).Kind() {
		case reflect.Slice:
			streams := reflect.ValueOf(v["streams"])
			for i := 0; i < streams.Len(); i++ {
				parsed.Streams = append(parsed.Streams, streams.Index(i).Interface().(string))
			}
		}
	case "unsubscribe":
		parsed.Method = "unsubscribe"
		streams, ok := v["streams"]
		if !ok {
			return parsed, fmt.Errorf("no streams provided")
		}
		switch reflect.TypeOf(streams).Kind() {
		case reflect.Slice:
			streams := reflect.ValueOf(v["streams"])
			for i := 0; i < streams.Len(); i++ {
				parsed.Streams = append(parsed.Streams, streams.Index(i).Interface().(string))
			}
		}
	case "order":
		parsed.Method = "order"

		if !isAuthClient {
			return parsed, errInvalidEvent
		}

		bb, err := json.Marshal(v["data"])
		if err != nil {
			return parsed, err
		}

		parsed.Message = bb
	default:
		return parsed, errInvalidEvent
	}

	return parsed, nil
}
