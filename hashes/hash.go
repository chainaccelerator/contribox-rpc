package hashes

import (
	"bc_node_api/api3/persistance"
	"strconv"
)

// Hash ...
type Hash struct {
	Left  interface{}
	Right interface{}
	Hash  string
	Data  string
}

// HashRef ...
type HashRef struct {
	Left  string
	Right string
	Hash  string
	Data  string
}

// BuildHash ...
func BuildHash(hParam map[string]interface{}) Hash {
	leftParam := hParam["left"]
	rightParam := hParam["right"]
	hash := hParam["hash"].(string)
	data := hParam["data"].(string)

	// If leaf
	if leftParam.(string) == "Type" || leftParam.(string) == "Value" {
		return Hash{
			Left:  leftParam.(string),
			Right: rightParam.(string),
			Hash:  hash,
			Data:  data,
		}
	}

	left := BuildHash(leftParam.(map[string]interface{}))

	// If single
	// if rightParam.(string) == "" {
	// 	return Hash{
	// 		Left:  left,
	// 		Right: "",
	// 		Hash:  hash,
	// 		Data:  data,
	// 	}
	// }

	// If normal
	return Hash{
		Left:  left,
		Right: BuildHash(rightParam.(map[string]interface{})),
		Hash:  hash,
		Data:  data,
	}
}

// ValidateHash ...
func ValidateHash(h Hash) bool {
	validHashAndData := len(h.Hash) <= 64 && (len(h.Data) >= 64 && len(h.Data) <= 6400000000)

	// If leaf
	if h.Left == "Type" || h.Left == "Value" {
		return validHashAndData
	}

	leftParam := h.Left.(map[string]interface{})

	// If single
	// if hash.Right == "" {
	// 	return validHashAndData && ValidateHash(BuildHash(leftParam))
	// }

	rightParam := h.Right.(map[string]interface{})
	return validHashAndData && ValidateHash(BuildHash(leftParam)) && ValidateHash(BuildHash(rightParam))
}

// SaveHash ...
func SaveHash(h Hash, dbConf persistance.DbConf) int {
	// If leaf
	if h.Left == "Type" || h.Left == "Value" {
		leafInserted := InsertLeaf(h, dbConf)
		if leafInserted == 0 {
			return 0
		}
	}

	// If normal node
	left := BuildHash(h.Left.(map[string]interface{}))
	right := BuildHash(h.Right.(map[string]interface{}))
	nodeInserted := InsertNode(left.Hash, right.Hash, h.Hash, h.Data, dbConf)
	leftInserted := SaveHash(left, dbConf)
	rightInserted := SaveHash(right, dbConf)

	if nodeInserted == 0 || leftInserted == 0 || rightInserted == 0 {
		return 0
	}

	return nodeInserted
}

// GetHash ...
func GetHash(hashID int, dbConf persistance.DbConf) Hash {
	hashRef := GetHashRef(hashID, dbConf)

	if hashRef.Left == "Type" || hashRef.Left == "Value" {
		return Hash{
			Left:  hashRef.Left,
			Right: hashRef.Right,
			Hash:  hashRef.Hash,
			Data:  hashRef.Data,
		}
	}

	hashRefLeft, _ := strconv.Atoi(hashRef.Left)
	hashRefRight, _ := strconv.Atoi(hashRef.Right)

	return Hash{
		Left:  GetHash(hashRefLeft, dbConf),
		Right: GetHash(hashRefRight, dbConf),
		Hash:  hashRef.Hash,
		Data:  hashRef.Data,
	}
}
