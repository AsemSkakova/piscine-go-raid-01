package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}

		z01.PrintRune(10)

		for i := 0; i < y-2; i++ {
			for j := 0; j < x; j++ {
				if j == 0 || j == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(10)
		}

		if y != 1 {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
		}
		z01.PrintRune(10)
	}
}
