package voice

import (
	_ "embed"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"log"
	"time"
)

//go:embed ding.mp3
var ding []byte
var context *oto.Context
var data []byte

func Init() {
	var dec *minimp3.Decoder
	var err error
	if dec, data, err = minimp3.DecodeFull(ding); err != nil {
		log.Fatal(err)
	}

	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024); err != nil {
		log.Fatal(err)
	}
}

func Play() {
	var player = context.NewPlayer()
	player.Write(data)

	<-time.After(time.Second)

	if err := player.Close(); err != nil {
		log.Fatal(err)
	}
}
