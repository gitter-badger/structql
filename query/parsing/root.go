package parsing

type RootState struct{}

func (s *RootState) Next() []State {
	return []State{
		&SelectState{},
	}
}

func (s *RootState) Parse(result Node, p *Parser) (Node, bool) {
	return nil, true
}
