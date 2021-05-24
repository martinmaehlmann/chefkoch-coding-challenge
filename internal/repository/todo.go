package repository

import (
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/todo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TodoRepository handles the operation of the Todo repository.
type TodoRepository struct {
	logger *zap.Logger
	config config.Dialector
	db     *gorm.DB
}

// Connect connects to the database or fails with a fatal error.
func (t *TodoRepository) Connect() {
	db, err := gorm.Open(t.config.Dialector(), &gorm.Config{})
	if err != nil {
		t.logger.Fatal("could not open connection to database")
	}

	t.db = db
}

// AutoMigrate Automatically migrate your schema, to keep your schema up to date.
// NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change
// existing column’s type if its size, precision, nullable changed. It WON’T delete unused columns to protect your data.
func (t *TodoRepository) AutoMigrate() error {
	err := t.db.AutoMigrate(&todo.Todo{})
	if err != nil {
		return err
	}

	err = t.db.AutoMigrate(&todo.Task{})
	if err != nil {
		return err
	}

	return nil
}

// FindAll returns all found Todos.
func (t *TodoRepository) FindAll() []*todo.Todo {
	todos := make([]*todo.Todo, 0)

	t.db.Preload(clause.Associations).Find(&todos)

	return todos
}

// Find returns a Todo by id, nil if none were found.
func (t *TodoRepository) Find(id uint) *todo.Todo {
	var toDo *todo.Todo

	if err := t.db.Preload(clause.Associations).Find(&toDo, id).Error; err != nil {
		return nil
	}

	return toDo
}

// Update updates an existing entry and returns the updated value, if it exists.
// If no entry exists, nil is returned.
func (t *TodoRepository) Update(toDo *todo.Todo) *todo.Todo {
	if foundTodo := t.Find(toDo.ID); foundTodo != nil {
		t.db.Save(toDo)

		return toDo
	}

	return nil
}

// Create creates the given entry with a new ID and returns the new entry.
func (t *TodoRepository) Create(toDo *todo.Todo) *todo.Todo {
	t.db.Create(toDo)

	return toDo
}

// Delete deletes the specified Todo and returns the rows affected.
func (t *TodoRepository) Delete(id uint) int64 {
	// the ID needs to be set here, otherwise the deletion hook will not know about the id
	result := t.db.Delete(&todo.Todo{Model: gorm.Model{ID: id}})

	return result.RowsAffected
}

// Close closes the connection to the database
func (t *TodoRepository) Close() {
	db, err := t.db.DB()
	if err != nil {
		t.logger.Fatal("could not get database from repository")
	}

	err = db.Close()
	if err != nil {
		t.logger.Fatal("could not close connection to database")
	}
}

// NewTodoRepository initializes a new TodoRepository.
func NewTodoRepository(config config.Dialector, logger *zap.Logger) *TodoRepository {
	return &TodoRepository{
		logger: logger,
		config: config,
	}
}
