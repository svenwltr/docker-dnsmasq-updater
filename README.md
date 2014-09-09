docker-dnsmasq-updater
======================

Poor mans service discovery for docker.

Requirements
------------

 * dnsmasq
 * systemd (to restart dnsmasq)
 * root access (to restart dnsmasq)
 * dnsmasq must read config file `/etc/dnsmasq.d/docker.conf`
 * Docker socket on `/var/run/docker.sock`

Usage
-----

    sudo docker-dnsmasq-updater

Programm will listen to docker events and updates dnsmasq config on host change.

Example
-------

    # docker ps
    CONTAINER ID        IMAGE                    COMMAND                CREATED             STATUS              PORTS                          NAMES
	627b0e442ba5        tutum/glassfish:latest   /run.sh                21 hours ago        Up 8 hours          4848/tcp, 8080/tcp, 8181/tcp   glassfish_as_1
	4d2ecef55680        mongo:latest             /usr/src/mongo/docke   4 days ago          Up 14 minutes       27017/tcp                      foo_db_1

This produces following dnsmasq config:

    address=/glassfish_as_1.test/172.17.0.5
    address=/foo_db_1.test/172.17.0.21


Libaries
--------

 * https://github.com/fsouza/go-dockerclient


Author
------

Sven Walter <sven@wltr.eu>
