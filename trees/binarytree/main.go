package main

import (
	"log"
	"strings"
	"time"
)

func main() {

	log.Println("A1" < "B3")
	// os.Exit(1)

	Item1 := Item{
		Tag:  "A1:A2:A3:A4",
		Name: "meow",
	}
	Item2 := Item{
		Tag:  "B1:B2:B3:B4",
		Name: "meow",
	}
	Item3 := Item{
		Tag:  "B1:B2:A1:A2",
		Name: "meow",
	}
	Item4 := Item{
		Tag:  "A1:A2:A3:A4",
		Name: "meow",
	}
	Item5 := Item{
		Tag:  "B1:B2:B3:B4",
		Name: "meow",
	}
	Item6 := Item{
		Tag:  "B1:B2:B3:B4",
		Name: "meow",
	}
	AddItemToTree(&Item1)
	AddItemToTree(&Item2)
	AddItemToTree(&Item3)
	AddItemToTree(&Item4)
	AddItemToTree(&Item5)
	AddItemToTree(&Item6)

	log.Println()
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	TREE.Traverse("x", 0, TREE.Root, func(parent string, level int, e *Edge) {
		lv := "       "
		rv := "       "
		if e.Left != nil {
			lv = e.Left.Value
		}
		if e.Right != nil {
			rv = e.Right.Value
		}
		log.Println(lv, " <<< ||| ", level, "("+parent+")", e.Value, "||| >>> ", rv, "     // ", e.ItemList)
	})
}

type Item struct {
	Tag  string
	Name string
}

func AddItemToTree(newItem *Item) {

	ss := time.Now()
	tag := strings.Split(newItem.Tag, ":")

	log.Println(tag)

	found := TREE.TreeFindAndInsert([]string{tag[0], tag[1], tag[2], tag[3]}, newItem)
	if found {

		log.Println("FOUND:", time.Since(ss).Milliseconds(), newItem.Name, tag)

	} else {

		TREE.Insert("root", nil).Insert(tag[0], newItem).Insert(tag[1], newItem).Insert(tag[2], newItem).Insert(tag[3], newItem)
		// TREE.Insert("root", nil).Insert(tag[0], newItem).Insert(tag[1], newItem).Insert(tag[2], newItem).Insert(tag[3], newItem)
		// TREE.Insert("root", nil).Insert(tag[1], newItem).Insert(tag[2], newItem).Insert("*", newItem).Insert(tag[4], newItem)

		log.Println("NEW:", time.Since(ss).Nanoseconds(), newItem.Name, newItem.Tag)
	}
}
