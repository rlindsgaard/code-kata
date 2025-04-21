package main

import "fmt"

type Nameval struct {
	Name  string
	Value rune
	left  *Nameval
	right *Nameval
}

func (treep *Nameval) Insert(newp *Nameval) *Nameval {
	if treep == nil {
		return newp
	}
	if newp.Name == treep.Name {
		fmt.Println("Duplicate name:", newp.Name)
	} else if newp.Name < treep.Name {
		treep.left = treep.left.Insert(newp)
	} else {
		treep.right = treep.right.Insert(newp)
	}
	return treep
}

func (treep *Nameval) Lookup(name string) *Nameval {
	if treep == nil {
		return nil
	}
	if name == treep.Name {
		return treep
	} else if name < treep.Name {
		return treep.left.Lookup(name)
	} else {
		return treep.right.Lookup(name)
	}
}

func (treep *Nameval) NrLookup(name string) *Nameval {
	for treep != nil {
		if name == treep.Name {
			return treep
		} else if name < treep.Name {
			treep = treep.left
		} else {
			treep = treep.right
		}
	}
	return nil
}

func (treep *Nameval) ApplyInOrder(fn func(*Nameval, any), arg any) {
	if treep == nil {
		return
	}
	treep.left.ApplyInOrder(fn, arg)
	fn(treep, arg)
	treep.right.ApplyInOrder(fn, arg)
}

func (treep *Nameval) ApplyPostOrder(fn func(*Nameval, any), arg any) {
	if treep == nil {
		return
	}
	treep.left.ApplyPostOrder(fn, arg)
	treep.right.ApplyPostOrder(fn, arg)
	fn(treep, arg)
}

func PrintItem(p *Nameval, arg any) {
	if p != nil {
		format := arg.(string)
		fmt.Printf(format, p.Name, p.Value)
	}
}

func main() {
	root := &Nameval{Name: "smiley", Value: 0x263A}
	root.Insert(&Nameval{Name: "Aacute", Value: 0x00c1})
	root.Insert(&Nameval{Name: "zeta", Value: 0x03b6})
	root.Insert(&Nameval{Name: "AElig", Value: 0x00c6})
	root.Insert(&Nameval{Name: "Aacute", Value: 0x00c1})
	root.Insert(&Nameval{Name: "Acirc", Value: 0x00c5})

	// Lookup some names
	if found := root.Lookup("Aacute"); found != nil {
		fmt.Printf("Found by Lookup: %s = %c\n", found.Name, found.Value)
	}
	if found := root.NrLookup("Aacute"); found != nil {
		fmt.Printf("Found by NrLookup: %s = %c\n", found.Name, found.Value)
	}

	root.ApplyInOrder(PrintItem, "Name: %s, Value: %c\n")
	root.ApplyPostOrder(PrintItem, "Name: %s, Value: %c\n")
}
