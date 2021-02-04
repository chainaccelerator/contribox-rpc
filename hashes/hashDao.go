package hashes

import (
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"
)

// InsertLeaf ...
func InsertLeaf(h Hash, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := InsertHashQuery(h.Left.(string), h.Right.(string), h.Hash, h.Data, dbConf.DbName)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	leafID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int(leafID)
}

// InsertNode ...
func InsertNode(left string, right string, hash string, data string, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := InsertHashQuery(left, right, hash, data, dbConf.DbName)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	nodeID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int(nodeID)
}

// GetHashRef ...
func GetHashRef(hashID int, dbConf persistance.DbConf) *HashRef {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := GetHashRefQuery(hashID, dbConf.DbName)
	fmt.Println(query)

	var hashRef HashRef

	err = db.QueryRow(query).Scan(&hashRef.Hash, &hashRef.Left, &hashRef.Right, &hashRef.Data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &hashRef
}
