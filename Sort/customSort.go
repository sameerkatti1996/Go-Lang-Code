//Aggregate data types can be sorted by using sort package.
//Three functions need to be defined: Len(), Swap(), Less()
package main

import (
	"fmt"
	"sort"
)

type footballer struct {
	age  int
	name string
	team string
}

//==============================SORT BY AGE=================================
//Custom type ByAge which has base type: slice of footballer
type ByAge []footballer

//Associate following functions to ByAge so that ByAge type
//can be sorted using sort package
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

//==============================SORT BY NAME=================================

//Custom type ByName which has base type: slice of footballer
type ByName []footballer

//Associate following functions to ByName so that ByName type
//can be sorted using sort package
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].name < a[j].name }

//==========================================================================

func main() {
	p1 := footballer{
		age:  35,
		name: "Cristiano Ronaldo",
		team: "Juventus",
	}

	p2 := footballer{
		age:  33,
		name: "Lionel Messi",
		team: "Barcelona",
	}

	p3 := footballer{
		age:  31,
		name: "Gareth Bale",
		team: "Real Madrid",
	}

	fmt.Println("Before sorting:")
	players := []footballer{p1, p2, p3}
	fmt.Println(players)

	//Convert slice of footballer into ByAge
	//Convertible as ByAge has base type []footballer
	//Since ByAge is associated to three functions needed for sorting,
	//we can sort using sort.Sort function

	//============== SORT BY AGE =====================
	sort.Sort(ByAge(players))
	fmt.Println("\nSort by age:\n", players)

	//Convert slice of footballer into ByName
	//Convertible as ByName has base type []footballer
	//Since ByName is associated to three functions needed for sorting,
	//we can sort using sort.Sort function

	//============== SORT BY NAME =====================
	sort.Sort(ByName(players))
	fmt.Println("\nSort by name:\n", players)
}

/*
Output:
Before sorting:
[{35 Cristiano Ronaldo Juventus} {33 Lionel Messi Barcelona} {31 Gareth Bale Real Madrid}]

Sort by age:
 [{31 Gareth Bale Real Madrid} {33 Lionel Messi Barcelona} {35 Cristiano Ronaldo Juventus}]

Sort by name:
 [{35 Cristiano Ronaldo Juventus} {31 Gareth Bale Real Madrid} {33 Lionel Messi Barcelona}]
*/
