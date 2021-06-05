const Web3 = require('web3');
const HDWalletProvider = require("@truffle/hdwallet-provider");
import { arbitrage } from "../../Build/Arbitrage.json"; //TODO: Find a way to bring a compiled contract from truffle

class Web3Connector {
  public async executeTransaction(
    privateKey: string, //TODO: Check more secure way
    transactionParameters: string, //For now, string, soon convert to an object or array with desired values
    contractAddress: string
  ) : Promise<any> {
    const currentNet = process.env.CURRENT_NET;
    const accountProvider = new HDWalletProvider({
      privateKeys: [privateKey],
      providerOrUrl: currentNet,
    });
    const web3 = new Web3(accountProvider);
    const owner = web3.eth.accounts.privateKeyToAccount(privateKey);
    const contract = new web3.eth.Contract(arbitrage as any, contractAddress);

    var transactionLogs = await contract.startArbitrage(transactionParameters);

    return transactionLogs
}

export default new Web3Connector();
