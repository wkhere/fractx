package fractx

type Colorset []struct {
	IterPercent uint
	Color       string
}

var (
	Colorset1 = Colorset{
		{70, "white"},
		{40, "yellow"},
		{24, "orange"},
		{12, "red"},
		{7, "purple"},
		{3, "darkblue"},
		{0, "blue"},
	}
)

func (a Colorset) Len() int { return len(a) }

func (a Colorset) Less(i, j int) bool {
	return a[i].IterPercent > a[j].IterPercent
}

func (a Colorset) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
