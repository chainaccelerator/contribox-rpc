package boarding

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"
)

const boardingTable = "boardings"
const proofTableName = "proofs"
const templateTableName = "templates"
const templateAndProofTableName = "templatesAndProofs"
const templateAndXPubTableName = "templatesAndXPubs"
const xPubTableName = "xPubs"

// GetTemplate ...
func GetTemplate(projectName string, licenceSPDX string, groupRoleName string, state string, dbConf commons.DbConf) persistance.Template {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Template{}
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT t.* FROM %v.%v t WHERE t.projectName = '%v'AND t.licenseSPDX = '%v' AND t.groupRoleName = '%v' AND t.state = '%v'",
		dbConf.DbName,
		templateTableName,
		projectName,
		licenceSPDX,
		groupRoleName,
		state,
	)
	fmt.Println(query)

	var template persistance.Template

	// Si pas possible alors Scan(&template.Id, &template.Hash, etc...)
	err = db.QueryRow(query).Scan(template)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Template{}
	}

	return template
}

// UpdateProofBoardings ...
func UpdateProofBoardings(templateID int, onBoarding commons.Boarding, outBoarding commons.Boarding, dbConf commons.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	// TODO
	return true
}

// GetXPubListByTemplateID ...
func GetXPubListByTemplateID(templateID int, dbConf commons.DbConf) []persistance.XPub {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT xp.* FROM %v.%v xp INNER JOIN %v.%v txp ON xp.id = txp.xPubId WHERE txp.templateId = %v",
		dbConf.DbName,
		xPubTableName,
		dbConf.DbName,
		templateAndXPubTableName,
		templateID,
	)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(query)

	var xPubList []persistance.XPub
	for results.Next() {
		var xPub persistance.XPub
		err = results.Scan(&xPub.Id, &xPub.XPub, &xPub.XPubType)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		xPubList = append(xPubList, xPub)
	}

	return xPubList
}

// GetProofByID ...
func GetProofByID(templateID int, dbConf commons.DbConf) persistance.Proof {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Proof{}
	}

	query := fmt.Sprintf(
		"SELECT p.* FROM %v.%v p INNER JOIN %v.%v tp ON p.Id = tp.proofId WHERE tp.templateId = '%v",
		dbConf.DbName,
		proofTableName,
		dbConf.DbName,
		templateAndProofTableName,
		templateID,
	)

	var proof persistance.Proof

	err = db.QueryRow(query).Scan(&proof.Id, &proof.LicenseSPDX, &proof.LicenseSPDXChange, &proof.GroupRoleName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Proof{}
	}

	return proof
}

// GetTemplateByTypeXPubAndState ...
func GetTemplateByTypeXPubAndState(_type string, xPubS string, state string, dbConf commons.DbConf) persistance.Template {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Template{}
	}

	query := fmt.Sprintf(
		"SELECT t.* FROM %v.%v t INNER JOIN %v.%v txp ON t.Id = txp.templateId INNER JOIN %v.%v xp ON txp.xPubId = xp.Id WHERE xp.xPub = '%v' AND t.state = '%v'",
		dbConf.DbName,
		templateTableName,
		dbConf.DbName,
		templateAndXPubTableName,
		dbConf.DbName,
		xPubTableName,
		xPubS,
		state,
	)

	var template persistance.Template

	err = db.QueryRow(query).Scan(
		&template.Id,
		&template.Hash,
		&template.ProjectName,
		&template.LicenceSPDX,
		&template.GroupRoleName,
		&template.State,
		&template.UserRequirement,
		&template.ProjectRequirement,
		&template.UserUser,
		&template.UserBackup,
		&template.UserLock,
		&template.UserWitness,
		&template.ProjectOld,
		&template.ProjectParent,
		&template.ProjectBoard,
		&template.ProjectMember,
		&template.ProjectCosigner,
		&template.ProjectWitness,
	)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Template{}
	}

	return template
}
