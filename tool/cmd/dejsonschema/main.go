package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed openapi.json
var openapi []byte

func main() {
	m := make(map[string]any)
	err := json.Unmarshal(openapi, &m)
	if err != nil {
		panic(err)
	}
	bs, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

func IterateJSONTree(m map[string]any) map[string]any {
	references := make(map[string]any)

	return references
}
