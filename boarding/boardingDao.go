package boarding

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"
)

const boardingsTable = "boardings"
const proofsTableName = "proofs"
const templatesTableName = "templates"
const templatesAndProofsTableName = "templatesAndProofs"
const templatesAndXPubsTableName = "templatesAndXPubs"
const xPubsTableName = "xPubs"

// GetTemplate ...
func GetTemplate(projectName string, licenceSPDX string, groupRoleName string, state string, dbConf persistance.DbConf) persistance.Template {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Template{}
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT t.Id, t.hash, t.projectName, t.licenceSPDX, t.groupRoleName, t.state, t.userRequirement, t.projectRequirement, "+
			"t.userUser, t.userBackup, t.userLock, t.userWitness, "+
			"t.projectOld, t.projectParent, t.projectBoard, t.projectMember, t.projectCosigner, t.projectWitness "+
			"FROM %v t WHERE t.projectName = '%v' AND t.licenceSPDX = '%v' AND t.groupRoleName = '%v' AND t.state = '%v'",
		dbConf.DbName+"."+templatesTableName,
		projectName,
		licenceSPDX,
		groupRoleName,
		state,
	)
	fmt.Println(query)

	var template persistance.Template

	// Si pas possible alors Scan(&template.Id, &template.Hash, etc...)
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

// InsertTemplate ...
func InsertTemplate(_type string, resource commons.Template, state string, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	return true
}

// UpdateProofBoardings ...
func UpdateProofBoardings(templateID int, onBoarding commons.Boarding, outBoarding commons.Boarding, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	// TODO
	return true
}

// GetTemplateByTypeXPubAndState ...
func GetTemplateByTypeXPubAndState(_type string, xPubS string, state string, dbConf persistance.DbConf) persistance.Template {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Template{}
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT t.Id, t.hash, t.projectName, t.licenceSPDX, t.groupRoleName, t.state, t.userRequirement, t.projectRequirement, "+
			"t.userUser, t.userBackup, t.userLock, t.userWitness, "+
			"t.projectOld, t.projectParent, t.projectBoard, t.projectMember, t.projectCosigner, t.projectWitness "+
			"FROM %v t INNER JOIN %v txp ON t.Id = txp.templateId INNER JOIN %v xp ON txp.xPubId = xp.Id "+
			"WHERE xp.xPub = '%v' AND t.state = '%v'",
		dbConf.DbName+"."+templatesTableName,
		dbConf.DbName+"."+templatesAndXPubsTableName,
		dbConf.DbName+"."+xPubsTableName,
		xPubS,
		state,
	)
	fmt.Println(query)

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
