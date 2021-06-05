# É o Arbitas

## Environment Requirements

* Install NodeJS v12.X
* Install Yarn `npm install -g yarn`

## Environment Setup

### Install dependencies

```sh
yarn install
```

### Build

Unix based systems:
```sh
yarn build
```

Windows:
```sh
yarn buildWindows
```

It will generate transpiled js files into the `dist` folder.

### Execute

Run the following command providing the arguments.

1. Quantity of units desired to arbitrate (Required)
2. Ticker of source token to arbitrate (Required)
3. Ticker of destination token to arbitrate (Required)
4. Config to show values in a human readable way (Optional)

Unix based systems:
```sh
yarn start <QUANTITY> <FROM_TOKEN> <TO_TOKEN> [<HUMAN_READABLE>]
yarn start 10 BNB DAI
```

Windows
```sh
yarn startWindows <QUANTITY> <FROM_TOKEN> <TO_TOKEN> [<HUMAN_READABLE>]
yarn start 10 BNB DAI
```

## Available Tokens

| Token | Address                                   |
| ----- | ------------------------------------------ |
| ADA   | 0x3ee2200efb3400fabb9aacf31297cbdd1d435d47 |
| BNB   | 0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee |
| BUSD  | 0xe9e7cea3dedca5984780bafc599bd69add087d56 |
| DAI   | 0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3 |
| ETH   | 0x2170ed0880ac9a755fd29b2688956bd959f933f8 |
| USDC  | 0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d |
| USDT  | 0x55d398326f99059ff775485246999027b3197955 |
| WBNB  | 0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c |
