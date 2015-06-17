#!/bin/bash
_CURRENT_FILE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"


GOVER=$(go version | cut -d" " -f3 | cut -d"o" -f2)

WORKSPACE=$_CURRENT_FILE_DIR/workspace
GODEP=$WORKSPACE/bin/godep
GOXC=$WORKSPACE/bin/goxc
GOX=$WORKSPACE/bin/gox
GONATIVE=$WORKSPACE/bin/gonative

GOTOOLCHAIN=$WORKSPACE/gonative-toolchain
RELEASE_DIR=$_CURRENT_FILE_DIR/release

export GOPATH=$WORKSPACE


ACTION=$1

case $ACTION in

	build)
		echo "** install godep and restore goconfig-cli dependencies"
		go get github.com/tools/godep
		$GODEP restore

		# copy source to workspace
		cp -r -f src/ $WORKSPACE/src/

		echo "** build goconfig-cli"
		go install goconfig-cli

		mkdir -p $RELEASE_DIR
		cp -f $WORKSPACE/bin/goconfig-cli $RELEASE_DIR
	;;

	clean)
		rm -Rf $WORKSPACE
		rm -Rf $RELEASE_DIR
		rm -Rf $GOTOOLCHAIN
	;;


	prepare-cross)
		echo "** install gonative"
		go get github.com/inconshreveable/gonative

		echo "** build toolchain"
		mkdir -p $GOTOOLCHAIN
		cd $GOTOOLCHAIN
		$GONATIVE build --version="$GOVER" --platforms="windows_386 windows_amd64 linux_386 linux_amd64 darwin_386 darwin_amd64"
	;;

	build-cross)
		
		echo "** install gox"
  		go get github.com/mitchellh/gox
  		#$GOX -build-toolchain

        echo "** install godep and restore goconfig-cli dependencies"
		go get github.com/tools/godep
		$GODEP restore

		# copy source to workspace
		cp -r -f src/ $WORKSPACE/src/


		echo "** cross-compile goconfig-cli"

		mkdir -p $RELEASE_DIR
		cd $RELEASE_DIR
		PATH=$GOTOOLCHAIN/go/bin/:$PATH $GOX -verbose -osarch="windows/386 windows/amd64 linux/386 linux/amd64 darwin/386 darwin/amd64" goconfig-cli


	;;

	*)
		echo "Usage :"
		echo "$0 build : build goconfig-cli for current platform"
		echo "$0 prepare-cross : prepare a cross-platform chain -- use once only berfore build-cross"
		echo "$0 build-cross : build cross-platform goconfig-cli for all plaform"
		echo "$0 clean : clean everything"
	;;
esac
