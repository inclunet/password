package password

type Rounds struct {
	Rounds []Round
}

func (r *Rounds) AddRound(solo bool) Round {
	round := NewRound(len(r.Rounds)+1, solo)
	r.Rounds = append(r.Rounds, round)
	return round
}

func (r *Rounds) GetRound(round int) *Round {
	if round >= 0 && round <= len(r.Rounds) {
		return &r.Rounds[round-1]
	}

	return &Round{}
}
