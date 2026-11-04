VERSION ?=

# make tag VERSION=v0.1.102
tag:
	@test -n "$(VERSION)" || (echo "VERSION is required, e.g. make tag VERSION=v0.1.102" && exit 1)
	git tag $(VERSION)
	git push origin $(VERSION)
