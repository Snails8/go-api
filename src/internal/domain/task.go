package domain

import (
	"errors"
)

type Task struct {
	ID      int
	Title   string
	Content string
}

func NewTask(title, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	task := &Task{
		Title:   title,
		Content: content,
	}

	return task, nil
}
