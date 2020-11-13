package controller

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
)

type DbControl struct {
	DataControl
}

type Dber interface {
	Controller
}

func BuildDb(rule ruler.DataRuler, level int) Dber {
	it := &DbControl{ }
	rule.BindControl(level, it)
	return it
}

func (it *DbControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := likdom.BuildTableClass("menu")
	row := tbl.BuildTr()
	it.menuItemText(rule, row,"База данных")
	it.menuItemText(rule, row, "|")
	it.menuTools(rule, row)
	return tbl
}

func (it *DbControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	tbl := div.BuildTableClass("api")
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("api_title").BuildString("Очистка базы данных")
		proc := fmt.Sprintf("purge_all_base('%s')", it.buildPart("purgebase"))
		row.BuildTdClass("api_info").AppendItem(it.linkItemProc("Очистить", proc, "cmd"))
	}
	return div
}

func (it *DbControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("purgebase") {
		fmt.Println("purgebase")
	} else {
		it.execute(rule)
	}
}

func (it *DbControl) Marshal(rule ruler.DataRuler) {
}

//func (it *DbControl) buildLinePurge(rule ruler.DataRuler) likdom.Domer {
//	proc := it.buildProc("purge")
//	line := LinkTe("api", "Очистить", proc)
//	return line
//}

func (it *DbControl) buildProc(part string) string {
	path := it.buildPart(part)
	return fmt.Sprintf("%s('%s')", "db_" + part, path)
}

