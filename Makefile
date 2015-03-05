all:
	go build
run: all
	./infor-you-mation-spider  -alsologtostderr=true
