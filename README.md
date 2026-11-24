# Sum of Two Integers

> NeetCode 150 · Bit Manipulation · Blind 75 — XOR is the sum without carries, AND<<1 is the carries; loop until the carries run out. Go, stdlib only, tested.

## Problem

Add two integers without using + or −. The point is understanding what an adder circuit does.

## Approach

Per bit, addition splits into two independent pieces: the sum-bit ignoring carry (XOR) and the carry-bit (AND, shifted left one position to where it applies). `a+b = (a^b) + ((a&b)<<1)` — still a +, so iterate: fold the carry back in the same way until it's zero. Each round pushes the lowest unresolved carry at least one bit higher, so fixed-width integers terminate in ≤ word-size rounds. Negatives need no cases — two's complement makes the same bit mechanics produce correctly signed results, which the random cross-check against native + confirms over thousands of pairs.

**Complexity:** O(word size) ≈ O(1); O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
