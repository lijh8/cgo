#ifndef cgo_h
#define cgo_h

#include <stdio.h>
#include <stddef.h>

#define PRINT(fmt, ...) do { \
    printf("%s:%d:%s: " fmt "\n", \
        __FILE__, __LINE__, __func__, ##__VA_ARGS__); \
} while (0)

char *strncpy2(char *dest, const char *src, size_t n);

int *intcpy(int *dest, int src);

#endif
