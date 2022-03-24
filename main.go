package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup
var URL = flag.String("url", "", "输入url")
var Urlfile = flag.String("urlfile", "", "输入url的文件名")

func main() {
	sign.Banner()
	if *URL != "" {
		Run.Run(*URL)
	}
	if *Urlfile != "" {
		file, err := os.Open(*Urlfile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		wg.Add(3)

		for _, line := range lines {

			go Run.Run(line)

		}
	}
	if *URL == "" && *Urlfile == "" {
		fmt.Println("请使用-url 或 -urlfile 来指定目标")
		os.Exit(0)
	}
	fmt.Println("Done")
}
