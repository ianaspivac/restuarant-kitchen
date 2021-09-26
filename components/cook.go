package components

type Cook struct {
	Rank        int
	Proficiency int
	Name        string
	CatchPhrase string
}
type Rank int

func HireCook(rank int) *Cook {
	return &Cook{
		Rank:        rank,
		Proficiency: rank,
		Name: "John Johnson",
		CatchPhrase: "No time for talking!",
	}
}
