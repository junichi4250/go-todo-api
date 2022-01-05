package repository

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {

}
