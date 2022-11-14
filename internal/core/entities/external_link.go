package entities

type ExternalLink struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	LinkID   uint   `json:"link_id"`
	Link     Link   `json:"-" gorm:"foreignKey:LinkID;references:ID"`
	URL      string `json:"url"`
	Provider string `json:"provider"`
}
