run:
	templ generate
	npx tailwindcss -o ./build/output.css 
	go run ./cmd/bbs/main.go

dev:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: npx tailwindcss -o ./web/static/output.css :: go run ./cmd/bbs/main.go
