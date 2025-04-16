# 🛠️ Setting Up Sentinel

## 🌟 Becoming a Provider

### 🪙 Create a Wallet Account for the Provider

To create a wallet account for the provider, run the following command:

```shell
uramd  add <provider-wallet-account> --keyring-backend test
```

### 🔍 Get the Provider Public Key

Retrieve the provider's public key with:

```bash
uramd  show <provider-wallet-account> -p --keyring-backend test | jq -r .key
```

Convert the result to Bech32 format:

```bash
uramd debug pubkey-raw <result-from-above-command> | grep "Bech32 Acc" | awk '{ print $NF }'
```

> **ℹ️ Note:** Request tokens from the faucet to bond the provider in the relevant 💬 Discord channel.

### 🤝 Bond the Provider

Bond your provider by executing the following command:

```shell
uramd tx arkeo bond-provider <provider-pubkey> <service-providing> <bond-amount> --from <provider-wallet> --keyring-backend 🧪 --fees 20uarkeo
```

## 🚀 Starting the Sentinel Service

### 🛠️ Build the Sentinel Binary

Compile the Sentinel binary by running:

```bash
TAG=testnet make install
```

### ⚙️ Set Environment Variables

Configure the environment variables as follows:

```bash
NET="testnet" \
MONIKER="<your-moniker>" \
WEBSITE="<website-address>" \
DESCRIPTION="<provider description>" \
LOCATION="<location>" \
PORT="<sentinel-port>" \
SOURCE_CHAIN="<arkeo chain address>" \
EVENT_STREAM_HOST="<arkeo event stream host (rpc address)>" \
FREE_RATE_LIMIT=<free tier rate limit> \
FREE_RATE_LIMIT_DURATION="<duration>" \
CLAIM_STORE_LOCATION="~/.arkeo/claims" \
CONTRACT_CONFIG_STORE_LOCATION="~/.arkeo/contract_configs" \
PROVIDER_PUBKEY="<Provider PubKey>" \
PROVIDER_CONFIG_STORE_LOCATION="~/.arkeo/provider"
```

### ▶️ Run Sentinel

Start the Sentinel service by executing:

```bash
sentinel
```

When Sentinel starts, you should see output similar to the following:

```bash
I[2024-10-28|11:58:20.056] Starting Sentinel (reverse proxy)....        
Moniker                          <your-moniker>
Website                          <website address>
Description                      <provider description>
Location                         <location>
Port                             <sentinel-port>
TLS Certificate                  
TLS Key                          
Source Chain                     <arkeo chain address>
Event Stream Host                <arkeo event stream host (rpc address)>
Provider PubKey                  <Provider Pubkey>
Claim Store Location             ~/.arkeo/claims
Contract Config Store Location   ~/.arkeo/contract_configs
Free Tier Rate Limit             <free tier rate limit> requests per <duration>
Provider Config Store Location   ~/.arkeo/provider
I[2024-10-28|11:58:20.057] service start                                msg="Starting WSEvents service" impl=WSEvents
```

## 📝 Add Provider Metadata

Once the Sentinel service is running, update the provider metadata by running:

```shell
uramd tx arkeo mod-provider <provider-pubkey> <service> "http://<sentineladdress>/metadata.json" <nonce> <status> <min-contract-duration> <max-contract-duration> <subscription-rates> <pay-as-you-go-rates> <settlement-duration> --from <provider-wallet> --keyring-backend  --fees 20uarkeo
```

## Sequence Diagram

```mermaid
sequenceDiagram
    participant Provider
    participant ARKEO_CLI
    participant ARKEO
    participant Sentinel

    Provider->>ARKEO_CLI: Bond Provider (with bond amount and service)
    ARKEO_CLI->>ARKEO: Validate and Execute Transaction
    ARKEO-->>ARKEO_CLI: Tx hash returned to user
    ARKEO_CLI-->>Provider: Verify Tx hash for success
    ARKEO-->>ARKEO_CLI: Tx Result
    ARKEO_CLI-->>Provider: Tx Result

    Provider->>Sentinel: Starts Sentinel with Provider PubKey
    Sentinel-->>Provider: URL for provider metadata and service data
    Sentinel->>ARKEO: Listen to Events via Websockets
    ARKEO->>Sentinel: Websocket Response

    Provider->>ARKEO_CLI: Create Tx Mod Provider (with Sentinel Address)
    ARKEO_CLI->>ARKEO: Validate and Execute Transaction
    ARKEO-->>ARKEO_CLI: Tx hash
    ARKEO-->>ARKEO_CLI: Tx Result
    ARKEO_CLI-->>Provider: Verify Tx hash for success
    ARKEO_CLI-->>Provider: Tx Result

```