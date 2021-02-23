
all: assets/css/main.css app

app:
	go build

assets/css/main.css: src/css/main.css
	NODE_ENV=production npx tailwindcss-cli@latest build ./src/css/main.css -o ./assets/css/tailwind.css 

.PHONY:
linux:
	GOOS=linux GOARCH=amd64 go build

.PHONY:
deploy: linux
	scp ./bktw content.brian.dev:~/bktwnew

