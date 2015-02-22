# goconfig-cli

[![baby-gopher](https://raw.github.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Command line tool reading ini file based on _goconfig_ https://github.com/Unknwon/goconfig

By default, it will preserve original end line characters of the ini file ("\n" or "\r\n" or "\r")

* Can be cross-compiled in this repository or just installed with standard go get
* see api-doc here : https://gowalker.org/github.com/StudioEtrange/goconfig-cli

# Install

	go get github.com/StudioEtrange/goconfig-cli

NOTE : the dependencies are not locked down, as if you built from github-source (see below)

# Usage

On Unix:


	$GOPATH/bin/goconfig-cli <ini-file> getkey <section> <key>
	$GOPATH/bin/goconfig-cli <ini-file> setkey [--pretty] <section> <key> <value>
	$GOPATH/bin/goconfig-cli <ini-file> delkey [--pretty] <section> <key> 


On Windows:

	%GOPATH%\bin\goconfig-cli.exe <ini-file> getkey <section> <key>
	%GOPATH%\bin\goconfig-cli.exe <ini-file> setkey [--pretty] <section> <key> <value>
	%GOPATH%\bin\goconfig-cli.exe <ini-file> delkey [--pretty] <section> <key>

--pretty option : consider ini file having " = " as key/value separator instead of "="


# Build from github-source


## On Unix

	git clone https://github.com/StudioEtrange/goconfig-cli
	cd goconfig-cli

	./tool.sh build

goconfig-cli binary built for your current-platform is in release folder

	./tool.sh clean

remove any release and temporary workspace

	./tool.sh prepare-cross
	./tool.sh build-cross

cross-compiled version of goconfig-cli are in release folder

NOTE : _godep_ is used to stick version of dependencies

NOTE : _gonative_ and _gox_ are used for cross-compile
* see : https://github.com/tools/godep
* see : https://github.com/mitchellh/gox
* see : https://github.com/inconshreveable/gonative


## On Windows

_soon_