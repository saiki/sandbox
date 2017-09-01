package main

import (
	//	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{300, 200},
		Layout: Grid{
			Columns:     3,
			MarginsZero: true,
			SpacingZero: true,
		},
		Children: []Widget{
			PushButton{
				Text: "A",
			},
			PushButton{
				Text: "B",
			},
			PushButton{
				Text: "C",
			},
			PushButton{
				Text: "D",
			},
			PushButton{
				Text: "1",
			},
			PushButton{
				Text: "2",
			},
			PushButton{
				Text: "3",
			},
			PushButton{
				Text: "4",
			},
		},
	}.Run()
}
