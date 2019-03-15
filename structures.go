package togoist

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Indent int    `json:"indent"`
	Order  int    `json:"order"`
}
