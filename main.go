package main

import (
	"bc_node_api/api3/blabla"
	"bc_node_api/api3/boarding"
	"bc_node_api/api3/commons"
	"bc_node_api/api3/config"
	"bc_node_api/api3/contribution"
	"bc_node_api/api3/key"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var appURL string
var dbURL string
var dbName string

var dbConf commons.DbConf

func main() {
	fmt.Println("Blockchain node API")
	// Import config
	config := config.GetConfig()
	appURL = config.AppURL
	dbURL = config.DbURL
	dbName = config.DbName

	dbConf = commons.DbConf{DbURL: dbURL, DbName: dbName}

	// Test d'import de package
	blabla.Blabla()

	requestHandler()
}

// Server initialization
func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(appURL, router))
}

// Unamarshal request body to read method address and method parameters
func handleRequest(w http.ResponseWriter, r *http.Request) {
	reqBodyJSON, _ := ioutil.ReadAll(r.Body)
	var reqBody Request
	json.Unmarshal(reqBodyJSON, &reqBody)
	success, error := handleAddress(reqBody.Method, reqBody.Params)
	response := Response{Success: success, Error: error}
	json.NewEncoder(w).Encode(response)
}

// Route method address to actual method
func handleAddress(address string, params []interface{}) (interface{}, bool) {
	switch address {

	case "keyShare":
		_type, keyShared, state := key.KeyShareParamConvert(params)
		return key.KeyShare(_type, keyShared, state, dbConf)

	case "keyShareGet":
		_type, xPubS, state := key.KeyShareGetParamConvert(params)
		return key.KeyShareGet(_type, xPubS, state, dbConf)

	case "keyShareConfirm":
		_type, resource, hash, state := key.KeyShareConfirmParamConvert(params)
		return key.KeyShareConfirm(_type, resource, hash, state, dbConf)

	case "keyShareConfirmGet":
		_type, hash, state := key.KeyShareConfirmGetParamConvert(params)
		return key.KeyShareConfirmGet(_type, hash, state, dbConf)

	case "boardingTemplateGet":
		projectName, licenceSPDX, groupRoleName, onBoarding, outBoarding, hash, state := boarding.BoardingTemplateGetParamConvert(params)
		return boarding.BoardingTemplateGet(projectName, licenceSPDX, groupRoleName, onBoarding, outBoarding, hash, state, dbConf)

	case "boarding":
		_type, resource, state := boarding.BoardingParamConvert(params)
		return boarding.Boarding(_type, resource, state, dbConf)

	case "boardingGet":
		_type, xPubS, state := boarding.BoardingGetParamConvert(params)
		return boarding.BoardingGet(_type, xPubS, state, dbConf)

	case "boardingBroadcast":
		_type, resourceList, hash, state := boarding.BoardingBroadcastParamConvert(params)
		return boarding.BoardingBroadcast(_type, resourceList, hash, state, dbConf)

	case "boardingBroadcastGet":
		_type, hash, state := boarding.BoardingBroadcastGetParamConvert(params)
		return boarding.BoardingBroadcastGet(_type, hash, state, dbConf)

	case "contribution":
		_type, _contribution, state := contribution.ContributionParamConvert(params)
		return contribution.Contribution(_type, _contribution, state, dbConf)

	case "contributionGet":
		_type, xPubS, state := contribution.ContributionGetParamConvert(params)
		return contribution.ContributionGet(_type, xPubS, state, dbConf)

	case "contributionConfirm":
		_type, sig, hash, state := contribution.ContributionConfirmParamConvert(params)
		return contribution.ContributionConfirm(_type, sig, hash, state, dbConf)

	case "contributionConfirmGet":
		_type, hash, state := contribution.ContributionConfirmGetParamConvert(params)
		return contribution.ContributionConfirmGet(_type, hash, state, dbConf)

	case "contributionBroadcast":
		_type, resourceList, hash, state := contribution.ContributionBroadcastParamConvert(params)
		return contribution.ContributionBroadcast(_type, resourceList, hash, state, dbConf)

	case "contributionBroadcastGet":
		_type, hash, state := contribution.ContributionBroadcastGetParamConvert(params)
		return contribution.ContributionBroadcastGet(_type, hash, state, dbConf)

	default:
		return "Address not found", true
	}
}

// Request ...
type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// Response ...
type Response struct {
	Success interface{}
	Error   bool
}
