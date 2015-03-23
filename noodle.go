package main

import "github.com/krig/go-sox"
import "log"

// Attempting to get a sox 'null' filestream open and apply the synth effect
// This spits out some undefined errors
// It's late and we must rest though, so here it lie for now
// TODO resolve the undefined errors
//      apply some other effects
//

func main() {
	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	defer sox.Quit()

	in := sox.OpenRead("null")
	if in == nil {
		log.Fatal("Failed to open test-input.wav")
	}

	out := sox.OpenWrite("default", in.Signal(), nil, "alsa")
	if out == nil {
		out = sox.OpenWrite("default", in.Signal(), nil, "pulseaudio")
		if out == nil {
			log.Fatal("Failed to open output device")
		}
	}

	chain := sox.CreateEffectsChain(in.Encoding(), out.Encoding())

	e := sox.CreateEffect(sox.FindEffect("synth"))
	e.Options(in)
	chain.AddEffect(e, in.Signal(), in.Signal())
	e.Free()

	chain.FlowEffects()

	chain.Delete()
	out.Close()
	in.Close()
}
