package front

import (
	"github.com/massarakhsh/lik"
)

func (rule *DataRule) Marshal() lik.Seter {
	rule.SeekPageSize()
	rule.marshalControls()
	rule.marshalLog()
	return rule.GetAllResponse()
}

func (rule *DataRule) marshalControls() {
	for lev := 0; lev < len(rule.ItPage.Controls); lev++ {
		if ctrl := rule.ItPage.Controls[lev]; ctrl != nil {
			ctrl.Marshal(rule)
		}
	}
}

func (rule *DataRule) marshalLog() {
	if index := lik.GetProtoIndex(); index != rule.ItPage.IndexProto {
		rule.ItPage.IndexProto = index
		rule.StoreItem(rule.showProtocol())
	}
}

/*
func doMarshal(rule *repo.DataRule) {
	if !rule.ItPage.GetTrust() {
		if rule.BindSession() {
			rule.SetResponse(rule.ItPage.GetPageId(), "_page")
			//PageReload(rule)
		}
	} else if rule.IsNeedGoPath {
		frontDoGoPath(rule)
	} else if rule.IsNeedReload {
		frontDoReload(rule)
	} else if _,control := rule.GetFront(); control != nil {
		control.RunMarshal(rule)
	}
}*/
