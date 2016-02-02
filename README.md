# SapiServer

This is SapiServer repository. SapiServer is Windows SAPI simple web API server.

## API

### Get Voice Index

Get available voice index on server machine.

`GET /sapi/voices`

#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
    {
        "index":0,
        "name":"Microsoft David Desktop",
        "gender":"Male",
        "language":"409",
        "vendor":"Microsoft",
        "age":"Adult",
        "description":"Microsoft David Desktop - English (United States)"
    },
    {
        "index":1,
        "name":"Microsoft Hazel Desktop",
        "gender":"Female",
        "language":"809",
        "vendor":"Microsoft",
        "age":"Adult",
        "description":"Microsoft Hazel Desktop - English (Great Britain)"
    },
    {
        "index":2,
        "name":"Microsoft Zira Desktop",
        "gender":"Female",
        "language":"409",
        "vendor":"Microsoft",
        "age":"Adult",
        "description":"Microsoft Zira Desktop - English (United States)"
    },
    {
        "index":3,
        "name":"Microsoft Haruka Desktop",
        "gender":"Female",
        "language":"411",
        "vendor":"Microsoft",
        "age":"Adult",
        "description":"Microsoft Haruka Desktop - Japanese"
    }
]
```


### Get Speech Data

Create speech wave data

`POST /sapi/create`


#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **message** | *string* | Speech sentence or xml | `"This is a pen."` |
| **voice_index** | *string* | Speech voice index | `"0"` |

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **sapi_id** | *string* | unique id | `"abcd12344"` |


#### Response Example

```
HTTP/1.1 200 OK
```

```
<wave binary>
```


## Install on server

Require go installation.

```ps1
> Set-ExecutionPolicy Unrestricted
```

```ps1
> ./register_service.ps1
```

```ps1
> Server-Start SapiServer
```

and check SapiServer port :9081 on Windows Firewall.
