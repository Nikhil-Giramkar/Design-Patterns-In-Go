package main

type Snapshot struct {
	state string
}

func (s *Snapshot) getSavedState() string {
	return s.state
}
