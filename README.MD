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
	│   └── ttmapper.xml
	├── ttmodel.go
	└── vars.go
    ```

* 通过datasource生成

    ```shell script
    mctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*"  -dir="./model"
    ```

* 生成代码示例
  
	```go

	package model

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

	```

* 生成mapper代码示例

	```xml
	<?xml version="1.0" encoding="utf-8"?>
	<!DOCTYPE mapper PUBLIC "gobatis"
			"https://raw.githubusercontent.com/wenj91/gobatis/master/gobatis.dtd">
	<mapper namespace="TtMapper">
	
	<sql id="Base_Column_List">
		id,aid,c,new_column
	</sql>


	<insert id="save">
		insert into tt (id, aid, c, new_column)
		values (#{id},#{aid},#{c},#{new_column})
	</insert>

	<insert id="saveSelective">
		insert into tt
		<trim prefix="(" suffix=")" suffixOverrides=",">
		<if test="Id != nil and Id != ''">
			id,  
		</if>
		<if test="Aid != nil and Aid != ''">
			aid,  
		</if>
		<if test="C != nil and C != ''">
			c,  
		</if>
		<if test="NewColumn != nil and NewColumn != ''">
			new_column,  
		</if>
		</trim>
		<trim prefix="values (" suffix=")" suffixOverrides=",">
		<if test="Id != nil and Id != ''">
			#{Id},
		</if>
		<if test="Aid != nil and Aid != ''">
			#{Aid},
		</if>
		<if test="C != nil and C != ''">
			#{C},
		</if>
		<if test="NewColumn != nil and NewColumn != ''">
			#{NewColumn},
		</if>
		</trim>
	</insert>


	<update id="update">
		update tt
		set aid = #{Aid},
		c = #{C},
		new_column = #{NewColumn}
		where id = #{Id}
	</update>

	<update id="updateSelective">
		update tt
		<set>
		<if test="Aid != nil and Aid != ''">
			aid = #{Aid},
		</if>
		<if test="C != nil and C != ''">
			c = #{C},
		</if>
		<if test="NewColumn != nil and NewColumn != ''">
			new_column = #{NewColumn},
		</if>
		</set>
		where id = #{Id}
	</update>


	<delete id="delete">
		delete from tt
		where id = #{Id}
	</delete>

	<select id="findOne" resultType="struct">
		select 
		<include refid="Base_Column_List" />
		from tt
		where id = #{Id}
		limit 1
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
| bool           | int64           | gobatis.NullInt64                          |
| boolean        | int64           | gobatis.NullInt64                          |
| tinyint        | int64           | gobatis.NullInt64                          |
| smallint       | int64           | gobatis.NullInt64                          |
| mediumint      | int64           | gobatis.NullInt64                          |
| int            | int64           | gobatis.NullInt64                          |
| integer        | int64           | gobatis.NullInt64                          |
| bigint         | int64           | gobatis.NullInt64                          |
| float          | float64         | gobatis.NullFloat64                        |
| double         | float64         | gobatis.NullFloat64                        |
| decimal        | float64         | gobatis.NullFloat64                        |
| date           | time.Time       | gobatis.NullTime                           |
| datetime       | time.Time       | gobatis.NullTime                           |
| timestamp      | time.Time       | gobatis.NullTime                           |
| time           | string          | gobatis.NullString                         |
| year           | time.Time       | gobatis.NullInt64                          |
| char           | string          | gobatis.NullString                         |
| varchar        | string          | gobatis.NullString                         |
| binary         | string          | gobatis.NullString                         |
| varbinary      | string          | gobatis.NullString                         |
| tinytext       | string          | gobatis.NullString                         |
| text           | string          | gobatis.NullString                         |
| mediumtext     | string          | gobatis.NullString                         |
| longtext       | string          | gobatis.NullString                         |
| enum           | string          | gobatis.NullString                         |
| set            | string          | gobatis.NullString                         |
| json           | string          | gobatis.NullString                         |
