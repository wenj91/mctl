# Mctl Model

mctl model 为go-zero生成github.com/wenj91/gobatis模板代码工具，主要代码实现来源于`github.com/tal-tech/go-zero/tools/goctl/model`，目前仅支持识别mysql ddl进行model层代码生成，通过命令行或者idea插件（即将支持）。

## 快速开始

* mctl安装

	```bash
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/wenj91/mctl
	```

* 通过ddl生成

    ```shell script
    mctl model mysql ddl -src="./*.sql" -dir="./sql/model" -c
    ```

    执行上述命令后即可快速生成CURD代码。

    ```Plain Text
    ./model
	├── mappers
	│   └── testuserinfomapper.xml
	├── testuserinfomodel.go
	└── vars.go
    ```

* 通过datasource生成

    ```shell script
    mctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*"  -dir="./model"
    ```

* 生成代码示例
  
	```go
	package nocache

	import (
		"encoding/json"
		"time"

		"github.com/wenj91/gobatis"
	)

	type (
		TestUserInfoFindResult struct {
			testUserInfos []*TestUserInfo
		}

		TestUserInfoModel interface {
			WithConn(conn gobatis.GoBatis) TestUserInfoModel
			Insert(data *TestUserInfo) (id int64, affected int64, err error)
			InsertSelective(data *TestUserInfo) (id int64, affected int64, err error)
			FindOne(id int64) (*TestUserInfo, error)
			FindOneByNanosecond(nanosecond int64) (*TestUserInfo, error)
			FindSelective(data *TestUserInfoSelective) (*TestUserInfoFindResult, error)
			Update(data *TestUserInfo) (affected int64, err error)
			UpdateSelective(data *TestUserInfo) (affected int64, err error)
			Delete(id int64) (affected int64, err error)
		}

		defaultTestUserInfoModel struct {
			conn  gobatis.GoBatis
			table string
		}

		TestUserInfo struct {
			Id         *int64     `field:"id" json:"id"`
			Nanosecond *int64     `field:"nanosecond" json:"nanosecond"`
			Data       *string    `field:"data" json:"data"`
			Content    *string    `field:"content" json:"content"`
			CreateTime *time.Time `field:"create_time" json:"createTime"`
			UpdateTime *time.Time `field:"update_time" json:"updateTime"`
		}

		TestUserInfoSelective struct {
			Id              *int64
			Nanosecond      *int64
			Data            *string
			Content         *string
			StartCreateTime *time.Time
			EndCreateTime   *time.Time
			CreateTime      *time.Time
			StartUpdateTime *time.Time
			EndUpdateTime   *time.Time
			UpdateTime      *time.Time
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

	func newTestUserInfoFindResult(testUserInfos []*TestUserInfo) *TestUserInfoFindResult {
		return &TestUserInfoFindResult{
			testUserInfos: testUserInfos,
		}
	}

	func (r *TestUserInfoFindResult) List() []*TestUserInfo {
		return r.testUserInfos
	}

	func (r *TestUserInfoFindResult) One() *TestUserInfo {
		if len(r.testUserInfos) == 0 {
			return nil
		}

		return r.testUserInfos[0]
	}

	func NewTestUserInfoModel(conn gobatis.GoBatis) TestUserInfoModel {
		return &defaultTestUserInfoModel{
			conn:  conn,
			table: "`test_user_info`",
		}
	}

	func (m *defaultTestUserInfoModel) method(mt string) string {
		return "TestUserInfoMapper." + mt
	}

	func (m *defaultTestUserInfoModel) WithConn(conn gobatis.GoBatis) TestUserInfoModel {
		return &defaultTestUserInfoModel{
			conn:  conn,
			table: "test_user_info",
		}
	}

	func (m *defaultTestUserInfoModel) Insert(data *TestUserInfo) (id int64, affected int64, err error) {
		id, affected, err = m.conn.Insert(m.method("save"), data)
		return
	}

	func (m *defaultTestUserInfoModel) InsertSelective(data *TestUserInfo) (id int64, affected int64, err error) {
		id, affected, err = m.conn.Insert(m.method("saveSelective"), data)
		return
	}

	func (m *defaultTestUserInfoModel) FindOne(id int64) (*TestUserInfo, error) {
		var resp *TestUserInfo
		err := m.conn.Select(m.method("findOne"), map[string]interface{}{
			"Id": id,
		})(&resp)
		return resp, err
	}

	func (m *defaultTestUserInfoModel) FindOneByNanosecond(nanosecond int64) (*TestUserInfo, error) {
		var resp *TestUserInfo
		err := m.conn.Select(m.method("findOneByNanosecond"), map[string]interface{}{
			"Nanosecond": nanosecond,
		})(&resp)
		return resp, err
	}

	func (m *defaultTestUserInfoModel) FindSelective(data *TestUserInfoSelective) (*TestUserInfoFindResult, error) {
		resp := make([]*TestUserInfo, 0)
		err := m.conn.Select(m.method("findSelective"), data)(&resp)
		return &TestUserInfoFindResult{
			testUserInfos: resp,
		}, err
	}

	func (m *defaultTestUserInfoModel) Update(data *TestUserInfo) (affected int64, err error) {
		affected, err = m.conn.Update(m.method("update"), data)
		return
	}

	func (m *defaultTestUserInfoModel) UpdateSelective(data *TestUserInfo) (affected int64, err error) {
		affected, err = m.conn.Update(m.method("updateSelective"), data)
		return
	}

	func (m *defaultTestUserInfoModel) Delete(id int64) (affected int64, err error) {
		affected, err = m.conn.Delete(m.method("delete"), map[string]interface{}{
			"Id": id,
		})
		return
	}


	```

