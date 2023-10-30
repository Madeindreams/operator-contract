require("@nomicfoundation/hardhat-toolbox");
require('dotenv').config();
const Operator = require('../artifacts/contracts/Operator.sol/Operator.json')
const { ethers, NonceManager } = require("ethers");

const addressProvider = "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb"  // polygon

// Define the task to deploy the Operator contract
task(
    "deploy",
    "Deploy The contract",
    async (_, { network }) => {
        const provider = new ethers.JsonRpcProvider(process.env.RPC);
        const owner = new NonceManager(new ethers.Wallet(process.env.OWNER_PRIVATE_KEY, provider));

        // deploy www3 share contract 
        const OperatorContract = new ethers.ContractFactory(Operator.abi, Operator.bytecode, owner);
        const operator = await OperatorContract.deploy(addressProvider);
        console.log("operator contract address", operator.target)

    })
