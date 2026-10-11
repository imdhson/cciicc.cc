## 2024-10-02 - Server crash DoS due to unchecked array slice access and XSS prevention via html.EscapeString
**Vulnerability:** A critical Denial of Service (DoS) existed in `service/AddChatFrom_space_id.go` where an iterator defaulted to index 0 if a space was missing, leading to an index out of bounds panic if no spaces existed. Also, XSS vulnerabilities existed across `controller/PostHandler_chat.go`, `controller/PostHandler_create_space.go`, and `controller/PostHandler_join_space.go` because user input (chats, usernames, space names) was not sanitized.
**Learning:** The previous implementation failed to handle edge cases properly when iterating over the global singleton struct slices. Go panics on out-of-bounds slice indexing which, if unrecovered, can crash the entire HTTP server in a DoS attack. Although the frontend offers basic XSS protection (`textContent`), server-side validation/sanitization (`html.EscapeString`) is essential for Defense in Depth.
**Prevention:** Check for loop exhaustion explicitly. Initialize loop result indexes to `-1` and verify they are not `-1` before using them as slice indexes. Always wrap external strings originating from HTML forms with `html.EscapeString()` directly as they are parsed from the `http.Request`.

## 2024-10-05 - Weak Session Key Generation
**Vulnerability:** The session key generation used a weak random source (16-bit integer) and an insecure hashing algorithm (MD5).
**Learning:** The entropy of the generated session keys was extremely low, making them predictable and susceptible to brute-force attacks, leading to potential session hijacking.
**Prevention:** Always use cryptographically secure random number generators (e.g., crypto/rand) for sensitive data like session keys, generating at least 16 random bytes and avoiding insecure hashing algorithms like MD5.
## 2023-10-10 - Security Headers Missing
**Vulnerability:** The application was missing basic security headers (e.g., X-Content-Type-Options, X-Frame-Options, X-XSS-Protection) in HTTP and HTTPS responses.
**Learning:** These headers provide defense-in-depth against attacks like MIME-type sniffing, Clickjacking, and Cross-Site Scripting. Since Go's `http.ServeMux` doesn't include these out of the box, they need to be added manually.
**Prevention:** Always implement a security header middleware to wrap the main router/handler before passing it to the server.
