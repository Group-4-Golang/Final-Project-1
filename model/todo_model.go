package model

import "time"

type Todo struct {
	Id          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"Assigntment 1"`
	Deadline    time.Time `json:"deadline" example:"2022-09-15T14:30:45.0000001+07:00"`
	Description string    `json:"description" example:"Something to do in the weekend"`
	Status      string    `json:"status" example:"New"`
}

type Message struct {
	Message string `json:"message" example:"Todo not found"`
}

type TodoRequest struct {
	Name        string    `json:"name" example:"Assigntment 1"`
	Deadline    time.Time `json:"deadline" example:"2022-09-15T14:30:45.0000001+07:00"`
	Description string    `json:"description" example:"Something to do in the weekend"`
	Status      string    `json:"status" example:"New"`
}
