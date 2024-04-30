package main

import (
	"fmt"
	"novachat_engine/pkg/etcd/leader"
	"novachat_engine/pkg/log"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("node name invalid")
		return
	}

	op := log.DefaultOptions()

	log.Configure(op)
	log.Sync()

	log.DebugEnabled()

	allScopes := log.Scopes()
	scopes := allScopes[log.DefaultScopeName]
	if scopes != nil {
		scopes.SetOutputLevel(log.StringToLevel("info"))
	}

	leaderClient, err := leader.NewLeader(leader.Config{
		EndPoints:   []string{"127.0.0.1:2379"},
		ClusterName: "leader",
		NodeName:    os.Args[1],
	})
	if err != nil {
		panic(err.Error())
	}

	result := make(chan leader.Result, 1)
	go func() {
		defer close(result)
		err = leaderClient.LeasesLock(result)
		if err != nil {
			panic(err.Error())
		}
	}()

	for v := range result {
		if v == leader.ResultNone {
			break
		}

		log.Infof("result:%v", v)
	}
}
