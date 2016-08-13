package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static void myPrint(const char* msg) {
	printf("myPrint: %s", msg);
}

void say() {
	printf("i say hi\n");
}

void shuxue() {
	int i, sum = 0;
	for(i = 1;i <=100;i++)
		sum = sum + i;
	printf("%d",sum);
}

void filew() {
	FILE* fstream;
	char msg[100] = "Hello!I have read this file.";
	char message[] = "\nfile content";
	fstream=fopen("test.txt","a+");
	if(fstream == NULL)
	{
		printf("open file test.txt failed!\n");
		exit(1);
	}
	else
	{
		printf("open file test.txt succeed!\n");
		fwrite(msg,strlen(message),1,fstream);
	}
	fclose(fstream);
}

*/
import "C"

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* cff() {
	char* str = "HelloWorld";
	return str;
}

//int* jk() {
//	int* aee = 30;
//	return aee;
//}

*/
import "C"
import "fmt"
import "unsafe"

func main() {
	//C.puts(C.CString("Hello, 世界\n"))
	//C.printf(C.CString("Hello, 世界\n"))
	// error
	C.puts(C.CString("puts hello \n"))
	C.myPrint(C.CString("Hello, world for C Language \n"))
	C.say()
	C.shuxue()
	C.filew()

	//	var cb string
	//	cb = *string(C.cf())
	fmt.Print(C.GoString(C.cff()))

	a := (*[10]byte)(unsafe.Pointer(C.cff()))
	fmt.Print(string(a[0]) + string(a[1]) + string(a[2]) + string(a[3]) + string(a[4]))
}
