package main

import "encoding/json"

const (
	users   = "users"
	content = "contents"
)

var toDeleteResources = make([]map[string]any, 0)

func toDelete(resource string, body []byte) {
	var bd = make(map[string]interface{})
	e := json.Unmarshal(body, &bd)
	if e != nil {
		return
	}
	toDeleteResources = append(toDeleteResources, map[string]any{resource: bd["ID"]})
}
