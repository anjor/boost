# Groups
* [Actor](#Actor)
  * [ActorSectorSize](#ActorSectorSize)
* [Auth](#Auth)
  * [AuthNew](#AuthNew)
  * [AuthVerify](#AuthVerify)
* [Boost](#Boost)
  * [BoostDagstoreInitializeAll](#BoostDagstoreInitializeAll)
  * [BoostDagstoreInitializeShard](#BoostDagstoreInitializeShard)
  * [BoostDeal](#BoostDeal)
  * [BoostDummyDeal](#BoostDummyDeal)
  * [BoostIndexerAnnounceAllDeals](#BoostIndexerAnnounceAllDeals)
  * [BoostOfflineDealWithData](#BoostOfflineDealWithData)
* [Deals](#Deals)
  * [DealsConsiderOfflineRetrievalDeals](#DealsConsiderOfflineRetrievalDeals)
  * [DealsConsiderOfflineStorageDeals](#DealsConsiderOfflineStorageDeals)
  * [DealsConsiderOnlineRetrievalDeals](#DealsConsiderOnlineRetrievalDeals)
  * [DealsConsiderOnlineStorageDeals](#DealsConsiderOnlineStorageDeals)
  * [DealsConsiderUnverifiedStorageDeals](#DealsConsiderUnverifiedStorageDeals)
  * [DealsConsiderVerifiedStorageDeals](#DealsConsiderVerifiedStorageDeals)
  * [DealsPieceCidBlocklist](#DealsPieceCidBlocklist)
  * [DealsSetConsiderOfflineRetrievalDeals](#DealsSetConsiderOfflineRetrievalDeals)
  * [DealsSetConsiderOfflineStorageDeals](#DealsSetConsiderOfflineStorageDeals)
  * [DealsSetConsiderOnlineRetrievalDeals](#DealsSetConsiderOnlineRetrievalDeals)
  * [DealsSetConsiderOnlineStorageDeals](#DealsSetConsiderOnlineStorageDeals)
  * [DealsSetConsiderUnverifiedStorageDeals](#DealsSetConsiderUnverifiedStorageDeals)
  * [DealsSetConsiderVerifiedStorageDeals](#DealsSetConsiderVerifiedStorageDeals)
  * [DealsSetPieceCidBlocklist](#DealsSetPieceCidBlocklist)
* [I](#I)
  * [ID](#ID)
* [Log](#Log)
  * [LogList](#LogList)
  * [LogSetLevel](#LogSetLevel)
* [Market](#Market)
  * [MarketDataTransferUpdates](#MarketDataTransferUpdates)
  * [MarketGetAsk](#MarketGetAsk)
  * [MarketGetRetrievalAsk](#MarketGetRetrievalAsk)
  * [MarketImportDealData](#MarketImportDealData)
  * [MarketListDataTransfers](#MarketListDataTransfers)
  * [MarketListRetrievalDeals](#MarketListRetrievalDeals)
  * [MarketRestartDataTransfer](#MarketRestartDataTransfer)
  * [MarketSetAsk](#MarketSetAsk)
  * [MarketSetRetrievalAsk](#MarketSetRetrievalAsk)
* [Net](#Net)
  * [NetAddrsListen](#NetAddrsListen)
  * [NetAgentVersion](#NetAgentVersion)
  * [NetAutoNatStatus](#NetAutoNatStatus)
  * [NetBandwidthStats](#NetBandwidthStats)
  * [NetBandwidthStatsByPeer](#NetBandwidthStatsByPeer)
  * [NetBandwidthStatsByProtocol](#NetBandwidthStatsByProtocol)
  * [NetBlockAdd](#NetBlockAdd)
  * [NetBlockList](#NetBlockList)
  * [NetBlockRemove](#NetBlockRemove)
  * [NetConnect](#NetConnect)
  * [NetConnectedness](#NetConnectedness)
  * [NetDisconnect](#NetDisconnect)
  * [NetFindPeer](#NetFindPeer)
  * [NetPeerInfo](#NetPeerInfo)
  * [NetPeers](#NetPeers)
## Actor


### ActorSectorSize
There are not yet any comments for this method.

Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `34359738368`

## Auth


### AuthNew


Perms: admin

Inputs:
```json
[
  [
    "write"
  ]
]
```

Response: `"Ynl0ZSBhcnJheQ=="`

### AuthVerify


Perms: read

Inputs:
```json
[
  "string value"
]
```

Response:
```json
[
  "write"
]
```

## Boost


### BoostDagstoreInitializeAll


Perms: admin

Inputs:
```json
[
  {
    "MaxConcurrency": 123,
    "IncludeSealed": true
  }
]
```

Response:
```json
{
  "Key": "string value",
  "Event": "string value",
  "Success": true,
  "Error": "string value",
  "Total": 123,
  "Current": 123
}
```

### BoostDagstoreInitializeShard


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response: `{}`

### BoostDeal


Perms: admin

Inputs:
```json
[
  "07070707-0707-0707-0707-070707070707"
]
```

Response:
```json
{
  "DealUuid": "07070707-0707-0707-0707-070707070707",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "ClientDealProposal": {
    "Proposal": {
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1032,
      "VerifiedDeal": true,
      "Client": "f01234",
      "Provider": "f01234",
      "Label": "string value",
      "StartEpoch": 10101,
      "EndEpoch": 10101,
      "StoragePricePerEpoch": "0",
      "ProviderCollateral": "0",
      "ClientCollateral": "0"
    },
    "ClientSignature": {
      "Type": 2,
      "Data": "Ynl0ZSBhcnJheQ=="
    }
  },
  "IsOffline": true,
  "ClientPeerID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "DealDataRoot": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "InboundFilePath": "string value",
  "Transfer": {
    "Type": "string value",
    "ClientID": "string value",
    "Params": "Ynl0ZSBhcnJheQ==",
    "Size": 42
  },
  "ChainDealID": 5432,
  "PublishCID": null,
  "SectorID": 9,
  "Offset": 1032,
  "Length": 1032,
  "Checkpoint": 1,
  "Err": "string value",
  "NBytesReceived": 9
}
```

### BoostDummyDeal


Perms: admin

Inputs:
```json
[
  {
    "DealUUID": "07070707-0707-0707-0707-070707070707",
    "IsOffline": true,
    "ClientDealProposal": {
      "Proposal": {
        "PieceCID": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "PieceSize": 1032,
        "VerifiedDeal": true,
        "Client": "f01234",
        "Provider": "f01234",
        "Label": "string value",
        "StartEpoch": 10101,
        "EndEpoch": 10101,
        "StoragePricePerEpoch": "0",
        "ProviderCollateral": "0",
        "ClientCollateral": "0"
      },
      "ClientSignature": {
        "Type": 2,
        "Data": "Ynl0ZSBhcnJheQ=="
      }
    },
    "DealDataRoot": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Transfer": {
      "Type": "string value",
      "ClientID": "string value",
      "Params": "Ynl0ZSBhcnJheQ==",
      "Size": 42
    }
  }
]
```

Response:
```json
{
  "Accepted": true,
  "Reason": "string value"
}
```

### BoostIndexerAnnounceAllDeals
There are not yet any comments for this method.

Perms: admin

Inputs: `null`

Response: `{}`

### BoostOfflineDealWithData


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response:
```json
{
  "Accepted": true,
  "Reason": "string value"
}
```

## Deals


### DealsConsiderOfflineRetrievalDeals


Perms: admin

Inputs: `null`

Response: `true`

### DealsConsiderOfflineStorageDeals


Perms: admin

Inputs: `null`

Response: `true`

### DealsConsiderOnlineRetrievalDeals


Perms: admin

Inputs: `null`

Response: `true`

### DealsConsiderOnlineStorageDeals
There are not yet any comments for this method.

Perms: admin

Inputs: `null`

Response: `true`

### DealsConsiderUnverifiedStorageDeals


Perms: admin

Inputs: `null`

Response: `true`

### DealsConsiderVerifiedStorageDeals


Perms: admin

Inputs: `null`

Response: `true`

### DealsPieceCidBlocklist


Perms: admin

Inputs: `null`

Response:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

### DealsSetConsiderOfflineRetrievalDeals


Perms: admin

Inputs:
```json
[
  true
]
```

Response: `{}`

### DealsSetConsiderOfflineStorageDeals


Perms: admin

Inputs:
```json
[
  true
]
```

Response: `{}`

### DealsSetConsiderOnlineRetrievalDeals


Perms: admin

Inputs:
```json
[
  true
]
```

Response: `{}`

### DealsSetConsiderOnlineStorageDeals


Perms: admin

Inputs:
```json
[
  true
]
```

Response: `{}`

### DealsSetConsiderUnverifiedStorageDeals


Perms: admin

Inputs:
```json
[
  true
]
```

Response: `{}`

### DealsSetConsiderVerifiedStorageDeals


Perms: admin

Inputs:
```json
[
  true
]
```

Response: `{}`

### DealsSetPieceCidBlocklist


Perms: admin

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    }
  ]
]
```

Response: `{}`

## I


### ID


Perms: read

Inputs: `null`

Response: `"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"`

## Log


### LogList


Perms: write

Inputs: `null`

Response:
```json
[
  "string value"
]
```

### LogSetLevel


Perms: write

Inputs:
```json
[
  "string value",
  "string value"
]
```

Response: `{}`

## Market


### MarketDataTransferUpdates


Perms: write

Inputs: `null`

Response:
```json
{
  "TransferID": 3,
  "Status": 1,
  "BaseCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "IsInitiator": true,
  "IsSender": true,
  "Voucher": "string value",
  "Message": "string value",
  "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Transferred": 42,
  "Stages": {
    "Stages": [
      {
        "Name": "string value",
        "Description": "string value",
        "CreatedTime": "0001-01-01T00:00:00Z",
        "UpdatedTime": "0001-01-01T00:00:00Z",
        "Logs": [
          {
            "Log": "string value",
            "UpdatedTime": "0001-01-01T00:00:00Z"
          }
        ]
      }
    ]
  }
}
```

### MarketGetAsk


Perms: read

Inputs: `null`

Response:
```json
{
  "Ask": {
    "Price": "0",
    "VerifiedPrice": "0",
    "MinPieceSize": 1032,
    "MaxPieceSize": 1032,
    "Miner": "f01234",
    "Timestamp": 10101,
    "Expiry": 10101,
    "SeqNo": 42
  },
  "Signature": {
    "Type": 2,
    "Data": "Ynl0ZSBhcnJheQ=="
  }
}
```

### MarketGetRetrievalAsk


Perms: read

Inputs: `null`

Response:
```json
{
  "PricePerByte": "0",
  "UnsealPrice": "0",
  "PaymentInterval": 42,
  "PaymentIntervalIncrease": 42
}
```

### MarketImportDealData


Perms: write

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "string value"
]
```

Response: `{}`

### MarketListDataTransfers


Perms: write

Inputs: `null`

Response:
```json
[
  {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    }
  }
]
```

### MarketListRetrievalDeals
There are not yet any comments for this method.

Perms: read

Inputs: `null`

Response:
```json
[
  {
    "PayloadCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "ID": 5,
    "Selector": {
      "Raw": "Ynl0ZSBhcnJheQ=="
    },
    "PieceCID": null,
    "PricePerByte": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "UnsealPrice": "0",
    "StoreID": 42,
    "ChannelID": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "PieceInfo": {
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "Deals": [
        {
          "DealID": 5432,
          "SectorID": 9,
          "Offset": 1032,
          "Length": 1032
        }
      ]
    },
    "Status": 0,
    "Receiver": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "TotalSent": 42,
    "FundsReceived": "0",
    "Message": "string value",
    "CurrentInterval": 42,
    "LegacyProtocol": true
  }
]
```

### MarketRestartDataTransfer


Perms: write

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### MarketSetAsk


Perms: admin

Inputs:
```json
[
  "0",
  "0",
  10101,
  1032,
  1032
]
```

Response: `{}`

### MarketSetRetrievalAsk


Perms: admin

Inputs:
```json
[
  {
    "PricePerByte": "0",
    "UnsealPrice": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42
  }
]
```

Response: `{}`

## Net


### NetAddrsListen


Perms: read

Inputs: `null`

Response:
```json
{
  "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Addrs": [
    "/ip4/52.36.61.156/tcp/1347/p2p/12D3KooWFETiESTf1v4PGUvtnxMAcEFMzLZbJGg4tjWfGEimYior"
  ]
}
```

### NetAgentVersion


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
]
```

Response: `"string value"`

### NetAutoNatStatus


Perms: read

Inputs: `null`

Response:
```json
{
  "Reachability": 1,
  "PublicAddr": "string value"
}
```

### NetBandwidthStats


Perms: read

Inputs: `null`

Response:
```json
{
  "TotalIn": 9,
  "TotalOut": 9,
  "RateIn": 12.3,
  "RateOut": 12.3
}
```

### NetBandwidthStatsByPeer


Perms: read

Inputs: `null`

Response:
```json
{
  "12D3KooWSXmXLJmBR1M7i9RW9GQPNUhZSzXKzxDHWtAgNuJAbyEJ": {
    "TotalIn": 174000,
    "TotalOut": 12500,
    "RateIn": 100,
    "RateOut": 50
  }
}
```

### NetBandwidthStatsByProtocol


Perms: read

Inputs: `null`

Response:
```json
{
  "/fil/hello/1.0.0": {
    "TotalIn": 174000,
    "TotalOut": 12500,
    "RateIn": 100,
    "RateOut": 50
  }
}
```

### NetBlockAdd


Perms: admin

Inputs:
```json
[
  {
    "Peers": [
      "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
    ],
    "IPAddrs": [
      "string value"
    ],
    "IPSubnets": [
      "string value"
    ]
  }
]
```

Response: `{}`

### NetBlockList


Perms: read

Inputs: `null`

Response:
```json
{
  "Peers": [
    "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
  ],
  "IPAddrs": [
    "string value"
  ],
  "IPSubnets": [
    "string value"
  ]
}
```

### NetBlockRemove


Perms: admin

Inputs:
```json
[
  {
    "Peers": [
      "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
    ],
    "IPAddrs": [
      "string value"
    ],
    "IPSubnets": [
      "string value"
    ]
  }
]
```

Response: `{}`

### NetConnect


Perms: write

Inputs:
```json
[
  {
    "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Addrs": [
      "/ip4/52.36.61.156/tcp/1347/p2p/12D3KooWFETiESTf1v4PGUvtnxMAcEFMzLZbJGg4tjWfGEimYior"
    ]
  }
]
```

Response: `{}`

### NetConnectedness


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
]
```

Response: `1`

### NetDisconnect


Perms: write

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
]
```

Response: `{}`

### NetFindPeer


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
]
```

Response:
```json
{
  "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Addrs": [
    "/ip4/52.36.61.156/tcp/1347/p2p/12D3KooWFETiESTf1v4PGUvtnxMAcEFMzLZbJGg4tjWfGEimYior"
  ]
}
```

### NetPeerInfo


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"
]
```

Response:
```json
{
  "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Agent": "string value",
  "Addrs": [
    "string value"
  ],
  "Protocols": [
    "string value"
  ],
  "ConnMgrMeta": {
    "FirstSeen": "0001-01-01T00:00:00Z",
    "Value": 123,
    "Tags": {
      "name": 42
    },
    "Conns": {
      "name": "2021-03-08T22:52:18Z"
    }
  }
}
```

### NetPeers


Perms: read

Inputs: `null`

Response:
```json
[
  {
    "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Addrs": [
      "/ip4/52.36.61.156/tcp/1347/p2p/12D3KooWFETiESTf1v4PGUvtnxMAcEFMzLZbJGg4tjWfGEimYior"
    ]
  }
]
```
