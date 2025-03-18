#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int i;
int *ptr;
int inputc;
int inputs[100];

void inc() { ptr[i] += 1; }

void dec() { ptr[i] -= 1; }

void movl() { i -= 1; }

void movr() { i += 1; }

void set(int v) { ptr[i] = v; }

int buildInputs(int argc, char *argv[]) {
  for (int i = 1; i < argc; i++) {
    char *str = argv[i];
    char *start = str;
    char *end;

    while (*start != '\0') {
      long num = strtol(start, &end, 10);
      if (start == end) {
        if (*start != '\0') {
          start++;
        }
      } else {
        if (inputc >= 100) {
            fprintf(stderr, "Too many numbers. Maximum 100.\n");
            return 1;
        }
        inputs[inputc++] = (int)num;
        start = end;
      }
    }
  }
  return 0;
}

int main(int argc, char *argv[]) {
  ptr = (int *)calloc(sizeof(int), 16384);

  int returnCode = buildInputs(argc, argv);
  if (returnCode != 0) {
    return returnCode;
  }

  for (int ii=0; ii < inputc; ii++) {
      printf("%d ", inputs[ii]);
  }
  printf("\n");
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
