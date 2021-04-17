package rand

type RandomNumberGenerator interface {
	NextUint32() (result uint32)
	NextUint32LessThan(n int) (result uint32)
	NextUint32InRange(min, max int) (result uint32)
	NextFloat32() (result float32)
	Event(chance int) bool
	Maybe() bool
}
