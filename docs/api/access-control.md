---
slug: 'access-control'
---

## Overview

Access Control in PMM can be used to restrict access to individual metrics.  
Access Control is in tech preview and needs to be enabled manually from the PMM settings before it can be used.

Once enabled, restricting access to metrics can be performed by:

1. Creating a Percona role
2. Assigning the role to a user

### Create a Percona role

```bash
curl -X POST "http://localhost/v1/management/Role/Create" \ 
     -H "Authorization: Basic xxx" \
     -H "Content-Type: application/json" \ 
     -d '{
        "title": "My custom role role",
        "filter": "{environment=\"staging\"}"
      }'
```

The `filter` parameter is a [PromQL query](https://prometheus.io/docs/prometheus/latest/querying/basics/) restricting access to the specified metrics.  
Full access can be provided by specifying an empty `filter` field.

### Assign a Percona role

Users can be assigned roles by using the `/v1/management/Role/Assign` API.  
The endpoint assigns new roles to a user. Other roles, that may have been assigned to the user previously, stay intact.

```bash
curl -X POST "http://localhost/v1/management/Role/Assign" \ 
     -H "Authorization: Basic xxx" \
     -H "Content-Type: application/json" \ 
     -d '{
        "user_id": 1,
        "role_ids": [2, 3, 4]
      }'
```
