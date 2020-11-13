package main

import (
	"github.com/massarakhsh/polyforum/api"
	"github.com/massarakhsh/polyforum/generate"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/front"
)

var (
	HostPort = 80
	HostServ = "localhost"
	HostBase = "polyforum"
	HostUser = "polyforum"
	HostPass = "Polyforum17"
)

func main() {
	lik.SetLevelInf()
	lik.SayError("System started")
	if !getArgs() {
		return
	}
	if !base.OpenDB(HostServ, HostBase, HostUser, HostPass) {
		return
	}
	http.HandleFunc("/", routerMain)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", HostPort), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getArgs() bool {
	args, ok := lik.GetArgs(os.Args[1:])
	if val := args.GetInt("port"); val > 0 {
		HostPort = val
	}
	if val := args.GetString("serv"); val != "" {
		HostServ = val
	}
	if val := args.GetString("base"); val != "" {
		HostBase = val
	}
	if val := args.GetString("user"); val != "" {
		HostUser = val
	}
	if val := args.GetString("pass"); val != "" {
		HostPass = val
	}
	if len(HostBase) <= 0 {
		fmt.Println("HostBase name must be present")
		ok = false
	}
	if !ok {
		fmt.Println("Usage: polyforum [-key val | --key=val]...")
		fmt.Println("port    - port value (80)")
		fmt.Println("serv    - Database server")
		fmt.Println("base    - Database name")
		fmt.Println("user    - Database user")
		fmt.Println("pass    - Database pass")
	}
	return ok
}

func routerMain(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PROPFIND" {
		return
	}
	isapi := lik.RegExCompare(r.RequestURI,"^/api")
	isfront := lik.RegExCompare(r.RequestURI,"^/front")
	ismarshal := lik.RegExCompare(r.RequestURI,"^/marshal")
	if match := lik.RegExParse(r.RequestURI, "/ean13/(\\d+)\\.png"); match != nil {
		path := generate.DirectEan13(match[1], r.RequestURI)
		likapi.ProbeRouteFile(w, r, path)
		return
	}
	if !isfront && !ismarshal &&
		lik.RegExCompare(r.RequestURI, "\\.(js|css|htm|html|ico|gif|png|jpg|jpeg|pdf|doc|docx|xls|xlsx)(\\?|$)") {
		likapi.ProbeRouteFile(w, r, r.RequestURI)
		return
	}
	var page *ruler.DataPage
	if sp := lik.StrToInt(likapi.GetParm(r, "_sp")); sp > 0 {
		if pager := likapi.FindPage(sp); pager != nil {
			page = pager.(ruler.DataPager).GetItPage()
		}
	}
	if page == nil {
		page = ruler.StartPage()
	}
	var rule ruler.DataRuler
	if isapi {
		rule = api.BuildRule(page)
	} else {
		rule = front.BuildRule(page)
	}
	rule.LoadRequest(r)
	if !ismarshal {
		rule.RuleLog()
	}
	if !rule.Authority() && !isfront && !ismarshal {

	}
	if isfront {
		json := rule.Execute()
		likapi.RouteJson(w, 200, json, false)
	} else if ismarshal {
		json := rule.Marshal()
		likapi.RouteJson(w, 200, json, false)
	} else if !rule.Authority() {
		likapi.Route401(w, 401, "realm=\"PolyForum\"")
	} else if isapi {
		json := rule.Execute()
		likapi.RouteJson(w, 200, json, false)
	} else {
		html := rule.ShowPage()
		likapi.RouteCookies(w, rule.GetAllCookies())
		likapi.RouteHtml(w, 200, html.ToString())
	}
}
