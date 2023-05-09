//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

// RawMessage is exported by common/json package.
type RawMessage = json.RawMessage

var (
	// Marshal is exported by common/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by common/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by common/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by common/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by common/json package.
	NewEncoder = json.NewEncoder
)
