package main

import (
	"fmt"
	"github.com/arykalin/training-app/pkg"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func main() {

	cl := pkg.Kubeclient()
	nodeList, err := cl.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	nodename := nodeList.Items[0].Name

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is Go HTTP server running on node %s.", nodename)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
