# ListFiller
A function that each invocation adds elements into a vector aiming to consume memory on purpose.

## Details
- **Number of elements**: [2^21 (2097152).](https://github.com/dfquaresma/openfaas-functions/blob/main-readme/listfiller/listfiller/src/main/java/com/openfaas/function/Handler.java#L53)
- **Memory usage**: about 16MB for each call - it depends on the number of elements and its size. 

    Estimation: (elements * size) / (1024 * 1024) = (2097152 * 8) / (1024 * 1024) = 16MB
