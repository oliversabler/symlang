# Symlang
This is a tiny programming language called The Symbol Language (Symlang for short).

Disclaimer: Programming in this language will be tedious, and slow to type in (a lot of copy paste).

I present to you recursive implementation of the Fibonacci sequence in Symlang below.
```
ƒ fib(n) {
    ¿ (n ≤ 1) ↵ n;
    ↵ fib(n - 2) + fib(n - 1);
}

• result ← 0;
• i ← 0;
∞ {
    ¿ (i ≥ 20) Ɵ;

    result ← fib(i);
    ✉ result;

    i ← i + 1;
}
```


If you want to give it a spin, try to run a file in the examples directory, for example
`go run main.go examples/lab.sym`.

## Symbols
```
-     // 2 - 1; (subtract)
+     // 2 + 3; (add)
÷     // 9 ÷ 3; (divide)
×     // 3 × 3; (multiply)

!     // !●;            (not)
=     // "abc" = "abc"; (equal)
≠     // "abc" ≠ "cba"; (not equal)
>     // 3 > 1;         (greater)
≥     // 3 ≥ 3;         (greater equal)
<     // 1 < 3;         (less)
≤     // 3 ≤ 3;         (less equal)

&     // ¿ (1 + 1 = 2 & ●) { ... }                 (and)
←     // existingVariable ← "new value";           (assign)
Ɵ     // ∞ { Ɵ; }                                  (break)
○     // false
ƒ     // ƒ functionName() { ... }                  (function)
¿     // ¿ (1 + 1 = 2) { ... }                     (if)
∞     // ∞ { ... }                                 (loop)
ø     // nil
|     // ¿ (1 + 1 = 2 | ●) { ... }                 (or)
✉     // ✉ "print me";                             (print)
↵     // ƒ functionName() { ↵ "string value"; }    (return)
●     // true
•     // • myVariable;                             (variable)
```
