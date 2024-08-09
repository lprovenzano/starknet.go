package utils

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core/felt"
	"strings"
)

func UnwrapJSON(data map[string]interface{}, tag string) (map[string]interface{}, error) {
	if data[tag] != nil {
		var unwrappedData map[string]interface{}
		dataInner, err := json.Marshal(data[tag])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(dataInner, &unwrappedData); err != nil {
			return nil, err
		}
		return unwrappedData, nil
	}
	return data, nil
}

func NormalizeAddress(address *felt.Felt) string {
	addressStr := address.String()
	addressChars := len(addressStr)
	if addressChars > 66 {
		panic("malformed address, it's too long")
	}
	if addressChars < 66 {
		charsDiff := 66 - addressChars
		var trailingZeros strings.Builder
		for i := 0; i < charsDiff; i++ {
			trailingZeros.WriteString("0")
		}
		addressStr = addressStr[:2] + trailingZeros.String() + addressStr[2:]
	}
	return addressStr
}
