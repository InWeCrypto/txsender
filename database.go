package txsender

import "time"

// Order .
type Order struct {
	ID          string     `xorm:"pk"`
	TX          string     `xorm:"index"`
	From        string     `xorm:"index(from_to)"`
	To          string     `xorm:"index(from_to)"`
	Asset       string     `xorm:"notnull"`
	Value       string     `xorm:"notnull"`
	Block       int64      `xorm:"notnull index default (-1)"`
	CreateTime  time.Time  `xorm:"TIMESTAMP notnull"`
	SendTime    *time.Time `xorm:"TIMESTAMP"`
	ConfirmTime *time.Time `xorm:"TIMESTAMP"`
}
