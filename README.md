# Symlang
This is a tiny programming language called The Symbol Language (Symlang for short).

Disclaimer: Programming in this language will be tedious, and slow.

This is obviously a work in progress but to give you a taste of what to expect in the future,
I present to you recursive implementation of the Fibonacci sequence in Symlang below.
```
ƒ fib(n) {
    ¿ (n ≤ 1) ↵ n;
    ↵ fib(n - 2) + fib(n - 1);
}

• result ← 0;
• i ← 0;
∞ {
    ¿ (i > 20) Ɵ;

    result ← fib(i);
    ✉ result;

    i ← i + 1;
}
```
