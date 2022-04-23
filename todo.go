package todo

// List describes a list of todo-items
type List struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Item describes an item of todo-list
type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

// UsersList links User and List object
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

// ListsItem links List and its Item
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}
