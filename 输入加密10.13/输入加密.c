#include <stdio.h>
#include <stdlib.h>
#include <conio.h>


int main()
{   int i=0;
  printf("请输入密码：");
char b[6];

while(i<6)
{
char a;

   a= getch();
   if(a==13){printf("\n"); break;}
   
printf("*");
if(i==5)
{
  printf("\n");
}

b[i]=a;
i++;
}   
printf("你输入的密码为：");
for(int c=0;c<6;c++)  
 {printf("%c",b[c]);}
 
 
    return 0;
}