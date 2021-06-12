import shell from "shelljs"


class GanacheNetwork {
  public async startLocalNetwork(
    blockNumber: number
  ): Promise<void> {

  }

  private startNetwork(rpcServerAddr: string, blockNumber: number): void{
    shell.exec(`ganache-cli --fork ${rpcServerAddr}@${blockNumber}`);
  }

}

export default new GanacheNetwork();