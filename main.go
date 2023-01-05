package main

import (
	"fmt"
	"github.com/engchina/golang-oracle-demo/router"
)

func main() {
	r := router.Router()
	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Errorf("error in start gin server %w", err))
	}
}
