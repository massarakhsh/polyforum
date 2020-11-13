package controller

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/ruler"
)

type TableControl struct {
	DataControl
	IsWork bool
}

type Tabler interface {
	Controller
}

func BuildTabler(rule ruler.DataRuler, level int, iswork bool) Tabler {
	it := &TableControl{ IsWork: iswork }
	rule.BindControl(level, it)
	return it
}

func (it *TableControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := it.menuPrepare(rule, false)
	row := tbl.BuildTr()
	var title string
	if it.IsWork {
		title = "База данных"
	} else {
		title = "Справочники"
	}
	it.menuItemCmd(rule, row,"", title, "seek")
	it.menuItemCmd(rule, row,"", "|", "")
	for _, table := range base.ListTables {
		if table.IsWork == it.IsWork {
			it.menuItemCmd(rule, row, table.Part, table.Title, "mode/"+table.Part)
		}
	}
	it.menuTools(rule, row)
	return tbl
}

func (it *TableControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	return div
}

func (it *TableControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("mode") {
		it.execMode(rule)
	} else {
		it.execute(rule)
	}
}

func (it *TableControl) execMode(rule ruler.DataRuler) {
	if it.Mode = rule.Shift(); it.Mode != "" {
		BuildGrider(rule, it.Level + 1, it.Mode)
	} else {
		rule.BindControl(it.Level + 1, nil)
	}
	rule.SetNeedRedraw()
}

func (it *TableControl) Marshal(rule ruler.DataRuler) {
}

