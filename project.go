package togoist

type Project struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Color      int    `json:"color"`
	Indent     int    `json:"indent"`
	Order      int    `json:"item_order"`
	Collapsed  int    `json:"collapsed"`
	Shared     bool   `json:"shared"`
	IsDeleted  int    `json:"is_deleted"`
	IsArchived int    `json:"is_archived"`
	IsFavorite int    `json:"is_favorite"`
}
