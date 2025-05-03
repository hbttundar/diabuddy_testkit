# Diabuddy Testkit

The **Diabuddy Testkit** is a reusable Go testing toolkit designed to support full-stack, HTTP-first, and database-backed testing across all DiaBuddy microservices.

---

## 🧩 Modules

### ✅ Suite
- `BaseSuite`, `IntegrationSuite`, `BrowserSuite`
- Provides `testing.T`, context, transaction, router, and HTTP client

### ✅ DB Helpers
- `db/connection.go`: creates DB connections from `.env`
- `db/migrater.go`: runs and rolls back migrations

### ✅ HTTP Testing
- `http/client.go`: `Send`, `Post`, `WithJSONBody`, `WithBearerToken`, etc.
- `http/response.go`: `MustStatus`, `AssertPaginationHeaders`, `AssertSortedBy`, etc.

### ✅ Faker
- `faker/string.go`: `RandomString`, `RandomEmail`
- `faker/number.go`: `RandomInt`, `RandomFloat`
- `faker/time.go`: `RandomPastTime`, `RandomTimeRange`

### ✅ Factory
- `Factory[T]` interface with `Make`, `Create`, `MakeMany`, `CreateMany`
- Utilities to bulk-generate test data with `GenerateMany`

---

## 🔍 Usage

### Create a test suite:
```go
s := suite.NewBrowserSuite(t, routerSetupFn)
defer s.Cleanup()
```

### Send a JSON POST request:
```go
resp := http.Post(t, "/users", http.WithJSONBody(map[string]any{
  "email": "test@local", "role": "admin",
}))
http.MustStatus(t, resp, 201)
```

### Generate fake data:
```go
email := faker.RandomEmail()
date := faker.RandomPastTime(72 * time.Hour)
```

### Use factories:
```go
user := userfactory.Create(ctx, tx, map[string]any{"role": "admin"})
```

---

## 📌 Best Practices
- Keep test-only logic in testkit
- Use `diabuddy-api-infra` for runtime infra only
- Avoid domain-specific factories in testkit — keep them in each service

---

## 🛠️ Coming Soon (optional)
- `AssertSortedByNumeric`
- JSON Schema validators
- Kafka test topic support

---
