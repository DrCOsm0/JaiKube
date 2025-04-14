package main

import (
	"fmt"
	//"flag"
	"os"
)


func createNodes() {
	fmt.Println("hi")
}

func stopNodes() {

}

func deleteNodes() {

}

func rebootNodes() {

}

func listNodes() {

}

func help() {

}


func main() {
    
	//argsWithProg := os.Args
    //argsWithoutProg := os.Args[1:]

	//key := os.Args[]

	//check for null input
	if len(os.Args) <= 1 {
		help()
		return
	}

	//check for JaiKube cluster operations
	if os.Args[1] == "create" {
		createNodes()
	} else if os.Args[1] == "stop" {
		stopNodes()
	} else if os.Args[1] == "delete" {
		deleteNodes()
	} else if os.Args[1] == "reboot" {
		rebootNodes()
	} else if os.Args[1] == "list" {
		listNodes()
	} else {
		fmt.Println("Unknown Paramater:", os.Args[1])
		help()
	}
	



	//fmt.Println(os.Args[1])



	// cpPtr := flag.String("cp", "", "Control Plane/Etcd Nodes")

	// flag.Parse()

	// fmt.Println("cp:", *cpPtr)


}
