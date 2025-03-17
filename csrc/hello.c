#include <stdio.h>
#include <stdlib.h>

int i;
int *ptr;

void inc() { ptr[i] += 1; }

void dec() { ptr[i] -= 1; }

void movl() { i -= 1; }

void movr() { i += 1; }

void set(int v) { ptr[i] = v; }

int main(void) {
  ptr = (int *)calloc(sizeof(int), 16384);
  inc();
  movr(); movr();
  inc(); inc();
  movr(); movr();
  inc(); inc(); inc();
  for (int ii = 0; ii < 6; ii++)
    printf("%d ", ptr[ii]);

  free(ptr);
  return 0;
}
