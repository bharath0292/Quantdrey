CARGO = cargo

build:
	$(CARGO) build --release

dev:
	$(CARGO) run

watch:
	$(CARGO) watch --ignore "*.csv" -w . -x run

fix deps:
	$(CARGO) machete --fix

clean:
	$(CARGO) clean

test:
	$(CARGO) test
