package parsing

import (
	"github.com/s2gatev/structql/query/ast"
)

type RootState struct{}

func (s *RootState) Next() []State {
	return []State{
		s.selectState(),
		s.updateState(),
	}
}

func (s *RootState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	return nil, true
}

func (s *RootState) selectState() *SelectState {
	offsetState := &OffsetState{}

	limitState := &LimitState{}
	limitState.SetNext(offsetState)

	whereState := &WhereState{}
	whereState.SetNext(limitState)

	fromState := &FromState{}
	fromState.SetNext(whereState, limitState)

	selectState := &SelectState{}
	selectState.SetNext(fromState)

	return selectState
}

func (s *RootState) updateState() *UpdateState {
	whereState := &WhereState{}

	setState := &SetState{}
	setState.SetNext(whereState)

	updateState := &UpdateState{}
	updateState.SetNext(setState)

	return updateState
}
