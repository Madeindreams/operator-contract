require("@nomicfoundation/hardhat-toolbox");
require("./tasks/deployment")
require('dotenv').config()
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {

  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
        chainId: 31337,
    },
    polygon:{
      url: process.env.RPC,
      chainId: 137,
      accounts: [process.env.OWNER_PRIVATE_KEY]

    }
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    artifacts: "./artifacts",
    tasks: "./tasks",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  solidity: "0.8.22",
};
