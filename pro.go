package main

type Pros struct {
	Names []string
	Index int
}

func (p *Pros) Next() (name string) {
	name = p.Names[p.Index]
	p.Index++
	if p.Index == len(p.Names) {
		p.Index = 0
	}
	return
}

var sunday_pros = Pros{
	Names: []string{
		"Chris Chatfield",
		"Phil Scapens",
		"Allan Carli",
		"John Gratton",
		"Norm Wallis",
		"Daryle Bampton",
		//	"Graham Ingerson",
		"Tim Arland",
		"David Lane",
		"John Inverarity",
		"Ecio Marcel",
		"Dean Crichton",
		"Kym Stringer",
		"John Brolese",
		"Kevin Bowden",
		"John Faturic",
		"Ian Dowsett",
		"Rob Lees",
		"Brian Purvis",
		"Jeff Watson",
		"Simon How",
		"Alan Gold",
		"Chris Juttner",
		"Steve Arthur",
		"Marcus Towell",
		"Warwick Bowen",
		"Peter Hendy",
		"Max Walker",
		"Alex Hayter",
		"Neil Martin",
		"Greg Peake",
		"Alex Scapens",
		"Ruth Williams",
	},
}

var tuesday_pros = Pros{
	Names: []string{
		"Alex Scapens",
		"Phil Scapens",
		"Tim Arland",
		"Chris Chatfield",
		"Norman Wallis",
		"David Lane",
		"Graham Ingerson",
		"John Gratton",
		"Peter Hendy",
		"Alan Gold",
		"Dean Crichton",
		"Ian Dowsett",
		"Rob Lees",
		"Daryle Bampton",
		"Neil Martin",
		"Kevin Bowden",
		"Warrick Bowen",
		"Ecio Marcel",
		"John Inverarity",
		"Graham Ingerson", // Double, no Sunday
		"Max Walker",
		"Alex Hayter",
		"John Brolese",
		"Chris Juttner",
		"John Faturic",
		"Johno Johnson",
	},
}

var wednesday_pros = Pros{
	Names: []string{
		"Garry Loughhead",
		"Tim Arland",
		"Kym Stringer",
		"Rob Lees",
		"Kevin Bartlett",
		"Peter Phillis",
		"Alex Hayter",
	},
}
