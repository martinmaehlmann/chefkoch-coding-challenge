package todo

import (
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestTodo_Valid(t1 *testing.T) {
	type fields struct {
		Model       gorm.Model
		Name        string
		Description string
		Tasks       []Task
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test valid todo",
			fields: fields{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name:        "test",
				Description: "test",
				Tasks: []Task{
					{
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
				},
			},
			want: true,
		},
		{
			name: "test valid todo",
			fields: fields{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name:        "test",
				Description: "test",
				Tasks:       []Task{},
			},
			want: true,
		},
		{
			name: "test invalid todo",
			fields: fields{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name:        "",
				Description: "test",
				Tasks: []Task{
					{
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
				},
			},
			want: false,
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
				Name:        "test",
				Description: "test",
				Tasks: []Task{
					{
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
				},
			},
			want: false,
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
				Name:        "test",
				Description: "test",
				Tasks: []Task{
					{
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
					{
						Model: gorm.Model{
							ID:        0,
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
							DeletedAt: gorm.DeletedAt{},
						},
						Name:        "",
						Description: "test",
						TodoID:      1,
					},
				},
			},
			want: false,
		},
		{
			name: "test nil task slice",
			fields: fields{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name:        "test",
				Description: "test",
				Tasks:       nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Todo{
				Model:       tt.fields.Model,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Tasks:       tt.fields.Tasks,
			}
			if got := t.Valid(); got != tt.want {
				t1.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
