# Uram 1 TESTNET

URAM Sandbox Blockchain Environment

## Uram app-chain binaries installation (uramd)

```
go: go version go1.24.3 linux/amd64
name: uramd
```

## Prerequisites

### Install go

```
sudo rm -rvf /usr/local/go/
wget https://golang.org/dl/go1.24.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz
rm go1.24.3.linux-amd64.tar.gz
```

### Put PATH to ~/.profile

```
nano .profile
```

```
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

### Use source ~/.profile

```
source .profile
```

### Check go

```
go version
```

### Install packages

```
sudo apt update
sudo apt upgrade
sudo apt install mc btop nano screen git make build-essential
```

## Binary building

### Clone source from repo

```
git clone https://github.com/semalis/uram.git
```

```

### Build binary

```
make install
```

## Network launch

### Init

```bash:
uramd init "<moniker-name>" --chain-id  uram-testnet-1
```

### Set minimum-gas-prices = "" in app.toml to minimum-gas-prices = "0.0025uuram"

```
sed -i -e "s|^minimum-gas-prices *=.*|minimum-gas-prices = \"0.0025uuram\"|" $HOME/.uram/config/app.toml
```

### Generate keys

```bash:
# To create new keypair - make sure you save the mnemonics!
uramd keys add <key-name>
```

or

```
# Restore existing odin wallet with mnemonic seed phrase.
# You will be prompted to enter mnemonic seed.
uramd keys add <key-name> --recover
```

or

```
# Add keys using ledger
uramd keys show <key-name> --ledger
```

Check your key:

```
# Query the keystore for your public address
uramd keys show <key-name> -a
```

### Create account to genesis

```
uramd genesis add-genesis-account <key-name> 1000000uodis
```

### Create GenTX

```
# Create the gentx.
# Note, your gentx will be rejected if you use any amount greater than 1000000uuram.
# Make sure that all participants built their gentx files without typos.

uramd genesis gentx <key-name> 1000000uuram \
  --pubkey=$(uramd tendermint show-validator) \
  --chain-id=uram-testnet-1 \
  --moniker="my-moniker" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01"
```

### Add all accounts to genesis

```
# Add account addresses of all participants before generating genesis.
# (whose Gentx files you're using to generate genesis)
uramd genesis add-genesis-account <account-address> 1000000uuram
```

### Generate genesis

```
uramd genesis collect-gentxs
```

### Start network

```
uramd start
```

### ****Set Up uram Service****

Set up a service to allow binary to run in the background as well as restart automatically if it runs into any problems:
```
sudo tee /etc/systemd/system/uramd.service > /dev/null << EOF
[Unit]
Description=Uram app chain daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=uramd"
Environment="DAEMON_HOME=${HOME}/.uram"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_LOG_BUFFER_SIZE=512"
Environment="UNSAFE_SKIP_BACKUP=true"
User=$USER
ExecStart=${HOME}/go/bin/uramd start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
EOF
```

And start service:
```
sudo systemctl daemon-reload
sudo systemctl enable uramd 
sudo systemctl restart uramd
```

How you can check the logs
```
sudo journalctl -u uramd -f --output cat
```

How you can check blocks sync
```
curl http://localhost:26657/status | jq -r ".result.sync_info"
```