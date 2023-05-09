//go:build jsoniter
// +build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

// RawMessage is exported by common/json package.
type RawMessage = jsoniter.RawMessage

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by common/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by common/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by commonn/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by common/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by common/json package.
	NewEncoder = json.NewEncoder
)
