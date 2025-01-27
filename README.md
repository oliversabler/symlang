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
-: minus
+: plus
÷: divide
×: multiply

!: not
=: equal
≠: not equal
>: greater
≥: greater equal
<: less
≤: less equal

&: and
←: assign
Ɵ: break
○: false
ƒ: function
¿: if
∞: loop
ø: nil
|: or
✉: print
↵: return
●: true
•: variable
