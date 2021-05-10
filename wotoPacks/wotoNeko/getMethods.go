package wotoNeko

import "strings"

func (n *NekoBase) IsPhoto() bool {
	tmp := strings.ToLower(n.Url)
	b1 := strings.HasSuffix(tmp, jpg)
	b2 := strings.HasSuffix(tmp, jpeg)
	b3 := strings.HasSuffix(tmp, png)
	return b1 || b2 || b3
}

func (n *NekoBase) IsGif() bool {
	tmp := strings.ToLower(n.Url)
	return strings.HasSuffix(tmp, gif)
}
