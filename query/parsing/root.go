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
	offsetState := &OffsetState{
		NextStates: []State{},
	}

	limitState := &LimitState{
		NextStates: []State{
			offsetState,
		},
	}

	whereState := &WhereState{
		NextStates: []State{
			limitState,
		},
	}

	fromState := &FromState{
		NextStates: []State{
			whereState,
			limitState,
		},
	}

	return &SelectState{
		NextStates: []State{
			fromState,
		},
	}
}

func (s *RootState) updateState() *UpdateState {
	whereState := &WhereState{
		NextStates: []State{},
	}

	setState := &SetState{
		NextStates: []State{
			whereState,
		},
	}

	return &UpdateState{
		NextStates: []State{
			setState,
		},
	}
}
