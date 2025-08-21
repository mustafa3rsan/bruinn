# Fluxx

This document describes how to set up and use the Fluxx integration.

## Connection Parameters

### Required Parameters

- `name`: Connection name
- `instance`: Instance URL for OAuth connection
- `client_id`: OAuth Client ID
- `client_secret`: OAuth Client Secret

## URI Format

```
fluxx://{instance}?{parameters}
```

## Example Configuration

```yaml
connections:
  fluxx:
    - name: "fluxx-prod"
      instance: "your-instance"
      client_id: "your-client-id"
      client_secret: "your-client-secret"
```

## Supported Tables

- All tables and endpoints supported by the Fluxx API
- Custom queries and filters available

For more information, refer to the Fluxx API documentation.
