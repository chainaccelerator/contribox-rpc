package contribution

import (
	"bc_node_api/api3/persistance"
	"database/sql"
	"fmt"
)

const contributionsTableName = "contributions"
const contributionsAndProofsTableName = "contributionsandproofs"
const contributionsAndPubKeysTableNames = "contributionsandpubkeys"
const contributionsAndTemplatesTableName = "contributionsandtemplates"
const proofsTableName = "proofs"
const proofsAndXPubTableName = "proofsandxpubs"
const pubKeysTableNames = "pubkeys"
const templatesTableName = "templates"
const xPubsTableName = "xpubs"

// InsertContribution ...
func InsertContribution(contribution persistance.Contribution) {

}

// GetContributionByXPubAndState ...
func GetContributionByXPubAndState(_type string, xPub string, state string, dbConf persistance.DbConf) persistance.Contribution {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Contribution{}
	}
	defer db.Close()

	var contribution persistance.Contribution

	query := fmt.Sprintf(
		"SELECT c.Id, c.hash, c.xPub, c.tx1Id, c.tx0IdAmount, c.tx0IdIssueAsset, c.tx0IdSigA "+
			"FROM %v.%v c INNER JOIN %v.%v ct on c.Id = ct.contributionId INNER JOIN %v.%v t on ct.templateId = t.Id "+
			"WHERE c.xPub = '%v' AND t.state = '%v'",
		dbConf.DbName,
		contributionsTableName,
		dbConf.DbName,
		contributionsAndTemplatesTableName,
		dbConf.DbName,
		templatesTableName,
		xPub,
		state,
	)
	fmt.Println(query)

	err = db.QueryRow(query).Scan(
		&contribution.Id,
		&contribution.Hash,
		&contribution.XPub,
		&contribution.Tx1Id,
		&contribution.Tx0IdAmount,
		&contribution.Tx0IdIssueAsset,
		&contribution.Tx0IdSigA,
	)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Contribution{}
	}

	return contribution
}

// GetContributionByHashAndState ...
func GetContributionByHashAndState(_type string, hash string, state string, dbConf persistance.DbConf) persistance.Contribution {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Contribution{}
	}
	defer db.Close()

	var contribution persistance.Contribution

	query := fmt.Sprintf(
		"SELECT c.Id, c.hash, c.xPub, c.tx1Id, c.tx0IdAmount, c.tx0IdIssueAsset, c.tx0IdSigA "+
			"FROM %v.%v c INNER JOIN %v.%v ct on c.Id = ct.contributionId INNER JOIN %v.%v t on ct.templateId = t.Id "+
			"WHERE c.hash = '%v' AND t.state = '%v'",
		dbConf.DbName,
		contributionsTableName,
		dbConf.DbName,
		contributionsAndTemplatesTableName,
		dbConf.DbName,
		templatesTableName,
		hash,
		state,
	)
	fmt.Println(query)

	err = db.QueryRow(query).Scan(
		&contribution.Id,
		&contribution.Hash,
		&contribution.XPub,
		&contribution.Tx1Id,
		&contribution.Tx0IdAmount,
		&contribution.Tx0IdIssueAsset,
		&contribution.Tx0IdSigA,
	)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Contribution{}
	}

	return contribution
}

// GetProofByContributionID ...
func GetProofByContributionID(contributionID int, dbConf persistance.DbConf) persistance.Proof {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Proof{}
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT p.Id, p.xPub, p.projectName, p.licenceSPDX, p.licenceSPDXChange, p.groupRoleName "+
			"FROM %.v%v p INNER JOIN %v.%v cp on p.Id = cp.proofId WHERE cp.contributionId = %v",
		dbConf.DbName,
		proofsTableName,
		dbConf.DbName,
		contributionsAndProofsTableName,
		contributionID,
	)
	fmt.Println(query)

	var proof persistance.Proof

	db.QueryRow(query).Scan(
		&proof.Id,
		&proof.XPub,
		&proof.ProjectName,
		&proof.LicenceSPDX,
		&proof.LicenceSPDXChange,
		&proof.GroupRoleName,
	)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Proof{}
	}

	return proof
}

// GetTemplateByContributionID ...
func GetTemplateByContributionID(contributionID int, dbConf persistance.DbConf) persistance.Template {
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
			"FROM %.v%v t INNER JOIN %v.%v ct on t.Id = ct.templateId WHERE ct.contributionId = %v",
		dbConf.DbName,
		templatesTableName,
		dbConf.DbName,
		contributionsAndTemplatesTableName,
		contributionID,
	)
	fmt.Println(query)

	var template persistance.Template

	db.QueryRow(query).Scan(
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

// GetPubKeysByContributionID ...
func GetPubKeysByContributionID(contributionID int, dbConf persistance.DbConf) []persistance.PubKey {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT pk.* FROM %v.%v pk INNER JOIN %v.%v cpk on pk.Id = cpk.pubkeyId WHERE cpk.contributionId = %v",
		dbConf.DbName,
		pubKeysTableNames,
		dbConf.DbName,
		contributionsAndPubKeysTableNames,
		contributionID,
	)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var pubKeyPair []persistance.PubKey
	for results.Next() {
		var pubKey persistance.PubKey
		results.Scan(&pubKey.PubKey, &pubKey.Base58Encoded, &pubKey.PubKeyType)
		pubKeyPair = append(pubKeyPair, pubKey)
	}

	return pubKeyPair
}
