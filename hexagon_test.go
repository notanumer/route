package main

import (
	"reflect"
	"testing"
)

func TestNewHexagonDrawer(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		want   *HexagonDrawer
	}{
		{
			name:   "basic hexagon",
			width:  3,
			height: 2,
			want:   &HexagonDrawer{width: 3, height: 2},
		},
		{
			name:   "minimal hexagon",
			width:  1,
			height: 1,
			want:   &HexagonDrawer{width: 1, height: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexagonDrawer(tt.width, tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexagonDrawer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexagonDrawer_Draw(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		want   []string
	}{
		{
			name:   "minimal hexagon 1x1",
			width:  1,
			height: 1,
			want: []string{
				"/ \\",
				"/___\\",
				"\\_/",
			},
		},
		{
			name:   "small hexagon 2x1",
			width:  2,
			height: 1,
			want: []string{
				"/  \\",
				"/____\\",
				"\\__/",
			},
		},
		{
			name:   "medium hexagon 3x2",
			width:  3,
			height: 2,
			want: []string{
				" /   \\",
				"/     \\",
				"/_______\\",
				"\\_____/",
				" \\___/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHexagonDrawer(tt.width, tt.height)
			if got := h.Draw(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexagonDrawer.Draw() = %v, want %v", got, tt.want)
				for i, line := range got {
					if i < len(tt.want) {
						t.Errorf("Line %d: got %q, want %q", i, line, tt.want[i])
					} else {
						t.Errorf("Line %d: got %q, want nothing", i, line)
					}
				}
			}
		})
	}
}

func BenchmarkHexagonDrawer_Draw(b *testing.B) {
	drawer := NewHexagonDrawer(10, 5)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		drawer.Draw()
	}
}