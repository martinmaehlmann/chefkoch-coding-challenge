package repository

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	config_mock "gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config/mock"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/todo"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
)

func TestTodoRepository_AutoMigrate(t *testing.T) {
	// get the test setup
	ctrl := gomock.NewController(t)
	tempDir, repository := newSQLliteTodoRepository(t, ctrl)

	defer cleanup(t, tempDir, ctrl, repository)

	// call automigrate
	err := repository.AutoMigrate()
	assert.NoError(t, err)
}

func TestTodoRepository_Create(t *testing.T) {
	// get the test setup
	ctrl := gomock.NewController(t)
	tempDir, repository := newSQLliteTodoRepository(t, ctrl)

	defer cleanup(t, tempDir, ctrl, repository)

	// create the schema
	err := repository.AutoMigrate()
	assert.NoError(t, err)

	// get a Todo and alter its id
	toDoToCreate := defaultTestReturnTodo()
	toDoToCreate.ID = 10

	// create the todo
	repository.Create(toDoToCreate)

	// check if anything was created
	todos := repository.FindAll()
	assert.Equal(t, 1, len(todos))

	// check if the todo was created and is still the same
	toDo := repository.Find(1)
	assert.True(t, todoEqualsWithoutTimeFields(t, toDo, defaultTestReturnTodo()))

	// check if the tasks are created
	var tasks []*todo.Task

	repository.db.Table("tasks").Find(&tasks)
	assert.Equal(t, 2, len(tasks))
	assert.Equal(t, 1, tasks[0].TodoID)
	assert.Equal(t, uint(2), tasks[1].ID)
}

func TestTodoRepository_SoftDelete(t *testing.T) {
	// get the test setup
	ctrl := gomock.NewController(t)
	tempDir, repository := newSQLliteTodoRepository(t, ctrl)

	defer cleanup(t, tempDir, ctrl, repository)

	// setup schema
	err := repository.AutoMigrate()
	assert.NoError(t, err)

	// create todo
	repository.Create(defaultTestReturnTodo())

	// check if todo was created
	toDo := repository.Find(1)
	assert.True(t, todoEqualsWithoutTimeFields(t, toDo, defaultTestReturnTodo()))

	// delete todo
	rowsAffected := repository.Delete(1)
	assert.Equal(t, int64(1), rowsAffected)

	// check if the todo was deleted
	todos := repository.FindAll()
	assert.Equal(t, 0, len(todos))

	// check if tasks were deleted cascading
	var tasks []*todo.Task

	repository.db.Table("tasks").Find(&tasks)
	assert.Equal(t, 0, len(tasks))

	// check if the todo can still be found if deleted_at is ignored
	var todosUnscoped []*todo.Todo

	repository.db.Unscoped().Find(&todosUnscoped)
	assert.Equal(t, 1, len(todosUnscoped))

	// check if the tasks can still be found if deleted_at is ignored
	var tasksUnscoped []*todo.Task

	repository.db.Table("tasks").Unscoped().Find(&tasksUnscoped)
	assert.Equal(t, 2, len(tasksUnscoped))
}

func TestTodoRepository_Find(t *testing.T) {
	// get the test setup
	ctrl := gomock.NewController(t)
	tempDir, repository := newSQLliteTodoRepository(t, ctrl)

	defer cleanup(t, tempDir, ctrl, repository)

	// setup schema
	err := repository.AutoMigrate()
	assert.NoError(t, err)

	// create todo
	repository.Create(defaultTestReturnTodo())

	// check if todo was created
	toDo := repository.Find(1)
	assert.True(t, todoEqualsWithoutTimeFields(t, toDo, defaultTestReturnTodo()))
}

func TestTodoRepository_FindAll(t *testing.T) {
	// get the test setup
	ctrl := gomock.NewController(t)
	tempDir, repository := newSQLliteTodoRepository(t, ctrl)

	defer cleanup(t, tempDir, ctrl, repository)

	// create the schema
	err := repository.AutoMigrate()
	assert.NoError(t, err)

	// create the todo
	repository.Create(defaultTestReturnTodo())
	repository.Create(defaultTestReturnTodo())
	repository.Create(defaultTestReturnTodo())

	// check if the todo can still be found if deleted_at is ignored
	var todosUnscoped []*todo.Todo

	repository.db.Unscoped().Find(&todosUnscoped)
	assert.Equal(t, 3, len(todosUnscoped))
}

