package types

import (
	"fmt"
	"strconv"
)

type Duration int

func (d Duration) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", d)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

type Envelope map[string]interface{}
