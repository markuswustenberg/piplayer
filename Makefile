.PHONY: build
build: build-css
	GOOS=linux GOARCH=arm go build -o piplayer ./cmd/piplayer

.PHONY: build-css
build-css: tailwindcss
	./tailwindcss -i tailwind.css -o public/styles/app.css --minify

.PHONY: lint
lint:
	golangci-lint run

.PHONY: start
start:
	go run ./cmd/piplayer

.PHONY: start-vlc
start-vlc:
	cd ~/Music/iTunes/iTunes\ Media/Music && /Applications/VLC.app/Contents/MacOS/VLC -I http --http-password hest123

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 tailwindcss

.PHONY: watch-css
watch-css: tailwindcss
	./tailwindcss -i tailwind.css -o public/styles/app.css --watch
