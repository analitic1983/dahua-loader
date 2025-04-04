#!/bin/bash
go build  -o ./dahua-loader-console koshmin/dahua-loader/console
go build  -o ./dahua-loader-server koshmin/dahua-loader/server

chmod 0777 ./dahua-loader-console
chmod 0777 ./dahua-loader-server
