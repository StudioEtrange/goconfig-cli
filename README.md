# goconfig-cli

command line tool reading ini file

based on goconfig https://github.com/Unknwon/goconfig

# Install

	go get github.com/StudioEtrange/goconfig-cli

# Usage

On Unix:


	$GOPATH/bin/goconfig-cli <ini-file> getkey <section> <key>
	$GOPATH/bin/goconfig-cli <ini-file> setkey <section> <key> <value>
	$GOPATH/bin/goconfig-cli <ini-file> delkey <section> <key>


On Windows:

	%GOPATH%\bin\goconfig-cli.exe <ini-file> getkey <section> <key>
	%GOPATH%\bin\goconfig-cli.exe <ini-file> setkey <section> <key> <value>
	%GOPATH%\bin\goconfig-cli.exe <ini-file> delkey <section> <key>