# Run sharness tests
#
# Copyright (c) 2014 Christian Couder
# MIT Licensed; see the LICENSE file in this repository.
#
# NOTE: run with TEST_VERBOSE=1 for verbose sharness tests.

T = $(sort $(wildcard t[0-9][0-9][0-9][0-9]-*.sh))
LIBDIR = lib
SHARNESSDIR = sharness
AGGREGATE = $(LIBDIR)/$(SHARNESSDIR)/aggregate-results.sh

# Binaries generated
BINS =
BINS += bin/fs-repo-migrations
BINS += bin/ipfs-0-to-1
BINS += bin/ipfs-1-to-2
BINS += bin/ipfs-2-to-3
BINS += bin/ipfs-3-to-4
BINS += bin/ipfs-4-to-5
BINS += bin/ipfs-5-to-6
BINS += bin/ipfs-6-to-7
BINS += bin/ipfs-update
BINS += bin/random-files
BINS += bin/go-sleep
BINS += bin/pollEndpoint
BINS += bin/iptb

# Source files location
FS_REPO_SRC = ../
IPFS_0_TO_1_SRC = ../ipfs-0-to-1
IPFS_1_TO_2_SRC = ../ipfs-1-to-2
IPFS_2_TO_3_SRC = ../ipfs-2-to-3
IPFS_3_TO_4_SRC = ../ipfs-3-to-4
IPFS_4_TO_5_SRC = ../ipfs-4-to-5
IPFS_5_TO_6_SRC = ../ipfs-5-to-6
IPFS_6_TO_7_SRC = ../ipfs-6-to-7
IPFS_UPDATE_SRC = ../Godeps/_workspace/src/github.com/ipfs/ipfs-update
POLL_ENDPOINT_SRC = ./dependencies/pollEndpoint
IPTB_SRC = ./dependencies/iptb

# User might want to override those on the command line
GOFLAGS =

all: aggregate

clean: clean-test-results
	@echo "*** $@ ***"
	-rm -rf $(BINS)

clean-test-results:
	@echo "*** $@ ***"
	-rm -rf test-results

$(T): clean-test-results deps
	@echo "*** $@ ***"
	./$@ -v

aggregate: clean-test-results $(T)
	@echo "*** $@ ***"
	ls test-results/t*-*.sh.*.counts | $(AGGREGATE)

# Needed dependencies.
deps: sharness $(BINS)

sharness:
	@echo "*** checking $@ ***"
	lib/install-sharness.sh

find_go_files = $(shell find $(1) -name "*.go")

bin/fs-repo-migrations: $(call find_go_files, $(FS_REPO_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(FS_REPO_SRC)

bin/ipfs-update: $(call find_go_files, $(IPFS_UPDATE_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_UPDATE_SRC)

bin/random-files:
	@echo "*** installing $@ ***"
	go get github.com/jbenet/go-random-files/random-files
	go build $(GOFLAGS) -o $@ github.com/jbenet/go-random-files/random-files

bin/go-sleep:
	@echo "*** installing $@ ***"
	go get github.com/chriscool/go-sleep
	go build $(GOFLAGS) -o $@ github.com/chriscool/go-sleep

bin/iptb: $(call find_go_files, $(IPTB_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPTB_SRC)

bin/ipfs-0-to-1: $(call find_go_files, $(IPFS_0_TO_1_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_0_TO_1_SRC)

bin/ipfs-1-to-2: $(call find_go_files, $(IPFS_1_TO_2_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_1_TO_2_SRC)

bin/ipfs-2-to-3: $(call find_go_files, $(IPFS_2_TO_3_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_2_TO_3_SRC)

bin/ipfs-3-to-4: $(call find_go_files, $(IPFS_3_TO_4_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_3_TO_4_SRC)

bin/ipfs-4-to-5: $(call find_go_files, $(IPFS_4_TO_5_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_4_TO_5_SRC)

bin/ipfs-5-to-6: $(call find_go_files, $(IPFS_5_TO_6_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_5_TO_6_SRC)

bin/ipfs-6-to-7: $(call find_go_files, $(IPFS_6_TO_7_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(IPFS_6_TO_7_SRC)

bin/pollEndpoint: $(call find_go_files, $(POLL_ENDPOINT_SRC)) BUILD-OPTIONS
	@echo "*** installing $@ ***"
	go build $(GOFLAGS) -o $@ $(POLL_ENDPOINT_SRC)

BUILD-OPTIONS: FORCE
	@bin/checkflags '$@' '$(GOFLAGS)' '*** new Go flags ***'

.PHONY: all clean clean-test-results $(T) aggregate deps sharness FORCE