func TestTodoRepository_Update(t *testing.T) {
	// get the test setup
	ctrl := gomock.NewController(t)
	tempDir, repository := newSQLliteTodoRepository(t, ctrl)

	defer cleanup(t, tempDir, ctrl, repository)

	// create the schema
	err := repository.AutoMigrate()
	assert.NoError(t, err)

	// create the todo
	repository.Create(defaultTestReturnTodo())

	// check if todo was created
	toDo := repository.Find(1)
	assert.True(t, todoEqualsWithoutTimeFields(t, toDo, defaultTestReturnTodo()))

	// update todo
	updatedTodo := *toDo
	updatedTodo.Name = "updated test"

	repository.Update(&updatedTodo)

	// check if todo was updated
	foundTodo := repository.Find(1)
	assert.True(t, todoEqualsWithoutTimeFields(t, &updatedTodo, foundTodo))
}

// newSQLliteTodoRepository returns a new todoRepository with an sqllite db as the backend.
func newSQLliteTodoRepository(t *testing.T, ctrl *gomock.Controller) (dbFile *os.File, repository *todoRepository) {
	t.Helper()

	tempDir := os.TempDir()

	sqlLiteDatabaseFile, err := os.CreateTemp(tempDir, "sqllitedatabase")
	assert.NoError(t, err)

	dialector := sqlite.Open(sqlLiteDatabaseFile.Name())

	mockDialector := config_mock.NewMockDialector(ctrl)
	mockDialector.EXPECT().Dialector().Return(dialector)

	logger, err := zap.NewProduction()
	assert.NoError(t, err)

	repository = &todoRepository{
		logger: logger,
		config: mockDialector,
		db:     nil,
	}

	repository.Connect()

	return sqlLiteDatabaseFile, repository
}

// cleanup cleans up all the created files during the test.
func cleanup(t *testing.T, dbFile *os.File, ctrl *gomock.Controller, repository *todoRepository) {
	t.Helper()
	repository.Close()

	err := os.Remove(dbFile.Name())
	assert.NoError(t, err)

	ctrl.Finish()
}

// todoEqualsWithoutTimeFields checks if the todos are equal, if one ignores fields like created_at.
// These fields are not checked, as they are set by gorm and there is and should be little to no control over them
// by the application.
func todoEqualsWithoutTimeFields(t *testing.T, t1 *todo.Todo, t2 *todo.Todo) bool {
	t.Helper()

	if unequal := t1.ID != t2.ID; unequal {
		return !unequal
	}

	if unequal := t1.Name != t2.Name; unequal {
		return !unequal
	}

	if unequal := t1.Description != t2.Description; unequal {
		return !unequal
	}

	if unequal := len(t1.Tasks) != len(t2.Tasks); unequal {
		return !unequal
	}

	for i := range t1.Tasks {
		taskEqualsWithoutTimeFields(t, t1.Tasks[i], t2.Tasks[i])
	}

	return true
}

// taskEqualsWithoutTimeFields checks if the tasks are equal, if one ignores fields like created_at.
// These fields are not checked, as they are set by gorm and there is and should be little to no control over them
// by the application.
func taskEqualsWithoutTimeFields(t *testing.T, t1 *todo.Task, t2 *todo.Task) bool {
	t.Helper()

	if unequal := t1.ID != t2.ID; unequal {
		return !unequal
	}

	if unequal := t1.Name != t2.Name; unequal {
		return !unequal
	}

	if unequal := t1.Description != t2.Description; unequal {
		return !unequal
	}

	return true
}

// defaultTestReturnTodo convenience function to return a struct to test with.
func defaultTestReturnTodo() *todo.Todo {
	tasks := make([]*todo.Task, 2)
	tasks[0] = &todo.Task{
		ID:          1,
		Name:        "test1",
		Description: "test1",
		TodoID:      1,
	}
	tasks[1] = &todo.Task{
		ID:          2,
		Name:        "test2",
		Description: "test2",
		TodoID:      1,
	}

	return &todo.Todo{
		ID:          1,
		Name:        "test",
		Description: "test",
		Tasks:       tasks,
	}
}
