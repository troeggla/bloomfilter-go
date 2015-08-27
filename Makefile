build/bloomfilter: bloomfilter.go main.go
	go build -o $@ $^

clean:
	rm build/bloomfilter
	rmdir build
