import fs from "fs"
import Web3 from "web3"
import Ganache from "ganache-core";

class LocalNetworkService {
  public async startLocalNetworkDeployContract(
    serverUrl: string,
    blockNumber: number,
    privateKey: string
  ): Promise<void> {
    const serverStartUpOptions = {
      accounts: [{
        balance: 2000,
        secretKey: privateKey
      }],
      fork: `${serverUrl}/@${blockNumber}`
    };
    const server = Ganache.server(serverStartUpOptions);

    //const web3Client = new Web3(ganache.provider()

  }


  private async deployContract(): void{
    const provider = Ganache.provider();
    provider.setMaxListeners(15);       // Suppress MaxListenersExceededWarning warning
    const web3 = new Web3(provider);
    const account = await web3.eth.getAccounts();


    // FALTA PEGAR O BYTECODE DO CÓDIGO, PEGAR O ABI E FAZER O DEPLOY COM A CONTA SERVIDA PELO GANACHE 
    
    // Read in the compiled contract code and fetch ABI description and the bytecode as objects
    const compiled = JSON.parse(fs.readFileSync("output/contracts.json"));
    const abi = compiled.contracts["ContractName.sol"]["ContractName"].abi;
    const bytecode = compiled.contracts['ContractName.sol']['ContractName'].evm.bytecode.object;

    // Deploy the contract and send it gas to run.
    this.contract = await new web3.eth.Contract(abi)
        .deploy({data:'0x'+ bytecode, arguments: []})
        .send({from: this.accounts[0], gas:'5000000'});

  }

}

export default new LocalNetworkService();