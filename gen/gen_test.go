package gen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/tools/goctl/config"
)

var (
	source = "CREATE TABLE `test_user_info` (\n  `id` bigint NOT NULL AUTO_INCREMENT,\n  `nanosecond` bigint NOT NULL DEFAULT '0',\n  `data` varchar(255) DEFAULT '',\n  `content` json DEFAULT NULL,\n  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,\n  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `nanosecond_unique` (`nanosecond`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
)

func TestModel(t *testing.T) {
	logx.Disable()
	_ = Clean()
	dir, _ := filepath.Abs("./testmodel")
	noCacheDir := filepath.Join(dir, "nocache")
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	g, err := NewDefaultGenerator(noCacheDir, "com.github.wenj91.test", &config.Config{
		NamingFormat: "goZero",
	})
	assert.Nil(t, err)

	err = g.StartFromDDL(source)
	assert.Nil(t, err)
	assert.True(t, func() bool {
		_, err := os.Stat(filepath.Join(noCacheDir, "TestUserInfoQuery.java"))
		return err == nil
	}())
}
