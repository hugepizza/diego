#!/bin/bash

set -ex;

if [ "${TRAVIS_PULL_REQUEST}" != "false" ]; then
  echo "this is a pull request, no need to build."
  exit 0;
fi

export CGO_ENABLED=0;

APP="diego"
PWD=`pwd`
PKG="github.com/ckeyer/$APP"
OS=`go env GOOS`
ARCH=`go env GOARCH`
VERSION=`cat VERSION.txt`
GIT_COMMIT=`git rev-parse --short HEAD`
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
BUILD_AT=`date "+%Y-%m-%dT%H:%M:%SZ%z"`
IMAGE="ckeyer/$APP"
PACKAGE_NAME="$APP.$VERSION.$OS-$ARCH"

LD_FLAGS="-X github.com/ckeyer/commons/version.version=$VERSION \
 -X github.com/ckeyer/commons/version.gitCommit=$GIT_COMMIT \
 -X github.com/ckeyer/commons/version.buildAt=$BUILD_AT -w"

echo "===== start diego building."

echo "=== start building ui."
echo "=== building ui successful."

echo "=== start building binary."
go build -a -ldflags="$LD_FLAGS"  -o bundles/$APP .
echo "=== building binary successful."

echo "=== start building dockerimage."
docker build -t $IMAGE:$GIT_COMMIT .
docker tag $IMAGE:$GIT_COMMIT $IMAGE:$GIT_BRANCH
docker tag $IMAGE:$GIT_COMMIT $IMAGE:$VERSION
docker push $IMAGE:$GIT_COMMIT
docker push $IMAGE:$GIT_BRANCH
docker push $IMAGE:$VERSION
echo "=== building dockerimage successful."

echo "===== diego built successful."

echo ${PACKAGE_NAME}
echo ${PKG}
