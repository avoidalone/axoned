## okp4d query gov proposals

Query proposals with optional filters

```
okp4d query gov proposals [flags]
```

### Examples

```
okp4d query gov proposals --depositor cosmos1...
okp4d query gov proposals --voter cosmos1...
okp4d query gov proposals --proposal-status (PROPOSAL_STATUS_DEPOSIT_PERIOD|PROPOSAL_STATUS_VOTING_PERIOD|PROPOSAL_STATUS_PASSED|PROPOSAL_STATUS_REJECTED|PROPOSAL_STATUS_FAILED)
```

### Options

```
      --depositor account address or key name                                                                        
      --grpc-addr string                                                                                             the gRPC endpoint to use for this chain
      --grpc-insecure                                                                                                allow gRPC over insecure channels, if not the server must use TLS
      --height int                                                                                                   Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help                                                                                                         help for proposals
      --no-indent                                                                                                    Do not indent JSON output
      --node string                                                                                                  <host>:<port> to CometBFT RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string                                                                                                Output format (text|json) (default "text")
      --page-count-total                                                                                             
      --page-key binary                                                                                              
      --page-limit uint                                                                                              
      --page-offset uint                                                                                             
      --page-reverse                                                                                                 
      --proposal-status ProposalStatus (unspecified | deposit-period | voting-period | passed | rejected | failed)    (default unspecified)
      --voter account address or key name                                                                            
```

### SEE ALSO

* [okp4d query gov](okp4d_query_gov.md)	 - Querying commands for the gov module
