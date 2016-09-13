package main

/*
#include <stdio.h>
#include <string.h>

void sayHi() {
    printf("Hii");
}

void writeFile() {
	FILE *pFile = fopen("mycgo.txt","w");
	char greeting[]="hellocgo ,I love you thank you !";
	fwrite(greeting,1,strlen(greeting),pFile);
	fflush(pFile);
}
*/
import "C"

func main() {
	C.sayHi()
	C.writeFile()
}
