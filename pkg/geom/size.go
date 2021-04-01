package geom

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/math"
)

type Size struct {
	W, H int
}

func (s Size) Scale(factor float32) Size {
	return Size{
		W: int(float32(s.W) * factor),
		H: int(float32(s.H) * factor),
	}
}

func (s Size) Union(other Size) Size {
	return Size{
		W: math.MaxI(s.W, other.W),
		H: math.MaxI(s.H, other.H),
	}
}

func (s Size) ToString() string {
	return fmt.Sprintf("{w:%d,h:%d}", s.W, s.H)
}
