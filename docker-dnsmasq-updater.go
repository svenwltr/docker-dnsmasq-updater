package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	containers, err := client.ListContainers(opts)

	if err != nil {
		log.Fatal("Unable to list containers: ", err)
	}

	var lines []string

	for _, container := range containers {
		inspect, _ := client.InspectContainer(container.ID)
		lines = append(lines, fmt.Sprintf("address=%s.test/%s", inspect.Name,
			inspect.NetworkSettings.IPAddress))

	}

	return strings.Join(lines, "\n")

}

func updateConfig(hosts string) {
	log.Print("Updating dnsmasq config")
	var err error

	err = ioutil.WriteFile(DNSMASQ_CONF, []byte(hosts), 0644)
	if err != nil {
		log.Print("ERROR: unable to write dnsmasq config: ", err)
	}

	err = exec.Command("systemctl", "restart", "dnsmasq").Run()
	if err != nil {
		log.Print("ERROR: unable to restart dnsmasq: ", err)
	}

}

func main() {
	var err error
	client, err := docker.NewClient(ENDPOINT)

	if err != nil {
		log.Fatal("Unable to initialize Docker client. ", err)
	}

	listener := make(chan *docker.APIEvents, 10)
	err = client.AddEventListener(listener)
	if err != nil {
		log.Fatal("Unable to initialize Docker event listener: ", err)
	}

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
