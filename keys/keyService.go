package keys

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/hashes"
	"bc_node_api/api3/persistance"

	// Used in conjunction with database/sql
	_ "github.com/go-sql-driver/mysql"
)

// KeyShareDb ...
// Create a new key share process
func KeyShareDb(pubKeyEncryptedList []commons.PubKeyEncrypted, pubKeySEncryptedList []commons.PubKeyEncrypted, keyEncryptedList []commons.KeyEncrypted, metadata commons.Metadata, dbConf persistance.DbConf) bool {

	// Insert the PubKeyEncrypted of the individual who wants to share keys
	from := pubKeyEncryptedList[0]
	pubKeyEncryptedID := InsertPubKeyEncrypted(from, dbConf)
	if pubKeyEncryptedID == 0 {
		return false
	}
	pubKeyEncryptedHashID := hashes.SaveHash(from.Hash, dbConf)
	if pubKeyEncryptedHashID == 0 {
		return false
	}
	hashAndPubKeyEncryptedInserted := InsertHashAndPubKeyEncrypted(pubKeyEncryptedHashID, pubKeyEncryptedID, dbConf)
	if !hashAndPubKeyEncryptedInserted {
		return false
	}

	// Insert the PubKeyEncrypted of the individuals with whom the keys are shared with, and these keys
	for i := range keyEncryptedList {
		pubKeySEncryptedID := InsertPubKeyEncrypted(pubKeySEncryptedList[i], dbConf)
		if pubKeySEncryptedID == 0 {
			return false
		}
		pubKeySEncryptedHashID := hashes.SaveHash(pubKeySEncryptedList[i].Hash, dbConf)
		if pubKeySEncryptedHashID == 0 {
			return false
		}
		hashAndPubKeySEncryptedInserted := InsertHashAndPubKeyEncrypted(pubKeySEncryptedHashID, pubKeySEncryptedID, dbConf)
		if !hashAndPubKeySEncryptedInserted {
			return false
		}

		keyEncryptedID := InsertKeyEncrypted(keyEncryptedList[i], dbConf)
		if keyEncryptedID == 0 {
			return false
		}
		keyEncryptedHashID := hashes.SaveHash(keyEncryptedList[i].Hash, dbConf)
		if keyEncryptedHashID == 0 {
			return false
		}
		hashAndKeyEncryptedInserted := InsertHashAndKeyEncrypted(keyEncryptedHashID, keyEncryptedID, dbConf)
		if !hashAndKeyEncryptedInserted {
			return false
		}

		keyID := InsertKey(keyEncryptedList[i].Key, keyEncryptedID, dbConf)
		if keyID == 0 {
			return false
		}
		keyHashID := hashes.SaveHash(keyEncryptedList[i].Key.Hash, dbConf)
		if keyHashID == 0 {
			return false
		}
		hashAndKeyInserted := InsertHashAndKey(keyHashID, keyID, dbConf)
		if !hashAndKeyInserted {
			return false
		}

		processID := InsertKeyShareProcess(pubKeyEncryptedID, pubKeySEncryptedID, keyEncryptedID, metadata.Hash, dbConf)
		if processID == 0 {
			return false
		}
		processHashID := hashes.SaveHash(metadata.Hash, dbConf)
		if processHashID == 0 {
			return false
		}
		hashAndKeyShareProcessInserted := InsertHashAndKeyShareProcess(processHashID, processID, dbConf)
		if !hashAndKeyShareProcessInserted {
			return false
		}
	}

	return true
}

// KeyShareGetDb ...
// Get an already existing key
func KeyShareGetDb(pubKeyEncryptedList []commons.PubKeyEncrypted, dbConf persistance.DbConf) (*hashes.Hash, []commons.PubKeyEncrypted, []commons.KeyEncrypted) {
	to := pubKeyEncryptedList[0]
	processHashID, _, _ := GetProcessData(to.Hash.Hash, dbConf)
	if processHashID == 0 {
		return nil, nil, nil
	}
	processHash := hashes.GetHash(processHashID, dbConf)
	if &processHash == nil {
		return nil, nil, nil
	}

	var pubKeyEncryptedFromList = make([]commons.PubKeyEncrypted, len(pubKeyEncryptedList))
	var keyEncryptedList = make([]commons.KeyEncrypted, len(pubKeyEncryptedList))
	for i := range pubKeyEncryptedList {
		_, fromID, keyID := GetProcessData(pubKeyEncryptedList[i].Hash.Hash, dbConf)
		if fromID == 0 || keyID == 0 {
			return nil, nil, nil
		}

		pubKeyEncrypted, pubKeyEncryptedHashID := GetPubKeyEncryptedByID(fromID, dbConf)
		if pubKeyEncrypted == nil || pubKeyEncryptedHashID == 0 {
			return nil, nil, nil
		}
		pubKeyEncryptedHash := hashes.GetHash(pubKeyEncryptedHashID, dbConf)
		if &pubKeyEncryptedHash == nil {
			return nil, nil, nil
		}
		pubKeyEncrypted.Hash = pubKeyEncryptedHash
		pubKeyEncryptedList[i] = *pubKeyEncrypted

	}

	return &processHash, pubKeyEncryptedFromList, keyEncryptedList
}

// KeyShareConfirmDb ...
// Permet au signataire de confirmer qu'il a reçu la clé
func KeyShareConfirmDb(_type string, xPubS string, Hash string, state string, dbConf persistance.DbConf) string {
	// db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return ""
	// }
	// defer db.Close()

	// query := fmt.Sprintf("UPDATE %v SET state = ? WHERE type = ? AND xPubS = ? AND hash = ?", dbConf.DbName+"."+keysTableName)

	// updateStatement, err := db.Prepare(query)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return ""
	// }

	// update, err := updateStatement.Exec(state, _type, xPubS, Hash)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return ""
	// }
	// fmt.Println(update.RowsAffected())

	// return state
	return ""
}

// KeyShareConfirmGetDb ...
// Allow a user to check whether the key has been well received by the other side
func KeyShareConfirmGetDb(_type string, hash string, state string, dbConf persistance.DbConf) string {
	// db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return ""
	// }
	// defer db.Close()

	// query := fmt.Sprintf("SELECT k.xPubS FROM %v k WHERE k.type = '%v' AND k.hash = '%v' AND k.state = '%v'", dbConf.DbName+"."+keysTableName, _type, hash, state)
	// fmt.Println(query)

	// var xPubS string

	// err = db.QueryRow(query).Scan(&xPubS)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return ""
	// }

	// return xPubS
	return ""
}
