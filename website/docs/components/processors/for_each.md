---
title: for_each
type: processor
status: stable
categories: ["Composition"]
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the corresponding source file under internal/impl/<provider>.
-->

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

A processor that applies a list of child processors to messages of a batch as though they were each a batch of one message.

```yml
# Config fields, showing default values
label: ""
for_each: []
```

This is useful for forcing batch wide processors such as [`dedupe`](/docs/components/processors/dedupe) or interpolations such as the `value` field of the `metadata` processor to execute on individual message parts of a batch instead.

Please note that most processors already process per message of a batch, and this processor is not needed in those cases.


