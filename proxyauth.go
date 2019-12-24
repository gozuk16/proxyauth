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
			"-disable-gpu",                // 暫定的に必要とのこと
			"--ignore-certificate-errors", // 認証関係のエラー無視
			"--no-sandbox",
			"--disable-xss-auditor",
		})

	driver := agouti.ChromeDriver(options)
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Start error: %s\n", err)
		return
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewPage error: %s\n", err)
		return
	}

	page.Navigate("https://xxx.xxx.xxx.xxx:99999/proxy?foo=bar&hogehoge=http://neverssl.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Navigate error: %s\n", err)
		return
	}

	// Navigateは成功しているのにページがロードできてない？後の認証も失敗するので待ってみる
	page.Session().SetImplicitWait(10000)

	t, err := page.Title()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Title error: %s\n", err)
		return
	}
	fmt.Println("Title: " + t)

	fmt.Println("set userid & password")
	page.FindByID("user").Fill("xxxxxxxx")
	page.FirstByName("passwd").Fill("xxxxxxxxxxxxxxxx")

	fmt.Println("authentication")
	page.FindByID("submit").Click()

	time.Sleep(time.Second * 5)
}
