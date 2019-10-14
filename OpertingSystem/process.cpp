#include <stdio.h>
#include <cstdlib>
#include <unistd.h>
#include <wait.h>
using namespace std; 
int main()
{
    int x, fd[2];
    char buf[30], s[30];
    pipe(fd);
    while ((x = fork()) == -1);
    if (x == 0)
    {
        printf("1");
        printf("3");
        sprintf(buf, "This is an example\n");
        write(fd[1], buf, 30);
        exit(0);
    }
    else
    {
        printf("2");
        wait(0);
        read(fd[0], s, 30);
        printf("%s", s);
    }
    return 0;
}
