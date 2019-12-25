package main

import (
	"fmt"

	"github.com/xgdwq/gomod"
	gomodv3 "github.com/xgdwq/gomod/v3"
	gomodv4 "github.com/xgdwq/gomod/v4"
	gomodv5 "github.com/xgdwq/gomod/v5"
)

func main() {
	fmt.Println(gomod.Hello())
	fmt.Println(gomodv3.HelloV3())
	fmt.Println(gomodv4.HelloV4())
	fmt.Println(gomodv5.HelloV5())
}
