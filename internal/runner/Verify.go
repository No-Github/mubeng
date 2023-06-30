package runner

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ClashConfig struct {
	ProxyGroups []struct {
		Proxies []string `yaml:"proxies"`
	} `yaml:"proxy-groups"`
}

func Verify_Yaml(file string) {
	// 读取配置文件
	configBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析配置文件
	var config ClashConfig
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	// 提取代理地址
	proxies := []string{}
	for _, group := range config.ProxyGroups {
		for _, proxy := range group.Proxies {
			proxies = append(proxies, proxy)
		}
	}

	// 输出代理地址
	fmt.Println(proxies)

}
