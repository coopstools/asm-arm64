#include <stdio.h>
#include <stdlib.h>

int i = 0;
int *ptr;

void inc() { ptr[i] += 1; }

void dec() { ptr[i] -= 1; }

void movl() { i -= 1; }

void movr() { i += 1; }

void set(int v) { ptr[i] = v; }

int main(void) {
  i = 0
  ptr = (int *)calloc(sizeof(int), 16384);
  movr();
  movr();
  inc();
  inc();
  movl();
  inc();
  movr();
  movr();
  inc();
  movl();
  dec();
  movl();
  movl();
  movr();
  movr();
  movr();
  movr();

  for (int ii = 0; ii <= 6; ii++)
    printf("%d ", ptr[ii]);
  free(ptr);
  return 0;
}
