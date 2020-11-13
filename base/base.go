package base

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likbase"
	"fmt"
)

const (
	Version = "1.0.5"
)

type TableBase struct {
	Part	string
	Title	string
	IsWork	bool
	Fields []FieldBase
}

type FieldBase struct {
	Key		string
	Title	string
	Proto	string
}

var (
	DB likbase.DBaser

	DictionParticipate = TableBase{"Participate", "Форма участия", false,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Code", "Код", "L"},
			{"Name", "Наименование", "S"},
		}}
	DictionDegree = TableBase{"Degree", "Ученая степень", false,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Code", "Код", "L"},
			{"Name", "Наименование", "S"},
		}}
	DictionSection = TableBase{"Section", "Секция", false,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Code", "Код", "L"},
			{"Name", "Наименование", "S"},
		}}
	DictionStatus = TableBase{"Status", "Статус", false,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Code", "Код", "L"},
			{"Name", "Наименование", "S"},
			{"Logo", "Логотип", "S"},
		}}
	DictionTypeComm = TableBase{"TypeComm", "Тип комитетчика", false,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Code", "Код", "L"},
			{"Name", "Наименование", "S"},
		}}

	TableExhibition = TableBase{"Exhibition", "Форум", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"AssociationId", "Ассоциация", "L"},
			{"Category", "Категория", "S"},
			{"Name", "Наименование", "S"},
			{"Address", "Адрес", "S"},
			{"BeginDt", "Начало", "S"},
			{"EndDt", "Окончание", "S"},
			{"Web", "Сайт", "S"},
			{"Badge", "Макет", "S"},
			{"ProgPdf", "Программа", "S"},
			{"Phone", "Телефон", "S"},
			{"Email", "Почта", "S"},
			{"Driveway", "Схема проезда", "S"},
			{"Welcome", "Приветствие", "S"},
			{"Info", "Информация", "S"},
			{"Plan", "План выставки", "S"},
			{"Plan2", "План 2 этажа", "S"},
			{"Logo", "Логотип", "S"},
			{"BarStart", "Старт штрих-кода", "S"},
			{"BarLength", "Количество штрих-кодов", "S"},
		}}
	TableAction = TableBase{"Action", "Мероприятие", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Category", "Категория", "S"},
			{"Name", "Наименование", "S"},
			{"ExhibitionId", "Форум", "L"},
			{"ZoneId", "Место", "L"},
			{"DateDt", "Дата", "S"},
			{"BeginAt", "Начало", "S"},
			{"EndAt", "Окончание", "S"},
			{"Busy", "Присутствует", "L"},
			{"FlagNMAO", "Флаг NMAO", "L"},
			{"SpeakerId", "Ведущий спикер", "L"},
			{"Code", "Code", "L"},
		}}
	TableZone = TableBase{"Zone", "Место", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"ExhibitionId", "Форум", "L"},
			{"Category", "Категория", "S"},
			{"Name", "Наименование", "S"},
			{"Capacity", "Мест", "L"},
			{"Code", "Code", "L"},
		}}
	TablePerson = TableBase{"Person", "Персона", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Surname", "Фамилия", "S"},
			{"Name", "Имя", "S"},
			{"Name2", "Отчество", "S"},
			{"Phone", "Телефон", "S"},
			{"Email", "Почта", "S"},
			{"Code", "Code", "L"},
			{"BirthDt", "Дата рождения", "S"},
			{"City", "Город", "S"},
			{"AppId", "Устройство", "S"},
			{"Login", "Логин", "S"},
			{"PW5", "Пароль", "S"},
		}}

	TableRequest = TableBase{"Request", "Заявка", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"PersonId", "Персона", "L"},
			{"ExhibitionId", "Форум", "L"},
			{"ParticipateId", "Форма участия", "L"},
			{"DegreeId", "Ученая степень", "L"},
			{"Post", "Должность", "S"},
			{"Firm", "Организация", "S"},
			{"BarCode", "Штрих-код", "S"},
		}}
	TableEvent = TableBase{"Event", "Событие", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"At", "Время", "L"},
			{"Category", "Категория", "S"},
			{"Name", "Наименование", "S"},
			{"ExhibitionId", "Форум", "L"},
			{"ZoneId", "Место", "L"},
			{"PersonId", "Персона", "L"},
		}}
	TableSpeaker = TableBase{"Speaker", "Спикер", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"FIO", "ФИО спикера", "S"},
			{"City", "Город", "S"},
			{"Info", "Информация", "S"},
			{"Photo", "Фотография", "S"},
			{"ExhibitionId", "Форум", "L"},
			{"IsMain", "Флаг ведущего", "L"},
		}}
	TableProgram = TableBase{"Program", "Программа", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"ExhibitionId", "Форум", "L"},
			{"ActionId", "Мероприятие", "L"},
			{"SpeakerId", "Ведущий спикер", "L"},
			{"Topic", "Тема выступления", "S"},
			{"Duration", "Длительность (мин.)", "L"},
			{"Authors", "Авторы", "S"},
		}}
	TableExhibitor = TableBase{"Exhibitor", "Экспонент", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"ExhibitionId", "Форум", "L"},
			{"Name", "Наименование", "S"},
			{"Stand", "Номер стенда", "S"},
			{"StatusId", "Статус", "L"},
			{"Info", "Информация", "S"},
			{"Logo", "Логотип", "S"},
		}}
	TableAssociation = TableBase{"Association", "Ассоциации", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Name", "Наименование", "S"},
			{"Web", "Адрес", "S"},
			{"Logo", "Логотип", "S"},
			{"Code", "Code", "L"},
		}}
	TableOrganizer = TableBase{"Organizer", "Организатор", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"Name", "Наименование", "S"},
			{"ExhibitionId", "Форум", "L"},
			{"Web", "Адрес", "S"},
			{"Logo", "Логотип", "S"},
		}}
	TableCommittee = TableBase{"Committee", "Комитет", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"TypeCommId", "Тип", "L"},
			{"ExhibitionId", "Форум", "L"},
			{"FIO", "ФИО", "S"},
			{"Info", "Информация", "S"},
			{"Web", "Адрес", "S"},
		}}
	TableMainDate = TableBase{"MainDate", "Основные даты", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"ExhibitionId", "Форум", "L"},
			{"Date", "Дата", "S"},
			{"Info", "Информация", "S"},
		}}
	TableShortProg = TableBase{"ShortProg", "Краткая программа", true,
		[]FieldBase{
			{"Id", "ID", "LP"},
			{"ExhibitionId", "Форум", "L"},
			{"Name", "Наименование", "S"},
			{"Date", "Дата", "S"},
			{"TimeBeg", "Время начала", "S"},
			{"TimeEnd", "Время окончания", "S"},
			{"Place", "Место провендения", "S"},
		}}

	ListTables = []*TableBase{
		&DictionParticipate, &DictionDegree,
		&DictionSection, &DictionStatus,
		&DictionTypeComm,
		&TableExhibition, &TableAction, &TableZone,
		&TablePerson,
		&TableRequest, &TableEvent,
		&TableSpeaker, &TableProgram, &TableExhibitor,
		&TableAssociation, &TableOrganizer,
		&TableCommittee, &TableMainDate, &TableShortProg,
	}
)

