package system

import ("github.com/AsynkronIT/protoactor-go/actor"
		"github.com/MeridianHoldings/learning/pkg/client"
		"github.com/MeridianHoldings/learning/pkg/command"
	"log"
	//"time"
)

type Actor struct {
	Client client.FileClient
}

func (state *Actor) Receive(context actor.Context)  {
	switch msg := context.Message().(type) {
	case *command.ReadFile:
		log.Println("Reading file ...")
		d, err := state.Client.Read(msg.Name)
		log.Printf("err: %v file Contents : %v" , err, d)
	case *command.WriteFile:
		log.Print("Received write...")
		err := state.Client.Write(msg.Data, msg.Name)
		log.Printf("err: %v", err)
	case *actor.Started:
		log.Print("actor started...")
		data := []map[string] interface{}{
			map[string]interface{}{
				"Customer": "Mikes Milkshakes",
				"Grading": "A+",
				"Location": "Alberton",
				"Customer No": "mk01",
			},
			map[string] interface{}{
				"Customer": "Spur",
				"Grading": "A++",
				"Location": "Greenstone",
				"Customer No": "sp03",
			},
			map[string] interface{}{
				"Customer": "Del Forno",
				"Grading": "B",
				"Location": "Kelvin",
				"Customer No": "df13",
			},
			map[string] interface{}{
				"Customer": "Mikes Kitchen",
				"Grading": "D",
				"Location": "Johannesburg",
				"Customer No": "mk13",
			},
			map[string] interface{}{
				"Customer": "Rocco Mammas",
				"Grading": "A",
				"Location": "ParkTown",
				"Customer No": "rm98",
			},
		}
		//time.Sleep(5 * time.Second)
		context.Self().Tell(&command.WriteFile{data, "test.csv"})
		log.Print("Sending message to write json file")
		context.Self().Tell(&command.ReadFile{"test.csv"})

		//actor.NewPID("127.0.0.1:9311", "Remote$system").Tell()
	}
}