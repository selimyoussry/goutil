package goutil

import (
	"fmt"
	"strings"
)

type CypherSet struct {
	Key          string `json:"key"`
	PropertyName string `json:"propertyName"`
}

type CypherSets struct {
	NodeName string                 `json:"nodeName"`
	Sets     []*CypherSet           `json:"sets"`
	Props    map[string]interface{} `json:"props"`
}

func (cs *CypherSets) MakeString() string {
	joined := []string{}
	for _, cypherSet := range cs.Sets {
		joined = append(joined, fmt.Sprintf(
			"%s.%s = {props}.%s", cs.NodeName, cypherSet.Key, cypherSet.PropertyName,
		))
	}

	return fmt.Sprintf("SET %s", strings.Join(joined, ", "))
}

// NewCypherSets instanciates
func NewCypherSets(nodeName string, props map[string]interface{}) *CypherSets {
	return &CypherSets{
		NodeName: nodeName,
		Sets:     []*CypherSet{},
		Props:    props,
	}
}

func (cs *CypherSets) AddString(key, propertyName string, value *string) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddBool(key, propertyName string, value *bool) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddInt(key, propertyName string, value *int) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddInt32(key, propertyName string, value *int32) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddInt64(key, propertyName string, value *int64) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddFloat32(key, propertyName string, value *float32) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddFloat64(key, propertyName string, value *float64) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) AddArray(key, propertyName string, value *[]interface{}) *CypherSets {
	if value == nil {
		return cs
	}
	return cs.ForceAdd(key, propertyName, value)
}

func (cs *CypherSets) ForceAdd(key, propertyName string, value interface{}) *CypherSets {
	cs.Sets = append(cs.Sets, &CypherSet{
		Key:          key,
		PropertyName: propertyName,
	})
	cs.Props[propertyName] = value
	return cs
}
