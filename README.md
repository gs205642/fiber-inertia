# Inertia.js Fiber Middleware

## Installation

Install the package using the `go get` command:

```go
go get github.com/gs205642/fiber-inertia
```

## Usage

### 1. Create Template

resources/views/app.html

```html
<!DOCTYPE html>
<html>
<head>
    <!-- Scripts -->
    {{.reactRefresh}}
    {{.scripts}}
</head>
<body>
    {{.inertia}}
</body>
</html>
```

### 2. Create Router

```go
import (
    inertia "github.com/gs205642/fiber-inertia"
)

app.Use(inertia.New())

app.Static("/assets", "public/build/assets")

app.Get("/", func(c *fiber.Ctx) error {
    return inertia.Render(c, "Welcome", fiber.Map{
        "Foo": "bar",
    })
})
```

### 3. Install Client Side

Use [official documentation](https://inertiajs.com/client-side-setup) to install client side.
