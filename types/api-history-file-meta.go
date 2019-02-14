package types

// APIHistoryFileMeta contains the meta data about a history file
type APIHistoryFileMeta struct {
	Date     string `json:"date"`
	DataType string `json:"date_type"`
	Size     int64  `json:"size"`
}
