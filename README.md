
## 分布式爬虫（以www.zhenai.com为例）


>存储使用的是 elasticsearch
docker run -d -p 9200:9200 elasticsearch

运行：
go run crawler/main.go

## crawler_distributed
在单机爬虫基本之上，扩展的并发版爬虫版本。
使用json-rpc实现多个节点调用
主要扩展内容：存储服务（一个），Work服务（网页抓取）多个，引擎（一个）
运行步骤：

1. 存储
> go run crawler_distributed/persist/server/saver_main.go （占用1234端口）

2，Worker
> go run crawler_distributed/worker/server/worker_main.go --port=9003

3, 引擎
> go run crawler_distributed/distributed.go --hosts="192.168.1.8:9002,192.168.1.8:9004,:9002,:9003"



