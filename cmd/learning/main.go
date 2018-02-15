package main

import ("log"
		"github.com/AsynkronIT/protoactor-go/actor"
		//"github.com/MeridianHoldings/learning/pkg/jsonclient"
		"github.com/MeridianHoldings/learning/pkg/system"
		"github.com/AsynkronIT/protoactor-go/remote"
		console "github.com/AsynkronIT/goconsole"
	"github.com/MeridianHoldings/learning/pkg/csvclient"
)

func main(){
	log.Printf("running")
	actorProps := actor.FromInstance(&system.Actor{
		Client: &csvclient.Client{},
	})

	_, _ = actor.SpawnNamed(actorProps, "system")
	console.ReadLine()
}




