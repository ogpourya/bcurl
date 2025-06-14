# bcurl

`bcurl` is a lightweight wrapper around `curl` that automatically adds browser-like headers to your HTTP requests. It simplifies making requests that mimic a browser's behavior while allowing full customization of headers and other `curl` options.

## Features

- Pre-configured with Firefox-like headers.
- Allows user-defined headers to override default ones.
- Supports all standard `curl` options.

## Installation

```bash
go install github.com/pzaeemfar/bcurl@latest
```

## Usage

Use `bcurl` just like `curl`. For example:
```bash
bcurl https://example.com
```

You can add custom headers or options:
```bash
bcurl -H "Authorization: Bearer <token>" https://example.com
```