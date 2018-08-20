#!/usr/bin/env bash

set -ex


echo '##########################################################################'
echo '################# About to run install-golang.sh script ##################'
echo '##########################################################################'

golang_archive_filename=go1.10.3.linux-amd64.tar.gz
curl -o ${golang_archive_filename} https://dl.google.com/go/${golang_archive_filename}
tar -C /usr/local -xzf ${golang_archive_filename}
cd /usr/local/bin/
cp -rfp /usr/local/bin/* /usr/bin/