package persistance

import (
	"database/sql"
	"fmt"
)

const keyValsTableName = "keyvals"
const proofsAndKeyValsTableName = "proofsandkeyvals"

// GetKeyValListByProofID ...
func GetKeyValListByProofID(proofID int, dbConf DbConf) []KeyVal {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT kv.Id, kv.key, kv.val, kv.keyvalType FROM %v.%v kv INNER JOIN %v.%v pkv on kv.Id = pkv.keyvalId WHERE pkv.proofId = %v",
		dbConf.DbName,
		keyValsTableName,
		dbConf.DbName,
		proofsAndKeyValsTableName,
		proofID,
	)
	fmt.Println(query)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var keyValList []KeyVal
	for results.Next() {
		var keyVal KeyVal
		err = results.Scan(&keyVal.Id, &keyVal.Key, &keyVal.Val, &keyVal.KeyValType)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		keyValList = append(keyValList, keyVal)
	}

	return keyValList
}
