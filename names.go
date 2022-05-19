package fractx

import "sort"

func ImageGeneratorNames() (ss []string) {
	ss = make([]string, len(ImageGenerators))
	i := 0
	for k := range ImageGenerators {
		ss[i] = k
		i++
	}
	sort.Strings(ss)
	return
}
