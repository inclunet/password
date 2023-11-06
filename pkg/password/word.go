package password

type Word struct {
	Count  int
	Points int
	Solved bool
	Tips   []string
	Word   string
}

func (w *Word) CalculatePoints() int {
	if len(w.Tips) == 0 {
		w.Points = 100
	} else {
		w.Points = 100 / len(w.Tips)
	}

	return w.Points
}

func (w *Word) Check(word string) bool {
	if w.Word == word {
		w.Solved = true
		return true
	}

	return false
}

func (w *Word) GetTip() string {
	if len(w.Tips) == 0 {
		return ""
	}

	if w.Count > len(w.Tips) {
		return w.Tips[len(w.Tips)-1]
	}

	tip := w.Count

	if w.Count <= len(w.Tips) {
		w.Count++
		w.CalculatePoints()
	}

	return w.Tips[tip]
}

func (w *Word) SetTip(tip string) {
	if tip != "" {
		w.Tips = append(w.Tips, tip)
	}
}
