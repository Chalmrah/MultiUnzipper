
$MODULE_PATH="main.go"

$GOOS="linux"
$GOARCH="amd64"
echo "Building for $GOOS/$GOARCH..."
go build -o D:/Source/publish/unzipper/multiunzipper $MODULE_PATH

$GOOS="windows"
$GOARCH="amd64"
echo "Building for $GOOS/$GOARCH..."
go build -o D:/Source/publish/unzipper/multiunzipper.exe $MODULE_PATH

echo "Done"