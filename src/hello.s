//
// hello world
//

.global _main
.align 2 // necessary for arm assembly

_main:
  b _printf
  b _term

_printf:
  mov X0, #1
  adr X1, helloworld
  mov X2, #12
  mov X16, #4
  svc 0

_term:
  mov X0, #0
  mov X16, #1
  svc 0

helloworld: .ascii "hello world\n"
