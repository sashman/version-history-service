FROM centos:6
MAINTAINER Oleksandr Stasyk <oleksandr.stasyk@logicnow.com>

RUN	yum -y update && yum clean all \
	&& yum -y install http://www.percona.com/downloads/percona-release/redhat/0.1-3/percona-release-0.1-3.noarch.rpm \
	&& yum -y install Percona-Server-server-56.x86_64 \
	&& yum -y install Percona-Server-client-56.x86_64 \
	&& yum -y install Percona-Server-devel-56.x86_64 \
	&& yum -y install Percona-Server-56-debuginfo.x86_64 \
	&& yum -y install Percona-Server-shared-56.x86_64 \
	&& yum -y install Percona-Server-shared-compat.x86_64

RUN	yum -y update && yum clean all \
	&& yum -y install wget curl tar rpm bzip2 \
	&& rpm -ivh http://dl.iuscommunity.org/pub/ius/stable/CentOS/6/x86_64/epel-release-6-5.noarch.rpm \
	&& rpm -ivh http://dl.iuscommunity.org/pub/ius/stable/CentOS/6/x86_64/ius-release-1.0-11.ius.centos6.noarch.rpm \
	&& yum -y groupinstall 'Development Tools' \
	&& yum -y install bzip2-devel libcurl-devel t1lib-devel mcrypt libmcrypt libmcrypt-devel \
	&& yum -y install openssl openssl-devel \
	&& yum -y install libxml2 libxml2-devel libtool-ltdl-devel \
	&& yum -y install libjpeg-turbo-devel libpng-devel libXpm-devel freetype-devel t1lib-devel \
	&& yum -y install gmp-devel mcrypt libmcrypt libmcrypt-devel libtidy-devel tidy bison libtool-ltdl-devel \
	&& yum -y install autoconf213 \
	&& yum -y install unixODBC unixODBC-devel libsodium libsodium-devel \
	&& yum -y install sqlite-devel xz-libs 

RUN	yum -y install http://ftp.riken.jp/Linux/fedora/epel/6/i386/epel-release-6-8.noarch.rpm
RUN yum -y install golang
RUN yum -y install mongodb-org

RUN	mkdir /Application
COPY	Application /Application