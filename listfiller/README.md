# ListFiller
A function that each invocation adds elements into a vector aiming to consume memory on purpose.

## Details
- **Number of elements**: [2^21 (2097152).](https://github.com/dfquaresma/openfaas-functions/blob/main-readme/listfiller/listfiller/src/main/java/com/openfaas/function/Handler.java#L53)
- **Memory usage**: about 8MB for each call - it depends on the number of elements and its size. 

    Estimation: (elements * data_size) / (1024 * 1024) = (2097152 * 4) / (1024 * 1024) = 8MB

### The standard Java integer data types are:
- byte 1 byte  -128 to 127
- short 2 bytes  -32,768 to 32,767
- int 4 bytes -2,147,483,648 to 2,147,483,647
- long 8 bytes -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
