package jug

type Op func(plan *ExecutionPlan, frame DataFrame) DataFrame

func Sum(plan *ExecutionPlan, frame DataFrame) DataFrame {
	r := make(DataFrame)
	r["result"] = make(map[string]interface{})
	for _, columns := range frame {
		for column, value := range columns {
			v := value.(ExecutionItem)
			switch plan.Type {
			case TypeFloat:
				if _, ok := r["result"][column]; !ok {
					r["result"][column] = 0.0
				}
				c := r["result"][column].(float64)
				c += v.f
				r["result"][column] = c
			}
		}
	}
	return r
}

func Count(plan *ExecutionPlan, frame DataFrame) DataFrame {
	r := make(DataFrame)
	r["result"] = make(map[string]interface{})
	count := 0.0
	for range frame {
		count++
	}
	r["result"]["count"] = count
	return r
}

func Average(plan *ExecutionPlan, frame DataFrame) DataFrame {
	sumFrame := Sum(plan, frame)
	countFrame := Count(plan, frame)
	for column, value := range sumFrame["result"] {
		v := value.(float64)
		sumFrame["result"][column] = v / countFrame["result"]["count"].(float64)
	}
	return sumFrame
}
