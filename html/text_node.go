package html

type textNode string

func T(t string) textNode {
	return textNode(t)
}

func (n textNode) String() string {
	return string(n)
}
