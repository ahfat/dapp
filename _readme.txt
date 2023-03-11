Infura.io

MAINNET
url
https://mainnet.infura.io/v3/0784777345ab4b78a0cdfe76c50e831a

current block number
0x1002fea

api key
0784777345ab4b78a0cdfe76c50e831a

SEPOLIA
url
https://sepolia.infura.io/v3/0784777345ab4b78a0cdfe76c50e831a

account
#1
Public address of the key:   0x09684c61Ad3B7bD5c45d5E905fA45B2e7e4C8bc9
Path of the secret key file: keystore/UTC--2023-03-09T08-50-53.301214000Z--09684c61ad3b7bd5c45d5e905fa45b2e7e4c8bc9
pw: ahfatahfat

// USE THIS ONE in the demo
#2 Self generated, imported to Account 3 of Metamask
Public(Wallet) address of the key:   0x74817361919f209aFb38B490958751953c91598f
Path of the secret key file: keystore/UTC--2023-03-10T03-53-46.618961000Z--74817361919f209afb38b490958751953c91598f
Private key: 6659afa8d01d1af51ef2b3036e13b3080fa8700f66f54952682cc323b93a6c78
pw: ahfatahfat

#3 Metamask wallet Account 1
//Public address of the key:   0x027a8Bf9bD2AD9108df149957f26FFDDF8Fcfe4c
//Path of the secret key file: keystore/UTC--2023-03-09T09-54-59.861293000Z--47962dc09f37a761560c8fc00fc7f45a454a4a91
Private key: 689dd1a25f76e388a766a99c771cc4e0849c8130e4b4a6a7c186d0b5000cf1e2
Wallet account address: 0x027a8Bf9bD2AD9108df149957f26FFDDF8Fcfe4c

get ether balance
// created by Alchemy
curl https://sepolia.infura.io/v3/0784777345ab4b78a0cdfe76c50e831a \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getBalance","params": ["0x027a8Bf9bD2AD9108df149957f26FFDDF8Fcfe4c", "latest"],"id":1}'

// self-created key
curl https://sepolia.infura.io/v3/0784777345ab4b78a0cdfe76c50e831a \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getBalance","params": ["0x47962DC09F37A761560c8fC00Fc7F45a454a4A91", "latest"],"id":1}'

get ether for sepolia
https://faucet.paradigm.xyz/?ref=infura-blog
https://www.alchemy.com/overviews/sepolia-eth
https://sepoliafaucet.com/


Or use local test server
ganache-cli