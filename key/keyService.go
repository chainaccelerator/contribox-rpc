package key

import (
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"
	"strings"

	// Used in conjunction with database/sql
	_ "github.com/go-sql-driver/mysql"
)

const keyTableName = "keys"

// KeyShareDb ...
// Create a new key
func KeyShareDb(_type string, xPubSList []string, key string, hash string, state string, dbConf persistance.DbConf) string {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO %v (`type`, `xPubS`, `key`, `hash`, `state`) VALUES ", dbConf.DbName+"."+keyTableName)
	for _, xPubS := range xPubSList {
		query += fmt.Sprintf("VALUES ('%v', '%v', '%v', '%v', '%v'), ", _type, xPubS, key, hash, state)
	}
	query = strings.TrimSuffix(query, ", ")
	fmt.Println(query)

	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	insert.Close()

	return state
}

// KeyShareGetDb ...
// Get an already existing key
func KeyShareGetDb(_type string, xPubS string, state string, dbConf persistance.DbConf) string {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer db.Close()

	// Si le même user a partagé plusieurs clés avec le même signataire ?
	query := fmt.Sprintf("SELECT k.key FROM %v k WHERE type = '%v' AND xPubS = '%v' AND state = '%v'", dbConf.DbName+"."+keyTableName, _type, xPubS, state)
	fmt.Println(query)

	var key string

	err = db.QueryRow(query).Scan(&key)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return key
}

// KeyShareConfirmDb ...
// Permet au signataire de confirmer qu'il a reçu la clé
func KeyShareConfirmDb(_type string, xPubS string, Hash string, state string, dbConf persistance.DbConf) string {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE %v SET state = ? WHERE type = ? AND xPubS = ? AND hash = ?", dbConf.DbName+"."+keyTableName)

	updateStatement, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	update, err := updateStatement.Exec(state, _type, xPubS, Hash)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println(update.RowsAffected())

	return state
}

// KeyShareConfirmGetDb ...
// Allow a user to check whether the key has been well received by the other side
func KeyShareConfirmGetDb(_type string, hash string, state string, dbConf persistance.DbConf) string {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT k.xPubS FROM %v k WHERE k.type = '%v' AND k.hash = '%v' AND k.state = '%v'", dbConf.DbName+"."+keyTableName, _type, hash, state)
	fmt.Println(query)

	var xPubS string

	err = db.QueryRow(query).Scan(&xPubS)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return xPubS

}
