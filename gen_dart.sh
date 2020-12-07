export PATH="`pwd`/script/:$PATH"
rm -rf ./dart/*

ls ./pb/google/protobuf/*.proto | xargs protoc -I=./pb/ --dart_out=./dart
ls ./pb/message/*.proto | xargs protoc -I=./pb/ --dart_out=./dart