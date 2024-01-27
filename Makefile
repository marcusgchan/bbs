run:
	templ generate
	go run ./cmd/bbs/main.go

dev:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run ./cmd/bbs/main.go
