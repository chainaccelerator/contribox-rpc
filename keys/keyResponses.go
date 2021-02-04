package keys

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/hashes"
)

// KeyShareResponseBody ...
type KeyShareResponseBody struct {
	Hash hashes.Hash `json:"hash"`
}

// KeyShareGetResponseBody ...
type KeyShareGetResponseBody struct {
	Hash                hashes.Hash               `json:"hash"`
	PubKeyEncryptedList []commons.PubKeyEncrypted `json:"pubKeyEncryptedList"`
	KeyEncryptedList    []commons.KeyEncrypted    `json:"keyEncryptedList"`
}
