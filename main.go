package main

import (
	"bc_node_api/api3/blockchain"
	"bc_node_api/api3/boarding"
	"bc_node_api/api3/config"
	"bc_node_api/api3/contribution"
	"bc_node_api/api3/git"
	"bc_node_api/api3/graph"
	"bc_node_api/api3/key"
	"bc_node_api/api3/persistance"
	"bc_node_api/api3/public"
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

var dbConf persistance.DbConf

func main() {
	fmt.Println("Blockchain node API")
	// Import config
	config := config.GetConfig()
	appURL = config.AppURL
	dbURL = config.DbURL
	dbName = config.DbName

	dbConf = persistance.DbConf{DbURL: dbURL, DbName: dbName}

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

	// Key
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

	// Boarding
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

	// Contribution
	case "contribution":
		_type, _contribution, state := contribution.ContributionParamConvert(params)
		return contribution.Contribution(_type, _contribution, state, dbConf)

	case "contributionGet":
		_type, xPubS, state := contribution.ContributionGetParamConvert(params)
		return contribution.ContributionGet(_type, xPubS, state, dbConf)

	case "contributionConfirm":
		_type, sig, hash, xPub, resourceEncrypted, state := contribution.ContributionConfirmParamConvert(params)
		return contribution.ContributionConfirm(_type, sig, hash, xPub, resourceEncrypted, state, dbConf)

	case "contributionConfirmGet":
		_type, hash, state := contribution.ContributionConfirmGetParamConvert(params)
		return contribution.ContributionConfirmGet(_type, hash, state, dbConf)

	case "contributionBroadcast":
		_type, resourceList, hash, state := contribution.ContributionBroadcastParamConvert(params)
		return contribution.ContributionBroadcast(_type, resourceList, hash, state, dbConf)

	case "contributionBroadcastGet":
		_type, hash, state := contribution.ContributionBroadcastGetParamConvert(params)
		return contribution.ContributionBroadcastGet(_type, hash, state, dbConf)

	// Public
	case "publicPeerListGet":
		tagList := public.PublicPeerListGetParamConvert(params)
		return public.PublicPeerListGet(tagList)

	// Git
	case "contriBox.backend.git.store":
		_type, resourceEncrypted, xPub, state := git.StoreParamConvert(params)
		return git.Store(_type, resourceEncrypted, xPub, state)

	case "contribox.backend.git.peerAsk":
		_type, xPub, depthMax, depth, state := git.PeerAskParamConvert(params)
		return git.PeerAsk(_type, xPub, depthMax, depth, state)

	case "contribox.backend.git.commitHashAsk":
		_type, hash, xPub, state := git.CommitHashAskParamConvert(params)
		return git.CommitHashAsk(_type, hash, xPub, state)

	// Graph
	case "contribox.backend.graph.store":
		_type, trace, state := graph.StoreParamConvert(params)
		return graph.Store(_type, trace, state)

	case "contribox.backend.graph.peerAsk":
		_type, xPub, depthMax, depth, state := graph.PeerAskParamConvert(params)
		return graph.PeerAsk(_type, xPub, depthMax, depth, state)

	// case "contribox.backend.graph.hashAsk":

	// Blockchain
	case "contribox.backend.blockchain.peerBlocValidation":
		_type, contributionTxID, state := blockchain.PeerValidationParamConvert(params)
		return blockchain.PeerBlocValidation(_type, contributionTxID, state)

	case "contribox.backend.blockchain.peerPegValidation":
		_type, contributionTxID, state := blockchain.PeerValidationParamConvert(params)
		return blockchain.PeerPegValidation(_type, contributionTxID, state)

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
