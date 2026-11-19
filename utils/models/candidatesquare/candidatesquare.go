package candidatesquare

import ()

type CandidateSquare struct {
	Rank int
	File int
}

func (c *CandidateSquare) IsValidSquare() bool {
	if c.Rank >= 0 && c.Rank < 7 && c.File >= 0 && c.File < 8 {
		return true
	}
	return false
}
