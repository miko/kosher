---
layout: default
title: "I press/click the \"__\" button/link"
parent: Step Definitions
---

# I press/click the "\_\_" button/link

Presses/clicks first element matched by parameter 1.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )(?:press|click) the "([^"]*)" (?:button|link)$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | element     | field id/name/label/selector |              |

## Examples

```gherkin
Given I click the "Contact Form" link
And press the "Submit" button
```