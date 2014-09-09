package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

const (
	ENDPOINT     = "unix:///var/run/docker.sock"
	DNSMASQ_CONF = "/etc/dnsmasq.d/docker.conf"
)

func getHosts(client *docker.Client) string {
	opts := docker.ListContainersOptions{}
	containers, _ := client.ListContainers(opts)

	var lines []string

	for _, container := range containers {
		inspect, _ := client.InspectContainer(container.ID)
		lines = append(lines, fmt.Sprintf("address=%s.test/%s", inspect.Name, inspect.NetworkSettings.IPAddress))

	}

	return strings.Join(lines, "\n")

}

func updateConfig(hosts string) {
	fmt.Println("Updating dnsmasq config")
	_ = ioutil.WriteFile(DNSMASQ_CONF, []byte(hosts), 0644)
	exec.Command("systemctl", "restart", "dnsmasq").Run()

}

func main() {
	client, _ := docker.NewClient(ENDPOINT)

	listener := make(chan *docker.APIEvents, 10)
	_ = client.AddEventListener(listener)

	var newHosts string
	var oldHosts string

	oldHosts = getHosts(client)
	updateConfig(oldHosts)

	for {
		select {
		case msg := <-listener:
			fmt.Printf("%d: %s %s\n", msg.Time, msg.Status, msg.ID)
			newHosts = getHosts(client)

			if newHosts != oldHosts {
				updateConfig(newHosts)
			}

			oldHosts = newHosts
		}
	}

}
