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
		"SELECT t.Id, t.tag FROM %v t INNER JOIN %v pt on t.Id = pt.tagId WHERE pt.proofId = %v",
		dbConf.DbName+"."+tagsTableName,
		dbConf.DbName+"."+proofsAndTagsTableName,
		proofID,
	)
	fmt.Println(query)

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
