package main

//go:generate rm -rf gen proto
//go:generate mkdir -p gen
//go:generate protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative -I. cgentron/api/annotations.proto
//go:generate protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative -I. cgentron/api/fields.proto
//go:generate protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative -I. cgentron/api/methods.proto
//go:generate protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative -I. cgentron/api/resolver.proto
//go:generate mv gen/cgentron/api proto
//go:generate rm -rf gen
