package keys

import "fmt"

// Main tables
const keyShareProcessesTableName = "keyshareprocesses"
const hashesTableName = "hashes"
const keysTableName = "keys"
const keysEncryptedTableName = "keysencrypted"
const pubKeysEncryptedTableName = "pubkeysencrypted"

// Join tables
const hashesAndKeysTableName = "hashesandkeys"
const hashesAndKeysEncryptedTableName = "hashesandkeysencrypted"
const hashesAndPubKeysEncryptedTableName = "hashesandpubkeysencrypted"
const hashesAndKeyShareProcessesTableName = "hasheskeyshareprocesses"
const keysEncryptedAndKeysTableName = "keysencryptedandkeys"

// InsertPubKeyEncryptedQuery ...
func InsertPubKeyEncryptedQuery(dbName string, encrypted string, state string) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`encrypted`, `state`) VALUES ('%v', '%v')",
		dbName+"."+pubKeysEncryptedTableName,
		encrypted,
		state,
	)
}

// InsertHashQuery ...
func InsertHashQuery(dbName string, hash string, left string, right string, data string) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`hash`, `left`, `right`, `data`) VALUES ('%v', '%v', '%v', '%v')",
		dbName+"."+hashesTableName,
		hash,
		left,
		right,
		data,
	)
}

// InsertHashAndPubKeyEncryptedQuery ...
func InsertHashAndPubKeyEncryptedQuery(dbName string, hashID int, pubKeyEncryptedID int) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`hashId`, `pubKeyEncryptedId`) VALUES (%v, %v)",
		dbName+"."+hashesAndPubKeysEncryptedTableName,
		hashID,
		pubKeyEncryptedID,
	)
}

// InsertKeyEncryptedQuery ...
func InsertKeyEncryptedQuery(dbName string, encrypted string, state string) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`encrypted`, `state`) VALUES ('%v', '%v')",
		dbName+"."+keysEncryptedTableName,
		encrypted,
		state,
	)
}

// InsertHashAndKeyEncryptedQuery ...
func InsertHashAndKeyEncryptedQuery(dbName string, hashID int, keyEncryptedID int) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`hashId`, `keyEncryptedId`) VALUES (%v, %v)",
		dbName+"."+hashesAndKeysEncryptedTableName,
		hashID,
		keyEncryptedID,
	)
}

// InsertKeyQuery ...
func InsertKeyQuery(dbName string, data string, state string) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`data`, `state`) VALUES (%v, %v)",
		dbName+"."+keysEncryptedAndKeysTableName,
		data,
		state,
	)
}

// InsertHashAndKeyQuery ...
func InsertHashAndKeyQuery(dbName string, hashID int, keyID int) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`hashId`, `keyId`) VALUES (%v, %v)",
		dbName+"."+hashesAndKeysTableName,
		hashID,
		keyID,
	)
}

// InsertKeyEncryptedAndKeyQuery ...
func InsertKeyEncryptedAndKeyQuery(dbName string, keyEncryptedID int, keyID int) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`keyEncryptedId`, `keyId`) VALUES (%v, %v)",
		dbName+"."+keysEncryptedAndKeysTableName,
		keyEncryptedID,
		keyID,
	)
}

// InsertKeyShareProcessQuery ...
func InsertKeyShareProcessQuery(dbName string, fromID int, toID int, keyID int, processState string) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`fromId`, `toId`, `keyId`, `processState`) VALUES (%v, %v, %v, '%v')",
		dbName+"."+keyShareProcessesTableName,
		fromID,
		toID,
		keyID,
		processState,
	)
}

// InsertHashAndKeyShareProcessQuery ...
func InsertHashAndKeyShareProcessQuery(dbName string, hashID int, processID int) string {
	return fmt.Sprintf(
		"INSERT INTO %v (`hashId`, `keyShareProcessId`) VALUES (%v, %v)",
		dbName+"."+hashesAndKeyShareProcessesTableName,
		hashID,
		processID,
	)
}

// GetProcessDataQuery ...
func GetProcessDataQuery(pubKeyEncryptedHash string, dbName string) string {
	return fmt.Sprintf(
		"SELECT hksp.hashId FROM %v hksp, ksp.fromId, ksp.keyId "+
			"INNER JOIN %v ksp on hksp.keyShareProcessId = ksp.Id "+
			"INNER JOIN %v pke on ksp.toId = pke.Id "+
			"INNER JOIN %v hpke on pke.Id = hpke.publicKeyEncrypted "+
			"INNER JOIN %v h on hpke.hashId = h.Id "+
			"WHERE h.hash = '%v'",
		dbName+"."+hashesAndKeyShareProcessesTableName,
		dbName+"."+keyShareProcessesTableName,
		dbName+"."+pubKeysEncryptedTableName,
		dbName+"."+hashesAndPubKeysEncryptedTableName,
		dbName+"."+hashesTableName,
		pubKeyEncryptedHash,
	)
}

// GetPubKeyEncryptedByIDQuery ...
func GetPubKeyEncryptedByIDQuery(pubKeyEncryptedID int, dbName string) string {
	return fmt.Sprintf(
		"SELECT pke.encrypted, pke.state, hpke.hashId FROM %v pke "+
			"INNER JOIN %v hpke on pke.Id = hpke.privateKeyEncryptedId "+
			"WHERE pke.Id = %v",
		dbName+"."+pubKeysEncryptedTableName,
		dbName+"."+hashesAndPubKeysEncryptedTableName,
		pubKeyEncryptedID,
	)
}

// GetKeyEncryptedByIDQuery ...
func GetKeyEncryptedByIDQuery(keyEncryptedID int, dbName string) string {
	return fmt.Sprintf(
		"SELECT ke.encrypted, ke.state, hke.hashId, kek.KeyId FROM %v ke "+
			"INNER JOIN %v hke on ke.Id = hke.keyEncryptedId "+
			"INNER JOIN %v kek on ke.Id = kek.KeyEncryptedId "+
			"WHERE ke.Id = %v",
		dbName+"."+keysEncryptedTableName,
		dbName+"."+hashesAndKeysEncryptedTableName,
		dbName+"."+keysEncryptedAndKeysTableName,
		keyEncryptedID,
	)
}
