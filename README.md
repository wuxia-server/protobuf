# protobuf

游戏前后端交互协议

参考文档

https://developers.google.com/protocol-buffers/docs/gotutorial

https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

## 1. 环境配置

go语言中使用protobuf，依赖两个文件`protoc`和`protoc-gen-go`，当前文件已经下载到 script 目录

1. protoc

* 版本信息：protoc-3.13.0-osx-x86_64
* 下载地址：https://github.com/protocolbuffers/protobuf/releases/tag/v3.13.0

2. protoc-gen-go

* 版本信息：protoc-gen-go.v1.25.0.darwin.amd64.tar.gz
* 下载地址：https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.25.0
* 命令下载：go install google.golang.org/protobuf/cmd/protoc-gen-go

其中 protoc-gen-go 还有一个 386 版本的，应该是用不到了，amd64的应该就可以了。proto-gen-go使用命令下载时，会下载到$GOBIN目录，我在readme环境变量中添加了该配置，若重新配置服务器，不影响使用

手动下载程序需要手动配置环境依赖，或拷贝到$GOBIN 目录

## 2. 前端protobuf编译方法

- 目前使用的是导入静态文件的方式
- 安装pbjs：`npm i -g protobufjs`
- 进入pb目录：`cd pb`
- 执行`pbjs -t static-module -w commonjs -o noname.js */*.proto` 生成js文件
- 再执行`pbts -o noname.d.ts noname.js`生成ts文件
- 将生成的js文件中第一句`var $protobuf = require("protobufjs/minimal");`修改为`var $protobuf = protobuf;`
- 将生成的ts文件中第一句`import * as $protobuf from "protobufjs";`修改为`import * as $protobuf from "./protobuf";`
- 替换工程中对应的js文件和ts文件即可。




## 3. 后端使用

参照脚本 [gen_proto.sh](../update_pb.sh)使用，在protoc命令中不支持使用通配符，要导出的proto文件需要手动指定

脚本说明
```shell
# 将项目 script 目录添加到环境变量
export PATH="`pwd`/script/:$PATH"

# protoc 生成命令
protoc -I=./pb/ --go_out=./golang/ match/match.proto
protoc -I=./pb/ --go_out=./golang/ player/player.proto
protoc -I=./pb/ --go_out=./golang/ player/data_server.proto
```

修改好gen_proto.sh后，在项目目录中可以直接执行文件
```sh
sh gen_proto.sh
```

### 3. 协议使用约定

1. 固定协议头

```protobuf
syntax = "proto3"; // 使用proto3，与proto2的语法上的最大不同点在于，不用写optional和required了
package pb; // 包名，不要使用proto，可能会导致一些重名需要在import时起另外一个名字，这里建议统一使用pb


option go_package = "match/pb"; // 生成到目录的路径，这里配置是可选的，但是删掉会有警告，建议添加上

// 这里可以导入其他依赖proto，对于google的一些公共依赖，目录待定
//import "google/protobuf/timestamp.proto";

```

2. 变量语法

这里要注意，protobuf在生成代码时，对于以小写字母开头、下划线开头、下划线连接方式命名的变量进行了转换，示例如下

```protobuf
message C2S_PlayerJoin{
  int32 configId = 1;
  int32 nick_name = 2;
  int32 _money_free = 3;
}
```

生成后的结构体如下

```go
type C2S_PlayerJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId   int32 `protobuf:"varint,1,opt,name=configId,proto3" json:"configId,omitempty"`
	NickName   int32 `protobuf:"varint,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	XMoneyFree int32 `protobuf:"varint,3,opt,name=_money_free,json=MoneyFree,proto3" json:"_money_free,omitempty"`
}
```

这样在命名中出现`configId`和`config_id`时GoLand可能会有错误提示，导出时也会失败，因为导出会对应的变量都是ConfigId。**建议不要以下划线开头命名变量**

3. Golang 代码中使用protobuf

参照[protocol-buffers教程](https://developers.google.com/protocol-buffers/docs/gotutorial) 中的 [写消息](https://developers.google.com/protocol-buffers/docs/gotutorial#writing_a_message)和 [读消息](https://developers.google.com/protocol-buffers/docs/gotutorial#reading_a_message)

## 4. Go编码中的使用

对于protobuf，有多个库提供支持，在GoLand中给出的代码提示有一下几种及更多

```
github.com/gogo/protobuf
github.com/golang/protobuf
google.golang.org/golang/protobuf
...
```

1. gogo/protobuf ，该库是这样说的

```
gogoprotobuf is a fork of golang/protobuf with extra code generation features.
```

2. golang/protobuf，该库是这样说的

```
It has been superseded by the google.golang.org/protobuf module, which contains an updated and simplified API, support for protobuf reflection, and many other improvements. We recommend that new code use the google.golang.org/protobuf module.
```

gogo的新增特性并不是很了解，golang又推荐使用google.golang.org，在代码中使用[google.golang.org/golang/protobuf](https://pkg.go.dev/mod/google.golang.org/protobuf) 的 v1.25.0 版本，当前最新。编码中要注意不要引入其他两个库

## 5. GoLand包依赖更新

使用go mod 包管理方式

pb依赖更新，使用git tag标识的方式进行，服务端自用的tag，目前不涉及到前端。tag版本号约定，v上线版本.迭代大版本.迭代小版本。

* 上线版本，在项目上线前使用0
* 迭代大版本，自定义吧
* 小版本，自定义吧

可使用下面命令打tag，打完tag需要主动推送到tygit，tag后面可以删除，不用担心写错，[参考链接](https://git-scm.com/book/zh/v2/Git-%E5%9F%BA%E7%A1%80-%E6%89%93%E6%A0%87%E7%AD%BE)

```
git tag -a v0.1.0 -m "demo"

git push origin v0.1.0
```

更新包是需要修改 go.mod 文件中的版本号，更新到指定的包

```
tygit.touch4.me/rich-joy/ludo-protobuf v0.1.0
```

然后使用tidy命令更新

```
go mod tidy
```

