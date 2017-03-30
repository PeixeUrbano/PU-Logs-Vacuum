DIR=./dist
if [ -d "$DIR" ]; then
    printf '%s\n' "Removing dir ($DIR)"
    rm -rf "$DIR"
fi

# Android
echo 'Generating Android builds'
GOOS=android GOARCH=arm go build -o dist/logsVacuum-android.386 main.go

# Linux
echo 'Generating Linux builds'
GOOS=linux GOARCH=386 go build -o dist/logsVacuum-linux.386 main.go
GOOS=linux GOARCH=amd64 go build -o dist/logsVacuum-linux.amd64 main.go
GOOS=linux GOARCH=arm go build -o dist/logsVacuum-linux.arm main.go

# MacOs
echo 'Generating MacOs builds'
GOOS=darwin GOARCH=386 go build -o dist/logsVacuum-macos.386 main.go
GOOS=darwin GOARCH=amd64 go build -o dist/logsVacuum-macos.amd64 main.go

# Windows
echo 'Generating Windows builds'
GOOS=windows GOARCH=386 go build -o dist/logsVacuum-win.386.exe main.go
GOOS=windows GOARCH=amd64 go build -o dist/logsVacuum-win.amd64.exe main.go