package wotoNeko

type sfwNeko string

// sfw endpoints for nekos.life api.
const (
	Tickle   sfwNeko = "tickle"
	Slap     sfwNeko = "slap"
	Poke     sfwNeko = "poke"
	Pat      sfwNeko = "pat"
	Neko     sfwNeko = "neko"
	Meow     sfwNeko = "meow"
	Lizard   sfwNeko = "lizard"
	Kiss     sfwNeko = "kiss"
	Hug      sfwNeko = "hug"
	Fox_Girl sfwNeko = "fox_girl"
	Feed     sfwNeko = "feed"
	Cuddle   sfwNeko = "cuddle"
)

var sfw []sfwNeko

func initSfw() {
	if sfw == nil {
		sfw = []sfwNeko{
			Tickle,
			Slap,
			Poke,
			Pat,
			Neko,
			Meow,
			Lizard,
			Kiss,
			Hug,
			Fox_Girl,
			Feed,
			Cuddle,
		}
	}
}
