---
page_title: "Common issues/FAQ"
subcategory: "Guides"
description: |-
  Common issues and frequently asked questions when using the provider.
---

# Mist Provider Common Issues/FAQ

## 429 Too Many Requests

```
Status: 429 Too Many Requests
Retry-After: 798
```

Mist is enforcing a Rate Limit on the API Requests.The current rate limiting is 5000 API calls per hour and is reset at the hourly boundary.


This threshold is applied:
- to the API Token for the Organization API Tokens
- to the User for the User API Tokens or Authenticated requests
