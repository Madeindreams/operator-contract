.DEFAULT_GOAL := build

#
# Variables and exports
#
export CGO_ENABLED=0
export GOPRIVATE=github.com/madeindreams/*

SHELL := /bin/bash

SMARTCONTRACT_ABI_DIR := tmp/smartcontracts-abi
SMARTCONTRACT_BIN_DIR := tmp/smartcontracts-bin
SMARTCONTRACT_GOLANG_DIR := go
SOLIDITY_DIR := ./contracts


SMARTCONTRACTS_IN_GOLANG := Operator.go IERC20.go IUniswapV2Router01.go IQuoter.go


#
# Functions
#
f_get_abi_dir = $(SMARTCONTRACT_ABI_DIR)/$(call f_get_packagename, $(1))
f_get_abi_filename = $(SMARTCONTRACT_ABI_DIR)/$(call f_get_packagename, $(1))/$(basename $(1)).abi
f_get_bin_dir = $(SMARTCONTRACT_BIN_DIR)/$(call f_get_packagename, $(1))
f_get_bin_filename = $(SMARTCONTRACT_BIN_DIR)/$(call f_get_packagename, $(1))/$(basename $(1)).bin
f_get_out_dir = $(SMARTCONTRACT_GOLANG_DIR)/$(call f_get_packagename, $(1))
f_get_out_filename = $(SMARTCONTRACT_GOLANG_DIR)/$(call f_get_packagename, $(1))/$(basename $(1)).go
f_get_packagename = $(shell tr '[:upper:]' '[:lower:]' <<<$(basename $(1)))
f_get_solidity_filepath = $(addprefix $(SOLIDITY_DIR)/, $(basename $(1)).sol )

copy_files:
	cp node_modules/@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol contracts/
	cp node_modules/@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol contracts/
	cp node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol contracts/

#
# Targets
#
build: copy_files contracts
	
clean: clean-contracts

clean-contracts:
	-rm -fr $(SMARTCONTRACT_ABI_DIR)
	-rm -fr $(SMARTCONTRACT_BIN_DIR)
	-rm -fr $(SMARTCONTRACT_GOLANG_DIR)
	-rm -fr tmp

contracts: $(SMARTCONTRACTS_IN_GOLANG)
$(SMARTCONTRACTS_IN_GOLANG):
	@echo "Generating Golang" $(call f_get_out_filename, $@) "from Smart Contract" $(call f_get_solidity_filepath, $@)

	mkdir -p $(call f_get_abi_dir, $@)
	mkdir -p $(call f_get_bin_dir, $@)
	mkdir -p $(call f_get_out_dir, $@)
	solc --abi $(call f_get_solidity_filepath, $@) -o $(call f_get_abi_dir, $@) --overwrite
	solc --bin $(call f_get_solidity_filepath, $@) -o $(call f_get_bin_dir, $@) --overwrite
	abigen --abi=$(call f_get_abi_filename, $@) --bin=$(call f_get_bin_filename, $@) \
		   --pkg=$(call f_get_packagename, $@)  --out=$(call f_get_out_filename, $@) \
