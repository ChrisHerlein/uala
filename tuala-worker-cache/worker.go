package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	contentEnums "github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	userEnums "github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
)

func main() {
	// start config
	cfg := &Config{}
	cfg.Init()

	// connect to beanstalk
	qh, err := setQueueHandler(cfg.BeanstalkHost)
	if err != nil {
		panic(err)
	}

	// connect to redis
	rh, err := setCacheHandler(cfg.RedisHost)
	if err != nil {
		panic(err)
	}

	// start processing engine
	eng, err := newEngine(cfg, rh)
	if err != nil {
		panic(err)
	}

	// loop for commands
	var chNewFollow = make(chan message)
	var chNewContent = make(chan message)
	var chReadPage = make(chan message)

	go qh.readFromTube(userEnums.QueueRecreateFeedFollow, chNewFollow)
	go qh.readFromTube(contentEnums.QueueRecreateFeedNewContent, chNewContent)
	go qh.readFromTube(contentEnums.QueueRecreateFeedPageRead, chReadPage)

	go eng.awaitAndProcess(chNewFollow, eng.recreateByFollow)
	go eng.awaitAndProcess(chNewContent, eng.recreateByContent)
	go eng.awaitAndProcess(chReadPage, eng.pageRead)

	fmt.Println("Ready for use")

	// await for end
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
