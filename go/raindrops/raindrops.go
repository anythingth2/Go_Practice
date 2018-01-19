package raindrops

import (
	"fmt"
)

type RainSound struct {
	num   int
	sound string
}

var rains []RainSound = []RainSound{
	{
		3,
		"Pling",
	}, {
		5,
		"Plang",
	}, {
		7,
		"Plong",
	},
}

func Convert(num int) string {
	var song string = ""
	for _, rain := range rains {
		if num%rain.num == 0 {
			song += rain.sound
		}
	}
	if song != "" {
		return song
	} else {
		return fmt.Sprintf("%d", num)
	}
}
