package repository

import (
	"sample/domain/model"
)
/**
 * ここにはDBとのやりとりを定義します（今回はCRUDを定義しています）
 * domain層には技術的関心事を実装してはいけないというルールがあります。
 * 技術的関心事というのは「DBにMySQLを使って〜」や「ORMを使って〜」などです。
 * これらをdomain層に記述しないことによってdomain層が特定の技術に依存しないようになります。
 * なのでここにはCRUDの各メソッドをinterfaceとして定義しておいて、実装が書いてあるinfrastructure層がこのdomain層に依存するようにしてあげます。（この部分が依存関係逆転の原則を適用しているところです）
 */
type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	FindById(id int) (*model.Task, error)
	Update(task * model.Task) (*model.Task, error)

	)
}