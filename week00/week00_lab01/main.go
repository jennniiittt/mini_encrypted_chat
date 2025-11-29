package main

import (
  "fmt"
  "math/big"
)

func main() {


// CTF BabyRSA
// Given n = 43941819371451617899582143885098799360907134939870946637129466519309346255747
// Exponent(e) = 65537
// Ciphertext(c) = 9002431156311360251224219512084136121048022631163334079215596223698721862766


  pStr := "205237461320000835821812139013267110933"
  qStr := "214102333408513040694153189550512987959"
  eStr := "65537"
  cStr := "9002431156311360251224219512084136121048022631163334079215596223698721862766"

  p := new(big.Int)
  q := new(big.Int)
  e := new(big.Int)
  c := new(big.Int)

  p.SetString(pStr, 10)
  q.SetString(qStr, 10)
  e.SetString(eStr, 10)
  c.SetString(cStr, 10)

  // ----------------------------------------------------
  // 2. Compute n = p * q
  // ----------------------------------------------------
  n := new(big.Int).Mul(p, q)

  // ----------------------------------------------------
  // 3. Compute phi(n) = (p-1) * (q-1)
  // ----------------------------------------------------
  pMinus1 := new(big.Int).Sub(p, big.NewInt(1))
  qMinus1 := new(big.Int).Sub(q, big.NewInt(1))
  phi := new(big.Int).Mul(pMinus1, qMinus1)

  // ----------------------------------------------------
  // 4. Compute d = e^{-1} mod phi(n)
  // ----------------------------------------------------
  d := new(big.Int).ModInverse(e, phi)
  if d == nil {
    panic("No modular inverse for e mod phi(n); invalid parameters")
  }

  // ----------------------------------------------------
  // 5. Decrypt: m = c^d mod n
  // ----------------------------------------------------
  m := new(big.Int).Exp(c, d, n)

  fmt.Println("p      =", p)
  fmt.Println("q      =", q)
  fmt.Println("n      =", n)
  fmt.Println("phi(n) =", phi)
  fmt.Println("e      =", e)
  fmt.Println("d      =", d)
  fmt.Println("\nCiphertext c =", c)
  fmt.Println("Plaintext m (as big int) =", m)

  // OPTIONAL: convert m to bytes (if it encodes text)
  plaintextBytes := m.Bytes()
  fmt.Println("Plaintext as bytes:", plaintextBytes)
  fmt.Println("Plaintext as string:", string(plaintextBytes))
}
