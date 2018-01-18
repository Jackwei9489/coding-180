package main

import (
	"imooc.com/ccmouse/learngo/crawler/engine"
	"imooc.com/ccmouse/learngo/crawler/persist"
	"imooc.com/ccmouse/learngo/crawler/scheduler"
	"imooc.com/ccmouse/learngo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
