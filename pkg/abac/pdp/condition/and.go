/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package condition

import (
	"fmt"
	"strings"

	"github.com/TencentBlueKing/bk-iam/pkg/abac/pdp/condition/operator"
	"github.com/TencentBlueKing/bk-iam/pkg/abac/pdp/types"
)

// AndCondition 逻辑AND
type AndCondition struct {
	baseLogicalCondition
}

func NewAndCondition(content []Condition) Condition {
	return &AndCondition{
		baseLogicalCondition{
			content: content,
		},
	}
}

func newAndCondition(field string, values []interface{}) (Condition, error) {
	if field != "content" {
		return nil, fmt.Errorf("and condition not support field %s", field)
	}

	conditions := make([]Condition, 0, len(values))

	for _, v := range values {
		condition, err := newConditionFromInterface(v)
		if err != nil {
			return nil, fmt.Errorf("and condition parser error: %w", err)
		}

		conditions = append(conditions, condition)
	}

	return &AndCondition{baseLogicalCondition{content: conditions}}, nil
}

// GetName 名称
func (c *AndCondition) GetName() string {
	return operator.AND
}

// Eval 求值
func (c *AndCondition) Eval(ctx types.EvalContextor) bool {
	for _, condition := range c.content {
		if !condition.Eval(ctx) {
			return false
		}
	}

	return true
}

func (c *AndCondition) Translate(withSystem bool) (map[string]interface{}, error) {
	content := make([]map[string]interface{}, 0, len(c.content))
	for _, ci := range c.content {
		ct, err := ci.Translate(withSystem)
		if err != nil {
			return nil, err
		}
		content = append(content, ct)
	}

	return map[string]interface{}{
		"op":      "AND",
		"content": content,
	}, nil
}

// PartialEval 使用传递的部分资源执行表达式, 并返回剩余的部分
func (c *AndCondition) PartialEval(ctx types.EvalContextor) (bool, Condition) {
	// NOTE: If allowed=False, condition should be nil
	// once got False=> return
	remainedContent := make([]Condition, 0, len(c.content))
	for _, condition := range c.content {
		switch condition.GetName() {
		case operator.AND, operator.OR:
			// if AND/OR, do PartialEval recursive
			ok, ci := condition.(LogicalCondition).PartialEval(ctx)
			// a AND b, if a false, return false
			if !ok {
				return false, nil
			}

			// 如果残留单独一个any, any always=True, 则没有必要append
			if ci.GetName() != operator.ANY {
				remainedContent = append(remainedContent, ci)
			}
		case operator.ANY:
			// if any, it's always true, just continue
			continue
		default:
			key := condition.GetKeys()[0]
			dotIdx := strings.LastIndexByte(key, '.')
			if dotIdx == -1 {
				// panic("should contain dot in key")
				return false, nil
			}
			_type := key[:dotIdx]

			if ctx.HasResource(_type) {
				// a AND b, if a false, return false
				if !condition.Eval(ctx) {
					return false, nil
				}
			} else {
				// request has no resource, so append to remain
				remainedContent = append(remainedContent, condition)
			}
		}
	}

	// all eval success, 全部执行成功, 理论上只剩any
	switch len(remainedContent) {
	case 0:
		// all true, return any
		return true, NewAnyCondition()
	case 1:
		return true, remainedContent[0]
	default:
		// more than one left, combine them all
		return true, NewAndCondition(remainedContent)
	}
}
