package password

import (
	"testing"
)

func TestWord_GetTip(t *testing.T) {
	tests := []struct {
		name string
		w    *Word
		want string
	}{
		{
			name: "Get tip 0",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{},
				Word:   "Word",
			},
			want: "",
		},
		{
			name: "Get tip 1",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{"Tip 1", "Tip 2", "Tip 3"},
				Word:   "Word",
			},
			want: "Tip 1",
		},
		{
			name: "Get tip 2",
			w: &Word{
				Count:  1,
				Points: 0,
				Tips:   []string{"Tip 1", "Tip 2", "Tip 3"},
				Word:   "Word",
			},
			want: "Tip 2",
		},
		{
			name: "Get tip 3",
			w: &Word{
				Count:  2,
				Points: 0,
				Tips:   []string{"Tip 1", "Tip 2", "Tip 3"},
				Word:   "Word",
			},
			want: "Tip 3",
		},
		{
			name: "Get tip 4",
			w: &Word{
				Count:  2,
				Points: 0,
				Tips:   []string{"Tip 1", "Tip 2", "Tip 3"},
				Word:   "Word",
			},
			want: "Tip 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.GetTip(); got != tt.want {
				t.Errorf("Word.GetTip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_CalculatePoints(t *testing.T) {
	tests := []struct {
		name string
		w    *Word
		want int
	}{
		{
			name: "0 tips return 100 points",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{},
				Word:   "Word",
			},
			want: 100,
		},
		{
			name: "1 tip return 100 points",
			w: &Word{
				Count:  1,
				Points: 0,
				Tips:   []string{"word1"},
				Word:   "Word",
			},
			want: 100,
		},
		{
			name: "2 tips return 50 points",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{"word1", "word2"},
				Word:   "Word",
			},
			want: 50,
		},
		{
			name: "3 tips return 33 points",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{"word1", "word2", "word3"},
				Word:   "Word",
			},
			want: 33,
		},
		{
			name: "4 tips return 25 points",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{"word1", "word2", "word3", "word4"},
				Word:   "Word",
			},
			want: 25,
		},
		{
			name: "5 tips return 20 points",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{"word1", "word2", "word3", "word4", "word5"},
				Word:   "Word",
			},
			want: 20,
		},
		{
			name: "6 tips return 16 points",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{"word1", "word2", "word3", "word4", "word5", "word6"},
				Word:   "Word",
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.CalculatePoints(); got != tt.want {
				t.Errorf("Word.CalculatePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_Check(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		w    *Word
		args args
		want bool
	}{
		{
			name: "Check true",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{},
				Word:   "Word",
			},
			args: args{
				word: "Word",
			},
			want: true,
		},
		{
			name: "Check false",
			w: &Word{
				Count:  0,
				Points: 0,
				Tips:   []string{},
				Word:   "Word",
			},
			args: args{
				word: "Wrong",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Check(tt.args.word); got != tt.want {
				t.Errorf("Word.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
