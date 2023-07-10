package entity

import (
	"fmt"
	"vientiane/server/consts"
)

type Condition struct {
	Field  string      `json:"field"`
	OpType string      `json:"op_type"`
	Value  interface{} `json:"value"`
}

type Conditions struct {
	Musts   []Conditions `json:"musts"`
	Should  []Conditions `json:"should"`
	MustNot []Conditions `json:"must_not"`
	Cond    *Condition   `json:"cond"`
}

type ESQueryParam struct {
	From   int64             `json:"from,omitempty"`
	Size   int64             `json:"size,omitempty"`
	Query  *ESQueryParamBool `json:"query,omitempty"`
	Sort   []string          `json:"sort,omitempty"`
	Source []string          `json:"_source,omitempty"`
}

type ESQueryParamBool struct {
	Bool *ESBoolQueryParamBool `json:"bool"`
}

type ESBoolQueryParamBool struct {
	Filter  []map[string]interface{} `json:"filter"`
	Should  []map[string]interface{} `json:"should"`
	MustNot []map[string]interface{} `json:"must_not"`
}

// 将条件解析为 es 查询条件
func ParseToES(cond *Conditions) *ESQueryParamBool {
	query := &ESQueryParamBool{
		Bool: &ESBoolQueryParamBool{
			Filter:  parseToESParam(cond.Musts),
			Should:  parseToESParam(cond.Should),
			MustNot: parseToESParam(cond.MustNot),
		},
	}
	return query
}

func parseToESParam(conds []Conditions) []map[string]interface{} {
	if len(conds) == 0 {
		return nil
	}

	esConds := make([]map[string]interface{}, 0)
	for _, cond := range conds {
		m := make(map[string]interface{})

		if len(cond.Musts) > 0 || len(cond.Should) > 0 || len(cond.MustNot) > 0 {
			m["bool"] = &ESBoolQueryParamBool{
				Filter:  parseToESParam(cond.Musts),
				Should:  parseToESParam(cond.Should),
				MustNot: parseToESParam(cond.MustNot),
			}
		}

		if cond.Cond != nil {
			switch cond.Cond.OpType {
			case consts.ESOpTypeEq:
				m["term"] = map[string]interface{}{
					cond.Cond.Field: cond.Cond.Value,
				}
			case consts.ESOpTypeExist:
				m["exists"] = map[string]interface{}{
					cond.Cond.Field: cond.Cond.Value,
				}
			case consts.ESOpTypeIn:
				m["terms"] = map[string]interface{}{
					cond.Cond.Field: cond.Cond.Value,
				}
			case consts.ESOpTypeGt, consts.ESOpTypeGte, consts.ESOpTypeLt, consts.ESOpTypeLte:
				m["range"] = map[string]interface{}{
					cond.Cond.Field: map[string]interface{}{cond.Cond.OpType: cond.Cond.Value},
				}
			case consts.ESOpTypeRegexp:
				m["regexp"] = map[string]interface{}{
					cond.Cond.Field: map[string]interface{}{"value": fmt.Sprintf(".*%v.*", cond.Cond.Value), "flags": "ALL"},
				}
			}
		}
		esConds = append(esConds, m)
	}

	return esConds
}
