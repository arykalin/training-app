package main

import (
	"fmt"
	"github.com/arykalin/training-app/pkg"
	"io/ioutil"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func main() {

	var err error
	//Create kuberentes client to get node list
	//cl, err := pkg.Kubeclient()
	//panic
	if err != nil {
		panic(err.Error())
	}

	//Get node list from kubernetes cluster
	//nodeList, err := cl.CoreV1().Nodes().List(metav1.ListOptions{})

	//panic
	if err != nil {
		panic(err.Error())
	}

	//nodename := nodeList.Items[0].Name
	nodename := "minikube"
	data, err := ioutil.ReadFile("Makefile")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//log.Println("Contents of file:", string(data))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is Go HTTP server running on node %s.\n Data files is:\n %s", nodename, string(data))
	})
	if err := http.ListenAndServe(pkg.ListenPort, nil); err != nil {
		panic(err)
	}
}
