package togoist

type User struct {
	StartPage         string   `json:"start_page"`
	Features          Features `json:"features"`
	CompletedToday    int      `json:"completed_today"`
	IsPremium         bool     `json:"is_premium"`
	SortOrder         int      `json:"sort_order"`
	FullName          string   `json:"full_name"`
	AutoReminder      int      `json:"auto_reminder"`
	Id                int64    `json:"id"`
	ShareLimit        int      `json:"share_limit"`
	DaysOff           []int    `json:"days_off"`
	NextWeek          int      `json:"next_week"`
	CompletedCount    int      `json:"completed_count"`
	DailyGoal         int      `json:"daily_goal"`
	Theme             int      `json:"theme"`
	TimeZoneInfo      TZInfo   `json:"tz_info"`
	Email             string   `json:"email"`
	StartDay          int      `json:"start_day"`
	WeeklyGoal        int      `json:"weekly_goal"`
	DateFormat        int      `json:"date_format"`
	WebsocketURL      string   `json:"websocket_url"`
	InboxProject      int64    `json:"inbox_project"`
	TimeFormat        int      `json:"time_format"`
	KarmaTrend        string   `json:"karma_trend"`
	BusinessAccountId int64    `json:"business_account_id"`
	ImageId           string   `json:"image_id"`
	MobileNumber      string   `json:"mobile_number"`
	MobileHost        string   `json:"mobile_host"`
	PremiumUntil      string   `json:"premium_until"`
	JoinDate          string   `json:"join_date"`
	Karma             int64    `json:"karma"`
	IsBizAdmin        bool     `json:"is_biz_admin"`
	DefaultReminder   string   `json:"default_reminder"`
	Token             string   `json:"token"`
}

type Features struct {
	KarmaDisabled    bool `json:"karma_disabled"`
	Restriction      int  `json:"restriction"`
	KarmaVacation    bool `json:"karma_vacation"`
	Beta             int  `json:"beta"`
	HasPushReminders bool `json:"has_push_reminders"`
}

type TZInfo struct {
	Hours     int    `json:"hours"`
	Timezone  string `json:"timezone"`
	IsDST     int    `json:"is_dst"`
	Minutes   int    `json:"minutes"`
	GMTString string `json:"gmt_string"`
}
