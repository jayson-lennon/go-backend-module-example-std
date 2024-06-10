```
.
├── go.mod
├── handler         // all handlers
│   ├── root.go     // handlers for /
│   └── user        // user handlers in own package
│       └── user.go
├── main.go         // entry point
├── router          // all routes
│   ├── router.go   // main router
│   └── user.go     // user routes
├── server
│   └── server.go   // runs the server
└── util
    └── util.go     // accessible to everything
```
