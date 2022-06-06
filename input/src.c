/* input file */
// #include <stdio.h>

int main() {
    printf("Hello World!\n");
    int a = 10, b = 20;
    int _sum = a + b;

    if (_sum > b)
        printf("_sum 大于 b");

    char s[20] = "Hello Lexer!\n";

    float PI = 3.1415926535;

    signed cnt = 3;
    for (int i = 0; i < cnt; i++)
        printf("%d\n", i);

    while (cnt--)
        printf("%d\n", cnt);

    return 0;
}