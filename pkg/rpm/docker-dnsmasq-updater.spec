%global debug_package %{nil}

Name:		docker-dnsmasq-updater
Version:	0.1
Release:	1%{?dist}
Summary:	Poor mans service discovery for docker.

Group:		Development/Tools
License:	MIT
URL:		https://github.com/svenwltr/docker-dnsmasq-updater
Source0:	https://github.com/svenwltr/docker-dnsmasq-updater/archive/%{name}-%{version}.tar.gz

BuildRequires:	golang

%description


%prep
%setup


%build


%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}/usr/bin
install -pm 755 docker-dnsmasq-updater %{buildroot}/usr/bin/%{name}

install -dm 755 %{buildroot}/usr/lib/systemd/system
install -pm 644 docker-dnsmasq-updater.service %{buildroot}/usr/lib/systemd/system/%{name}.service


%files
#%defattr(-,root,root,-)
#%{_datadir}/%{name}
#%doc README.md
/usr/bin/%{name}
/usr/lib/systemd/system/%{name}.service



%changelog

