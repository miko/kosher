---
layout: default
title: "the \"__\" element should contain \"__\""
parent: Step Definitions
---

# the "\_\_" element should contain "\_\_"

Verifies that an element matched by the selector, ID, name, or label contains the specified text.
{: .fs-6 .fw-300 }

## Pattern

```
^(?|the )"([^"]*)" (?|element )should contain "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                            | Restrictions |
| :------: | ----------- | ------------------------------------- | ------------ |
|    1     | field       | field id/name/label/selector          |              |
|    2     | value       | string/[variables](../variables.html) |              |

## Examples

```gherkin
Given I am on the "home" page
Then the "HomeCenterMessage" element should contain "We have listed most of the components that are used by developers to build web applications."
```