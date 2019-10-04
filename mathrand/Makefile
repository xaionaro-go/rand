
default:
	echo 'no action'

generate-report:
	@echo '|func|passed|weak|failed|score|'
	@echo '|----|------|----|------|-----|'
	@bash -ec 'for FN in docs/dieharder-results/*.txt; do f="$${FN##*prngSample-}"; f="$${f%%.bin.txt}"; LINK="$$(grep -A 5 "$$f" *.go | grep "See also:" | awk -F "also: " "{print \$$2}" | head -1)"; if [ "$$LINK" = "" ]; then NAME="$$f"; else NAME="[$$f]($$LINK)"; fi; grep -E "(PASSED|WEAK|FAILED)" "$$FN" | awk -F "|" "{print \$$6}" | sort | uniq -c | awk "{if(\$$2==\"FAILED\"){failed=\$$1}; if(\$$2==\"PASSED\"){passed=\$$1}; if(\$$2==\"WEAK\"){weak=\$$1};} END {print \"|$$NAME|\"passed\"|\"weak\"|\"failed\"|[\"(passed + weak*0.5)\"](./$$FN)|\"}"; done'

dieharder-tests:
	go run ../internal/autogen/
	go fmt ./...
	go test ./...
	bash -exc 'for f in /tmp/prngSample-*bin; do FNAME=docs/dieharder-results/"$$(basename $$f)".txt; if [ -f "$$FNAME" ]; then continue; fi; timeout 7200 dieharder -g 201 -f "$$f" -a > $$FNAME; done'
