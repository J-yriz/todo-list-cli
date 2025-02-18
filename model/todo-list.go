package model

type DataTodo struct {
	ID        int16
	Feild     string
	CreatedAt string
}

type DataAll struct {
	DataTodo []DataTodo
	Username string
}