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
	"os"

	"github.com/gorilla/mux"
)

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

var appURL string
var dbURL string
var dbName string

var dbConf persistance.DbConf

func main() {
	fmt.Println("Blockchain node API")

	// We're excluding the path to the program
	args := os.Args[1:]
	env := args[0]

	// Import config
	config := config.GetConfig(env)
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
	case "contriBox.sdk.endpoint.keyShare":
		_type, keyShared, state := key.KeyShareParamConvert(params)
		return key.KeyShare(_type, keyShared, state, dbConf)

	case "contriBox.sdk.endpoint.keyShareGet":
		_type, xPubS, state := key.KeyShareGetParamConvert(params)
		return key.KeyShareGet(_type, xPubS, state, dbConf)

	case "contriBox.sdk.endpoint.keyShareConfirm":
		_type, resource, hash, state := key.KeyShareConfirmParamConvert(params)
		return key.KeyShareConfirm(_type, resource, hash, state, dbConf)

	case "contriBox.sdk.endpoint.keyShareConfirmGet":
		_type, hash, state := key.KeyShareConfirmGetParamConvert(params)
		return key.KeyShareConfirmGet(_type, hash, state, dbConf)

	// Boarding
	case "contriBox.sdk.endpoint.boardingTemplateGet":
		projectName, licenceSPDX, groupRoleName, onBoarding, outBoarding, hash, state := boarding.BoardingTemplateGetParamConvert(params)
		return boarding.BoardingTemplateGet(projectName, licenceSPDX, groupRoleName, onBoarding, outBoarding, hash, state, dbConf)

	case "contriBox.sdk.endpoint.boarding":
		_type, resource, state := boarding.BoardingParamConvert(params)
		return boarding.Boarding(_type, resource, state, dbConf)

	case "contriBox.sdk.endpoint.boardingGet":
		_type, xPubS, state := boarding.BoardingGetParamConvert(params)
		return boarding.BoardingGet(_type, xPubS, state, dbConf)

	case "contriBox.sdk.endpoint.boardingBroadcast":
		_type, resourceList, hash, state := boarding.BoardingBroadcastParamConvert(params)
		return boarding.BoardingBroadcast(_type, resourceList, hash, state, dbConf)

	case "contriBox.sdk.endpoint.boardingBroadcastGet":
		_type, hash, state := boarding.BoardingBroadcastGetParamConvert(params)
		return boarding.BoardingBroadcastGet(_type, hash, state, dbConf)

	// Contribution
	case "contriBox.sdk.endpoint.contribution":
		_type, _contribution, state := contribution.ContributionParamConvert(params)
		return contribution.Contribution(_type, _contribution, state, dbConf)

	case "contriBox.sdk.endpoint.contributionGet":
		_type, xPubS, state := contribution.ContributionGetParamConvert(params)
		return contribution.ContributionGet(_type, xPubS, state, dbConf)

	case "contriBox.sdk.endpoint.contributionConfirm":
		_type, sig, hash, xPub, resourceEncrypted, state := contribution.ContributionConfirmParamConvert(params)
		return contribution.ContributionConfirm(_type, sig, hash, xPub, resourceEncrypted, state, dbConf)

	case "contriBox.sdk.endpoint.contributionConfirmGet":
		_type, hash, state := contribution.ContributionConfirmGetParamConvert(params)
		return contribution.ContributionConfirmGet(_type, hash, state, dbConf)

	case "contriBox.sdk.endpoint.contributionBroadcast":
		_type, resourceList, hash, state := contribution.ContributionBroadcastParamConvert(params)
		return contribution.ContributionBroadcast(_type, resourceList, hash, state, dbConf)

	case "contriBox.sdk.endpoint.contributionBroadcastGet":
		_type, hash, state := contribution.ContributionBroadcastGetParamConvert(params)
		return contribution.ContributionBroadcastGet(_type, hash, state, dbConf)

	// Public
	case "contriBox.sdk.endpoint.publicPeerListGet":
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

	case "contribox.backend.graph.hashAsk":
		_type, trace, state := graph.HashAskParamConvert(params)
		return graph.HashAsk(_type, trace, state)

	// Blockchain
	case "contriBox.backend.blockchain.broadcast":
		_type, transaction, state := blockchain.BroadcastParamConvert(params)
		return blockchain.Broadcast(_type, transaction, state)

	case "contriBox.backend.blockchain.peerAsk":
		_type, txID, state := blockchain.PeerParamConvert(params)
		return blockchain.PeerAsk(_type, txID, state)

	case "contribox.backend.blockchain.peerBlocValidation":
		_type, contributionTxID, state := blockchain.PeerParamConvert(params)
		return blockchain.PeerBlocValidation(_type, contributionTxID, state)

	case "contribox.backend.blockchain.peerPegValidation":
		_type, contributionTxID, state := blockchain.PeerParamConvert(params)
		return blockchain.PeerPegValidation(_type, contributionTxID, state)

	default:
		return "Address not found", true
	}
}
