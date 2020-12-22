package persistance

import (
	"database/sql"
	"fmt"
)

const hashesTableName = "hashes"
const proofsAndHashesTableName = "proofsandhashes"

// GetHashListByProofID ...
func GetHashListByProofID(proofID int, dbConf DbConf) []Hash {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT h.* FROM %v.%v h INNER JOIN %v.%v ph on h.Id = ph.hashId WHERE ph.proofId = %v",
		dbConf.DbName,
		hashesTableName,
		dbConf.DbName,
		proofsAndHashesTableName,
		proofID,
	)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var hashList []Hash
	for results.Next() {
		var hash Hash
		err = results.Scan(&hash.Id, &hash.Hash, &hash.HashType)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		hashList = append(hashList, hash)
	}

	return hashList
}
