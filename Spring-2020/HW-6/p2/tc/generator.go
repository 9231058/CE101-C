package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var words = []string{
	"Tampa Bay Rays", "New York Yankees", "Toronto Blue Jays",
	"Baltimore Orioles", "Boston Red Sox", "Kansas City Royals",
	"Minnesota Twins", "Detroit Tigers", "Cleveland Indians",
	"Chicago White Sox", "Houston Astros", "Texas Rangers",
	"Los Angeles Angels", "Seattle Mariners", "Oakland Athletics",
	"saucepan", "double boiler", "dutch oven", "stock pot", "pressure cooker",
	"steamer", "steamer basket", "couscous kettle", "dripping pan", "terrine",
	"roasting pan", "fish poacher", "tajine", "fondue pot", "wok", "frying pan",
	"diable", "egg poacher", "blender", "mixer",
	"accordion", "trumpet", "clarinet", "bass guitar", "concertina", "drums",
	"saxophone", "violin", "banjo", "mandolin", "guitar", "harmonica", "cello",
	"dulcimer", "harpsichord", "flute", "harp", "piano", "ukulele", "xylophone",
	"lute",
	"anise", "allspice", "basil", "bay leaves", "caraway seeds", "cardamom",
	"cayenne pepper", "chili powder", "chives", "cilantro", "cinnamon",
	"cloves", "coriander", "cumin", "curry powder", "dill weed", "dill seed",
	"fennel seed", "fenugreek", "garlic", "ginger", "ginseng", "mace",
	"marjoram", "mint", "mustard", "nutmeg", "oregano", "paprika", "parsley",
	"peppercorns", "poppy seeds", "red pepper flakes", "rosemary", "saffron",
	"sage", "savory", "tarragon", "thyme", "turmeric",
	"Mary", "Patricia", "Ashley", "Elizabeth", "Linda", "Barbara", "Susan",
	"Margaret", "Jessica", "Sarah", "Dorothy", "Karen", "Nancy", "Betty",
	"Lisa", "Sandra", "Jennifer", "Katrina", "Donna", "Helen", "Carol",
	"Michelle", "Emily", "Amanda", "Melissa", "Deborah", "Laura", "Stephanie",
	"Rebecca", "Sharon", "Cynthia", "Kathleen", "Anna", "Shirley", "Ruth",
	"Amy", "Andrea", "Brenda", "Virginia", "Pamela", "Katherine", "Nicole",
	"Christine", "Samantha", "Carolyn",
}

func main() {
	for i := 0; i < 100; i++ {
		sLen := rand.Intn(len(words)-20) + 20
		sb := new(strings.Builder)
		for j := 0; j < sLen; j++ {
			index := rand.Intn(len(words))
			sb.WriteString(words[index])
			sb.WriteString(" ")
		}

		subString := words[rand.Intn(len(words))]
		replacement := words[rand.Intn(len(words))]

		fmt.Printf("%s\n%s\n%s;", sb.String(), subString, replacement)
	}
}
