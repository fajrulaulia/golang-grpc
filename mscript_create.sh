#!/bin/bash 
cat <<EOF >${proto}.proto
syntax= "proto3";

message ${proto^} {
    .....
}

service ${proto^}Service {
    rpc SayHello(${proto^}) returns (${proto^}) {}
}
EOF