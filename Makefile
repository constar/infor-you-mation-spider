all:
	go build -v
run: 
	./infor-you-mation-spider -expire=72 -alsologtostderr=true
debug:
	./infor-you-mation-spider  -alsologtostderr=true -v=5
test:
	go test ./...
