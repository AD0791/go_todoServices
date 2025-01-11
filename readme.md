# Get serious with go web programming

----

Get solid at Backend Engineering

Launch hmr server

```bash
air
```

update swagger docs

```bash
swag init
```

## Notes

best way to add swagger for fiber:

- [fiber wrapper for swagger](https://github.com/gofiber/swagger)

follow this documentation: `swag init` and `swag`

- [the declarative comment formats](https://github.com/swaggo/swag#declarative-comments-format)

## notions


In this project, we will implement JWT


### basic authentication

```mermaid
sequenceDiagram
    autonumber
    participant U as User
    participant F as Fiber (App)
    participant DB as Database

    Note over U: Enter username + password
    U->>F: POST /login (username, password)
    F->>DB: SELECT user by username
    alt user not found
        F->>U: 401 Unauthorized (User not found)
    else user found
        F->>F: Compare hashed password with provided password
        alt password mismatch
            F->>U: 401 Unauthorized (Invalid credentials)
        else match
            F->>F: Create session record (in store/memory)
            F->>U: 200 OK + Session Cookie
        end
    end

    Note over F: On subsequent requests<br>Use session ID to look up user

```


### JWT

```mermaid
sequenceDiagram
    autonumber
    participant U as User
    participant F as Fiber (App)
    participant DB as Database
    participant S as JWT Signer

    Note over U: Enter username + password
    U->>F: POST /login (username, password)
    F->>DB: Validate user credentials
    alt invalid credentials
        F->>U: 401 Unauthorized
    else valid credentials
        F->>S: Generate JWT (user claims, expiry)
        S->>F: Return signed JWT
        F->>U: 200 OK + JWT in response
    end

    Note over U: Store JWT (e.g., localStorage)

    Note over U: On subsequent calls
    U->>F: GET /protected <br> Authorization: Bearer JWT
    F->>S: Verify JWT signature + check expiry
    alt valid
        F->>U: 200 OK (Protected Resource)
    else invalid/expired
        F->>U: 401 Unauthorized
    end
```


### Oauth2

```mermaid
sequenceDiagram
    autonumber
    participant U as User
    participant C as Client (Your App)
    participant AS as Authorization Server
    participant RS as Resource Server (User Data)

    Note over U: Wants to login/authorize via external provider
    U->>C: Click "Login with X"
    C->>AS: GET /authorize?client_id=&redirect_uri=&state=xyz
    AS->>U: Prompt login/consent
    U->>AS: Login + consent to share resources
    AS->>C: Redirect (302) with auth code + state=xyz
    C->>AS: POST /token (auth code, client_secret, etc.)
    AS->>C: 200 OK {access_token, refresh_token, ...}

    Note over C: Access token is stored (server or client side)

    C->>RS: GET /user (Authorization: Bearer access_token)
    RS->>C: 200 OK { user info }
```



### MFA/OTP

```mermaid
sequenceDiagram
    autonumber
    participant U as User
    participant F as Fiber (App)
    participant DB as Database
    participant T as TOTP Library

    Note over U: Already has an account + password
    U->>F: Request MFA Setup
    F->>T: Generate TOTP secret + provisioning URI
    T->>F: Return secret + QR code data
    F->>DB: Store TOTP secret for user
    F->>U: Show QR code (scanned by Authenticator app)

    Note over U: Next login with password + TOTP code
    U->>F: POST /mfa/verify (token)
    F->>DB: Retrieve user's TOTP secret
    F->>T: Validate code with stored secret
    alt valid code
        F->>U: 200 OK (MFA success)
    else invalid code
        F->>U: 401 Unauthorized
    end

```


### ERD

```mermaid
erDiagram
    USER {
        bigint ID PK
        string username
        string password_hash
        string totp_secret
        bool enabled_2fa
    }

    SESSION {
        bigint ID PK
        bigint user_id FK
        string session_data
        datetime created_at
        datetime expires_at
    }

    JWT_TOKEN {
        bigint ID PK
        bigint user_id FK
        string token
        datetime expires_at
    }

    OAUTH_TOKEN {
        bigint ID PK
        bigint user_id FK
        string provider
        string access_token
        string refresh_token
        datetime expires_at
    }

    USER ||--|{ SESSION : "has many"
    USER ||--|{ JWT_TOKEN : "has many"
    USER ||--|{ OAUTH_TOKEN : "can have many"

```

## more notions


Look at these structures

### HTTP Request

```mermaid
flowchart TB
    A[HTTP Request] --> B(Method)
    A --> C(URI/Path)
    A --> D(Protocol Version)
    A --> E(Query Parameters)
    A --> F(Headers)
    A --> G(Body)

    B --> B1[Example: GET, POST, PUT...]
    C --> C1[Example: /api/v1/users]
    D --> D1[Example: HTTP/1.1]
    F --> F1[Content-Type, Accept, Host, etc]
    G --> G1[Optional: JSON, Form Data, etc]

```

### HTTP response

```mermaid
flowchart TB
    A[HTTP Response] --> B(Status Line)
    A --> C(Protocol Version)
    A --> D(Status Code)
    A --> E(Status Message)
    A --> F(Headers)
    A --> G(Body)

    C --> C1[Example: HTTP/1.1]
    D --> D1[Example: 200, 404, 500]
    E --> E1[Example: OK, Not Found, Internal Server Error]
    F --> F1[Content-Type, Content-Length, etc]
    G --> G1[JSON, HTML, etc]

```

### Session

```mermaid
flowchart TB
    A((Session)) --> B[Session ID]
    A --> C[User ID or user data reference]
    A --> D[Created At]
    A --> E[Expires At]
    A --> F[Session Data key-value pairs]

    B --> B1[Random or hashed token]
    C --> C1[Link to user record in DB]
    F --> F1[CSRF token, preferences, etc.]
```

### JWT

```mermaid
flowchart TB
    A[JWT Token] --> B[Header JSON]
    A --> C[Payload JSON]
    A --> D[Signature]

    B --> B1[alg: HS256, RS256, etc]
    B --> B2[typ: JWT]
    C --> C1[sub, iss, exp, iat, custom claims...]
    D --> D1[HMAC or RSA signature over Header + Payload]
```

### TOTP

```mermaid
flowchart TB
    A((TOTP Setup)) --> B[User ID]
    A --> C[TOTP Secret]
    A --> D[Issuer App name]
    A --> E[Digits e.g., 6]
    A --> F[Period e.g., 30s]

    C --> C1[Shared secret stored securely]
    F --> F1[Time window for each code]

```