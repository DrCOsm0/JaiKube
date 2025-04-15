package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var colorOrange string = "\033[38;5;214m"
var colorReset string = "\033[0m"

func createNodes(MasterNodeList string, WorkerNodeList string) {
	fmt.Println("create")
}

func stopNodes(nodeList string) {
	nodeNamesArray := strings.Split(nodeList, ",")

	for _, node := range nodeNamesArray {

		out, err := exec.Command("limactl", "stop", node).CombinedOutput()

		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
    }
}

func deleteNodes(nodeList string) {
	nodeNamesArray := strings.Split(nodeList, ",")

	//stop nodes first
	stopNodes(nodeList)

	for _, node := range nodeNamesArray {

		out, err := exec.Command("limactl", "delete", node).CombinedOutput()

		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
    }
}

func rebootNodes(nodeList string) {
	nodeNamesArray := strings.Split(nodeList, ",")

	for _, node := range nodeNamesArray {

		out, err := exec.Command("limactl", "stop", node).CombinedOutput()

		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
    }

	for _, node := range nodeNamesArray {

		out, err := exec.Command("limactl", "start", node).CombinedOutput()

		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
    }
}

func listNodes() {

	out, err := exec.Command("limactl", "list").Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}

func help() {
	fmt.Printf("Commands:\n\n")
	fmt.Printf("Create a cluster:\n")
	fmt.Printf("	./jaikube create server,nodes agent,nodes\n\n")
	fmt.Printf("Start node(s):\n")
	fmt.Printf("	./jaikube start nodes,list\n\n")
	fmt.Printf("Stop node(s):\n")
	fmt.Printf("	./jaikube stop nodes,list\n\n")
	fmt.Printf("Delete node(s):\n")
	fmt.Printf("	./jaikube delete nodes,list\n\n")
	fmt.Printf("Reboot node(s):\n")
	fmt.Printf("	./jaikube reboot nodes,list\n\n")
	fmt.Printf("List node(s):\n")
	fmt.Printf("	./jaikube list nodes,list\n\n")
}

func printLogo() {
	fmt.Println(colorOrange + `
             ____.      .__ ____  __.    ___.              
            |    |____  |__|    |/ _|__ _\_ |__   ____     
            |    \__  \ |  |      < |  |  \ __ \_/ __ \    
        /\__|    |/ __ \|  |    |  \|  |  / \_\ \  ___/    
        \________(____  /__|____|__ \____/|___  /\___  >   
                      \/           \/         \/     \/    
                         --Cluster Management, Made Easy

    Well Howdy, Welcome to JaiKube!
    ` + colorReset)
}


func main() {
    
	
	//argsWithProg := os.Args
    //argsWithoutProg := os.Args[1:]

	//key := os.Args[]

	//check for null input
	if len(os.Args) <= 1 {
		printLogo()
		help()
		return
	}

	//check for JaiKube cluster operations
	if os.Args[1] == "create" && len(os.Args) == 4 {
		MasterNodeList := os.Args[2]
		WorkerNodeList := os.Args[3]
		createNodes(MasterNodeList, WorkerNodeList)
	} else if os.Args[1] == "stop" && len(os.Args) == 3 {
		nodeList := os.Args[2]
		stopNodes(nodeList)
	} else if os.Args[1] == "delete" && len(os.Args) == 3 {
		nodeList := os.Args[2]
		deleteNodes(nodeList)
	} else if os.Args[1] == "reboot" && len(os.Args) == 3 {
		nodeList := os.Args[2]
		rebootNodes(nodeList)
	} else if os.Args[1] == "list" {
		listNodes()
	} else {
		fmt.Println("Incorrect Arguments")
		help()
	}
}
