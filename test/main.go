package main

import (
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"strconv"
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

type MovieInfo struct {
	Link    string `json:"link"`
	Name    string `json:"name"`
	Score   string `json:"score"`
	Command string `json:"command"`
}

func spiderDouBanMovies() {
	moviesInfo := make([]MovieInfo, 0)
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (iPad; CPU OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.77 Mobile/15E148 Safari/604.1",
	}

	for idx := 0; idx < 10; idx++ {
		subLink := "https://movie.douban.com/top250?start=" + strconv.Itoa(idx*25)
		source, err := soup.Get(subLink)
		if err != nil {
			fmt.Println(idx)
			fmt.Println(err.Error())
			return
		}
		doc := soup.HTMLParse(source)
		infos := doc.FindAll("div", "class", "info")
		for _, node := range infos {
			hd := node.Find("div", "class", "hd")
			bd := node.Find("div", "class", "bd")
			link := hd.Find("a").Attrs()["href"]
			name := hd.Find("span", "class", "title").Text()
			score := bd.Find("span", "class", "rating_num").Text()
			commandRoot := bd.Find("span", "class", "inq")
			command := ""
			if commandRoot.Pointer != nil {
				command = commandRoot.Text()
			}
			moviesInfo = append(moviesInfo, MovieInfo{
				Link:    link,
				Name:    name,
				Score:   score,
				Command: command,
			})
		}
	}

	for idx, infos := range moviesInfo {
		fmt.Println(idx+1, infos)
	}

	// save to json file
	filePtr, err := os.Create("./test/DouBanMoviesTop250.json")
	if err != nil {
		fmt.Println("create json file failed...")
		return
	}
	defer filePtr.Close()
	err = json.NewEncoder(filePtr).Encode(moviesInfo)
	if err != nil {
		fmt.Println("save to json file failed", err.Error())
		return
	}
	fmt.Println("save to json file successfully!")
}
func main() {
	//testSnowflake()
	//testMd5()
	spiderDouBanMovies()
}
