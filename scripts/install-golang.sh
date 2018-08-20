#!/usr/bin/env bash

set -ex


echo '##########################################################################'
echo '################# About to run install-golang.sh script ##################'
echo '##########################################################################'

golang_archive_filename=go1.10.3.linux-amd64.tar.gz
curl -o ${golang_archive_filename} https://dl.google.com/go/${golang_archive_filename}
tar -C /usr/local -xzf ${golang_archive_filename}
cd /usr/local/go/bin 
cp -rfp /usr/local/go/bin/* /usr/bin/


cp -f /vagrant/files/.go_config /root/

echo "source /root/.go_config" >> /root/.bash_profile


mkdir /root/go_workspaces
cd /root/go_workspaces

git clone https://github.com/mmcgrana/gobyexample.git

cd /root/go_workspaces/gobyexample
mv /root/go_workspaces/gobyexample/examples /root/go_workspaces/
rm -rf /root/go_workspaces/gobyexample/*
mv /root/go_workspaces/examples /root/go_workspaces/gobyexample/src
mkdir /root/go_workspaces/gobyexample/{bin,pkg}
cd /root/go_workspaces/gobyexample/src
rm -rf */*.hash */*.sh




exit 