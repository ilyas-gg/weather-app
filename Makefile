.PHONY: build deb install clean

build:
	go build -o weather-app ./cmd/linux/cli/main.go

deb:
	@chmod +x scripts/build-deb.sh
	@./scripts/build-deb.sh

install:
	sudo dpkg -i weather-app_1.0.0-1.deb

uninstall:
	sudo apt remove weather-app

purge:
	sudo apt purge weather-app

clean:
	rm -f weather-app
	rm -rf weather-app_1.0.0-1
	rm -f *.deb
