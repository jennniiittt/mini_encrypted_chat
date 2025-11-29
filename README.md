# mini_encrypted_chat
A small encrypted chat app running on **localhost**, built with Go.

It demonstrates how to create a secure communication channel using ECDH key exchange, AES-GCM encryption, and TCP sockets.

---

## ** Overview**

This project shows how two programs (server & client) can:
* Exchange public keys (ECDH)
* Derive the same shared secret
* Convert that secret into an AES key
* Encrypt every message using AES-GCM
* Chat securely over TCP

You can also watch the encrypted traffic using Wireshark.

---

## ** How it works (in simple words)**

1. Start the **server**
2. Start the **client**
3. They exchange public keys
4. Both calculate the same shared secret (using ECDH)
5. SHA-256 turns that into a 32-byte AES key
6. Every message is encrypted with AES-GCM
7. Chat securely🧩

---

## ** Installation / Setup**

1. Install Go

   ```
   https://go.dev/dl/
   ```

2. Clone the project folder from Git

3. Open the project directory:

   ```
   cd mini_encrypted_chat
   ```
4. Create each file, work on them.
5. Run the whole project:

   ```
   go run .
   ```

---

## ** Usage**

### **Start the server**

```
go run .
Mode: server
```

Server will listen on `localhost:9000`.

### **Start the client (in another terminal)**

```
go run .
Mode: client
```

Now both sides can chat securely.

### **Exit**

Type:

```
exit
```

to close the connection.

---

## ** What’s happening under the hood**

* Curve25519 ECDH → creates shared secret
* SHA-256 → turns it into AES key
* AES-GCM → encrypt + authenticate each message
* Length-prefix → helps read exact message size

Nothing fancy, just enough to understand secure communication zzZZZ.

---

## ** Monitor Encrypted Traffic (Optional)**

If you want to see encrypted packets in Wireshark:

Filter:

```
tcp.port == 9000
```

You’ll see the packets, but the contents will be unreadable because of AES-GCM.

---

## ** Dependencies / Libraries Used**

All dependencies are standard Go modules + one crypto library:

* `net` — TCP sockets
* `crypto/aes` — AES block cipher
* `crypto/cipher` — GCM mode
* `crypto/rand` — secure randomness
* `crypto/sha256` — key derivation
* `golang.org/x/crypto/curve25519` — ECDH (Curve25519 key exchange)

Install the curve library (if missing):

```
go get golang.org/x/crypto/curve25519
```
---

## ** Files**

main.go      # lets you choose server/client mode
server.go    # server logic
client.go    # client logic
crypto.go    # ECDH + AES-GCM functions

---

## ** Security Practices in This Project**

This project follows basic encryption best practices:

### **1. ECDH (Curve25519) Key Exchange**

* Both client and server generate random 32-byte private keys
* Public keys exchanged over TCP
* Shared secret is generated securely

### **2. AES-GCM Encryption**

* AES-256 key (`32 bytes`) derived from SHA-256(shared_secret)
* Each message uses a fresh random nonce
* GCM ensures:

  * Confidentiality (encrypted)
  * Integrity (tamper-protected)
  * Authentication (no forgery)

### **3. Secure Randomness**

* All private keys and nonces come from `crypto/rand`

### **4. Length-Prefixed Messages**

* Prevents partial reads and misaligned packets

---

### **Note**

This is a learning project — not production-grade.
Real-world systems require:

* Mutual authentication
* Identity verification
* Replay protection
* Secure key rotation
* TLS or Noise Protocol
  …but this project could be a starting point.

======== See you on Presentation Day លោកគ្រូ🙏🏼🙏🏼 ========