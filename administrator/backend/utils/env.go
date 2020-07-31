package utils

import (
	"os"
)

// GetEnvWithDefault は環境変数を取得する.
// 環境変数を取得できた場合にはそのまま返す.
// 取得出来なかった場合にはdefaultStrを返す.
func GetEnvWithDefault(key string, defaultStr string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultStr
}
