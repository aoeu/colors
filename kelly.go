package colors

var KellySafe map[string]RGB = map[string]RGB{
	"VividYellow":   *NewRGB(255, 179, 0),
	"StrongPurple":  *NewRGB(128, 62, 117),
	"VividOrange":   *NewRGB(255, 104, 0),
	"VeryLightBlue": *NewRGB(166, 189, 215),
	"VividRed":      *NewRGB(193, 0, 32),
	"GrayishYellow": *NewRGB(206, 162, 98),
	"MediumGray":    *NewRGB(129, 112, 102),
}

// Kelly colors that are not good for people
// with difficulty seeing some colors.
var KellyUnsafe map[string]RGB = map[string]RGB{
	"VividGreen":          *NewRGB(0, 125, 52),
	"StrongPurplishpink":  *NewRGB(246, 118, 142),
	"Strongblue":          *NewRGB(0, 83, 138),
	"StrongYellowishpink": *NewRGB(255, 122, 92),
	"StrongViolet":        *NewRGB(83, 55, 122),
	"VividOrangeYellow":   *NewRGB(255, 142, 0),
	"StrongPurplishRed":   *NewRGB(179, 40, 81),
	"VividGreenishYellow": *NewRGB(244, 200, 0),
	"StrongReddishBrown":  *NewRGB(127, 24, 13),
	"VividYellowishGreen": *NewRGB(147, 170, 0),
	"DeepYellowishBrown":  *NewRGB(89, 51, 21),
	"VividReddishOrange":  *NewRGB(241, 58, 19),
	"DarkOliveGreen":      *NewRGB(35, 44, 22),
}

func Kelly() map[string]RGB {
	k := make(map[string]RGB, len(KellySafe) + len(KellyUnsafe))
	for n, c := range KellySafe {
		k[n] = c
	}
	for n, c := range KellyUnsafe {
		k[n] = c
	}
	return k
}
