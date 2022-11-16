package main

import (
	"fmt"
	"tiny-bbs/dao/mysql"
	"tiny-bbs/pkg/snowflake"
)

func testSnowflake() {
	_ = snowflake.Init("2022-01-01", 1)
	for i := 0; i < 4; i++ {
		fmt.Println(snowflake.GenID())
	}
}

func testMd5() {
	fmt.Println(mysql.Md5Psw("123a"))
}

func main() {
	//testSnowflake()
	testMd5()
}
