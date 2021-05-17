#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>

int main()
{
    char stringTemplate[] = "Message from process %d\n";
    int fd[2];
    pipe(fd);
    pid_t p1 = fork();
    if (p1 == 0) {
        char string[23];
        sprintf(string, stringTemplate, 1);
        lockf(fd[1], 1, 0);
        write(fd[1], string, strlen(string) + 1);
        lockf(fd[1], 0, 0);
        return 0;
    }
    pid_t p2 = fork();
    if (p2 == 0) {
        char string[23];
        sprintf(string, stringTemplate, 2);
        lockf(fd[1], 1, 0);
        write(fd[1], string, strlen(string) + 1);
        lockf(fd[1], 0, 0);
        return 0;
    }
    pid_t p3 = fork();
    if (p3 == 0) {
        char string[23];
        sprintf(string, stringTemplate, 3);
        lockf(fd[1], 1, 0);
        write(fd[1], string, strlen(string) + 3);
        lockf(fd[1], 0, 0);
        return 0;
    }
    int status = 0;
    pid_t wpid;
    int count = 0;
    do {
    } while ((wpid = waitpid(-1, NULL, 0)) > 0);
    char readBuffer[24];
    for (int i = 0; i < 3; i++) {
        read(fd[0], readBuffer, sizeof(readBuffer));
        printf("%s", readBuffer);
    }
}