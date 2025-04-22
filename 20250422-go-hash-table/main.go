package main

import "fmt"

type Nameval struct {
	Name  string
	Value int32
	Next  *Nameval
}

const NHASH = 1337
const MULTIPLIER = 31

var symtab [NHASH]*Nameval

/* lookup: find name in symtab, with optional create */
func lookup(name string, create bool, value int32) *Nameval {
	var np *Nameval
	var h int

	/* hash the name */
	h = hash(name)

	/* search the list */
	for np = symtab[h]; np != nil; np = np.Next {
		if name == np.Name {
			return np
		}
	}

	if create {
		np = &Nameval{Name: name, Value: value, Next: symtab[h]}
		symtab[h] = np
	}
	return np
}

/* hash: compute hash value of string */
func hash(s string) int {
	var h int
	for _, c := range s {
		h = (h * MULTIPLIER) + int(c)
	}
	return h % NHASH
}

func main() {
	println("Hash table example")
	fmt.Printf("correct horse battery staple: %d\n", hash("correct horse battery staple"))
	fmt.Printf("example: %d\n", hash("example"))
	// Example usage
	fmt.Printf("Lookup returned: %d\n", lookup("example", true, 42).Value)
	fmt.Printf("Lookup returned: %d\n", lookup("example", false, 0).Value)

}
