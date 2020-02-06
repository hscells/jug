package jug

import (
	"errors"
	"strings"
)

type StatementPosition int

const (
	StatementOp StatementPosition = iota
	StatementType
)

type OpToken string
type TypeToken string

var (
	OpSum   OpToken = "sum"
	OpCount OpToken = "count"
	OpAvg   OpToken = "avg"
)

var (
	TypeFloat TypeToken = "float"
)

func Parse(query string) (*ExecutionPlan, error) {
	plan := new(ExecutionPlan)

	// Default to float
	plan.Type = TypeFloat

	tokens := strings.Split(query, " ")
	for i, token := range tokens {
		switch StatementPosition(i) {
		case StatementOp:
			switch OpToken(token) {
			case OpSum:
				plan.Op = Sum
			case OpCount:
				plan.Op = Count
			case OpAvg:
				plan.Op = Average
			default:
				return nil, errors.New("invalid operation")
			}
		case StatementType:
			switch t := TypeToken(token); t {
			case TypeFloat:
				plan.Type = t
			default:
				return nil, errors.New("invalid type")
			}
		default:
			plan.Columns = append(plan.Columns, token)
		}
	}
	return plan, nil
}
