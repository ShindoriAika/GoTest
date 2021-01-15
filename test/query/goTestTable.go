package query // 独自のクエリパッケージ

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// マスタからSELECTしたデータをマッピングする構造体
type Item struct {
	Id    string `db:"ID"`
	Name  string `db:"NAME"`
	Price string `db:"PRICE"`
}

// データ登録関数
func InsertItem(name, price string, db *sql.DB) (id int64, err error) {

	// プリペアードステートメント
	stmt, err := db.Prepare("insert into item (name,price) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// クエリ実行
	result, err := stmt.Exec(name, price) //更新系SQL文の実行関数
	if err != nil {
		return 0, err
	}

	// オートインクリメントのIDを取得
	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// INSERTされたIDを返す
	return insertedId, nil
}

// 単一行データ取得関数
func SelectItemById(id int64, db *sql.DB) (iteminfo Item, err error) {

	// 構造体item型の変数itemを宣言
	var item Item

	// プリペアードステートメント
	stmt, err := db.Prepare("SELECT id,name,price from item where id = ?")
	if err != nil {
		return item, err
	}

	// クエリ実行
	rows, err := stmt.Query(id) //参照系SQL文の実行関数
	if err != nil {
		return item, err
	}
	defer rows.Close()

	// SELECTした結果を構造体にマップ
	rows.Next()
	err = rows.Scan(&item.Id, &item.Name, &item.Price)
	if err != nil {
		return item, err
	}

	// 取得データをマッピングしたitem構造体を返す
	return item, nil
}
