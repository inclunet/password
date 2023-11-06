package password

type Round struct {
	currentPlayer int
	currentWord   int
	Players       []Player
	Playing       bool
	Round         int
	Solo          bool
	Words         []Word
}

func (r *Round) AddPlayer() *Player {
	if !r.Playing {
		player := NewPlayer(r.Round, len(r.Players)+1)
		r.Players = append(r.Players, player)
		return &player
	}

	return &Player{}
}

func (r *Round) GetPlayer(index int) *Player {
	if index >= 0 && index <= len(r.Players) {
		player := &r.Players[index-1]

		if player.Current {
			Word, ok := r.GetWord(r.currentWord)

			if ok {
				player.GetWord().SetTip(Word.GetTip())

			}
		}

		return player
	}

	return &Player{}
}

func (r *Round) GetWord(index int) (word *Word, ok bool) {
	if index >= 0 && index < len(r.Words) {
		word = &r.Words[index]
		ok = true
	}

	return word, ok
}

func NewRound(round int, solo bool) (newRound Round) {
	newRound.Round = round
	newRound.Solo = solo

	if solo {
		newRound.AddPlayer()
		newRound.Playing = true
	}

	return newRound
}
