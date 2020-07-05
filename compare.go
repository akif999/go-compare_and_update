package compare

type bar struct {
	foos []foo
}

type foo struct {
	id  int
	val int
}

func CompareAndUpdate(src, comparision *bar) (*bar, error) {
	// make map of comparision
	idMap := make(map[int]int, 100000)
	for _, f := range comparision.foos {
		idMap[f.id] = f.val
	}

	// compare and update val
	newBar := &bar{}
	for _, f := range src.foos {
		val, find := idMap[f.id]
		newFoo := foo{}
		newFoo.id = f.id
		newFoo.val = f.val
		if find {
			newFoo.val = val
		}
		newBar.foos = append(newBar.foos, newFoo)
	}
	return newBar, nil
}
