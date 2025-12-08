package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DimitriLaPoudre/InazuScrap/internal/model"
	"github.com/PuerkitoBio/goquery"
)

func Scrap() []model.Player {
	var players []model.Player
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
			var newPlayer model.Player
			name := s.Find("div.nameBox span.name ruby").First().Text()
			if name == "" {
				return
			}
			newPlayer.Name = name

			s.Find("div.detailBox div.rBox ul.param li").Each(func(i int, s *goquery.Selection) {
				statName := s.Find("dl:not([class]) dt").Text()
				switch statName {
				case "Position":
					position := s.Find("dl:not([class]) dd p").Text()
					if position == "" {
						return
					}
					newPlayer.Position = position

					element := s.Find("dl.box dd p").Text()
					if element == "" {
						return
					}
					switch element {
					case "Fire":
						element = "🔥"
					case "Mountain":
						element = "🏔️"
					case "Wind":
						element = "🌪️"
					case "Forest":
						element = "🌿"
					default:
						element = "?"
					}
					newPlayer.Element = element
				case "Kick":
					kick := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(kick, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Kick = uint(value)
				case "Control":
					control := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(control, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Control = uint(value)
				case "Technique":
					technique := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(technique, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Technique = uint(value)
				case "Pressure":
					pressure := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(pressure, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Pressure = uint(value)
				case "Physical":
					physical := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(physical, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Physical = uint(value)
				case "Agility":
					agility := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(agility, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Agility = uint(value)
				case "Intelligence":
					intelligence := s.Find("dl dd table tbody tr").Eq(1).Find("td").Text()
					value, err := strconv.ParseUint(intelligence, 10, strconv.IntSize)
					if err != nil {
						return
					}
					newPlayer.Intelligence = uint(value)

				}
			})
			newPlayer.Sum = newPlayer.Kick +
				newPlayer.Control +
				newPlayer.Technique +
				newPlayer.Pressure +
				newPlayer.Physical +
				newPlayer.Agility +
				newPlayer.Intelligence

			players = append(players, newPlayer)
		})
		if result.Length() == 0 {
			break
		}

	}
	return players
}
