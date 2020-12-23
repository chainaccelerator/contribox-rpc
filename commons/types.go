package commons

// BlindingKey ...
type BlindingKey struct {
	Key string
}

// BlindingKeyEncrypted ...
type BlindingKeyEncrypted struct {
	XPub   XPub
	String string
}

// Boarding ...
type Boarding struct {
	XPubList            []XPub
	In                  bool
	GroupRoleNameList   []GroupRoleName
	GroupActionNameList []GroupActionName
}

// Contribution ...
type Contribution struct {
	Hash            Hash
	XPub            XPub
	Proof           Proof
	BlindKeyList    []BlindingKeyEncrypted
	RangeList       []RangeEncrypted
	OnBoarding      Boarding
	OutBoarding     Boarding
	Tx1Id           TxId
	Tx0IdAmount     int
	Tx0IdIssueAsset IssueAsset
	Tx0IdSigA       SigData
	Vout0PubKA      PubKey
	Vout1PubKS      PubKey
}

// FullContribution ...
type FullContribution struct {
	Template Template
	Contribution
}

// FullTemplate ...
type FullTemplate struct {
	Contribution Contribution
	Template
}

// GroupActionName ...
type GroupActionName struct {
	Name string
}

// GroupRoleName ...
type GroupRoleName struct {
	Name string
}

// Hash ...
type Hash struct {
	Hash string
}

// IssueAsset ...
type IssueAsset struct {
	IssueAsset Hash
}

// Key ...
type Key struct {
	Key string
}

// KeyShared ...
type KeyShared struct {
	XPubSList []XPub
	Key       Key
	Hash      Hash
}

// KeyVal ...
type KeyVal struct {
	Key string
	Val map[string]interface{}
}

// Licence ...
type Licence struct {
	SPDX string
}

// ProjectName ...
type ProjectName struct {
	Name string
}

// Proof ...
type Proof struct {
	XPub                  XPub
	ProjectName           ProjectName
	LicenseSPDX           Licence
	LicenseSPDXChange     Licence
	GroupRoleName         GroupRoleName
	DescriptionPublicList []KeyVal
	IdentityList          []KeyVal
	Event                 Hash
	EventList             []KeyVal
	EnvironmentList       []KeyVal
	QualityList           []KeyVal
	ContributeList        []KeyVal
	OriginList            []Hash
	ParentList            []Hash
	PreviousList          []Hash
	LeftList              []Hash
	NdaList               []KeyVal
	ConfidentialDataList  []KeyVal
	MetaDataList          []KeyVal
	OfficerList           []KeyVal
	EditList              []KeyVal
	CertificateList       []KeyVal
	ExportControlList     []KeyVal
	KeyValueList          []KeyVal
	ForList               []XPub
	ToList                []XPub
	TagList               []string
}

// PubKey ...
type PubKey struct {
	PubKey        string
	Base58Encoded bool
}

// Range ...
type Range struct {
	Min int
	Max int
}

// RangeEncrypted ...
type RangeEncrypted struct {
	XPub   XPub
	String string
}

// Script ...
type Script struct {
	Script string
}

// Sig ...
type Sig struct {
	Sig   SigData
	XPub  XPub
	XPubS XPub
}

// SigData ...
type SigData struct {
	Sig string
}

// StateReason ...
type StateReason struct {
	Reason string
}

// Template ...
type Template struct {
	Hash                    Hash
	Proof                   Proof
	ProjectName             ProjectName
	LicenceSPDX             Licence
	GroupRoleName           GroupRoleName
	State                   StateReason
	UserRequirement         bool
	ProjectRequirement      bool
	UserUser                TemplateUser
	UserBackup              TemplateUser
	UserLock                TemplateUser
	UserWitness             TemplateUser
	UserUserXPubList        []XPub
	UserBackupXPubList      []XPub
	UserLockXPubList        []XPub
	UserWitnessXPubList     []XPub
	ProjectOld              TemplateProject
	ProjectParent           TemplateProject
	ProjectBoard            TemplateProject
	ProjectMember           TemplateProject
	ProjectCosigner         TemplateProject
	ProjectWitness          TemplateProject
	ProjectOldXPubList      []XPub
	ProjectParentXPubList   []XPub
	ProjectBoardXPubList    []XPub
	ProjectMemberXPubList   []XPub
	ProjectCosignerXPubList []XPub
	ProjectWitnessXPubList  []XPub
	ScriptTemplate          string
}

// TemplateProject ...
type TemplateProject struct {
	Quorum string
}

// TemplateUser ...
type TemplateUser struct {
	Quorum string
}

// TimeStamp ...
type TimeStamp struct {
	Time  int
	Delay int
}

// TxId ...
type TxId struct {
	Id string
}

// UTXO ...
type UTXO struct {
	Hash   Hash
	Tx0Id  TxId
	UTXO   UTXOData
	Script Script
}

// UTXOData ...
type UTXOData struct {
	UTXO string
}

// XPub ...
type XPub struct {
	XPub string
}
