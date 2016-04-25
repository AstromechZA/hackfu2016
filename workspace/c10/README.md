file:

```
decrypt: ELF 32-bit LSB  executable, MIPS, MIPS-II version 1, dynamically linked (uses shared libs), for GNU/Linux 2.6.26, BuildID[sha1]=a7f844334054ed2326d95a7dad803a6d671a9893, not stripped
```

strings:

```
...
Pass the 16 character key
e.g. ./decrypt example123456789
messages.enc
...
```

Need to run it on MIPS?

I have no idea how to read whatever I can dissassemble...

https://retdec.com/
https://www.onlinedisassembler.com/odaweb/
http://shadow-file.blogspot.co.za/2013/05/running-debian-mips-linux-in-qemu.html
http://www.mrc.uidaho.edu/mrc/people/jff/digital/MIPSir.html
http://darkdust.net/files/GDB%20Cheat%20Sheet.pdf

Learning a lot!

looks like this is just an XOR cipher..

wrote a simple simulator

an example key that I chose to generate lowercase letters
$ ./simulate b8F7d0A0a0E0h0A0
payyriexwtzarqfttzctztdcfxzyvtnxxree

