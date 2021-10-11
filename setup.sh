runenvy() {
}

rungateway() {

}

runapp() {
  pushd app && go build . -o && popd
  app/main
}

