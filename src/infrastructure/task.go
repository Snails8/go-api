package infrastructure

import (
	"github.com/jinzhu/gorm"
)

// 特定の技術基盤にアクセスする層

// ここで先程のrepositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装していきます。
// domain層に定義、infrastructure層に実装、と分けることでプロジェクトで使っているDBやORMが変更したとしてもinfrastructure層の実装のみを修正すればいいだけになります。
// 今回はDBにMySQL、ORMにgormを使っています
type taskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

// Create taskの保存
func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// FindByID taskをIDで取得
func (tr *TaskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := tr.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Update taskの更新
func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Model(&task).Update(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete taskの削除
func (tr *TaskRepository) Delete(task *model.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
