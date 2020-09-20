/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

// 来自题解
func deserialize(s string) *NestedInteger {
	// 如果开头是数字
	if s[0] != '[' {
		var res NestedInteger
		num, _ := strconv.Atoi(s)
		res.SetInteger(num)
		return &res
	}

	stack := make([]NestedInteger, 0)
	sign, num, isNum := 1,0,false
	for _, char := range s {
		if char == '[' {
			stack = append(stack, NestedInteger{})
		} else if char == ']' {
			// 处理一个[]
			if isNum {
				var tmp NestedInteger
				tmp.SetInteger(sign*num)
				stack[len(stack)-1].Add(tmp)
				sign, num, isNum = 1, 0, false
			}

			if len(stack) > 1 {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].Add(pop)
			}
		} else if char == ',' {
			if isNum {
				var tmp NestedInteger
				tmp.SetInteger(sign*num)
				stack[len(stack)-1].Add(tmp)
				sign,  num, isNum = 1, 0, false
			}
		} else if char == '-' {
			sign = -1
		} else {
			num = int(char-'0') + num*10
			isNum = true
		}
	}
	return &stack[0]
}