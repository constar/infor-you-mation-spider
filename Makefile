all:
	go build
run: 
	./infor-you-mation-spider  -alsologtostderr=true
debug:
	./infor-you-mation-spider  -alsologtostderr=true -v=5
