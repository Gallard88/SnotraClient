# Thomas Burns 2014
# thomasburns88@gmail.com
#

all:
	go build -v -o SnotraClient

install: SnotraClient
	install SnotraClient /usr/local/bin

uninstall:
	rm /usr/local/bin/SnotraClient

clean:
	rm -rf SnotraClient

