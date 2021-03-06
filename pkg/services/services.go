package kubebot

import (
    "fmt"
    "log"
  
    "github.com/go-chat-bot/bot"

    "k8s.io/kubernetes/pkg/api"
    "k8s.io/kubernetes/pkg/client/restclient"
    client "k8s.io/kubernetes/pkg/client/unversioned"
)

func svc(command *bot.Cmd) (msg string, err error) {

    config := &restclient.Config{
        Host: "http://localhost:8080",
    }
    c, err := client.New(config)
    if err != nil {
        log.Fatalln("Can't connect to Kubernetes API:", err)
    }

    s, err := c.Services(api.NamespaceDefault).Get("kubernetes")
    if err != nil {
        log.Fatalln("Can't get service:", err)
    }
    fmt.Println("Name:", s.Name)
    for p, _ := range s.Spec.Ports {
        fmt.Println("Port:", s.Spec.Ports[p].Port)
        fmt.Println("NodePort:", s.Spec.Ports[p].NodePort)
    }
    fmt.Println("ClusterIP:", s.Spec.ClusterIP)
    fmt.Println("type:", s.Spec.Type)
    msg = fmt.Sprintln(s.Name)
    return msg, nil
}

func init() {
	bot.RegisterCommand(
		"svc",
		"Get k8s Services",
		"",
		svc)
}
