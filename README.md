# Mctl Model

mctl model 为[mybatis-plus](https://github.com/baomidou/mybatis-plus)生成QueryWrapper代码工具

主要代码实现来源于`github.com/tal-tech/go-zero/tools/goctl/model`，目前仅支持识别mysql ddl进行model层代码生成，通过命令行或者idea插件（即将支持）。

## 快速开始

* mctl安装

	```bash
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/wenj91/mctl
	```

* 通过ddl生成

    ```shell script
    mctl model mysql ddl -src="./*.sql" -pkg="com.github.wenj91.test" -dir="./sql/model" -c
    ```

    执行上述命令后即可快速生成CURD代码。

    ```Plain Text
    ./model
	└──TestUserInfoQuery.java
    ```

* 通过datasource生成

    ```shell script
    mctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -pkg="com.github.wenj91.test" -table="*"  -dir="./model"
    ```

* 生成代码示例
  
	```java
    // Code generated by [mctl-zzinfo](https://github.com/wenj91/mctl/tree/zzinfo), DO NOT EDIT.
    package com.github.wenj91.test;
    
    import cn.zzstc.sbp.common.db.AbstractQuery;
    import cn.zzstc.sbp.common.db.Cond;
    import cn.zzstc.sbp.common.db.Op;
    
    
    
    public class TestUserInfoQuery extends AbstractQuery {
    
        private TestUserInfo Query() {}
    
        @Override
        public String table() {
            return "test_user_info";
        }
    
    
        public static TestUserInfoQuery query() {
            return new TestUserInfoQuery();
        }
    
    
        public TestUserInfoQuery id(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.EQ, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idEQ(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.EQ, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idNEQ(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.NEQ, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idIn(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.In, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idNotIn(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.NotIn, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idGT(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.GT, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idGTE(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.GTE, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idLT(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.LT, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idLTE(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.LTE, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idIsNull(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.IsNull, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idNotNull(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.NotNull, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idOrderByAsc(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.OrderByAsc, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery idOrderByDesc(Long id) {
            if (null != id) {
                this.push(Cond.of("id", Op.OrderByDesc, id));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecond(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.EQ, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondEQ(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.EQ, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondNEQ(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.NEQ, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondIn(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.In, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondNotIn(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.NotIn, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondGT(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.GT, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondGTE(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.GTE, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondLT(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.LT, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondLTE(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.LTE, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondIsNull(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.IsNull, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondNotNull(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.NotNull, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondOrderByAsc(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.OrderByAsc, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery nanosecondOrderByDesc(Long nanosecond) {
            if (null != nanosecond) {
                this.push(Cond.of("nanosecond", Op.OrderByDesc, nanosecond));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery data(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.EQ, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataEQ(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.EQ, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataNEQ(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.NEQ, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataIn(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.In, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataNotIn(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.NotIn, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataGT(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.GT, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataGTE(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.GTE, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataLT(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.LT, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataLTE(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.LTE, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataContains(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.Contains, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataHasPrefix(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.HasPrefix, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataHasSuffix(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.HasSuffix, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataIsNull(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.IsNull, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataNotNull(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.NotNull, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataOrderByAsc(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.OrderByAsc, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery dataOrderByDesc(String data) {
            if (null != data) {
                this.push(Cond.of("data", Op.OrderByDesc, data));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery content(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.EQ, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentEQ(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.EQ, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentNEQ(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.NEQ, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentIn(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.In, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentNotIn(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.NotIn, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentGT(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.GT, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentGTE(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.GTE, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentLT(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.LT, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentLTE(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.LTE, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentContains(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.Contains, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentHasPrefix(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.HasPrefix, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentHasSuffix(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.HasSuffix, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentIsNull(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.IsNull, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentNotNull(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.NotNull, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentOrderByAsc(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.OrderByAsc, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery contentOrderByDesc(String content) {
            if (null != content) {
                this.push(Cond.of("content", Op.OrderByDesc, content));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTime(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.EQ, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeEQ(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.EQ, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeNEQ(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.NEQ, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeIn(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.In, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeNotIn(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.NotIn, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeGT(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.GT, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeGTE(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.GTE, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeLT(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.LT, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeLTE(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.LTE, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeContains(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.Contains, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeHasPrefix(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.HasPrefix, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeHasSuffix(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.HasSuffix, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeIsNull(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.IsNull, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeNotNull(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.NotNull, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeOrderByAsc(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.OrderByAsc, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery createTimeOrderByDesc(String createTime) {
            if (null != createTime) {
                this.push(Cond.of("create_time", Op.OrderByDesc, createTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTime(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.EQ, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeEQ(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.EQ, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeNEQ(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.NEQ, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeIn(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.In, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeNotIn(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.NotIn, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeGT(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.GT, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeGTE(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.GTE, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeLT(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.LT, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeLTE(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.LTE, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeContains(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.Contains, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeHasPrefix(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.HasPrefix, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeHasSuffix(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.HasSuffix, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeIsNull(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.IsNull, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeNotNull(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.NotNull, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeOrderByAsc(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.OrderByAsc, updateTime));
    		}
    
            return this;
        }
    
        public TestUserInfoQuery updateTimeOrderByDesc(String updateTime) {
            if (null != updateTime) {
                this.push(Cond.of("update_time", Op.OrderByDesc, updateTime));
    		}
    
            return this;
        }
        
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
	mctl model mysql -src={patterns} -pkg={pkg} -dir={dir} 
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
       --pkg value, -p value  the java package
       --style value          the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]
       --idea                 for idea plugin [optional]
	```

  * datasource

	```shell script
	mctl model mysql datasource -url={datasource} -pkg={pkg} -table={patterns}  -dir={dir}
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
       --pkg value, -p value  the java package
       --style value            the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]
       --idea                   for idea plugin [optional]


	```

	示例用法请参考[用法](./example/generator.sh)
  
	> NOTE: mctl model mysql ddl/datasource 均新增了一个`--style`参数，用于标记文件命名风格。


* 不带缓存模式

  * ddl
  
      ```shell script
        mctl model -src={patterns} -dir={dir} -pkg={pkg}
      ```

  * datasource
  
      ```shell script
        mctl model mysql datasource -url={datasource} -table={patterns}  -dir={dir} -pkg={pkg}
      ```

  or
  * ddl

      ```shell script
        mctl model -src={patterns} -dir={dir} -pkg={pkg}
      ```

  * datasource
  
      ```shell script
        mctl model mysql datasource -url={datasource} -table={patterns}  -dir={dir} -pkg={pkg}
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
