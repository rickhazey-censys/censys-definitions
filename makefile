all: go python cpp

go: proto/*.proto
	./build.py go

python: proto/*.proto
	./build.py python

cpp: proto/*.proto
	./build.py cpp

deploy: python
	cd python && python setup.py sdist upload

clean:
	rm -f go/*.go python/zsearch_definitions/*.py cpp/*.h cpp/*.cc

.PHONY: deploy clean cpp python go
