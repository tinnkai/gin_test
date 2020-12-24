package mysql_models

import (
	"time"
)

type TicketRecord struct {
	Id         int       `gorm:"column(id);auto"`
	TicketId   int       `gorm:"column(ticket_id);null"`
	UserId     int       `gorm:"column(user_id);null"`
	GetTime    time.Time `gorm:"column(get_time);type(datetime)"`
	StartTime  time.Time `gorm:"column(start_time);type(datetime);null"`
	EndTime    time.Time `gorm:"column(end_time);type(datetime);null"`
	OrderId    uint      `gorm:"column(order_id);null"`
	UseTime    time.Time `gorm:"column(use_time);type(datetime);null"`
	Status     string    `gorm:"column(status);size(20);null" description:"NOTUSED | USED"`
	UpdateTime time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time `gorm:"column(create_time);type(datetime);null"`
}

func (t *TicketRecord) TableName() string {
	return "gin_ticket_record"
}

// 查询优惠券记录
func (t *TicketRecord) GetRecordListByUserId(userId int64) ([]TicketRecord, error) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	var ticketRecordList []TicketRecord
	err := db.Table(t.TableName()).Where("user_id = ? AND start_time <= ? AND end_time >= ? AND status=?", userId, nowTime, nowTime, "NOTUSED").
		Find(&ticketRecordList).Error

	return ticketRecordList, err
}

// 查询优惠券记录
func (t *TicketRecord) GetRecordInfoById(id int, userId int64) (TicketRecord, error) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	var ticketRecord TicketRecord
	err := db.Where("user_id = ? AND user_id = ? AND start_time <= ? AND end_time >= ? AND status=?", id, userId, nowTime, nowTime, "NOTUSED").
		First(&ticketRecord).Error

	return ticketRecord, err
}

// 更新优惠券使用记录
func (t *TicketRecord) UpdateRecordInfoById(id int, userId int64) error {
	if err := db.Model(t).Where("id = ? AND user_id = ? ", id, userId).Updates(t).Error; err != nil {
		return err
	}

	return nil
}
