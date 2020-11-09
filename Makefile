proto-build:
	cd ../../../../../ && protoc -I ./ --go_out=plugins=grpc:./build/go \
    		./svr/cal/v1/*.proto \
    		./svr/chart/v1/*.proto \
    		./svr/cell/v1/*.proto \
    		./svr/opt/v1/*.proto \
    		./svr/pkg/v1/*.proto \
    		./svr/ctx/v1/*.proto