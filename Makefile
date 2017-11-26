run:
	go run cmd/ad.go &
	go run cmd/dsp.go &
	go run cmd/imp.go
	pkill -f "cmd/"

fmt:
	go fmt ./...
	
stop:
	pkill -f 'cmd/'