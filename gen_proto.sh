export PATH="`pwd`/script/:$PATH"

rm -rf ./golang/*
#protoc -I=./pb/ --go_out=./golang/ match/match.proto --go_opt=paths=source_relative

ls ./pb/message/*.proto | xargs protoc -I=./pb/ --go_out=./golang --go_opt=paths=source_relative


sh gen_uri.sh
