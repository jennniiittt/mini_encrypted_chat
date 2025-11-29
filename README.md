# mini_encrypted_chat
A small encrypted chat app running on **localhost**, built with:

* ECDH key exchange (Curve25519)
* AES-GCM encryption
* TCP sockets

## ** How it works (in simple words)**

1. Start the **server**
2. Start the **client**
3. They exchange public keys
4. Both calculate the same shared secret (using ECDH)
5. SHA-256 turns that into a 32-byte AES key
6. Every message is encrypted with AES-GCM
7. Chat securely🧩

## ** How to run**
    ### 1. Start server

    ```
    go run .
    Mode: server
    ```

    ### 2. Start client (open another terminal)

    ```
    go run .
    Mode: client
    ```

    Once connected, you can type messages on both ends.

    Type **exit** to quit (on either side).

## **📁 Files**

main.go      # lets you choose server/client mode
server.go    # server logic
client.go    # client logic
crypto.go    # ECDH + AES-GCM functions

## ** Monitoring the traffic (Wireshark)**

If you want to see the encrypted packets, open Wireshark and capture on `localhost`:

Filter: *tcp.port == 9000*

You’ll see the messages going through, but the contents will be unreadable because everything is encrypted.

## ** What’s happening under the hood**

* Curve25519 ECDH → creates shared secret
* SHA-256 → turns it into AES key
* AES-GCM → encrypt + authenticate each message
* Length-prefix → helps read exact message size

Nothing fancy, just enough to understand secure communication zzZZZ.


## ** Why this project exists**

To show how encryption works in real-time chat without libraries doing everything for you.
Good practice for:

* Go networking
* Cryptography basics
* Secure key exchange
* AES-GCM usage

---

If you want, I can also generate:
✨ a diagram for the key exchange
✨ a simple flowchart
✨ color-coded README version

Just tell me!
