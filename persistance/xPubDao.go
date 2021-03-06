package persistance

import (
	"database/sql"
	"fmt"

	// Used in conjunction with database/sql
	_ "github.com/go-sql-driver/mysql"
)

const proofsAndXPubsTableName = "proofsandxpubs"
const templatesAndXPubsTableName = "templatesandxpubs"
const xPubsTableName = "xpubs"

// GetXPubListByProofID ...
func GetXPubListByProofID(proofID int, dbConf DbConf) []XPub {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT xp.* FROM %v xp INNER JOIN %v pxp on xp.Id = pxp.xpubId WHERE pxp.proofId = %v",
		dbConf.DbName+"."+xPubsTableName,
		dbConf.DbName+"."+proofsAndXPubsTableName,
		proofID,
	)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var xPubList []XPub
	for results.Next() {
		var xPub XPub
		err = results.Scan(&xPub.Id, &xPub.XPub, &xPub.XPubType)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		xPubList = append(xPubList, xPub)
	}

	return xPubList
}

// GetXPubListByTemplateID ...
func GetXPubListByTemplateID(templateID int, dbConf DbConf) []XPub {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT xp.Id, xp.xPub, xp.xPubType FROM %v xp INNER JOIN %v txp on xp.Id = txp.xpubId WHERE txp.templateId = %v",
		dbConf.DbName+"."+xPubsTableName,
		dbConf.DbName+"."+templatesAndXPubsTableName,
		templateID,
	)
	fmt.Println(query)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var xPubList []XPub
	for results.Next() {
		var xPub XPub
		err = results.Scan(&xPub.Id, &xPub.XPub, &xPub.XPubType)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		xPubList = append(xPubList, xPub)
	}

	return xPubList
}
