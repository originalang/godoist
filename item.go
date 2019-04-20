package togoist

type Item struct {
	AllDay            bool   `json:"all_day"`
	AssignedBy        int64  `json:"assigned_by_uid"`
	Checked           int    `json:"checked"`
	Collapsed         int    `json:"collapsed"`
	Content           string `json:"content"`
	DateAdded         string `json:"date_added"`
	DateCompleted     string `json:"date_completed"`
	DateString        string `json:"date_string"`
	DayOrder          int    `json:"day_order"`
	DueDate           string `json:"due_date_utc"`
	Id                int64  `json:"id"`
	InHistory         int    `json:"in_history"`
	Indent            int    `json:"indent"`
	IsArchived        int    `json:"is_archived"`
	IsDeleted         int    `json:"is_deleted"`
	ItemOrder         int    `json:"item_order"`
	ParentId          int64  `json:"parent_id"`
	Priority          int    `json:"priority"`
	ProjectId         int64  `json:"project_id"`
	ResponsibleUserId int64  `json:"responsible_uid"`
	SyncId            string `json:"sync_id"`
	UserId            int64  `json:"user_id"`
}
