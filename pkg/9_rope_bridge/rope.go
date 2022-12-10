package rope_bridge

import "github.com/hvpeteet/advent22/pkg/fuck_go"

type Rope struct {
	Knots []Vector2
}

func (r *Rope) Move(move Vector2) {
	r.Knots[0] = r.Knots[0].Add(move)
	for i := 1; i < len(r.Knots); i++ {
		r.Knots[i] = adjustTail(r.Knots[i-1], r.Knots[i])
	}
}

func NewRope(length int) Rope {
	rope := Rope{
		Knots: []Vector2{},
	}
	for i := 0; i < length; i++ {
		rope.Knots = append(rope.Knots, Vector2{})
	}
	return rope
}

func adjustTail(head Vector2, tail Vector2) Vector2 {
	diff := head.Subtract(tail)
	tailMoves := Vector2{}
	if fuck_go.IntAbs(diff.X) <= 1 && fuck_go.IntAbs(diff.Y) <= 1 {
		return tail
	}
	if fuck_go.IntAbs(diff.X)+fuck_go.IntAbs(diff.Y) > 2 {
		// Tail moves diagonally
		tailMoves.X = fuck_go.IntSign(diff.X)
		tailMoves.Y = fuck_go.IntSign(diff.Y)
	} else {
		// Tail moves straight
		if fuck_go.IntAbs(diff.X) == 2 {
			tailMoves.X = diff.X / 2
		}
		if fuck_go.IntAbs(diff.Y) == 2 {
			tailMoves.Y = diff.Y / 2
		}
	}
	if fuck_go.IntAbs(diff.X) > 2 || fuck_go.IntAbs(diff.Y) > 2 {
		panic("Move was made that was illegal")
	}
	return tail.Add(tailMoves)
}
