# gRPC error

:star: Star us on GitHub — it motivates us a lot!

gRPC error library for Go/Golang.

If you are a server developer, you should generate errors with enough information to help client developers understand and resolve the problem.

## Table of Content

- [Error Payloads](#error-payloads)
- [Usage/Examples](#usageexamples)
    - [NewNotFoundError](#newnotfounderror)
- [License](#license)
- [Links](#links)

## Error Payloads

The `google.rpc` package defines a set of standard error payloads, which are preferred to custom error payloads. The
following table lists each error code and its matching standard error payload, if applicable. We recommend advanced
applications look for these error payloads in `google.rpc.Status` when they handle errors.

| HTTP  | gRPC                  | Error detail                     |
|:------|:----------------------|:---------------------------------|
| `400` | `INVALID_ARGUMENT`    | `google.rpc.BadRequest`          |
| `400` | `FAILED_PRECONDITION` | `google.rpc.PreconditionFailure` |
| `400` | `OUT_OF_RANGE`        | `google.rpc.BadRequest`          |
| `401` | `UNAUTHENTICATED`     | `google.rpc.ErrorInfo`           |
| `403` | `PERMISSION_DENIED`   | `google.rpc.ErrorInfo`           |
| `404` | `NOT_FOUND`           | `google.rpc.ResourceInfo`        |
| `409` | `ABORTED`             | `google.rpc.ErrorInfo`           |
| `409` | `ALREADY_EXISTS`      | `google.rpc.ResourceInfo`        |
| `429` | `RESOURCE_EXHAUSTED`  | `google.rpc.QuotaFailure`        |
| `499` | `CANCELLED`           |                                  |
| `500` | `DATA_LOSS`           | `google.rpc.DebugInfo`           |
| `500` | `UNKNOWN`             | `google.rpc.DebugInfo`           |
| `500` | `INTERNAL`            | `google.rpc.DebugInfo`           |
| `501` | `NOT_IMPLEMENTED`     |                                  |
| `503` | `UNAVAILABLE`         | `google.rpc.DebugInfo`           |
| `504` | `DEADLINE_EXCEEDED`   | `google.rpc.DebugInfo`           |

## Usage/Examples

#### NewNotFoundError

```go
package main

import (
  "google.golang.org/genproto/googleapis/rpc/errdetails"

  "github.com/LipsarHQ/go-grpc-error"
)

func main() {
  // ...
  if err != nil {
    // Return error in case of resource not found.
    return nil, grpcerror.NewNotFoundError(&errdetails.ResourceInfo{
      ResourceType: "", // string(new(examplepb.Example).ProtoReflect().Descriptor().FullName())
      ResourceName: "", // ID.
      Owner:        "", // ID.
      Description:  err.Error(),
    })
  }
}

```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Links

* [Google Cloud API design guide](https://cloud.google.com/apis/design)
  * [Error Payloads](https://cloud.google.com/apis/design/errors#error_payloads)
