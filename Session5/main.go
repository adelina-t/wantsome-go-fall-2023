package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card interface {
	fmt.Stringer
	Name() string
}

type CarteDeJoc struct {
	Culoare string
	Numar   string
}

func NewCarte(c string, n string) *CarteDeJoc {
	return &CarteDeJoc{
		Culoare: c,
		Numar:   n,
	}
}

func (c *CarteDeJoc) String() string {
	return fmt.Sprintf("%s de tipul %s", c.Numar, c.Culoare)
}

func (c *CarteDeJoc) Name() string {
	return c.String()
}

type CarteDePokemon struct {
	Power int
	Nume  string
}

func NewCarteDePokemon(p int, n string) *CarteDePokemon {
	return &CarteDePokemon{
		Nume:  n,
		Power: p,
	}
}

func (cp *CarteDePokemon) String() string {
	return fmt.Sprintf("Caught %s with power %d", cp.Nume, cp.Power)
}

func (cp *CarteDePokemon) Name() string {
	return cp.String()
}

type Pachet[T Card] struct {
	carti []T
}

func NewPachetDeCarti() *Pachet[*CarteDeJoc] {
	culori := []string{"InimaRosie", "Trefla", "Romb", "InimaNeagra"}
	numere := []string{"2", "3", "4", "10", "J", "Q", "K", "A"}

	pachet := &Pachet[*CarteDeJoc]{}
	for _, c := range culori {
		for _, n := range numere {
			pachet.AdaugaCarte(NewCarte(c, n))
		}
	}
	return pachet
}

func NewPachetDePokemoni() *Pachet[*CarteDePokemon] {
	nume := []string{"pika", "salazar"}
	puteri := []int{100, 10}

	pachet := &Pachet[*CarteDePokemon]{}
	for i, _ := range nume {
		pachet.AdaugaCarte(NewCarteDePokemon(puteri[i], nume[i]))
	}
	return pachet
}

func (p *Pachet[T]) AdaugaCarte(carte T) {
	p.carti = append(p.carti, carte)
}

func (p *Pachet[T]) CarteRandom() T {

	r := rand.New(rand.NewSource(time.Now().UnixMicro()))

	return p.carti[r.Intn(len(p.carti))]

}

func main() {

	pachet := NewPachetDeCarti()
	pachetDePokemon := NewPachetDePokemoni()

	carte := pachet.CarteRandom()
	fmt.Println(carte)

	carte1 := pachetDePokemon.CarteRandom()
	fmt.Println(carte1)

}
