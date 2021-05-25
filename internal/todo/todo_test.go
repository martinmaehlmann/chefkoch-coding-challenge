package todo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodo_Valid(t1 *testing.T) {
	type fields struct {
		ID          uint
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
				ID:          1,
				Name:        "test",
				Description: "test",
				Tasks: []Task{
					{
						ID:          1,
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
				ID:          1,
				Name:        "test",
				Description: "test",
				Tasks:       []Task{},
			},
			want: true,
		},
		{
			name: "test invalid todo",
			fields: fields{
				ID:          1,
				Name:        "",
				Description: "test",
				Tasks: []Task{
					{
						ID:          1,
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
				ID:          1,
				Name:        "test",
				Description: "test",
				Tasks: []Task{
					{
						ID:          1,
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
				ID:          1,
				Name:        "test",
				Description: "test",
				Tasks: []Task{
					{
						ID:          1,
						Name:        "test",
						Description: "test",
						TodoID:      1,
					},
					{
						ID:          1,
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
				ID:          1,
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
				ID:          tt.fields.ID,
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

func TestTodo_PrettyString(t *testing.T) {
	todo := &Todo{}

	marschalIdent, err := json.MarshalIndent(todo, "", "  ")
	assert.NoError(t, err)

	prettyString, err := todo.PrettyString("  ")
	assert.NoError(t, err)

	assert.Equal(t, string(marschalIdent), prettyString)
}
