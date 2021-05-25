package todo

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Todo this struct represents a todo item.
type Todo struct {
	ID          uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt   time.Time      `json:"-" gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time      `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Tasks       []*Task        `json:"tasks" gorm:"foreignKey:TodoID"`
}

// PrettyString pretty prints the struct to json.
func (t *Todo) PrettyString(indentation string) (string, error) {
	jsonString, err := json.MarshalIndent(t, "", indentation)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

// BeforeDelete gorm hook to cascade the soft deletion of tasks.
func (t *Todo) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Table("tasks").Where("todo_id = ?", t.ID).Delete(&Task{})

	return
}

// BeforeCreate cleans up all accidentally set ids.
func (t *Todo) BeforeCreate(_ *gorm.DB) (err error) {
	t.ID = 0

	for _, task := range t.Tasks {
		task.ID = 0
	}

	return
}

// Valid checks if the Todo is valid and if all Tasks are valid.
func (t *Todo) Valid() bool {
	if len(t.Name) == 0 {
		return false
	}

	for _, task := range t.Tasks {
		if !task.Valid() {
			return false
		}
	}

	return true
}
