package pgengine

import "math"

type Tween struct {
	isFinished    bool
	current       Float
	startValue    Float
	distance      Float
	timePassed    Float
	totalDuration Float
	tweenF        TweenFunc
}

func NewTween(start, finish, totalDuration Float, tweenF TweenFunc) *Tween {
	return &Tween{
		tweenF:        tweenF,
		distance:      finish - start,
		startValue:    start,
		current:       start,
		totalDuration: totalDuration,
		isFinished:    false,
	}
}

func (t *Tween) IsFinished() bool {
	return t.isFinished
}

func (t *Tween) Value() Float {
	return t.current
}

func (t *Tween) FinishValue() Float {
	return t.startValue + t.distance
}

func (t *Tween) Update(elapsedTime Float) {
	t.timePassed = t.timePassed + (elapsedTime)
	t.current = t.tweenF(t.timePassed, t.startValue, t.distance, t.totalDuration)

	if t.timePassed > t.totalDuration {
		t.current = t.startValue + t.distance
		t.isFinished = true
	}
}

type TweenFunc func(t, b, c, d Float) Float

func EaseInQuat(t, b, c, d Float) Float {
	t = t / d
	return c*t*t + b
}

func EaseOutQuad(t, b, c, d Float) Float {
	t = t / d
	return -c*t*(t-2) + b
}

func EaseInCirc(t, b, c, d Float) Float {
	t = t / d
	return -c*(math.Sqrt(1-t*t)-1) + b
}

func EaseOutCirc(t, b, c, d Float) Float {
	t = t/d - 1
	return c*math.Sqrt(1-t*t) + b
}

func EaseOutInCirc(t, b, c, d Float) Float {
	if t < d/2 {
		return EaseOutCirc(t*2, b, c/2, d)
	}
	return EaseInCirc((t*2)-d, b+c/2, c/2, d)
}

func EaseInExpo(t, b, c, d Float) Float {
	return c*math.Pow(2, 10*(t/d-1)) + b
}

func EaseInBounce(t, b, c, d Float) Float {
	return c - EaseOutBounce(d-t, 0, c, d) + b
}

func EaseOutBounce(t, b, c, d Float) Float {
	t = t / d
	if t < (1 / 2.75) {
		return c*(7.5625*t*t) + b
	}
	if t < (2 / 2.75) {
		t = t - (1.5 / 2.75)
		return c*(7.5625*t*t+.75) + b
	}
	if t < (2.5 / 2.75) {
		t = t - (2.25 / 2.75)
		return c*(7.5625*t*t+.9375) + b
	}
}

func EaseInOutBounce(t, b, c, d Float) Float {
	if t < d/2 {
		return EaseInBounce(t*2, 0, c, d)*.5 + b
	} else {
		EaseOutBounce(t*2-d, 0, c, d)*.5 + c*.5 + b
	}
}

func Linear(t, b, c, d Float) Float {
	return c*t/d + b
}
