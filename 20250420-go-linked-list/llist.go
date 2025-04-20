package main

import "fmt"

type Nameval struct {
	Name  string
	Value rune
	Next  *Nameval /* in list */
}

type Applier interface {
	string
}

func NewItem(name string, value rune) *Nameval {
	return &Nameval{
		Name:  name,
		Value: value,
	}
}

// Add newp to the front of the list
func AddFront(listp *Nameval, newp *Nameval) *Nameval {
	newp.Next = listp
	return newp
}

func AddEnd(listp *Nameval, newp *Nameval) *Nameval {
	if listp == nil {
		return newp
	}
	p := listp
	for p.Next != nil {
		p = p.Next
	}
	p.Next = newp
	return listp
}

func Lookup(listp *Nameval, name string) *Nameval {
	for p := listp; p != nil; p = p.Next {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func Apply(listp *Nameval, f func(*Nameval, any), arg any) {
	for p := listp; p != nil; p = p.Next {
		f(p, arg)
	}
}

func PrintItem(p *Nameval, arg any) {
	if p != nil {
		format := arg.(string)
		fmt.Printf(format, p.Name, p.Value)
	}
}

func IncCount(p *Nameval, arg any) {
	ip := arg.(*int)
	(*ip)++
}

/* delete first "name" from listp */
func DeleteItem(listp *Nameval, name string) *Nameval {
	var prev *Nameval
	for p := listp; p != nil; p = p.Next {
		if p.Name == name {
			if prev == nil {
				return p.Next
			}
			prev.Next = p.Next
			return listp
		}
		prev = p
	}
	return listp
}

func main() {
	var nvlist *Nameval
	nvlist = AddFront(nvlist, NewItem("smiley", 0x263A))

	Apply(nvlist, PrintItem, "Name: %s, Value: %c\n")
	n := 0
	Apply(nvlist, IncCount, &n)
	fmt.Printf("Count: %d elements in nvlist\n", n)
	// // Print the list
	// for p := nvlist; p != nil; p = p.Next {
	// 	fmt.Printf("%s: %c (%U)\n", p.Name, p.Value, p.Value)
	// }
}
