package task1

/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func IsValidCharacter(s string) bool {
	flag := false
	if s == "" {
		return false
	}
	stack := []rune{}
	strMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, value := range s {
		switch value {
		case '(', '{', '[':
			//入栈
			stack = append(stack, value)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != strMap[value] {
				flag = false
				return flag
			}
			// 出栈
			stack = stack[:len(stack)-1]
		default:
			flag = false
			return flag
		}
	}
	flag = len(stack) == 0
	return flag
}
