package bracket

import Stack "algorithms-in-go/structure/stack"

type bracket struct {
	content string
}

func (b *bracket) isBalanced() bool {
	var stack = Stack.NewStack()
	for _, r := range b.content {
		c := string(r)
		if isOpenBracket(c) {
			stack.Push(c)
		} else if isCloseBracket(c) {
			tmp, _ := stack.Pop().(string)
			if !isPair(tmp, c) {
				stack.Push(tmp)
			}
		}
	}
	return stack.Size() == 0
}

var pairBracketsMap = map[string]string{
	"[": "]",
	"{": "}",
	"(": ")",
}

func isOpenBracket(c string) bool {
	return (c == "(" || c == "{" || c == "[")
}

func isCloseBracket(c string) bool {
	return (c == ")" || c == "}" || c == "]")
}

func isPair(a string, b string) bool {
	return (pairBracketsMap[a] == b)
}
