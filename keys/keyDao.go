package keys

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/hashes"
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"

	// Used in conjunction with database/sql
	_ "github.com/go-sql-driver/mysql"
)

// InsertPubKeyEncrypted ...
func InsertPubKeyEncrypted(pubKeyEncrypted commons.PubKeyEncrypted, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := InsertPubKeyEncryptedQuery(dbConf.DbName, pubKeyEncrypted.Encrypted, pubKeyEncrypted.State)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	pubKeyEncryptedID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int(pubKeyEncryptedID)
}

// InsertHashAndPubKeyEncrypted ...
func InsertHashAndPubKeyEncrypted(hashID int, pubKeyEncryptedID int, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	query := InsertHashAndPubKeyEncryptedQuery(dbConf.DbName, hashID, pubKeyEncryptedID)
	fmt.Println(query)

	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	insert.Close()

	return true
}

// InsertKeyEncrypted ...
func InsertKeyEncrypted(keyEncrypted commons.KeyEncrypted, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := InsertKeyEncryptedQuery(dbConf.DbName, keyEncrypted.Encrypted, keyEncrypted.State)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	keyEncryptedID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int(keyEncryptedID)
}

// InsertHashAndKeyEncrypted ...
func InsertHashAndKeyEncrypted(hashID int, keyEncryptedID int, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	query := InsertHashAndKeyEncryptedQuery(dbConf.DbName, hashID, keyEncryptedID)
	fmt.Println(query)

	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	insert.Close()

	return true
}

// InsertKey ...
func InsertKey(key commons.Key, keyEncryptedID int, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := InsertKeyQuery(dbConf.DbName, key.Data, key.State)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	keyID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int(keyID)
}

// InsertHashAndKey ...
func InsertHashAndKey(hashID int, keyID int, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	query := InsertHashAndKeyQuery(dbConf.DbName, hashID, keyID)
	fmt.Println(query)

	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	insert.Close()

	return true
}

// InsertKeyShareProcess ...
func InsertKeyShareProcess(fromID int, toID int, keyID int, processHash hashes.Hash, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := InsertKeyShareProcessQuery(dbConf.DbName, fromID, toID, keyID, "todo")
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	keyShareProcessID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int(keyShareProcessID)
}

// InsertHashAndKeyShareProcess ...
func InsertHashAndKeyShareProcess(hashID int, processID int, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	query := InsertHashAndKeyShareProcessQuery(dbConf.DbName, hashID, processID)
	fmt.Println(query)

	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	insert.Close()

	return true
}

// GetProcessData ...
func GetProcessData(pubKeyEncryptedHash string, dbConf persistance.DbConf) (int, int, int) {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, 0
	}
	defer db.Close()

	query := GetProcessDataQuery(pubKeyEncryptedHash, dbConf.DbName)
	fmt.Println(query)

	var hashID, fromID, keyID int
	err = db.QueryRow(query).Scan(&hashID, &fromID, keyID)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, 0
	}

	return hashID, fromID, keyID
}

// GetPubKeyEncryptedByID ...
func GetPubKeyEncryptedByID(pubKeyEncryptedID int, dbConf persistance.DbConf) (*commons.PubKeyEncrypted, int) {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}
	defer db.Close()

	query := GetPubKeyEncryptedByIDQuery(pubKeyEncryptedID, dbConf.DbName)
	fmt.Println(query)

	var pubKeyEncrypted commons.PubKeyEncrypted
	var pubKeyEncryptedHashID int

	err = db.QueryRow(query).Scan(&pubKeyEncrypted.Encrypted, &pubKeyEncrypted.State, &pubKeyEncryptedHashID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}

	return &pubKeyEncrypted, pubKeyEncryptedHashID
}

// GetKeyEncryptedByID ...
func GetKeyEncryptedByID(keyEncryptedID int, dbConf persistance.DbConf) (*commons.KeyEncrypted, int, int) {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, 0
	}
	defer db.Close()

	query := GetKeyEncryptedByIDQuery(keyEncryptedID, dbConf.DbName)
	fmt.Println(query)

	var keyEncrypted commons.KeyEncrypted
	var keyEncryptedHashID int
	var keyID int

	err = db.QueryRow(query).Scan(&keyEncrypted.Encrypted, &keyEncrypted.State, &keyEncryptedHashID, &keyID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, 0
	}

	return &keyEncrypted, keyEncryptedHashID, keyID
}
