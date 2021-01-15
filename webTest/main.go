package main

import (
	"fmt"
	"io"
	"net/http"
)

type fugaHandler string

// String に ServeHTTP 関数を追加
func (s fugaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// HTMLテキストをhttp.ResponseWriterへ書き込み
	fmt.Fprint(w, s)
}

func hogeHandler(w http.ResponseWriter, req *http.Request) {
	// HTMLテキストをhttp.ResponseWriterへ書き込む
	io.WriteString(w, `
		<!DOCTYPE html>
		<html lang="ja">
			<head>
				<meta charset="UTF-8">
				<title>Go webアプリ test</title>
			</head>
			<body>
				<h1>Hello, World!</h1>
				<p>http.HandleFunc関数でハンドラを定義した場合</p>
			</body>
		</html>
	`)
}

func main() {
	// "/hoge"へのリクエストを関数で処理する
	http.HandleFunc("/hoge", hogeHandler)
	// DuckTyping的に、ServeHTTP関数があれば良い.
	http.Handle("/fuga", fugaHandler(`
		<!DOCTYPE html>
		<html lang="ja">
			<head>
				<meta charset="UTF-8">
				<title>Go webアプリ test</title>
			</head>
			<body>
				<h1>Hello, World!</h1>
				<p>http.Handle関数でハンドラを定義した場合</p>
			</body>
		</html>
  	`))
	// ルートへのリクエストを"html"ディレクトリ内のHTMLファイルで処理する
	http.Handle("/", http.FileServer(http.Dir("html")))

	// localhost:8080でサーバー処理開始
	http.ListenAndServe(":8080", nil)
}
