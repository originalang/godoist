package togoist

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Indent int    `json:"indent"`
	Order  int    `json:"order"`
}

type Task struct {
	Id        int     `json:"id"`
	Indent    int     `json:"indent"`
	Order     int     `json:"order"`
	Priority  int     `json:"priority"`
	ProjectId int     `json:"project_id"`
	Completed bool    `json:"completed"`
	Content   string  `json:"content"`
	Due       DueInfo `json:"due"`
}

type DueInfo struct {
	Date       string `json:"date"`
	Recurring  bool   `json:"recurring"`
	DateTime   string `json:"datetime"`
	StringDate string `json:"string"`
	Timezone   string `json:"timezone"`
}
