package kdy

import "calculus"

const Eps = 1e-7

type TwoVariable func(float64, float64) float64

func (f TwoVariable) Partial(x, y float64, direction int) float64 {
	switch direction {
	case 0:
		return (f(x+Eps, y) - f(x, y)) / Eps
	case 1:
		return (f(x, y+Eps) - f(x, y)) / Eps
	default:
		panic("TwoVariable.Partial: Direction must be 0 or 1")
	}
}

var _ calculus.RealValued = TwoVariable(nil)

func (f TwoVariable) Eval(p calculus.Point) float64 {
	return f(p[0], p[1])
}

func (f TwoVariable) Dim() int {
	return 2
}

func (f TwoVariable) Bounds() []float64 {
	return []float64{0, 1, 0, 1}
}

func (f TwoVariable) Contains(p calculus.Point) bool {
	return len(p) == 2
}

func (f TwoVariable) Map(p calculus.Point) calculus.Point {
	return calculus.Point{f.Eval(p)}
}

func Partial(f calculus.RealValued, p calculus.Point, direction int) float64 {
	originalValue := p[direction]               //p의 direction 좌표 값을 originalValue에 저장
	p[direction] = originalValue + Eps          // p[direction]에 eps를 더하여 좌표 값을 eps만큼 증가
	fxPlusEps := f.Eval(p)                      // f.Map(p): p에서의 함수 값을 계산하여 calculus.Point 타입으로 반환. (calculus.Point)[0]을 사용-> 함수 값의 첫 번째 요소를 fxPlusEps에 저장합니다.
	p[direction] = originalValue - Eps          // p[direction]에 eps를 더하여 좌표 값을 eps만큼 감소
	fxMinusEps := f.Eval(p)                     // f.Map(p): p에서의 함수 값을 계산하여 calculus.Point 타입으로 반환. (calculus.Point)[0]을 사용-> 함수 값의 첫 번째 요소를 fxPlusEps에 저장합니다.
	p[direction] = originalValue                //p[direction]을 originalValue로 복원-> p를 원래 상태 복귀
	return (fxPlusEps - fxMinusEps) / (2 * Eps) //두 함수 값의 차이를 2 * eps로 나누어 편미분을 근사
}
