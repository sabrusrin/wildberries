package sushi

type visitor interface {
}

type buySushi interface {
	buySushi
}

type SushiBar interface {
	Accept(v visitor) string
}

// SushiBar implements the Place interface.
type sushiBar struct {
}

// Accept implementation.
func (s *sushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi implementation.
func (s *sushiBar) BuySushi() string {
	return "Buy sushi..."
}

func NewBar() SushiBar {
	return &sushiBar{}
}
