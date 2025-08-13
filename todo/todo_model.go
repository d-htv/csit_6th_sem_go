package todo

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title" binding:"required"`
	Completed bool `json:"completed"`
}
// dummy data for now, alternative to database
var TodoList = []Todo{
	{ID: 1,Title: "Buy Milk",Completed: false,},
	{ID: 2,Title: "Buy Eggs",Completed: false,},
	{ID: 3,Title: "Buy Bread",Completed: false,},
}