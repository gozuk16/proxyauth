package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"os"
	"time"
)

func main() {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"-disable-gpu", // 暫定的に必要とのこと
			"--no-sandbox",
			"--disable-xss-auditor",
		})

	driver := agouti.ChromeDriver(options)
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page.Navigate("https://xxx.xxx.xxx.xxx:99999/proxy?foo=bar&hogehoge=http://neverssl.com")
	fmt.Println(page.Title())

	fmt.Println("set userid & password")
	page.FindByID("user").Fill("xxxxxxxx")
	page.FirstByName("passwd").Fill("xxxxxxxxxxxxxxxx")

	fmt.Println("authentication")
	page.FindByID("submit").Click()

	time.Sleep(time.Second * 5)
}
