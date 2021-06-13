package fgeom

type Insets struct {
	Top, Right, Bottom, Left float32
}

func HomogeneousInsets(inset float32) Insets {
	return Insets{Top: inset, Right: inset, Bottom: inset, Left: inset}
}
