.PHONY: build build-css build-css-dev lint start start-vlc

build: build-css
	GOOS=linux GOARCH=arm go build -o piplayer cmd/piplayer/*.go

build-css:
	NODE_ENV=production ./node_modules/.bin/tailwindcss build views/app.css -o public/styles/app.css

build-css-dev:
	NODE_ENV=development ./node_modules/.bin/tailwindcss build views/app.css -o public/styles/app.css

lint:
	golangci-lint run

start:
	go run cmd/piplayer/*.go

start-vlc:
	cd ~/Music/iTunes/iTunes\ Media/Music && /Applications/VLC.app/Contents/MacOS/VLC -I http --http-password hest123
