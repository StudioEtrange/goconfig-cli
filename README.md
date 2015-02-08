# goconfig-cli

[![baby-gopher](https://raw.github.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

command line tool reading ini file

based on goconfig https://github.com/Unknwon/goconfig

By default, it will preserve original end line characters of the ini file ("\n" or "\r\n" or "\r")

see api-doc here : https://gowalker.org/github.com/StudioEtrange/goconfig-cli

# Install

	go get github.com/StudioEtrange/goconfig-cli

# Usage

On Unix:


	$GOPATH/bin/goconfig-cli <ini-file> getkey <section> <key>
	$GOPATH/bin/goconfig-cli <ini-file> setkey [--pretty] <section> <key> <value>
	$GOPATH/bin/goconfig-cli <ini-file> delkey [--pretty] <section> <key> 


On Windows:

	%GOPATH%\bin\goconfig-cli.exe <ini-file> getkey <section> <key>
	%GOPATH%\bin\goconfig-cli.exe <ini-file> [--pretty] setkey <section> <key> <value>
	%GOPATH%\bin\goconfig-cli.exe <ini-file> [--pretty] delkey <section> <key>



--pretty option : consider ini file having " = " as key/value separator instead of "="
