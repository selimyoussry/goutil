package goutil

import "encoding/json"

// Typer is used for serialization of union types
type Typer struct {
	Type     string          `json:"type"`
	RawValue json.RawMessage `json:"value"`
}
