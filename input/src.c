/* #include <stdio.h> */

int main() {
    printf("Hello World!\n");
    int a = 10, b = 20; // 无符号整数的识别
    int _sum = a + b; // _符号打头的标识符的识别、算术运算符的识别

    if (_sum > b) // 保留字的识别
        printf("_sum 大于 b");

    char s[20] = "Hello Lexer!\n";

    float PI = 3.1415926535;

    signed cnt = 3;
    for (int i = 0; i < cnt; i++)
        printf("%d\n", i);

    while (cnt--)
        printf("%d\n", cnt);

    return 0; // 保留字的识别
}