# try-connect

## Learnings

1. Running `buf lint` shows it has strict requirements on
   - Folder structure (where to put the protos)
   - Version number (e.g. "v1" and "v1alpha1" are ok but not "v0")
   - Service proto must be suffixed with "Service"

2. Adding server/client is simple enough. The interface seems well defined.
   Deployed the server to Cloud Run (needs to set port protocol to be `h2c`
   since we used `h2c` server)

3. Using the Connect client is probably the easier. Using gRPC client is the
   same, which requires dialing the connection first. Using HTTP is probably the
   most troublesome one because 1) there is no strong type 2) need to use http2.

4. Connect [interceptors](https://connectrpc.com/docs/go/interceptors) are not
   with the same gRPC interceptor interfaces. So any existing gRPC interceptors
   can't be reused without modification.

5. Connect [routes](https://connectrpc.com/docs/go/routing#constructing-routes)
   need POST as the method. This isn't "classic" REST style.

6. Connect [errors](https://connectrpc.com/docs/go/common-errors) have its own
   error codes which are not gRPC native codes. If fully bought into the
   framework, it might actually work better because we no longer need to worry
   about HTTP status code vs. gRPC status code ourselves (Connect takes care of
   that).

7. Buf has a pretty rich [plugin selections](https://buf.build/plugins) to
   generate code. The majority of the base are from existing protobuf codegens.
   But being able to aggregate these plugins is something the vanilla protobuf
   community failed to do well. Some plugins are actually interesting, e.g.
   https://buf.build/community/google-gnostic-openapi

## TODOs

- Explore streaming
- Actually implement an interceptor
- Try Open Telemetry
