package front

import (
	"github.com/massarakhsh/polyforum/controller"
	"github.com/massarakhsh/polyforum/ruler"
	"github.com/massarakhsh/lik/likdom"
)

func (rule *DataRule) ShowRedraw() {
	rule.StoreItem(rule.showMainGen())
}

func (rule *DataRule) showMainGen() likdom.Domer {
	div := likdom.BuildDivClassId("main_page", "page")
	if len(rule.ItPage.Controls) == 0 {
		controller.BuildRoot(rule, 0)
	}
	dat := div.BuildDivClass("main_data fill")
	dat.AppendItem(rule.showControlGen(rule.ItPage.Controls[0]))
	return div
}

func (rule *DataRule) showControlGen(controller ruler.Controller) likdom.Domer {
	tbl := likdom.BuildTableClass("main_data")
	tbl.BuildTrTdClass("main_data").AppendItem(controller.ShowMenu(rule))
	tbl.BuildTrTdClass("main_space")
	dat := tbl.BuildTrTdClass("main_info")
	lev := controller.GetLevel()
	if lev + 1 < len(rule.ItPage.Controls) {
		dat.AppendItem(rule.showControlGen(rule.ItPage.Controls[lev + 1]))
	} else {
		dat.AppendItem(controller.ShowInfo(rule))
	}
	return tbl
}

