package fractx

import "sort"

func ImageBuilderNames() (ss []string) {
	ss = make([]string, len(ImageBuilders))
	i := 0
	for k := range ImageBuilders {
		ss[i] = k
		i++
	}
	sort.Strings(ss)
	return
}
