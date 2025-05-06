package models

type ES_Document struct {
	Content           any         `json:"content"`
	FondCode          string      `json:"fondCode"`
	FondCodePartyCode string      `json:"fondCode_partyCode"`
	ID                string      `json:"id"`
	Metadata          ES_Metadata `json:"metadata"`
}

type ES_Metadata struct {
	ArcDocCode         string           `json:"arcDocCode"`
	ArcFileCode        string           `json:"arcFileCode"`
	ArchivesNumber     string           `json:"archivesNumber"`
	Autograph          string           `json:"autograph"`
	CodeNotation       string           `json:"codeNotation"`
	CodeNumber         string           `json:"codeNumber"`
	Colour             string           `json:"colour"`
	ConfidenceLevel    string           `json:"confidenceLevel"`
	Description        string           `json:"description"`
	DocAttached        string           `json:"docAttached"`
	DocID              string           `json:"docId"`
	EndDate            int64            `json:"endDate"`
	EndDateRaw         string           `json:"endDate_raw"`
	EventName          string           `json:"eventName"`
	FileExtension      string           `json:"fileExtension"`
	FilmSize           string           `json:"filmSize"`
	Format             string           `json:"format"`
	ImageTitle         string           `json:"imageTitle"`
	InforSign          string           `json:"inforSign"`
	IssuedDate         int64            `json:"issuedDate"`
	IssuedDateRaw      string           `json:"issuedDate_raw"`
	Keyword            string           `json:"keyword"`
	Language           []string         `json:"language"`
	Maintenance        string           `json:"maintenance"`
	Mode               string           `json:"mode"`
	MovieTitle         string           `json:"movieTitle"`
	NumberOfPage       string           `json:"numberOfPage"`
	NumberOfPaper      string           `json:"numberOfPaper"`
	OrganName          string           `json:"organName"`
	PaperFileCode      string           `json:"paperFileCode"`
	PhotoPlace         string           `json:"photoPlace"`
	PhotoTime          int64            `json:"photoTime"`
	PhotoTimeRaw       string           `json:"photoTime_raw"`
	Photographer       string           `json:"photographer"`
	PlayTime           string           `json:"playTime"`
	Process            string           `json:"process"`
	Quality            string           `json:"quality"`
	RecordDate         int64            `json:"recordDate"`
	RecordDateRaw      string           `json:"recordDate_raw"`
	RecordPlace        string           `json:"recordPlace"`
	Recorder           string           `json:"recorder"`
	RiskRecovery       string           `json:"riskRecovery"`
	RiskRecoveryStatus string           `json:"riskRecoveryStatus"`
	SchemaID           string           `json:"schemaId"`
	StartDate          int64            `json:"startDate"`
	StartDateRaw       string           `json:"startDate_raw"`
	Subject            string           `json:"subject"`
	Title              string           `json:"title"`
	TotalDoc           string           `json:"totalDoc"`
	TypeMedia          string           `json:"typeMedia"`
	TypeName           string           `json:"typeName"`
	TypePic            string           `json:"typePic"`
	Attachments        []ES_Attachments `json:"attachments"`
}

type ES_Attachments struct {
	FileExtension string `json:"fileExtension"`
	FileID        string `json:"fileId"`
	FileName      string `json:"fileName"`
	FilePath      string `json:"filePath"`
	MimeType      string `json:"mimeType"`
}
