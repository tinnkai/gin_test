package mysql_models

import (
	"gin_test/pkg/utils"
	"sort"
	"time"
)

type Ticket struct {
	Id           int       `gorm:"column(id);auto"`
	Name         string    `gorm:"column(name);size(200)"`
	GetStartTime time.Time `gorm:"column(get_start_time);type(datetime)"`
	GetEndTime   time.Time `gorm:"column(get_end_time);type(datetime);null"`
	StartTime    time.Time `gorm:"column(start_time);type(datetime);null"`
	EndTime      time.Time `gorm:"column(end_time);type(datetime);null"`
	Amount       float64   `gorm:"column(amount);null;digits(5);decimals(2)" description:"优惠券面额"`
	MinAmount    float64   `gorm:"column(min_amount);null;digits(10);decimals(2)" description:"最低使用金额"`
	Limit        int       `gorm:"column(limit);null"`
	Scope        string    `gorm:"column(scope);size(20);null" description:"ALL | GOODS"`
	Describe     string    `gorm:"column(describe);null" description:"描述"`
	Status       string    `gorm:"column(status);size(20);null" description:"ENABELD | DISABLED"`
	IsDel        string    `gorm:"column(is_del);size(1);null" description:"Y|N"`
	UpdateTime   time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime   time.Time `gorm:"column(create_time);type(datetime);null"`
}

// 优惠券
type OrderTicket struct {
	Id                     int            `json:"-"`
	Name                   string         `json:"name"`
	Amount                 float64        `json:"amount"`
	MinAmount              float64        `json:"-"`
	RecordId               int            `json:"record_id"`
	IsSelect               string         `json:"is_select"`
	Status                 string         `json:"status"`
	Scope                  string         `json:"-"`
	TicketGoodsIds         map[int]int    `json:"-"`
	TicketGoodsCategoryIds map[int]int    `json:"-"`
	StartTime              utils.Datetime `json:"start_time"`
	EndTime                utils.Datetime `json:"end_time"`
}

func (t *Ticket) TableName() string {
	return "gin_ticket"
}

// 根据id获取优惠券列表
func GetTicketInfoById(id int) (Ticket, error) {
	var ticket Ticket
	err := db.Where("id= ?", id).First(&ticket).Error
	return ticket, err
}

// 根据用户id获取优惠券列表
func GetTicketListByUserId(userId int64) ([]OrderTicket, error) {
	var orderTicket []OrderTicket

	// 获取优惠券记录
	var ticketRecord *TicketRecord
	ticketRecordList, err := ticketRecord.GetRecordListByUserId(userId)
	if err != nil {
		return orderTicket, err
	}

	if len(ticketRecordList) < 1 {
		return orderTicket, err
	}

	for _, v := range ticketRecordList {
		ticketInfo, err := GetTicketInfoById(v.TicketId)
		if err != nil {
			return orderTicket, err
		}

		// 查询优惠券商品
		var ticketGoods TicketGoods
		ticketGoodsIds, ticketGoodsCategoryIds, err := ticketGoods.GetTicketGoodsListByTicketId(v.TicketId)
		if err != nil {
			return orderTicket, err
		}

		orderTicket = append(orderTicket, OrderTicket{
			Id:                     ticketInfo.Id,
			Name:                   ticketInfo.Name,
			Amount:                 ticketInfo.Amount,
			MinAmount:              ticketInfo.MinAmount,
			RecordId:               v.Id,
			IsSelect:               "no",
			Scope:                  ticketInfo.Scope,
			Status:                 "unavailable",
			TicketGoodsIds:         ticketGoodsIds,
			TicketGoodsCategoryIds: ticketGoodsCategoryIds,
			StartTime:              ticketInfo.StartTime,
			EndTime:                ticketInfo.EndTime,
		})

	}
	// 按价格倒序排序
	sort.SliceStable(orderTicket, func(i, j int) bool {
		return orderTicket[i].Amount > orderTicket[j].Amount
	})
	return orderTicket, err
}

// 根据优惠券记录id和用户id获取优惠券信息
func (t *Ticket) GetTicketInfoById(ticketId int, userId int64) (OrderTicket, error) {
	var orderTicket OrderTicket

	// 获取优惠券记录
	var ticketRecord *TicketRecord
	ticketRecordInfo, err := ticketRecord.GetRecordInfoById(ticketId, userId)
	if err != nil {
		return orderTicket, err
	}
	// 获取优惠券主表信息
	ticketInfo, err := GetTicketInfoById(ticketRecordInfo.TicketId)
	if err != nil {
		return orderTicket, err
	}

	// 查询优惠券商品
	var ticketGoods TicketGoods
	ticketGoodsIds, ticketGoodsCategoryIds, err := ticketGoods.GetTicketGoodsListByTicketId(ticketRecordInfo.TicketId)
	if err != nil {
		return orderTicket, err
	}

	orderTicket = OrderTicket{
		Id:                     ticketInfo.Id,
		Name:                   ticketInfo.Name,
		Amount:                 ticketInfo.Amount,
		MinAmount:              ticketInfo.MinAmount,
		RecordId:               ticketRecordInfo.Id,
		IsSelect:               "no",
		Scope:                  ticketInfo.Scope,
		Status:                 "unavailable",
		TicketGoodsIds:         ticketGoodsIds,
		TicketGoodsCategoryIds: ticketGoodsCategoryIds,
		StartTime:              ticketInfo.StartTime,
		EndTime:                ticketInfo.EndTime,
	}

	return orderTicket, err
}
