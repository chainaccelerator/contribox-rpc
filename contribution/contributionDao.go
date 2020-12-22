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

// GetContributionByXPub ...
func GetContributionByXPub(_type string, xPub string, state string, dbConf persistance.DbConf) persistance.Contribution {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return persistance.Contribution{}
	}
	defer db.Close()

	var contribution persistance.Contribution

	query := fmt.Sprintf(
		"SELECT c.* FROM %v.%v c INNER JOIN %v.%v cp on c.Id = cp.contributionId INNER JOIN %v.%v pxp on cp.proofId = pxp.proofId INNER JOIN %v.%v xp on pxp.xpubId = xp.Id WHERE xp.xpub = '%v'",
		dbConf.DbName,
		contributionsTableName,
		dbConf.DbName,
		contributionsAndProofsTableName,
		dbConf.DbName,
		proofsAndXPubTableName,
		dbConf.DbName,
		xPubsTableName,
		xPub,
	)

	err = db.QueryRow(query).Scan(&contribution)
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
		"SELECT p.* FROM %.v%v p INNER JOIN %v.%v cp on p.Id = cp.proofId WHERE cp.contributionId = %v",
		dbConf.DbName,
		proofsTableName,
		dbConf.DbName,
		contributionsAndProofsTableName,
		contributionID,
	)

	var proof persistance.Proof

	db.QueryRow(query).Scan(&proof)
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
		"SELECT t.* FROM %.v%v t INNER JOIN %v.%v ct on t.Id = ct.templateId WHERE ct.contributionId = %v",
		dbConf.DbName,
		templatesTableName,
		dbConf.DbName,
		contributionsAndTemplatesTableName,
		contributionID,
	)

	var template persistance.Template

	db.QueryRow(query).Scan(&template)
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
