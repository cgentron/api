package api

//go:generate mkdir -p proto
//go:generate protoc --go_out=plugins=grpc:. --go_opt=module=github.com/cgentron -I. cgentron/api/annotations.proto
//go:generate protoc --go_out=plugins=grpc:. --go_opt=module=github.com/cgentron  -I. cgentron/api/fields.proto
//go:generate protoc --go_out=plugins=grpc:. --go_opt=module=github.com/cgentron  -I. cgentron/api/methods.proto
//go:generate protoc --go_out=plugins=grpc:. --go_opt=module=github.com/cgentron  -I. cgentron/api/resolver.proto
