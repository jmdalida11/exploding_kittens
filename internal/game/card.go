package game

type Card string

const (
	ExplodingKitten    Card = "ExplodingKitten"
	Defuse             Card = "Defuse"
	Nope               Card = "Nope"
	SeeTheFuture       Card = "SeeTheFuture"
	Attack             Card = "Attack"
	Skip               Card = "Skip"
	Favor              Card = "Favor"
	Shuffle            Card = "Shuffle"
	Tacocat            Card = "Tacocat"
	Cattermelon        Card = "Cattermelon"
	HairyPotatoCat     Card = "HairyPotatoCat"
	BeardCat           Card = "BeardCat"
	RainbowRalphingCat Card = "RainbowRalphingCat"
)

func CardEffect(card Card, game *Game) {
	switch card {
	case SeeTheFuture:
		println(SeeTheFuture)
	case Attack:
		println(Attack)
	case Skip:
		println(Skip)
	case Favor:
		println(Favor)
	case Shuffle:
		println(Shuffle)
	default:
		println("not a valid card to activate")
	}
}
