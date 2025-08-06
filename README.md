# NAAS Guide: Orchestrator Node Setup

Welcome! This guide helps **Node-as-a-Service (NAAS)** providers to set up and run their own **Orchestrator Node**

## 📦 Binary Release

Download the latest release for your platform from GitHub:

➡️ [Orchestrator Node Software v0.0.1](https://github.com/NodeOps-app/uno-node-software/releases/tag/v0.0.1)

---

## 🛠 Prerequisites

Before deploying the node, ensure your environment meets the minimum specs:

- **RAM**: 64 MB
- **CPU**: 1 x86 Core >2.5GHz
- **Internet**: 10 Mbps stable connection
- **OS**: Linux recommended

---

## ⚙️ Environment Configuration

Set the following environment variables before running the binary:

```bash
export RPC_URL="https://nodeops-orchestrator-network.calderachain.xyz/http"
export CONTRACT_ADDRESS="0xa234B8924Bb2707195664e4C4Cf17668db9b7286"
export PRIVATE_KEY="<your-operator-wallet-private-key>"
```

---

## 📍 NFT Details

The following are the **UNO NFT contract addresses** deployed across supported chains. These contracts represent the NFT licenses that grant permission to operate as a node

You can query the NFTs held by a wallet address on any of these networks using the appropriate chain and contract address below. To programmatically check the licenses held by a user, use the `checkUserLicenses()` function

Make sure the user wallet **owns** or **is delegated** at least one license NFT from any of the above chains. This is required for the node to start participating in epochs and reporting activity

| Chain    | UNO NFT Contract Address                     |
| -------- | -------------------------------------------- |
| Arbitrum | `0x2C8515B14Fb9feBCe6d3845d5A0e643D84205Ba1` |
| Ethereum | `0x2Eab5E11cfe22Db4dA9F63b0B5d3F620ADa9a4bf` |
| Base     | `0xC0d2aF6d32240494742Ae486B9b73ec6dCC54Aa1` |

---

### 🔍 How to Fetch a Wallet’s UNO NFTs

You can use Alchemy’s NFT API to check which UNO licenses a wallet holds.

**🔗 Example API Call on Arbitrum(Alchemy)**

```
GET https://arb-mainnet.g.alchemy.com/nft/v3/<ALCHEMY_API_KEY>/getNFTsForOwner?owner=<WALLET_ADDRESS>&contractAddresses[]=<UNO_CONTRACT_ADDRESS>&pageSize=100
```

As we attempt to retrieve the NFTs on the Arbitrum chain, we will use the 0x2C8515B14Fb9feBCe6d3845d5A0e643D84205Ba1 as the parameter for <UNO_CONTRACT_ADDRESS>

Replace placeholders:

- <ALCHEMY_API_KEY> — Your Alchemy API Key
- <WALLET_ADDRESS> — The address to check
- <UNO_CONTRACT_ADDRESS> — One of the supported license contracts (see table above)

---

### 🧠 NFT Contract ABI

To programmatically interact with the UNO License NFT (e.g. to read ownership, metadata, or validate delegation), use this ABI:

mapping(uint256 => uint256) public canonicalTokenId;

This maps tokenId ➝ canonical ID to support cross-chain license resolution.

You can find the full ABI used by the NFT contract [here](./uno-nft-abi.json)

## 🧩 Delegation Workflow

The delegation process involves signing license-specific messages and sending them to the orchestrator contract.

### 🔄 Delegate a License

```solidity
delegateLicense(address _operator, uint256 _licenseId, bytes signature)
```

• The operator signs a message containing license ID and their address.
• This signed message is passed to the orchestrator contract.
• The license becomes bound to the operator address.

### ❌ Undelegate a License

```solidity
undelegateLicense(uint256 _licenseId)
```

• Call this function to remove the operator’s assignment from a license.

## 🔎 Checking Delegation Status

Use the following read-only functions:

```solidity
// View current epoch
function currentEpoch() public view returns (uint256)

// Check assigned operator of a license
function getOperator(uint256 _licenseId) public view returns (address)

// Fetch operator's full info
function getOperatorInfo(address _operator) public view returns (OperatorInfo)
```

### 📡 Fetch Delegated Licenses for an Owner

You can fetch all licenses delegated by an address using:

```solidity
function getDelegatedLicensesForOwner(address _owner) public view returns (uint256[])
```

Or fetch multiple operators under the same owner:

```solidity
function getOperatorsForOwner(address _owner) public view returns (address[])
```

## 🧾 ABI Reference

The Orchestrator contract exposes functions and events to support:
• Delegation
• Undelegation
• Role-based access control
• Epoch tracking
• Operator activity reporting

You can find the full ABI [here](./orchestrator-abi.json)