* 生成mapper代码示例

	```xml
	<?xml version="1.0" encoding="utf-8"?>
	<!DOCTYPE mapper PUBLIC "gobatis"
			"https://raw.githubusercontent.com/wenj91/gobatis/master/gobatis.dtd">
	<mapper namespace="TestUserInfoMapper">
		<sql id="Base_Column_List">
			id,nanosecond,data,content,create_time,update_time
		</sql>
		<insert id="save">
			insert into test_user_info (id, nanosecond, data, content, create_time, update_time)
			values (#{id},#{nanosecond},#{data},#{content},#{create_time},#{update_time})
		</insert>
		<insert id="saveSelective">
			insert into test_user_info
			<trim prefix="(" suffix=")" suffixOverrides=",">
				<if test="Id != nil">
					id,  
				</if>
				<if test="Nanosecond != nil">
					nanosecond,  
				</if>
				<if test="Data != nil">
					data,  
				</if>
				<if test="Content != nil">
					content,  
				</if>
				<if test="CreateTime != nil">
					create_time,  
				</if>
				<if test="UpdateTime != nil">
					update_time,  
				</if>
			</trim>
			<trim prefix="values (" suffix=")" suffixOverrides=",">
				<if test="Id != nil">
					#{Id},
				</if>
				<if test="Nanosecond != nil">
					#{Nanosecond},
				</if>
				<if test="Data != nil">
					#{Data},
				</if>
				<if test="Content != nil">
					#{Content},
				</if>
				<if test="CreateTime != nil">
					#{CreateTime},
				</if>
				<if test="UpdateTime != nil">
					#{UpdateTime},
				</if>
			</trim>
		</insert>
		<update id="update">
			update test_user_info
			set nanosecond = #{Nanosecond},
			data = #{Data},
			content = #{Content}
			where id = #{Id}
		</update>
		<update id="updateSelective">
			update test_user_info
			<set>
				<if test="Nanosecond != nil">
					nanosecond = #{Nanosecond},
				</if>
				<if test="Data != nil">
					data = #{Data},
				</if>
				<if test="Content != nil">
					content = #{Content},
				</if>
			</set>
			where id = #{Id}
		</update>
		<delete id="delete">
			delete from test_user_info
			where id = #{Id}
		</delete>
		<select id="findOne" resultType="struct">
			select 
				<include refid="Base_Column_List" />
			from test_user_info
			where id = #{Id}
			limit 1
		</select>
		<select id="findOneByNanosecond" resultType="struct">
			select 
				<include refid="Base_Column_List" />
			from test_user_info
			where nanosecond = #{Nanosecond}
			limit 1
		</select>
		<select id="findSelective" resultType="structs">
			select 
				<include refid="Base_Column_List" />
			from test_user_info
			<where>
				<if test="Nanosecond != nil">
					and nanosecond = #{Nanosecond}
				</if>
				<if test="Data != nil">
					and data = #{Data}
				</if>
				<if test="Content != nil">
					and content = #{Content}
				</if>
				<if test="StartCreateTime != nil">
					and create_time >= #{StartCreateTime}
				</if>
				<if test="EndCreateTime != nil">
					and create_time <![CDATA[<=]]> #{EndCreateTime}
				</if>
				<if test="CreateTime != nil">
					and create_time = #{CreateTime}
				</if>
				<if test="StartUpdateTime != nil">
					and update_time >= #{StartUpdateTime}
				</if>
				<if test="EndUpdateTime != nil">
					and update_time <![CDATA[<=]]> #{EndUpdateTime}
				</if>
				<if test="UpdateTime != nil">
					and update_time = #{UpdateTime}
				</if>
			</where>
		</select>
	</mapper>
	```

