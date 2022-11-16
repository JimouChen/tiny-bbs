package main

import (
	"fmt"
	"tiny-bbs/pkg/snowflake"
)

func testSnowflake() {
	_ = snowflake.Init("2022-01-01", 1)
	for i := 0; i < 4; i++ {
		fmt.Println(snowflake.GenID())
	}
}

func main() {
	testSnowflake()
}
