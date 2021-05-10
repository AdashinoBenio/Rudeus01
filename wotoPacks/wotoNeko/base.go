package wotoNeko

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
)

const (
	hostUrl = "https://nekos.life/api/v2"
	img     = hostUrl + "/img/"
	jpg     = ".jpg"
	jpeg    = ".jpeg"
	gif     = ".gif"
	png     = ".png"
)

type NekoBase struct {
	Url string `json:"url"`
}

func GetRandomSfw() (*NekoBase, error) {
	ep := getRandomSfw()
	return getRandomThing(string(ep))
}

func GetRandomNsfw() (*NekoBase, error) {
	ep := getRandomNSfw()
	return getRandomThing(string(ep))
}

func GetRandomTickle() (*NekoBase, error) {
	return getRandomThing(string(Tickle))
}

func GetRandomSlap() (*NekoBase, error) {
	return getRandomThing(string(Slap))
}

func GetRandomPoke() (*NekoBase, error) {
	return getRandomThing(string(Poke))
}

func GetRandomPat() (*NekoBase, error) {
	return getRandomThing(string(Pat))
}

func GetRandomNeko() (*NekoBase, error) {
	return getRandomThing(string(Neko))
}

func GetRandomMeow() (*NekoBase, error) {
	return getRandomThing(string(Meow))
}

func GetRandomLizard() (*NekoBase, error) {
	return getRandomThing(string(Lizard))
}

func GetRandomKiss() (*NekoBase, error) {
	return getRandomThing(string(Kiss))
}

func GetRandomHug() (*NekoBase, error) {
	return getRandomThing(string(Hug))
}

func GetRandomFoxGirl() (*NekoBase, error) {
	return getRandomThing(string(Fox_Girl))
}

func GetRandomFeed() (*NekoBase, error) {
	return getRandomThing(string(Feed))
}

func GetRandomCuddle() (*NekoBase, error) {
	return getRandomThing(string(Cuddle))
}

func getRandomSfw() sfwNeko {
	initSfw()
	return sfw[rand.Intn(len(sfw))]
}

func getRandomNSfw() nsfwNeko {
	initNsfw()
	return nsfw[rand.Intn(len(nsfw))]
}

func getRandomThing(endpoint string) (base *NekoBase, err error) {
	addr := img + endpoint
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	b, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		// log.Println("In Resp.Body ERR!" + err.Error())
		return nil, errBody
	}

	neko := NekoBase{}
	err = json.Unmarshal(b, &neko)
	if err != nil {
		// log.Println("In UnMarshal ERR!" + err.Error())
		return nil, err
	}

	// log.Println("Final neko url is:" + neko.Url)
	return &neko, nil
}
