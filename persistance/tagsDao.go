package persistance

import (
	"database/sql"
	"fmt"
)

const tagsTableName = "tags"
const proofsAndTagsTableName = "proofsandtags"

// GetTagListByProofID ...
func GetTagListByProofID(proofID int, dbConf DbConf) []Tag {
	db, err := sql.Open("mysql", dbConf.DbURL+dbConf.DbName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT t.* FROM %v.%v t INNER JOIN %v.%v pt on t.Id = pt.tagId WHERE pt.proofId = %v",
		dbConf.DbName,
		tagsTableName,
		dbConf.DbName,
		proofsAndTagsTableName,
		proofID,
	)

	results, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var tagList []Tag
	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.Id, &tag.Tag)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		tagList = append(tagList, tag)
	}

	return tagList
}
