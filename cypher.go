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

func CypherSetsMakeString(css ...*CypherSets) string {
	joined := []string{}

	for _, cs := range css {

		for _, cypherSet := range cs.Sets {
			joined = append(joined, fmt.Sprintf(
				"%s.%s = {props}.%s", cs.NodeName, cypherSet.Key, cypherSet.PropertyName,
			))
		}

	}

	if len(joined) == 0 {
		return ""
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

func (cs *CypherSets) AddIfNotNull(key, propertyName string, value interface{}) *CypherSets {
	isNil, err := NilCheck(value)
	if isNil || (err != nil) {
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
