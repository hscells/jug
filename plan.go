package jug

type ExecutionPlan struct {
	Op      Op
	Type    TypeToken
	Columns []string

	f float64
	s string
}

type ExecutionItem struct {
	f float64
	s string
}

type ExecutionResult struct {
	Type  TypeToken
	Frame DataFrame
}

type DataFrame map[string]map[string]interface{}

func (p *ExecutionPlan) Execute(frame DataFrame) (*ExecutionResult, error) {
	result := new(ExecutionResult)
	r := make(DataFrame)
	for row, v := range frame {
		if _, ok := r[row]; !ok {
			r[row] = make(map[string]interface{})
		}
		for column, obj := range v {
			doOp := false
			if len(p.Columns) == 0 {
				doOp = true
			} else {
				for _, col := range p.Columns {
					if col == column {
						doOp = true
						break
					}
				}
			}

			if doOp {
				p.f = 0

				var item ExecutionItem

				switch o := obj.(type) {
				case float64:
					if _, ok := r[row][column]; !ok {
						r[row][column] = 0.0
					}
					if p.Type == TypeFloat {
						item.f = o
					}
				}
				r[row][column] = item
			}
		}
	}

	result.Frame = p.Op(p, r)

	result.Type = p.Type
	return result, nil
}
