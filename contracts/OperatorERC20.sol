/*SPDX-License-Identifier: MIT
   _____              .___     .___       ________                                       
  /     \ _____     __| _/____ |   | ____ \______ \_______   ____ _____    _____   ______
 /  \ /  \\__  \   / __ |/ __ \|   |/    \ |    |  \_  __ \_/ __ \\__  \  /     \ /  ___/
/    Y    \/ __ \_/ /_/ \  ___/|   |   |  \|    `   \  | \/\  ___/ / __ \|  Y Y  \\___ \ 
\____|__  (____  /\____ |\___  >___|___|  /_______  /__|    \___  >____  /__|_|  /____  >
        \/     \/      \/    \/         \/        \/            \/     \/      \/     \/ 

author: Ian Dorion

This is the OperatorERC20 Contract:

**/

import "../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";

pragma solidity 0.8.22;

contract OperatorERC20 is ERC20 {


    constructor() ERC20("Operator","Op") {
   
    }

   

}