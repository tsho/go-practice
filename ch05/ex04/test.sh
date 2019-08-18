go build main.go
if [ $? != 0 ]; then
    echo 'build error'
    exit 1
fi
./fetch https://golang.org | ./main