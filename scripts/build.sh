#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/vars.sh

BUILDDIR="$2"

build_local_and_docker() {
   module="$1"
   folder="$2"
   title=$(printf "$module" | awk '{ print toupper($0) }')

   printf '%s' "Building $title Locally...  "
   cwd=$PWD
   cd $folder
   go build -mod=readonly -trimpath -o $BUILDDIR ./... 2>&1 | grep -v -E "deprecated|keychain";
   local_build_succeeded=$?
   cd $cwd
   echo "Done" 

   echo "Building $title Docker...  "
   docker build --tag stridezone:$module -f Dockerfile.$module . 
   docker_build_succeeded=$?

   return $docker_build_succeeded && $local_build_succeeded
}

ADMINS_FILE=${SCRIPT_DIR}/../utils/admins.go
ADMINS_FILE_BACKUP=${SCRIPT_DIR}/../utils/admins.go.main

replace_admin_address() {
   cp $ADMINS_FILE $ADMINS_FILE_BACKUP
   sed -i -E "s|stride1k8c2m5cn322akk5wy8lpt87dd2f4yh9azg7jlh|$STRIDE_ADMIN_ADDRESS|g" $ADMINS_FILE
}

revert_admin_address() {
   mv $ADMINS_FILE_BACKUP $ADMINS_FILE
   rm -f ${ADMINS_FILE}-E
}

# build docker images and local binaries
while getopts sgojthir flag; do
   case "${flag}" in
      # For stride, we need to update the admin address to one that we have the seed phrase for
      s) replace_admin_address
         if (build_local_and_docker stride .) ; then
            revert_admin_address
         else
            revert_admin_address
            exit 1
         fi
         ;;
      g) build_local_and_docker gaia deps/gaia ;;
      j) build_local_and_docker juno deps/juno ;;
      o) build_local_and_docker osmo deps/osmosis ;;
      t) build_local_and_docker stars deps/stargaze ;;
      i) build_local_and_docker icq deps/interchain-queries ;;
      r) build_local_and_docker relayer deps/relayer ;;  
      h) echo "Building Hermes Docker... ";
         docker build --tag stridezone:hermes -f Dockerfile.hermes . ;

         printf '%s' "Building Hermes Locally... ";
         cd deps/hermes; 
         cargo build --release --target-dir $BUILDDIR/hermes; 
         cd ../..
         echo "Done" ;;
   esac
done