## 用法

```Plain Text
mctl model mysql -h
```

```Plain Text
NAME:
   mctl model mysql - generate mysql model"

USAGE:
   mctl model mysql command [command options] [arguments...]

COMMANDS:
   ddl         generate mysql model from ddl"
   datasource  generate model from datasource"

OPTIONS:
   --help, -h  show help
```

## 生成规则

* 默认规则
  
  我们默认用户在建表时会创建createTime、updateTime字段(忽略大小写、下划线命名风格)且默认值均为`CURRENT_TIMESTAMP`，而updateTime支持`ON UPDATE CURRENT_TIMESTAMP`，对于这两个字段生成`insert`、`update`时会被移除，不在赋值范畴内，当然，如果你不需要这两个字段那也无大碍。
* 带缓存模式
  * ddl

	```shell script
	mctl model mysql -src={patterns} -dir={dir} 
	```

	help

	```
	NAME:
       mctl model mysql ddl - generate mysql model from ddl
    
    USAGE:
       mctl model mysql ddl [command options] [arguments...]
    
    OPTIONS:
       --src value, -s value  the path or path globbing patterns of the ddl
       --dir value, -d value  the target dir
       --style value          the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]
       --idea                 for idea plugin [optional]
	```

  * datasource

	```shell script
	mctl model mysql datasource -url={datasource} -table={patterns}  -dir={dir}
	```

	help

	```
	NAME:
       mctl model mysql datasource - generate model from datasource
    
    USAGE:
       mctl model mysql datasource [command options] [arguments...]
    
    OPTIONS:
       --url value              the data source of database,like "root:password@tcp(127.0.0.1:3306)/database
       --table value, -t value  the table or table globbing patterns in the database
       --dir value, -d value    the target dir
       --style value            the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]
       --idea                   for idea plugin [optional]


	```

	示例用法请参考[用法](./example/generator.sh)
  
	> NOTE: mctl model mysql ddl/datasource 均新增了一个`--style`参数，用于标记文件命名风格。


* 不带缓存模式

  * ddl
  
      ```shell script
        mctl model -src={patterns} -dir={dir}
      ```

  * datasource
  
      ```shell script
        mctl model mysql datasource -url={datasource} -table={patterns}  -dir={dir}
      ```

  or
  * ddl

      ```shell script
        mctl model -src={patterns} -dir={dir}
      ```

  * datasource
  
      ```shell script
        mctl model mysql datasource -url={datasource} -table={patterns}  -dir={dir}
      ```
  
生成代码仅基本的CURD结构。

# 类型转换规则
| mysql dataType | golang dataType | golang dataType(if null&&default null) |
|----------------|-----------------|----------------------------------------|
| bool           | int64           | *int64                                 |
| boolean        | int64           | *int64                                 |
| tinyint        | int64           | *int64                                 |
| smallint       | int64           | *int64                                 |
| mediumint      | int64           | *int64                                 |
| int            | int64           | *int64                                 |
| integer        | int64           | *int64                                 |
| bigint         | int64           | *int64                                 |
| float          | float64         | *float64                               |
| double         | float64         | *float64                               |
| decimal        | float64         | *float64                               |
| date           | time.Time       | *time.Time                             |
| datetime       | time.Time       | *time.Time                             |
| timestamp      | time.Time       | *time.Time                             |
| time           | string          | *string                                |
| year           | time.Time       | *int64                                 |
| char           | string          | *string                                |
| varchar        | string          | *string                                |
| binary         | string          | *string                                |
| varbinary      | string          | *string                                |
| tinytext       | string          | *string                                |
| text           | string          | *string                                |
| mediumtext     | string          | *string                                |
| longtext       | string          | *string                                |
| enum           | string          | *string                                |
| set            | string          | *string                                |
| json           | string          | *string                                |
