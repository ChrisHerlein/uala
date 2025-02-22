package models

type FeedPage struct {
	UserName string    `json:"userName"` // user who will receive the content
	Content  []Content `json:"content"`
	Order    int       `json:"order"`
}

type Control struct {
	UserName   string `json:"userName"`
	MostRecent int    `json:"mostRecent"` // #order of last page
	SizeOfLast int    `json:"sizeOfLast"`
}
