package password

import (
	"reflect"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		round  int
		player int
	}
	tests := []struct {
		name          string
		args          args
		wantNewPlayer Player
	}{
		{
			name:          "NewPlayer",
			args:          args{round: 1, player: 1},
			wantNewPlayer: Player{Player: 1, Round: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewPlayer := NewPlayer(tt.args.round, tt.args.player); !reflect.DeepEqual(gotNewPlayer, tt.wantNewPlayer) {
				t.Errorf("NewPlayer() = %v, want %v", gotNewPlayer, tt.wantNewPlayer)
			}
		})
	}
}

func TestPlayer_AddWord(t *testing.T) {
	type args struct {
		index int
		word  Word
	}
	tests := []struct {
		name string
		p    *Player
		args args
		Want Player
	}{
		{
			name: "AddWord",
			p:    &Player{Player: 1, Round: 1},
			args: args{index: 0, word: Word{Word: "test"}},
			Want: Player{Player: 1, Round: 1, Words: [5]Word{{Word: "test"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.AddWord(tt.args.index, tt.args.word)
			if !reflect.DeepEqual(tt.p, &tt.Want) {
				t.Errorf("AddWord() = %v, want %v", tt.p, tt.Want)
			}
		})
	}
}
