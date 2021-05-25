package todo

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Task represents one task.
type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt   time.Time      `json:"-" gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time      `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	TodoID      int            `json:"-"`
}

// PrettyString pretty prints JSON.
func (t *Task) PrettyString(indentation string) (string, error) {
	jsonString, err := json.MarshalIndent(t, "", indentation)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

// Valid checks, if the Name has been set.
func (t *Task) Valid() bool {
	return len(t.Name) != 0
}
