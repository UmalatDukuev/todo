package todo

type TodoList struct {
	Id          int    `json:"-"`
	Title       string `json."title"  binding:"required"`
	Description string `json."description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json."title"`
	Description string `json."description"`
	Done        bool   `json."done"`
}

type ListItem struct {
	Id     int
	UserId int
	ListId int
}
