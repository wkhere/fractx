package main

type colorset []struct {
	iterPercent uint
	color       string
}

var (
	colorset1 = colorset{
		{70, "white"},
		{40, "yellow"},
		{24, "orange"},
		{12, "red"},
		{7, "purple"},
		{3, "darkblue"},
		{0, "blue"},
	}
)

func (a colorset) Len() int { return len(a) }

func (a colorset) Less(i, j int) bool {
	return a[i].iterPercent > a[j].iterPercent
}

func (a colorset) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
