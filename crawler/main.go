package main

import (
	"distributed_crawler/crawler/engine"
	"distributed_crawler/crawler/scheduler"
	"distributed_crawler/crawler/persist"
	"distributed_crawler/crawler/zhengai/parser"
)

// 单机，并发
func main() {
	itemChan, err := persist.ItemSaver("profiles")
	if err != nil {
		panic(err)
	}
	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://www.zhenai.com/zhenghun/beijing",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
		//{
		//	Url:       "http://www.zhenai.com/zhenghun/henan",
		//	ParseFunc: parser.ParseCity,
		//},
		//{
		//	Url:   "http://www.zhenai.com/zhenghun",
		//	Parse: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
		//},
	}
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		MaxWorkerCount:   100,
		Scheduler:        &scheduler.QueuedScheduler{},
		ItemChan:         itemChan,
		RequestWorker: engine.Worker,
	}
	e.Run(seed...)
}
