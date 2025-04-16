package main

import (
	"fmt"
	//"os"
	"os/exec"
	"strings"
	"flag"
)

var colorOrange string = "\033[38;5;214m"
var colorReset string = "\033[0m"

func createNodes(serverNodeList string, agentNodeList string, configFile string) {

	serverNodeArray := strings.Split(serverNodeList, ",")
	agentNodeArray := strings.Split(agentNodeList, ",")

	fmt.Println(colorOrange + "JaiKube: Starting server node configuration" + colorReset)


	//create server nodes
	for _, node := range serverNodeArray {
		//check if server nodes is empty
		if node == "" {
			continue
		}
		fmt.Println(colorOrange + "JaiKube: Creating server node: " + node + colorReset)
		out, err := exec.Command("limactl", "create", configFile, "--name="+node).CombinedOutput()
		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
	}


	//create agent nodes
	for _, node := range agentNodeArray {
		//check if agent nodes is empty
		if node == "" {
			continue
		}
		fmt.Println(colorOrange + "JaiKube: Creating agent node: " + node + colorReset)
		out, err := exec.Command("limactl", "create", configFile, "--name="+node).CombinedOutput()
		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
	}


	//start server nodes
	for _, node := range serverNodeArray {
		//check if server nodes is empty
		if node == "" {
			continue
		}
		fmt.Println(colorOrange + "JaiKube: Starting server node: " + node + colorReset)
		out, err := exec.Command("limactl", "start", node).CombinedOutput()
		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
	}
	//start agent nodes
	for _, node := range agentNodeArray {
		//check if server agent is empty
		if node == "" {
			continue
		}
		fmt.Println(colorOrange + "JaiKube: Starting agent node: " + node + colorReset)
		out, err := exec.Command("limactl", "start", node).CombinedOutput()
		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
	}

	//master node setup

	//initial server node configuration
	fmt.Println(colorOrange + "JaiKube: Configuring server node: " + serverNodeArray[0] + colorReset)

	//capture initial server node IP
	command := fmt.Sprintf("limactl shell %s ip addr show lima0 | grep 'inet ' | awk '{print $2}' | cut -d/ -f1", serverNodeArray[0])
	masterIP, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		fmt.Printf("%s", string(masterIP))
	}

	//delete eth0 route
	command = fmt.Sprintf("limactl shell %s sudo ip route del default dev eth0", serverNodeArray[0])
	exec.Command("bash", "-c", command).CombinedOutput()

	//install k3s server and set permissions for kubectl
	command = fmt.Sprintf("limactl shell %s \"curl https://get.k3s.io | sh -s - server --write-kubeconfig-mode=644 --node-ip=%s --flannel-iface=lima0 --cluster-init\"", serverNodeArray[0], masterIP)
	exec.Command("bash", "-c", command).CombinedOutput()
	command = fmt.Sprintf("limactl shell %s sudo chmod 644 /etc/rancher/k3s/k3s.yaml", serverNodeArray[0])
	exec.Command("bash", "-c", command).CombinedOutput()

	//install helm and set up kubeconfig location for helm to use


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
	//stop nodes
	for _, node := range nodeNamesArray {
		out, err := exec.Command("limactl", "stop", node).CombinedOutput()
		if err != nil {
			fmt.Printf("%s", string(out))
		} else {
			fmt.Printf("%s", string(out))
		}
    }
	//start nodes
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
	fmt.Printf("	./jaikube create server,nodes agent,nodes configFile\n\n")
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
    
	serverPtr := flag.String("server", "", "comma seperated string of server node names")
	agentPtr := flag.String("agent", "", "comma seperated string of agent node names")
	configPtr := flag.String("config", "", "lima configuration file")
	nodePtr := flag.String("nodes", "" ,"comma seperated string of node names")
	jobPrt := flag.String("job", "" ,"what function to run")

	flag.Parse()

	//check for JaiKube cluster operations
	if *jobPrt == "create" && (*serverPtr != "") && (*configPtr != "") {
		createNodes(*serverPtr, *agentPtr, *configPtr)
	} else if *jobPrt == "stop" && *nodePtr != "" {
		stopNodes(*nodePtr)
	} else if *jobPrt == "delete" && *nodePtr != "" {
		deleteNodes(*nodePtr)
	} else if *jobPrt == "reboot" && *nodePtr != "" {
		rebootNodes(*nodePtr)
	} else if *jobPrt == "list" {
		listNodes()
	} else {
		fmt.Println("Incorrect Arguments")
		help()
	}
}
