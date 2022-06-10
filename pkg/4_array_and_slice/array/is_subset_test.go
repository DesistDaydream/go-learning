package array

import (
	"log"
	"testing"
)

func TestIsSubset(t *testing.T) {
	type args struct {
		subset   []string
		superset []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "测试1",
			args: args{
				subset:   []string{"a", "b", "c"},
				superset: []string{"a", "b", "c", "d", "e"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSubset(tt.args.subset, tt.args.superset)
			if got != tt.want {
				t.Errorf("IsSubset() = %v, want %v", got, tt.want)
			}
			log.Printf("命题: %v 是 %v 的子集.\n结论: %v\n", tt.args.subset, tt.args.superset, got)
		})
	}
}
