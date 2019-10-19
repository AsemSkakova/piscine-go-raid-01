package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {

	for i := 0; i < x-1; i++ {
		if i > 0 && i < x-2 {
			z01.PrintRune('-')
		} else {
			z01.PrintRune('o')
		}
	}

	for j := 0; j < y-1; j++ {
		if j > 0 && j < y-2 {
			z01.PrintRune('|')
		} else {
			z01.PrintRune('o')
		}
	}

	// for i := 0; i < x; i++ {
	// 	var aRune string = "o---o"
	// 	z01.PrintRune(rune(aRune[i]))
	// }

	// for j := 0; j < y-2; j++ {
	// 	z01.PrintRune(10)
	// 	var aRune string = "|||"
	// 	z01.PrintRune(rune(aRune[j]))
	// 	for k := 0; k <= x-2; k++ {
	// 		if k <= x-2 && k >= x-2 {
	// 			z01.PrintRune('|')
	// 		} else {
	// 			z01.PrintRune(' ')
	// 		}
	// 	}
	// }

	// z01.PrintRune(10)

	// for i := 0; i < x; i++ {
	// 	var aRune string = "o---o"
	// 	z01.PrintRune(rune(aRune[i]))
	// }

	// z01.PrintRune(10)

}
