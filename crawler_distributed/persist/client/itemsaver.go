package client

import (
	"log"
	"distributed_crawler/crawler/engine"
	"distributed_crawler/crawler_distributed/rpcsupport"
)
// json rpc client
// 给客户端使用的代码 将数据通过 rpc 传送至 rpc-server
func ItemSaver(host string) (chan engine.Item, error) {
	ch := make(chan engine.Item, 1024)

	rpc, err := rpcsupport.NewClient(host)
	go func() {
		itemCount := 0
		for item := range ch {
			itemCount++
			log.Printf("Item Saver: Got Item #%d: %v", itemCount, item)

			result := ""
			rpc.Call("ItemSaverService.Save", item, &result)

			if err != nil {
				log.Printf("Item Saver: Save error: %s", err)
			}
		}
	}()
	return ch, err
}
