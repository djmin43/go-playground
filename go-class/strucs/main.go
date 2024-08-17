package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	r := &Response{Page: 1, Words: []string{"up", "in", "out"}}

	j, _ := json.Marshal(r)

	fmt.Printf(string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response

	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

}
