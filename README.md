# proxyauth
Web認証のプロキシを通す。

Web認証のURLを書き換えます。(httpサイトにリダイレクトするために `neverssl.com` を指定しています。なくてもいいかも)

```go
page.Navigate("https://xxx.xxx.xxx.xxx:99999/proxy?foo=bar&hogehoge=http://neverssl.com")
```

ユーザIDを書き換えます。

```go
page.FindByID("user").Fill("xxxxxxxx")
```

パスワードを書き換えます。
```go
page.FirstByName("passwd").Fill("xxxxxxxxxxxxxxxx")
```

パスワードは暗号化したものを外部から渡されて、復号してセットするほうがいいと思いますけど、とりあえず・・・  
パスワード変えたり、ユーザID変えたりすることも当然ありますからね。

## build
MacでLinux用のクロスコンパイルを行うために、以下のようにしています。

```bash
$ mkdir linux-amd64
$ GOOS=linux GOARCH=amd64 go build -o linux-amd64/proxyauth ./proxyauth.go
```

## run
ChromeDriverを入れておく必要があります。

Mac
```
$ brew cask install chromedriver
```

Linux
```
$ sudo apt install chromium-chromedriver
```
