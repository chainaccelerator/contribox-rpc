package persistance

import (
	"database/sql"
	"fmt"

	// Used in conjunction with database/sql
	_ "github.com/go-sql-driver/mysql"
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
		"SELECT p.Id, p.xPub, p.projectName, p.licenceSPDX, p.licenceSPDXChange, p.groupRoleName FROM %v p INNER JOIN %v tp ON p.Id = tp.proofId WHERE tp.templateId = %v",
		dbConf.DbName+"."+proofTableName,
		dbConf.DbName+"."+templateAndProofTableName,
		templateID,
	)

	fmt.Println(query)

	var proof Proof

	err = db.QueryRow(query).Scan(&proof.Id, &proof.XPub, &proof.ProjectName, &proof.LicenceSPDX, &proof.LicenceSPDXChange, &proof.GroupRoleName)
	if err != nil {
		fmt.Println(err.Error())
		return Proof{}
	}

	return proof
}
