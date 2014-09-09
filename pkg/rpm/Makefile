build:
	cd ../.. && make
	tar -cf docker-dnsmasq-updater-0.1.tar docker-dnsmasq-updater.spec
	tar \
		--append \
		--directory ../.. \
		--transform "s#^#docker-dnsmasq-updater-0.1/#" \
		--file docker-dnsmasq-updater-0.1.tar \
		README.md
	tar \
		--append \
		--directory ../.. \
		--transform "s#^#docker-dnsmasq-updater-0.1/#" \
		--file docker-dnsmasq-updater-0.1.tar \
		docker-dnsmasq-updater
	tar \
		--append \
		--transform "s#^#docker-dnsmasq-updater-0.1/#" \
		--file docker-dnsmasq-updater-0.1.tar \
		docker-dnsmasq-updater.service
	gzip -c docker-dnsmasq-updater-0.1.tar  > docker-dnsmasq-updater-0.1.tar.gz
	rpmbuild -ta docker-dnsmasq-updater-0.1.tar.gz



