package jsonutil

import (
	"fmt"
	"strings"

	"github.com/CoderI421/gframework/pkg/common/json"
)

type JSONRawMessage []byte

func (m JSONRawMessage) Find(key string) JSONRawMessage {
	var objmap map[string]json.RawMessage
	err := json.Unmarshal(m, &objmap)
	if err != nil {
		fmt.Printf("Resolve JSON Key failed, find key =%s, err=%s",
			key, err)
		return nil
	}
	return JSONRawMessage(objmap[key])
}

func (m JSONRawMessage) ToList() []JSONRawMessage {
	var lists []json.RawMessage
	err := json.Unmarshal(m, &lists)
	if err != nil {
		fmt.Printf("Resolve JSON List failed, err=%s",
			err)
		return nil
	}
	var res []JSONRawMessage
	for _, v := range lists {
		res = append(res, JSONRawMessage(v))
	}
	return res
}

func (m JSONRawMessage) ToString() string {
	res := strings.ReplaceAll(string(m[:]), "\"", "")
	return res
}
