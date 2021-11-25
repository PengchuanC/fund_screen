package manager

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/dao/utils"
	"fund_screen/db/models/funds"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"sync"
)

var (
	db *gorm.DB
)

type Manager struct {
	Secucode string
}

func (m *Manager) FundInstance() funds.Fund {
	var fund funds.Fund
	fund = funds.Fund{Secucode: m.Secucode}
	db = connector.GetDB()
	db.Find(&fund)
	return fund
}

func (m *Manager) Maincode() string {
	return utils.FundAssociate(m.Secucode)
}

// EstablishedDate 基金成立日期，转型基金转型前成立日期
func (m *Manager) EstablishedDate() carbon.ToDateString {
	var (
		main funds.FMainCode
	)
	db = connector.GetDB()
	sub := db.Model(main).Select("max(id)").Where("secucode = ?", m.Secucode)
	db.Find(&main, "id = (?)", sub)
	if main.EstablishDateii.ToDateString() != "" {
		return main.EstablishDateii
	}
	return main.LaunchDate
}

// CurrentManager 当前基金经理的个人编码
func (m *Manager) CurrentManager() int64 {
	var (
		manager funds.FManager
	)
	manager = funds.FManager{Postname: 1, Incumbent: 1}
	db = connector.GetDB()
	db.Where("secucode = ? and postname = 1 and incumbent = 1", m.Secucode).First(&manager)
	return manager.Personalcode
}

// ServeInfo 管理当前基金的基金经理信息
func (m *Manager) ServeInfo() (string, float64, int, string) {
	var (
		relate       string
		manager      funds.FManager
		start        string
		name         string
		days         int
		performance  float64
		personalCode int64
		err          error
	)
	performance = 0
	days = 0
	db = connector.GetDB()
	relate = m.Maincode()
	personalCode = m.CurrentManager()
	db.Last(&manager, "secucode = ? and personalcode = ?", m.Secucode, personalCode)
	performance = manager.Performance
	start = manager.Accessiondate.ToDateString()
	days += manager.Managementtime
	name = manager.Name
	if relate != "" {
		if err = db.Last(&manager, "secucode = ? and personalcode = ?", relate, personalCode).Error; err != nil {
			return name, performance, days, start
		}
		start = manager.Accessiondate.ToDateString()
		performance = (1+performance)*(1+manager.Performance) - 1
		days += manager.Managementtime
	}
	return name, performance, days, start
}

// Managed 同一基金经理管理的基金
func (m *Manager) Managed() []string {
	var (
		personalCode int64
		managed      []string
		manager      funds.FManager
		maincode     funds.FMainCode
	)
	db = connector.GetDB()
	personalCode = m.CurrentManager()
	db.Model(&manager).Select("secucode").Where("personalcode = ? and postname = 1 and incumbent = 1", personalCode).Order("secucode").Distinct().Find(&managed)
	db.Model(&maincode).Select("left(maincode,6) as maincode").Where("secucode in ?", managed).Order("maincode").Distinct().Find(&managed)
	return managed
}

func (m *Manager) Scale() float64 {
	var (
		scale funds.AnalysisScale
	)
	db = connector.GetDB()
	db.Last(&scale, "secucode = ?", m.Maincode())
	return scale.CombNvi / 1e8
}

func (m *Manager) ManagedScale() float64 {
	var (
		scale   float64
		managed []string
		table   string
		Scale   funds.AnalysisScale
	)
	db = connector.GetDB()
	managed = m.Managed()
	table = Scale.TableName()
	subQuery := db.Select("max(date)").Where("secucode in ('000001', '110011')").Table(table)
	subQuery = db.Table(table).Select("max(id)").Where("secucode in ? and date = (?)", managed, subQuery).Group("secucode")
	db.Table(table).Select("SUM(nvi) as scale").Where("id in (?)", subQuery).Find(&scale)
	return scale / 1e8
}

func (m *Manager) Nature(nature string) string {
	var (
		style string
		table string
		maincode string
	)
	switch nature {
	case "scale":
		table = funds.ScaleNature{}.TableName()
	case "style":
		table = funds.StyleNature{}.TableName()
	case "industry":
		table = funds.IndustryNature{}.TableName()
	case "industry_style":
		table = funds.IndustryStyle{}.TableName()
	default:
		table = funds.ScaleNature{}.TableName()
	}
	db = connector.GetDB()
	maincode = m.Maincode()
	sub := db.Table(table).Select("max(id)").Where("secucode = ?", maincode)
	db.Table(table).Select("style").Where("id = (?)", sub).Find(&style)
	return style
}

func (m Manager) Stock() float64 {
	var (
		stock float64
		al funds.FAllocate
	)
	db = connector.GetDB()
	sub := db.Model(&al).Select("id").Where("secucode = ?", m.Maincode()).Order("-date").Limit(1)
	db.Model(&al).Select("stock").Where("id = (?)", sub).Find(&stock)
	return stock
}

func (m Manager) Classify() string {
	var classify funds.FClassifyNOI
	db = connector.GetDB()
	db.Last(&classify, "secucode = ?", m.Maincode())
	return classify.Second
}

// RankAfterServe 任职后回报排名
func (m Manager) RankAfterServe() string {
	var (
		rank string
		rp funds.RankPerformance
	)
	db = connector.GetDB()
	sub := db.Model(&rp).Select("max(id)").Where("secucode = ? and cycle = 999", m.Secucode).Order("-date").Limit(1)
	db.Model(&rp).Select("rank").Where("id = (?)", sub).Find(&rank)
	return rank
}

func (m *Manager) Info() (resp map[string]interface{}) {
	resp = make(map[string]interface{})
	setup := m.EstablishedDate()
	name, _, _, start := m.ServeInfo()
	resp["setup"] = setup
	resp["start"] = start
	resp["manager"] = name
	resp["managed_scale"] = m.ManagedScale()
	resp["scale"] = m.Scale()
	resp["scale_type"] = m.Nature("scale")
	resp["style"] = m.Nature("style")
	resp["industry"] = m.Nature("industry")
	resp["industry_style"] = m.Nature("industry_style")
	resp["stock"] = m.Stock()
	resp["classify"] = m.Classify()
	return
}


func Managed(ctx context.Context, secucode string) []map[string]interface{} {
	var (
		m Manager
		managed []string
		ret chan map[string]interface{}
		resp []map[string]interface{}
		wg sync.WaitGroup
	)
	db = connector.GetDB()
	m = Manager{Secucode: secucode}
	managed = m.Managed()
	ret = make(chan map[string]interface{}, len(managed))
	for _, fund := range managed {
		wg.Add(1)
		go func(f string) {
			fund := f
			var info = map[string]interface{}{}
			var m = Manager{Secucode: fund}
			_, perf, _, d := m.ServeInfo()
			info["secucode"] = fund
			info["serve_date"] = d
			info["return_after"] = perf
			info["scale"] = m.Scale()
			ins := m.FundInstance()
			info["secuabbr"] = ins.Secuabbr
			info["launch"] = m.EstablishedDate()
			info["classify"] = m.Classify()
			info["rank"] = utils.PeriodPerformanceRank(ctx, fund, d)
			ret <- info
			wg.Done()
		}(fund)

	}
	wg.Wait()
	close(ret)
	resp = []map[string]interface{}{}
	for r := range ret {
		resp = append(resp, r)
	}
	return resp
}