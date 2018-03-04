package main

import (
	"fmt"

	"github.com/ghostec/sccg-server/api"
)

func main() {
	a, err := api.NewApp("localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	a.ListenAndServe()
}
