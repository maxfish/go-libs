package geom

type Alignment uint32

const (
	AlignmentHCenter Alignment = 1 << iota
	AlignmentHLeft
	AlignmentHRight
	AlignmentVCenter
	AlignmentVTop
	AlignmentVBottom

	AlignmentNone Alignment = 0
	AlignmentCenter = AlignmentHCenter | AlignmentVCenter
)

type FitMode int

const (
	FitModeAlign FitMode = iota
	FitModeFill
	FitModeAspectFit
	FitModeAspectFill
)
