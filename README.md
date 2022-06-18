# Backend Engineer Take Home Assessment

## Assumptions

1. Fintech clients have a valid authentication token for API access
1. Buy/Sell operations are performed immediately (no trading flows with orders for MKT/LMT, etc.)
1. Currency exchange is performed against the user's USD balance
1. Currency list returns the list of currencies in the specified customer portfolio (An endpoint returning the full list of currencies is out of this assignment scope)

## Planned REST API:

### Authentication
Each API endpoint takes the `Authentication: Bearer` token identifying/authorizing a specific fintech client.

### Endpoints and sample payload
1. ### Currency list: For a given user, get the currencies they can transact
>Endpoint:
>
>`GET /currencies?userId=<user_id>&limit=10&offset=0`
>
>Query Parameters:
>- `user_id: string`   - Required. Unique user identifier
>- `limit: number`     - Optional. Maximum number of currencies to return (defaults to 10)
>- `offset: number`    - Optional. Offset from the start of the list (defaults to 0)
>
>
>Response payload:
>
>```JS
>{
>    meta: {
>        pagination: {
>            total: 30,
>            limit: 10,
>            offset: 0,
>        }
>    },
>    currencies: [
>        "USD",
>        "EUR",
>        "BTC",
>        "ETH",
>        "USDC",
>        ...
>    ]
>}
>```

2. ### Buy/Sell a selected currency for a given user
> Endpoint
>
>`POST /exchange?userId=<user_id>`
>
>Query Parameters:
>- `user_id: string`   - Required. Unique user identifier
> 
>Request body payload:
>
>```JS
>{
>    direction: "BUY" | "SELL",
>    currency: 'BTC',
>    amount: 100,
>}
>```
>
>Response payload:
>
>
>```JS
>{
>    success: true,
>    balance: {
>        currency: 'BTC',
>        amount: 150
>    }
>}
>```
>


3. ### Balances: For a given user, get a list of balances

>Endpoint
>
>`GET /balances?userId=<user_id>?limit=10&offset=0`
>
>Query Parameters:
>- `user_id: string`   - Required. Unique user identifier
>- `limit: number`     - Optional. Maximum number of currencies to return (defaults to 10)
>- `offset: number`    - Optional. Offset from the start of the list (defaults to 0)
> 
>Response payload:
>
>
>```JS
>{
>    meta: {
>        pagination: {
>            total: 2,
>            limit: 10,
>            offset: 0,
>        }
>    },
>    balances: [
>        {
>            currency: 'BTC',
>            balance: 13435435        
>        },
>        {
>            currency: 'ETH',
>            balance: 2.34444
>        }
>    ]
>}
>```
>