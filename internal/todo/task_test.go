package todo

import (
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestTask_Valid(t1 *testing.T) {
	type fields struct {
		Model       gorm.Model
		Name        string
		Description string
		TodoID      int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test valid task",
			fields: fields{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name:        "test",
				Description: "test",
				TodoID:      1,
			},
			want: true,
		},
		{
			name: "test invalid task",
			fields: fields{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name:        "",
				Description: "test",
				TodoID:      1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Task{
				Model:       tt.fields.Model,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				TodoID:      tt.fields.TodoID,
			}
			if got := t.Valid(); got != tt.want {
				t1.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
