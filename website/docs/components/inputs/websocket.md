---
title: websocket
type: input
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the contents of:
     lib/input/websocket.go
-->


Connects to a websocket server and continuously receives messages.


import Tabs from '@theme/Tabs';

<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

import TabItem from '@theme/TabItem';

<TabItem value="common">

```yaml
# Common config fields, showing default values
input:
  websocket:
    url: ws://localhost:4195/get/ws
```

</TabItem>
<TabItem value="advanced">

```yaml
# All config fields, showing default values
input:
  websocket:
    url: ws://localhost:4195/get/ws
    open_message: ""
    oauth:
      access_token: ""
      access_token_secret: ""
      consumer_key: ""
      consumer_secret: ""
      enabled: false
      request_url: ""
    basic_auth:
      enabled: false
      password: ""
      username: ""
```

</TabItem>
</Tabs>

It is possible to configure an `open_message`, which when set to a
non-empty string will be sent to the websocket server each time a connection is
first established.

## Fields

### `url`

The URL to connect to.


Type: `string`  
Default: `"ws://localhost:4195/get/ws"`  

```yaml
# Examples

url: ws://localhost:4195/get/ws
```

### `open_message`

An optional message to send to the server upon connection.


Type: `string`  
Default: `""`  

### `oauth`

Allows you to specify open authentication.


Type: `object`  
Default: `{"access_token":"","access_token_secret":"","consumer_key":"","consumer_secret":"","enabled":false,"request_url":""}`  

```yaml
# Examples

oauth:
  access_token: baz
  access_token_secret: bev
  consumer_key: foo
  consumer_secret: bar
  enabled: true
  request_url: http://thisisjustanexample.com/dontactuallyusethis
```

### `basic_auth`

Allows you to specify basic authentication.


Type: `object`  
Default: `{"enabled":false,"password":"","username":""}`  

```yaml
# Examples

basic_auth:
  enabled: true
  password: bar
  username: foo
```

