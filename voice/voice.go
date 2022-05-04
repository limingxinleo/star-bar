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

func Play() {
	var dec *minimp3.Decoder
	var data []byte
	var err error
	if dec, data, err = minimp3.DecodeFull(ding); err != nil {
		log.Fatal(err)
		return
	}

	var context *oto.Context
	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024); err != nil {
		log.Fatal(err)
	}

	var player = context.NewPlayer()
	player.Write(data)

	<-time.After(time.Second)

	dec.Close()
	if err = player.Close(); err != nil {
		log.Fatal(err)
	}
}
