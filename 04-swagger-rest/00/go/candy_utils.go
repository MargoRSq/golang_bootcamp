package main

func fillCandies() (candies map[string]Candy) {
	candies = map[string]Candy{"CE": {"Cool Eskimo", 10},
		"AA": {"Apricot Aardvark", 15},
		"NT": {"Natural Tiger", 17},
		"D":  {"Dazzling", 21},
		"YR": {"Yellow Rambutan", 23},
	}
	return
}
