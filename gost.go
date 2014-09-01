package gost

import (
	"encoding/json"
	"fmt"
)

type Gost struct {
	values map[string][]byte
}

type GostKind string

const (
	DefaultKind GostKind = "DEFAULT"
)

func GetGost(kind ...GostKind) (*Gost, error) {
	if len(kind) == 0 {
		kind = []GostKind{DefaultKind}
	}
	if len(kind) > 1 {
		return nil, fmt.Errorf("too many kinds")
	}

	switch kind[0] {
	case DefaultKind:
		return &Gost{values: map[string][]byte{}}, nil
	default:
		return nil, fmt.Errorf("not a GostKind")
	}
}

func (g *Gost) Put(value interface{}, key string) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	g.values[key] = val
	return nil
}

func (g *Gost) Get(resp interface{}, key string) error {
	val, ok := g.values[key]
	if !ok {
		return fmt.Errorf("no value for key %q", key)
	}
	return json.Unmarshal(val, &resp)
}
