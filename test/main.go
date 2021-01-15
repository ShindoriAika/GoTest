package main

import (
	"database/sql"
	"fmt"
	"test/conf"  // 実装した設定パッケージの読み込み (srcから見たパス)
	"test/query" // 実装したクエリパッケージの読み込み

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 設定ファイルを読み込む
	confDB, err := conf.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 設定値から接続文字列を生成
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)

	// データベース接続
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	// deferで処理終了前に必ず接続をクローズする
	defer db.Close()

	// INSERTの実行
	id, err := query.InsertItem("りんご", "300", db)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("insert結果 id :【%d】\n", id)

	// SELECTの実行
	item, err := query.SelectItemById(id, db)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("select結果\n")
	fmt.Printf("[ID] %s\n", item.Id)
	fmt.Printf("[商品名] %s\n", item.Name)
	fmt.Printf("[価格] %s\n", item.Price)

}
