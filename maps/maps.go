package maps

import "fmt"

func ShowMaps()  {
	countries := make(map[string]string)
	countries["Mexico"] = "DF"
	countries["Argentina"] = "Buenos Aires"
	fmt.Println(countries)
	fmt.Println(countries["Argentina"])
	fmt.Println()

	championship := map[string]int{
		"Lakers": 56,
		"Bulls": 12,
		"Knicks": 1,
		"Mavericks": 0,
	}
	fmt.Println(championship)
	fmt.Println()

	for team, points := range championship {
		fmt.Printf("Team %s has %d points.\n", team, points)
	}
	fmt.Println()

	delete(championship, "Mavericks")
	fmt.Println(championship)
	fmt.Println()

	points, exist := championship["Heat"]
	fmt.Printf("Team exists = %t with %d points.\n", exist, points)

	points, exist = championship["Lakers"]
	fmt.Printf("Team exists = %t with %d points.\n", exist, points)
}
