package main

import (
	"gopkg.in/olivere/elastic.v5"
	"distributed_crawler/crawler_distributed/rpcsupport"
	"distributed_crawler/crawler_distributed/persist"
	"log"
	"distributed_crawler/crawler_distributed/config"
	"fmt"
)

func main() {
	fmt.Println("Imte Saver Start At:", config.ItemSaverPort)
	log.Fatal(serveRpc(config.ItemSaverPort, config.ElasticIndex))
}
// 启动一个 存储节点
func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		},
	)
}
