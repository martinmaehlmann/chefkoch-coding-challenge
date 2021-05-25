package todo

import (
	"testing"
)

func TestTask_Valid(t1 *testing.T) {
	type fields struct {
		ID          uint
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
				ID:          1,
				Name:        "test",
				Description: "test",
				TodoID:      1,
			},
			want: true,
		},
		{
			name: "test invalid task",
			fields: fields{
				ID:          1,
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
				ID:          tt.fields.ID,
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
