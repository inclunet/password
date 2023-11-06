package password

import "time"

type Player struct {
	Current    bool
	CurrentTip int
	Duration   time.Duration
	Player     int
	Points     int
	Round      int
	StartTime  time.Time
	Solved     bool
	Words      [5]Word
}

func (p *Player) AddWord(index int, word Word) {
	if index >= 0 && index < 5 {
		p.Words[index] = word
	}
}

func (p *Player) GetWord(word int) *Word {
	if word >= 0 && word < 5 {
		return &p.Words[word]
	}

	return nil
}

func (p *Player) SetTip(tip string) {
	if tip != "" {
		p.GetWord(p.CurrentTip)
	}
}

func NewPlayer(round, player int) (newPlayer Player) {
	newPlayer.Player = player
	newPlayer.Round = round
	newPlayer.StartTime = time.Now()
	return newPlayer
}
