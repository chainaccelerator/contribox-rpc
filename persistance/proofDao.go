package persistance

import (
	"database/sql"
	"fmt"
)

const proofTableName = "proofs"
const templateAndProofTableName = "templatesAndProofs"

// GetProofByTemplateID ...
func GetProofByTemplateID(templateID int, dbConf DbConf) Proof {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return Proof{}
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT p.* FROM %v.%v p INNER JOIN %v.%v tp ON p.Id = tp.proofId WHERE tp.templateId = '%v",
		dbConf.DbName,
		proofTableName,
		dbConf.DbName,
		templateAndProofTableName,
		templateID,
	)

	var proof Proof

	err = db.QueryRow(query).Scan(&proof.Id, &proof.LicenseSPDX, &proof.LicenseSPDXChange, &proof.GroupRoleName)
	if err != nil {
		fmt.Println(err.Error())
		return Proof{}
	}

	return proof
}
