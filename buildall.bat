@echo off 
goto comment
    Build the command lines and tests in Windows.
    Must install gcc tool before building.
:comment

echo on

go build -ldflags "-s -w" -o ./build/dashboard-api.exe ./cmd/api
@echo "Done dashboard-api building release"

pause
