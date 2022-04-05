package urls

import (
	"fmt"
	"slink/src/rds"
	"slink/src/scales"
	"slink/src/sids"
)

func Shorten(originUrl string) (string, error) {
	codeKey := "slink:code:" + originUrl

	// 先尝试从缓存中获取
	cmd := rds.Client.Get(codeKey)
	code, err := cmd.Result()
	if code != "" {
		return wrap(code), nil
	}

	// 如果缓存中没有再生成
	// 首先从发号器获取一个十进制ID
	sid, err := sids.Gen()
	if err != nil {
		return "", fmt.Errorf("generate sid error: %v", err)
	}

	// 将此ID转换为62进制字符串code
	code = scales.DecimalToAny(sid, 62)

	// 保存code与原始链接的映射关系
	rds.Client.Set(codeKey, code, -1)
	urlKey := "slink:url:" + code
	rds.Client.Set(urlKey, originUrl, -1)

	// 拼接域名后返回
	return wrap(code), nil
}

func Expand(code string) (string, error) {
	urlKey := "slink:url:" + code
	cmd := rds.Client.Get(urlKey)
	url, err := cmd.Result()
	if err != nil {
		return "", fmt.Errorf("get result of %v error: %v", cmd, err)
	} else {
		return url, nil
	}
}

func wrap(code string) string {
	return "localhost:8080/" + code
}