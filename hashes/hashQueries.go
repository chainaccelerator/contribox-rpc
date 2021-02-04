package hashes

import "fmt"

const hashesTableName = "hashes"

// InsertHashQuery ...
func InsertHashQuery(left string, right string, hash string, data string, dbName string) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`left`, `right`, `hash`, `data`) VALUES ('%v', '%v', '%v', '%v')",
		dbName+"."+hashesTableName,
		left,
		right,
		hash,
		data,
	)
}

// GetHashRefQuery ...
func GetHashRefQuery(hashID int, dbName string) string {
	return fmt.Sprintf(
		"SELECT h.hash, h.left, h.right, h.data from %v h WHERE h.Id = %v",
		dbName+"."+hashesTableName,
		hashID,
	)
}
