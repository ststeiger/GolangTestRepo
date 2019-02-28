
#include <stdio.h>
#include <stdlib.h>

#include "goc.h"


// gcc hellow.c goc.a -o hellow -lpthread
int main(int argc, char* argv[])
{
    printf("Hello world!\n");
    PrintInt(123);
    return EXIT_SUCCESS;
}
