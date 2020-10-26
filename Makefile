protoc-generate:
	mkdir ${proto}
	protoc --go_out=plugins=grpc:${proto} ${proto}.proto

create-proto:
	./mscript_create.sh 