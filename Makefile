proto-build:
	cd ../../../../../ && protoc -I ./ --go_out=plugins=grpc:./build/go \
    		./svr/cal/v1/*.proto \
    		./svr/ctx/v1/*.proto