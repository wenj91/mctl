# Mctl Model

mctl model 为[ent](https://github.com/ent/ent)生成模板代码工具

主要代码实现来源于`github.com/tal-tech/go-zero/tools/goctl/model`，目前仅支持识别mysql ddl进行model层代码生成，通过命令行或者idea插件（即将支持）。

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
	└──testuserinfomodel.go
    ```

* 通过datasource生成

    ```shell script
    mctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*"  -dir="./model"
    ```

* 生成代码示例
  
	```go
    package schema
    
    import (
        "entgo.io/ent"
        "entgo.io/ent/schema/field"
    )
    
    type Account struct {
        ent.Schema
    }
    
    func (Account) Fields() []ent.Field {
        return []ent.Field{
            field.Int("id"),
            field.String("accountNumber"),
            field.Int("userId"),
            field.Float("available"),
            field.Float("dailyLimit"),
            field.String("currency"),
            field.Time("createAt"),
        }
    }
    
    func (Account) Edges() []ent.Edge {
        return nil
    }

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
