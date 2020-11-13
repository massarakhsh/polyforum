package ruler

import (
	"github.com/massarakhsh/lik/likdom"
	"fmt"
)

type DataControl struct {
	Level	int
	Mode	string
}

type ControlExecuter interface {
	Run(rule DataPager)
}

type ControlMarshaler interface {
	Run(rule DataPager)
}

type Controller interface {
	SetLevel(lev int)
	GetLevel() int
	GetMode() string
	ShowMenu(rule DataRuler) likdom.Domer
	ShowInfo(rule DataRuler) likdom.Domer
	Execute(rule DataRuler)
	Marshal(rule DataRuler)
}

func (it *DataControl) SetLevel(lev int) {
	it.Level = lev
}

func (it *DataControl) GetLevel() int {
	return it.Level
}

func (it *DataControl) GetMode() string {
	return it.Mode
}

func GetIdLevel(lev int) string {
	return fmt.Sprintf("c%d", lev)
}

