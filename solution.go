package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type shape int

const (
	SidesCircle   shape = 0
	SidesTriangle shape = 3
	SidesSquare   shape = 4
)

func CalcSquare(sideLen float64, sidesNum shape) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case 3:
		return ((math.Sqrt(3) / 2) * sideLen * sideLen) / 2
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
