## okp4d query ibc-transfer denom-hash

Query the denom hash info from a given denom trace

### Synopsis

Query the denom hash info from a given denom trace

```
okp4d query ibc-transfer denom-hash [trace] [flags]
```

### Examples

```
okp4d query ibc-transfer denom-hash transfer/channel-0/uatom
```

### Options

```
      --grpc-addr string   the gRPC endpoint to use for this chain
      --grpc-insecure      allow gRPC over insecure channels, if not the server must use TLS
      --height int         Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help               help for denom-hash
      --node string        <host>:<port> to CometBFT RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string      Output format (text|json) (default "text")
```

### SEE ALSO

* [okp4d query ibc-transfer](okp4d_query_ibc-transfer.md)	 - IBC fungible token transfer query subcommands
