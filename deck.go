package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := make([]string, 0)

	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clover"}
	cardValues := []string{"As" /*"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen"*/}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%v of %v", value, suit))
		}
	}

	return cards
}

func (d *deck) print() {
	for i, v := range *d {
		fmt.Printf("deck[%v]: %v\n", i, v)
	}
}

func (d *deck) toString() string {
	return strings.Join((*d)[:], ",")
}

func (d *deck) saveToAFile(fileName string) {
	err := os.WriteFile(fmt.Sprintf("files/%v", fileName), []byte((*d).toString()), 0666)

	if err != nil {
		fmt.Println(err)
	}
}

func (d *deck) shuffle() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	r.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })

	// for index := range *d {
	// 	newPosition := r.Intn(len(*d) - 1)
	// 	(*d)[index], (*d)[newPosition] = (*d)[newPosition], (*d)[index]
	// }
}

func deal(d *deck, handSize uint8) (deck, deck) {

	return (*d)[:handSize], (*d)[handSize:]
}

func fromStringToDeck(s string) *deck {
	d := deck(strings.Split(s, ","))
	return &d
}

func newDeckFromFile(fileName string) *deck {
	bs, err := os.ReadFile(fmt.Sprintf("files/%v", fileName))

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	d := deck(strings.Split(string(bs), ","))
	return &d
}
