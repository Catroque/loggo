#--------------------------------------------------------[]
#
#   Makefile
#
#   - Aplicacao para testes da biblioteca zerolog.
#  
#--------------------------------------------------------
#
#   Data....: 11/01/2021
#   Usuario.: Rinaldo Catroque Bersi
#  
#--------------------------------------------------------[]

APP = main

all:
	go build -o dist/$(APP) cmd/main.go 

ex:
	go build -o dist/$(APP) examples/main.go 

exec:
	dist/$(APP)

clean:
	@find . -name '*~' -delete
	@find . -name '.*~' -delete
	@rm -f .*~    
	@rm -f *~
	@rm -f dist/$(APP)