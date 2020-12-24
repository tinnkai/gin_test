package mysql_models

import (
	"time"
)

type OrderTicketSnapshot struct {
	Id         int       `gorm:"column(id);auto"`
	OrderId    uint      `gorm:"column(order_id)"`
	UserId     int64     `gorm:"column(user_id);null"`
	RecordId   int       `orgormm:"column(record_id);null"`
	Content    string    `gorm:"column(content);type(text);null"`
	UpdateTime time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time `gorm:"column(create_time);type(datetime);null"`
}

func (t *OrderTicketSnapshot) TableName() string {
	return "gin_order_ticket_snapshot"
}

// 保存优惠券快照
func (o *OrderTicketSnapshot) SaveOrderTicketSnapshot() error {
	return db.Create(o).Error
}
