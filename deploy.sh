cd $GOPATH/src/github.com/TIBCOSoftware
git clone https://github.com/TIBCOSoftware/mashling-cicd.git mashling-cicd
cp -r $GOPATH/src/github.com/TIBCOSoftware/mashling-cicd/dependency-config/* $GOPATH/src/github.com/TIBCOSoftware/mashling
cd mashling
go install ./...
mashling version
cd $GOPATH/src/github.com/TIBCOSoftware
pushd mashling-cicd/sample-recipes
rm -rf builds
mkdir -p builds/latest
chmod ugo+x ./build-recipes.sh ./sanity.sh
./build-recipes.sh
./sanity.sh
popd
if [ "$TRAVIS_OS_NAME" == "linux" ]; then
  pushd mashling-cicd/docker/mashling
  chmod ugo+x ./build-mashling.sh
  ./build-mashling.sh
  popd ;
fi
