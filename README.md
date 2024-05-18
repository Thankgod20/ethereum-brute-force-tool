# Ethereum Address Brute Force Tool

## Overview

This project is a Go implementation designed to brute force Ethereum addresses with balances using low-level mnemonic generation. The mnemonics used are from early versions of Trust Wallet, which were less secure before their upgrades.

## Features

- **Low-Level Mnemonic Generation**: Utilizes the same low-level mnemonics that were used in early Trust Wallet versions.
- **Parallel Processing**: Leverages Go's concurrency features to maximize the brute-forcing speed.
- **Address Balance Checking**: Integrates with Ethereum and BSC nodes to check the balance of generated addresses.
- **Customizable Parameters**: Allows customization of path, number of addresses to checked, and more.

## Disclaimer

This tool is intended for educational purposes only. Unauthorized access to Ethereum wallets is illegal and unethical. Use this tool responsibly and only on wallets you own or have explicit permission to test.

## Prerequisites

- Go 1.16 or higher
- An Ethereum and BSC node or a service like Infura for balance checking
- Basic understanding of Ethereum and mnemonics

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/thankgod20/ethereum-brute-force-tool.git
   cd ethereum-brute-force-tool
   ```

2. **Install Dependencies**

   ```sh
   go mod tidy
   ```

3. **Build the Project**
   ```sh
   go build -o eth-brute-force
   ```

## Usage

1. **Set Up Configuration**
   Create a configuration file `config.json` with the following structure:

   ```json
   {
     "bsc_node_url": "https://bsc-dataseed1.binance.org/",
     "eth_node_url": "https://mainnet.infura.io/v3/API",
     "iteration": 1000,
     "path": "m/44'/60'/0'/0/0"
   }
   ```

2. **Run the Tool**

   ```sh
   ./eth-brute-force -config config.json
   ```

3. **Monitor Output**
   The tool will generate mnemonics, derive addresses, and check their balances,etc. Results will be displayed in the console.

## Configuration Options

- `num_addresses`: Number of addresses to generate and check.
- `eth_node_url`: URL of the Ethereum node to use for balance checking.
- `bsc_node_url`: URL of the Binance node to use for balance checking.

## Example Output

```
=========== Current Thread 49===========
[*] Mnemonics:- pitch analyst diary olive melt skin dirt eye fog interest buyer gym
[*] Address:- 0xec0ea226cB112813FaD1A865773DcbDB3a7EA909
[*] Balance Eth:- 0
[*] Nonce Eth:- 0
[*] Balance BSC:- 0
[*] Nonce BSC:- 0
[*] Limit:- (49/1000)
======================
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure that your code adheres to the existing style and passes all tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or issues, please open an issue on GitHub or contact the project maintainer.

---

**Warning:** Use of this tool against wallets that do not belong to you without permission is illegal and can result in severe penalties. The authors are not responsible for any misuse of this software.
