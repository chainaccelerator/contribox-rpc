package commons

import (
	"bc_node_api/api3/hashes"
)

// Metadata ...
type Metadata struct {
	Hash hashes.Hash
}

// ValidateMetadata ...
func ValidateMetadata(metadata Metadata) bool {
	return hashes.ValidateHash(metadata.Hash)
}
