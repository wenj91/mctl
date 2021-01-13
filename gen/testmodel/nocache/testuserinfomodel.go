package nocache

import (
	"encoding/json"
	"time"

	"github.com/wenj91/gobatis"
)

type (
	TestUserInfoModel interface {
		Insert(conn gobatis.GoBatis, data *TestUserInfo) (id int64, affected int64, err error)
		InsertSelective(conn gobatis.GoBatis, data *TestUserInfo) (id int64, affected int64, err error)
		FindOne(conn gobatis.GoBatis, id int64) (*TestUserInfo, error)
		FindOneByNanosecond(conn gobatis.GoBatis, nanosecond int64) (*TestUserInfo, error)
		FindSelective(conn gobatis.GoBatis, data *TestUserInfo) ([]*TestUserInfo, error)
		Update(conn gobatis.GoBatis, data *TestUserInfo) (affected int64, err error)
		UpdateSelective(conn gobatis.GoBatis, data *TestUserInfo) (affected int64, err error)
		Delete(conn gobatis.GoBatis, id int64) (affected int64, err error)
	}

	defaultTestUserInfoModel struct {
		table string
	}

	TestUserInfo struct {
		Id         int64              `field:"id" json:"id"`
		Nanosecond int64              `field:"nanosecond" json:"nanosecond"`
		Data       string             `field:"data" json:"data"`
		Content    gobatis.NullString `field:"content" json:"content"`
		CreateTime time.Time          `field:"create_time" json:"createTime"`
		UpdateTime time.Time          `field:"update_time" json:"updateTime"`
	}
)

func (m *TestUserInfo) ToString() string {
	str := ""

	bs, err := json.Marshal(m)
	if nil == err {
		str = string(bs)
	}

	return str
}

func NewTestUserInfoModel() TestUserInfoModel {
	return &defaultTestUserInfoModel{
		table: "`test_user_info`",
	}
}

func (m *defaultTestUserInfoModel) method(mt string) string {
	return "TestUserInfoMapper." + mt
}

func (m *defaultTestUserInfoModel) Insert(conn gobatis.GoBatis, data *TestUserInfo) (id int64, affected int64, err error) {
	id, affected, err = conn.Insert(m.method("save"), data)
	return
}

func (m *defaultTestUserInfoModel) InsertSelective(conn gobatis.GoBatis, data *TestUserInfo) (id int64, affected int64, err error) {
	id, affected, err = conn.Insert(m.method("saveSelective"), data)
	return
}

func (m *defaultTestUserInfoModel) FindOne(conn gobatis.GoBatis, id int64) (*TestUserInfo, error) {
	var resp *TestUserInfo
	err := conn.Select(m.method("findOne"), map[string]interface{}{
		"id": id,
	})(&resp)
	return resp, err
}

func (m *defaultTestUserInfoModel) FindOneByNanosecond(conn gobatis.GoBatis, nanosecond int64) (*TestUserInfo, error) {
	var resp *TestUserInfo
	err := conn.Select(m.method("findOneByNanosecond"), map[string]interface{}{
		"Nanosecond": nanosecond,
	})(&resp)
	return resp, err
}

func (m *defaultTestUserInfoModel) FindSelective(conn gobatis.GoBatis, data *TestUserInfo) ([]*TestUserInfo, error) {
	resp := make([]*TestUserInfo, 0)
	err := conn.Select(m.method("findSelective"), data)(&resp)
	return resp, err
}

func (m *defaultTestUserInfoModel) Update(conn gobatis.GoBatis, data *TestUserInfo) (affected int64, err error) {
	affected, err = conn.Update(m.method("update"), data)
	return
}

func (m *defaultTestUserInfoModel) UpdateSelective(conn gobatis.GoBatis, data *TestUserInfo) (affected int64, err error) {
	affected, err = conn.Update(m.method("updateSelective"), data)
	return
}

func (m *defaultTestUserInfoModel) Delete(conn gobatis.GoBatis, id int64) (affected int64, err error) {
	affected, err = conn.Delete(m.method("delete"), map[string]interface{}{
		"id": id,
	})
	return
}
