#include <string.h>
#include "cgo.h"

char *strncpy2(char *dest, const char *src, size_t n)
{
    strncpy(dest, src, n);
    return dest;
}

int *intcpy(int *dest, int src)
{
    *dest = src;
    return dest;
}
