package model

import (
	"errors"
)

type Task struct {
	ID      int
	Title   string
	Content string
}

/**
 * domain層 model
 * modelにはそのmodelの構造体やビジネスロジックが置かれています。
 * taskがどういうものかを定義していきます。（今回はコンストラクタとセッターのみ）
 * 理想はこのファイルを見ればtaskがどういうmodelかわかること(この場合、titleが必須ということがわかる)
 */
// 引数二つがわからない -> 型定義

// NewTask taskのコンストラクタ
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

// Set task のセッター
func (t *Task) Set(title, content string) error {
	if title == "" {
		return errors.New("titleを入力して下さい")
	}

	t.Title = title
	t.Content = content

	return nil
}
