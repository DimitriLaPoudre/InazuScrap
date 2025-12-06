package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Player struct {
	name         string
	position     string
	element      string
	kick         uint
	control      uint
	technique    uint
	pressure     uint
	physical     uint
	agility      uint
	intelligence uint
	sum          uint
}

func (p Player) String() string {
	return fmt.Sprint(p.name,
		"\n", p.position,
		"\n", p.element,
		"\nTotal: ", p.sum,
		"\n\tKick: ", p.kick,
		"\n\tControl: ", p.control,
		"\n\tTechnique: ", p.technique,
		"\n\tPressure: ", p.pressure,
		"\n\tPhysical: ", p.physical,
		"\n\tAgility: ", p.agility,
		"\n\tIntelligence: ", p.intelligence)
	// return "test"
}

// func (p *Player) String() string {
// 	return fmt.Sprintf(p.name, "\t", p.position,
// 		"\n\tKick: ", p.kick,
// 		"\n\tControl: ", p.control,
// 		"\n\tTechnique: ", p.technique,
// 		"\n\tPressure: ", p.pressure,
// 		"\n\tPhysical: ", p.physical,
// 		"\n\tAgility: ", p.agility,
// 		"\n\tIntelligence: ", p.intelligence)
// }

func main() {
	var GK []Player
	var DF []Player
	var MF []Player
	var FW []Player
	for page := 1; ; page++ {
		resp, err := http.Get(fmt.Sprintf("https://zukan.inazuma.jp/en/chara_param/?page=%d&per_page=1000", page))
		if err != nil {
			log.Fatal(err)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		resp.Body.Close()

		result := doc.Find("ul.charaListBox li").Each(func(i int, s *goquery.Selection) {
			var newPlayer Player
			name := s.Find("div.nameBox span.name ruby").First().Text()
			if name == "" {
				return
			}
			newPlayer.name = name

			s.Find("div.detailBox div.rBox ul.param li").Each(func(i int, s *goquery.Selection) {
				statName := s.Find("dl:not([class]) dt").Text()
				switch statName {
				case "Position":
					position := s.Find("dl:not([class]) dd p").Text()
					if position == "" {
						return
					}
					newPlayer.position = position

					element := s.Find("dl.box dd p").Text()
					if element == "" {
						return
					}
					newPlayer.element = element
				case "Kick":
					kick := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(kick, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.kick = uint(value)
				case "Control":
					control := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(control, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.control = uint(value)
				case "Technique":
					technique := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(technique, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.technique = uint(value)
				case "Pressure":
					pressure := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(pressure, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.pressure = uint(value)
				case "Physical":
					physical := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(physical, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.physical = uint(value)
				case "Agility":
					agility := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(agility, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.agility = uint(value)
				case "Intelligence":
					intelligence := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(intelligence, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.intelligence = uint(value)

				}
			})
			newPlayer.sum = newPlayer.kick + newPlayer.control + newPlayer.technique + newPlayer.pressure + newPlayer.physical + newPlayer.agility + newPlayer.intelligence
			switch newPlayer.position {
			case "GK":
				GK = append(GK, newPlayer)
			case "DF":
				DF = append(DF, newPlayer)
			case "MF":
				MF = append(MF, newPlayer)
			case "FW":
				FW = append(FW, newPlayer)
			}
		})
		if result.Length() == 0 {
			break
		}
	}
	slices.SortFunc(GK, func(a, b Player) int {
		return int(a.sum) - int(b.sum)
	})
	slices.SortFunc(DF, func(a, b Player) int {
		return int(a.sum) - int(b.sum)
	})
	slices.SortFunc(MF, func(a, b Player) int {
		return int(a.sum) - int(b.sum)
	})
	slices.SortFunc(FW, func(a, b Player) int {
		return int(a.sum) - int(b.sum)
	})

	fmt.Println("test")
	for _, player := range MF {
		// if player.name == "David Samford" {
		fmt.Println(player)
		// }
	}
}
