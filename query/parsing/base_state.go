package parsing

// BaseState provides base functionality to parse states.
type BaseState struct {
	nextStates []State
}

func (s *BaseState) SetNext(nextStates ...State) {
	s.nextStates = nextStates
}

func (s *BaseState) Next() []State {
	return s.nextStates
}
