package fuck_go

type Vector2 struct {
	X int
	Y int
}

var Up = Vector2{
	Y: 1,
}

var Down = Vector2{
	Y: -1,
}

var Left = Vector2{
	X: -1,
}

var Right = Vector2{
	X: 1,
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vector2) Manhattan() int {
	return IntAbs(v1.X) + IntAbs(v1.Y)
}

func (v1 Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}
