# Salesforce

This document describes how to set up and use the Salesforce integration.

## Connection Parameters

### Required Parameters

- `name`: Connection name
- `instance`: Instance URL for OAuth connection
- `client_id`: OAuth Client ID
- `client_secret`: OAuth Client Secret

## URI Format

```
salesforce://{instance}?{parameters}
```

## Example Configuration

```yaml
connections:
  salesforce:
    - name: "salesforce-prod"
      instance: "your-instance"
      client_id: "your-client-id"
      client_secret: "your-client-secret"
```

## Supported Tables

- All tables and endpoints supported by the Salesforce API
- Custom queries and filters available

For more information, refer to the Salesforce API documentation.