func OpenDB(serv string, name string, user string, pass string) bool {
	likbase.FId = "Id"
	logon := user + ":" + pass
	addr := "tcp(" + serv + ":3306)"
	if DB = likbase.OpenDBase("mysql", logon, addr, name); DB == nil {
		lik.SayError(fmt.Sprint("DB not opened"))
		return false
	}
	initializeTables()
	return true
}

func initializeTables() {
	for _, table := range ListTables {
		fields := []likbase.DBField{}
		for _,fld := range table.Fields {
			fields = append(fields, likbase.DBField{ Name: fld.Key, Proto: fld.Proto })
		}
		DB.ControlTable(table.Part, fields)
	}
}

func RestartDataBase() {
	for _, table := range ListTables {
		DB.Execute(fmt.Sprintf("DELETE FROM `%s`", table.Part))
	}
	initializeTables()
}

func GetTable(part string) *TableBase {
	var table *TableBase
	for _,tbl := range ListTables {
		if tbl.Part == part {
			table = tbl
			break
		}
	}
	return table
}

func GetElm(part string, id lik.IDB) lik.Seter {
	return DB.GetOneById(part, id)
}

func InsertElm(part string, sets lik.Seter) lik.IDB {
	return DB.InsertElm(part, sets)
}

func UpdateElm(part string, id lik.IDB, sets lik.Seter) bool {
	return DB.UpdateElm(part, id, sets)
}

func GetList(part string, sort string) lik.Lister {
	order := ""
	if sort == "Id" {
		order = "Id"
	} else if sort == "_Id" {
		order = "Id DESC"
	}
	return DB.GetListElm("*", part, "", order)
}

func DeleteElm(part string, id lik.IDB) bool {
	return DB.DeleteElm(part, id)
}

func GetLastId(part string) lik.IDB {
	id,_ := DB.CalculeIDB(DB.PrepareSql("MAX(Id)", part, "", ""))
	return id
}

