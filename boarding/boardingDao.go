package boarding

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"

	// Used in conjunction with database/sql
	_ "github.com/go-sql-driver/mysql"
)

const boardingsTable = "boardings"
const proofsTableName = "proofs"
const templatesTableName = "templates"
const templatesAndProofsTableName = "templatesAndProofs"
const templatesAndXPubsTableName = "templatesAndXPubs"
const xPubsTableName = "xPubs"

// InsertBoarding ...
// Insert a boarding request. The state will be "todo"
func InsertBoarding(boarding commons.Boarding, state string, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	query := fmt.Sprintf(
		"INSERT INTO %v (`in`, `state`) VALUES(%v, '%v')",
		dbConf.DbName+"."+boardingsTable,
		boarding.In,
		state,
	)
	fmt.Println(query)

	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	insert.Close()

	return true
}

// GetTemplate ...
// Get signature template
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
// Insert a signature template associated with either a boarding request or a contribution request
func InsertTemplate(_type string, template persistance.Template, state string, dbConf persistance.DbConf) int {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()

	query := fmt.Sprintf(
		"INSERT INTO %v "+
			"(`hash`, `projectName`, `licenceSPDX`, `state`, `userRequirement`, `projectRequirement`, "+
			"`userUser`, `userBackup`, `userLock`, `userWitness`, "+
			"`projectOld`, `projectParent`, `projectBoard`, `projectMember`, `projectCosigner`, `projectWitness`) "+
			"VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)",
		dbConf.DbName+"."+templatesTableName,
		template.Hash,
		template.ProjectName,
		template.LicenceSPDX,
		template.State,
		template.UserRequirement,
		template.ProjectRequirement,
		template.UserUser,
		template.UserBackup,
		template.UserLock,
		template.UserWitness,
		template.ProjectOld,
		template.ProjectParent,
		template.ProjectBoard,
		template.ProjectMember,
		template.ProjectCosigner,
		template.ProjectWitness,
	)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	lid, err := insert.LastInsertId()

	return int(lid)
}

// Insert one XPub associated with a signature template, specifying its type
func insertTemplateXPub(templateID int, xPub string, xPubType string, dbConf persistance.DbConf) bool {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	query := fmt.Sprintf(
		"INSERT INTO %v (`xPub`, `xPubType`) VALUES ('%v', '%v')",
		dbConf.DbName+"."+xPubsTableName,
		xPub,
		xPubType,
	)
	fmt.Println(query)

	insert, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	xPubID, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	jointQuery := fmt.Sprintf(
		"INSERT INTO %v (`templateId`, `xpubId`) VALUES (%v, %v)",
		dbConf.DbName+"."+templatesAndXPubsTableName,
		templateID,
		xPubID,
	)
	fmt.Println(jointQuery)

	jointInsert, err := db.Query(jointQuery)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	jointInsert.Close()

	return true
}

// InsertTemplateXPubs ...
func InsertTemplateXPubs(
	templateID int,
	userUserXPubList []commons.XPub,
	userBackupXPubList []commons.XPub,
	userLockXPubList []commons.XPub,
	userWitnessXPubList []commons.XPub,
	projectOldXPubList []commons.XPub,
	projectParentXPubList []commons.XPub,
	projectBoardXPubList []commons.XPub,
	projectMemberXPubList []commons.XPub,
	projectCosignerXPubList []commons.XPub,
	projectWitnessXPubList []commons.XPub,
	dbConf persistance.DbConf,
) bool {

	for _, xPub := range userUserXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "USERUSER", dbConf)
	}
	for _, xPub := range userBackupXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "USERBACKUP", dbConf)
	}
	for _, xPub := range userLockXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "USERLOCK", dbConf)
	}
	for _, xPub := range userWitnessXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "USERWITNESS", dbConf)
	}
	for _, xPub := range projectOldXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "PROJECTOLD", dbConf)
	}
	for _, xPub := range projectParentXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "PROJECTPARENT", dbConf)
	}
	for _, xPub := range projectBoardXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "PROJECTBOARD", dbConf)
	}
	for _, xPub := range projectMemberXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "PROJECTMEMBER", dbConf)
	}
	for _, xPub := range projectCosignerXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "PROJECTCOSIGNER", dbConf)
	}
	for _, xPub := range projectWitnessXPubList {
		insertTemplateXPub(templateID, xPub.XPub, "PROJECTWITNESS", dbConf)
	}

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
