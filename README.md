## mubeng 修改(lazypool)

- 修改验证,默认超时时间从 30s 减至 2s
- 验证地址新增 ifconfig.me
- 添加选项 -u 从远端下载配置
  ```
  ./lazypool -u "x.x.x.x/config.yaml"
  ```
- 添加选项 -verify 验证yaml格式和服务器地址
  ```
  ./lazypool -verify
  ```
- 添加选项 -loop 无限验证网络可用性
  ```
  ./lazypool -f proxies.txt --loop
  ```
