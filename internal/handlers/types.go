package handlers

// table related endpoints types
type TableInfo struct {
	Name      string `json:"name"`
	RowCount  int64  `json:"rowCount"`
	CreatedAt string `json:"createdAt"`
}

type ColumnInfo struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Nullable  string `json:"nullable,omitempty"`
	Default   string `json:"default,omitempty"`
	Extra     string `json:"extra,omitempty"`
	Key       string `json:"key,omitempty"`
	Reference string `json:"reference,omitempty"`
}
