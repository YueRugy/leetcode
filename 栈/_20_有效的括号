package 栈
//link https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if (len(s) & 1) == 1 {
		return false
	}
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)
	for _, c := range s {
		if ch, ok := m[c]; ok { //说明是右括号
			//判断stack是否为空
			if len(stack) == 0 {
				return false
			}
			//判断括号是否匹配
			if ch != stack[len(stack)-1] {
				return false
			}
			//弹出栈顶元素
			stack = stack[:len(stack)-1]
			continue
		}
		//左边括号入栈
		stack = append(stack, c)
	}
	//如果栈空说明合法不为空 还有没有匹配到的括号不合法
	return len(stack) == 0
}

