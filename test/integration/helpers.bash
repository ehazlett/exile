#!/bin/bash

# Root directory of the repository.
ROOT=${BATS_TEST_DIRNAME}/../..
PROJECT=exile

build() {
    pushd $ROOT >/dev/null
    godep go build
    popd >/dev/null
}

# build machine binary if needed
if [ ! -e $MACHINE_ROOT/$PROJECT ]; then
    build
fi

exile() {
    ${ROOT}/exile "$@"
}
