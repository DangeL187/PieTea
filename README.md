# About
PieTea is a command-line tool that makes it easier to use the [httpie](https://github.com/httpie/cli) by allowing you to specify requests in YAML files.

It also displays the responses from HTTP requests in a pretty-formatted and colorized way in the terminal window.

# Features
- Define HTTP requests using simple YAML files.
- Pretty-print JSON response bodies with syntax highlighting.
- Display HTTP response headers and bodies in a clean, bordered terminal UI.
- Modular design for easy extension and customization.

# Installation
Make sure you have Go installed and [httpie](https://github.com/httpie/cli) available on your system.

```bash
git clone https://github.com/DangeL187/PieTea.git
cd PieTea
go build -o ptea ./cmd/main.go
```

# Usage
Create a YAML file describing your HTTP request, for example:

```yaml
Method: POST
URL: localhost:8000/users/register

Headers:
  name: value

Body:
  email: test@example.com
  full_name: John Doe
  password: secret

QueryParams:
  name: value
```

Your YAML request file should define the following fields:
- Method: HTTP method (e.g., GET, POST, PUT).
- URL: Request URL.
- Headers: (Optional) Key-value pairs of HTTP headers.
- Body: (Optional) Key-value pairs for the request JSON body.
- QueryParams: (Optional) Key-value pairs for URL query parameters.

Run PieTea with your YAML file:
```bash
./ptea request.yaml
```

Example Output:
```
┌──────────────────────────────┐
│HTTP/1.1 200 OK               │
│Content-Type: application/json│
│                              │
│{                             │
│  "id": 123,                  │
│  "name": "PieTea"            │
│}                             │
└──────────────────────────────┘
```

# Contributing
Contributions are welcome!

Feel free to open issues or submit pull requests for improvements, bug fixes, or new features.
