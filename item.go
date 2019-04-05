package togoist

type Item struct {
	AllDay            bool   `json:"all_day"`
	AssignedBy        int64  `json:"assigned_by_uid"`
	DateCompleted     string `json:"date_completed"`
	ParentId          int64  `json:"parent_id"`
	ResponsibleUserId int64  `json:"responsible_uid"`
	SyncId            string `json:"sync_id"`
	Checked           bool   `json:"checked"`
	Collapsed         int    `json:"collapsed"`
	Content           string `json:"content"`
	DateAdded         string `json:"date_added"`
	DateString        string `json:"date_string"`
	DayOrder          int    `json:"day_order"`
	DueDate           string `json:"due_date_utc"`
	ID                int64  `json:"id"`
	InHistory         bool   `json:"in_history"`
	Indent            int    `json:"indent"`
	IsArchived        bool   `json:"is_archived"`
	IsDeleted         bool   `json:"is_deleted"`
	ItemOrder         int    `json:"item_order"`
	Priority          int    `json:"priority"`
	ProjectId         int64  `json:"project_id"`
	UserId            int64  `json:"user_id"`
}
