module github.com/hahaha3w/3w3_Ai_Server/memoir

go 1.23.6

replace github.com/hahaha3w/3w3_Ai_Server/rpc-gen => ../../rpc-gen

require (
	github.com/hahaha3w/3w3_Ai_Server/rpc-gen v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/cloudwego/fastpb v0.0.5 // indirect
	golang.org/x/sys v0.19.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
