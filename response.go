package togoist

type Response struct {
	DayOrders     map[string]int   `json:"day_orders"`
	FullSync      bool             `json:"full_sync"`
	Items         []Item           `json:"items"`
	Projects      []Project        `json:"projects"`
	SyncToken     string           `json:"sync_token"`
	TempIdMapping map[string]int64 `json:"temp_id_mapping"`
	User          User             `json:"user"`
}
