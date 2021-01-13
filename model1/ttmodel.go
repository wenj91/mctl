package model1

import (
	"encoding/json"

	"github.com/wenj91/gobatis"
)

type (
	TtModel interface {
		Insert(conn gobatis.GoBatis, data *Tt) (id int64, affected int64, err error)
		InsertSelective(conn gobatis.GoBatis, data *Tt) (id int64, affected int64, err error)
		FindOne(conn gobatis.GoBatis, id int64) (*Tt, error)
		Update(conn gobatis.GoBatis, data *Tt) (affected int64, err error)
		UpdateSelective(conn gobatis.GoBatis, data *Tt) (affected int64, err error)
		Delete(conn gobatis.GoBatis, id int64) (affected int64, err error)
	}

	defaultTtModel struct {
		table string
	}

	Tt struct {
		Id        int64 `field:"id" json:"id"`
		Aid       int64 `field:"aid" json:"aid"`
		C         int64 `field:"c" json:"c"`
		NewColumn int64 `field:"new_column" json:"newColumn"`
	}
)

func (m *Tt) ToString() string {
	str := ""

	bs, err := json.Marshal(m)
	if nil == err {
		str = string(bs)
	}

	return str
}

func NewTtModel() TtModel {
	return &defaultTtModel{
		table: "`tt`",
	}
}

func (m *defaultTtModel) method(mt string) string {
	return "TtMapper." + mt
}

func (m *defaultTtModel) Insert(conn gobatis.GoBatis, data *Tt) (id int64, affected int64, err error) {
	id, affected, err = conn.Insert(m.method("save"), data)
	return
}

func (m *defaultTtModel) InsertSelective(conn gobatis.GoBatis, data *Tt) (id int64, affected int64, err error) {
	id, affected, err = conn.Insert(m.method("saveSelective"), data)
	return
}

func (m *defaultTtModel) FindOne(conn gobatis.GoBatis, id int64) (*Tt, error) {
	var resp *Tt
	err := conn.Select(m.method("findOne"), map[string]interface{}{
		"id": id,
	})(&resp)
	return resp, err
}

func (m *defaultTtModel) Update(conn gobatis.GoBatis, data *Tt) (affected int64, err error) {
	affected, err = conn.Update(m.method("update"), data)
	return
}

func (m *defaultTtModel) UpdateSelective(conn gobatis.GoBatis, data *Tt) (affected int64, err error) {
	affected, err = conn.Update(m.method("updateSelective"), data)
	return
}

func (m *defaultTtModel) Delete(conn gobatis.GoBatis, id int64) (affected int64, err error) {
	affected, err = conn.Delete(m.method("delete"), map[string]interface{}{
		"id": id,
	})
	return
}
