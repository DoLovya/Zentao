package TerminalUI

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"math"
	"os"
)

type Option struct {
	id    int
	value string
}

type List struct {
	scn           tcell.Screen
	title         string
	options       []string
	selectedIndex int
	topIndex      int
	bottomIndex   int
}

func NewList() *List {
	list := &List{}
	list.scn, _ = tcell.NewScreen()
	list.scn.Init()
	list.topIndex = 0
	list.bottomIndex = -1
	return list
}

func (this *List) SetTitle(title string) {
	this.title = title
}

func (this *List) SetOptions(options []string) {
	this.options = options
}

func (this *List) DrawLine(row int, value string, style tcell.Style) {
	x := 0
	for _, ch := range value {
		width := runewidth.RuneWidth(ch)
		this.scn.SetContent(x, row, rune(ch), []rune(""), style)
		x += width
	}
}

func (this *List) Draw() {
	this.DrawLine(0, this.title, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	for row := this.topIndex; row <= this.bottomIndex; row++ {
		if row == this.selectedIndex {
			style := tcell.StyleDefault.
				Foreground(tcell.ColorPink)
			this.DrawLine(row-this.topIndex+1, this.options[row], style)
		} else {
			this.DrawLine(row-this.topIndex+1, this.options[row], tcell.StyleDefault)
		}
	}
}

func (this *List) Exec(selectedIndex *int) error {
	this.bottomIndex = int(math.Min(float64(len(this.options)), float64(20))) - 1
	this.selectedIndex = 0

	this.scn.Clear()
	this.Draw()
	this.scn.Show()

	events := make(chan tcell.Event)
	go func() {
		for {
			event := this.scn.PollEvent()
			events <- event
		}
	}()
	for {
		select {
		case ev := <-events:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlC, tcell.KeyEsc:
					this.scn.Fini()
					os.Exit(0)
				case tcell.KeyUp:
					if this.selectedIndex != 0 {
						this.selectedIndex -= 1
					}
					if this.selectedIndex < this.topIndex {
						this.topIndex -= 1
						this.bottomIndex -= 1
					}
				case tcell.KeyDown:
					if this.selectedIndex != len(this.options)-1 {
						this.selectedIndex += 1
					}
					if this.selectedIndex > this.bottomIndex {
						this.topIndex += 1
						this.bottomIndex += 1
					}
				case tcell.KeyEnter:
					this.scn.Fini()
					*selectedIndex = this.selectedIndex
					return nil
				default:
				}
			}
		}
		this.scn.Clear()
		this.Draw()
		this.scn.Show()
	}
}
