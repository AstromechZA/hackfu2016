//
// This file was generated by the Retargetable Decompiler
// Website: https://retdec.com
// Copyright (c) 2016 Retargetable Decompiler <info@retdec.com>
//

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

// ------------------- Function Prototypes --------------------

int32_t ___libc_start_main(int32_t a1);
int32_t __do_global_ctors_aux(int32_t a1);
void __do_global_dtors_aux(void);
int32_t __libc_csu_fini(int32_t a1);
int32_t __libc_csu_init(int32_t a1, int32_t a2, int32_t a3, int32_t a4, int32_t a5);
void __ZNSt8ios_base4InitD1Ev__GLIBCXX_3_4(void);
void __ZSt4endlIcSt11char_traitsIcEERSt13basic_ostreamIT_T0_ES6___GLIBCXX_3_4(void);
int32_t _fini(int32_t a1);
int32_t _ftext(void);
int32_t _GLOBAL__sub_I__Z7decryptSsSs(void);
int32_t _init(void);
void _PROCEDURE_LINKAGE_TABLE_(void);
int32_t _Z41__static_initialization_and_destruction_0ii(int32_t a1, int32_t a2, int32_t a3, int32_t a4);
int32_t _Z7decryptSsSs(int32_t a1, int32_t a2, int32_t a3, int32_t a4, int32_t a5, int32_t a6, int32_t a7);
int32_t frame_dummy(int32_t a1);
int32_t function_400b78(void);
int32_t function_400b88(int32_t a1);
int32_t function_400d1c(void);
int32_t function_400e3c(int32_t a1);
int32_t function_401290(int32_t a1, int32_t a2, int32_t a3);
int32_t function_40143c(int32_t a1);
int32_t function_4014f8(int32_t a1);
int32_t function_401554(int32_t a1);
int32_t function_4015a0(int32_t a1, int32_t a2);
int32_t unknown_ffffffff(void);

// --------------------- Global Variables ---------------------

int32_t g1 = 0; // $a0
int32_t g2 = 0; // $a1
int32_t g3 = 0; // $a2
bool g4 = false; // $a3
int32_t g5 = 0; // $at
int32_t g6 = 0; // $fp
int32_t g7 = 0; // $gp
int32_t g8 = 0; // $ra
int32_t g9 = 0; // $s0
int32_t g10 = 0; // $t9
int32_t g11 = 0; // $v0
int32_t g12 = -1; // 0x411780
int32_t g13 = 0; // 0x411790
int32_t g14 = 0; // 0x41184c
char g15 = 0; // 0x4118e0
int32_t g16 = 0; // 0x4118f0

// ------------------------ Functions -------------------------

// Address range: 0x400b38 - 0x400b77
int32_t _init(void) {
    int32_t v1 = g10; // 0x400b40
    g7 = v1 + 0x18cd8;
    int32_t * v2 = (int32_t *)(v1 + 0x10d0c); // 0x400b50_0
    int32_t v3 = *v2; // 0x400b50
    g11 = v3;
    if (v3 != 0) {
        int32_t v4 = *v2; // 0x400b60
        g10 = v4;
        g8 = 0x400b70;
        ((int32_t (*)())v4)();
        // branch -> 0x400b70
    }
    // 0x400b70
    function_400b78();
    int32_t v5 = g11; // 0x400b74_1
    return v5;
}

// Address range: 0x400b78 - 0x400b87
int32_t function_400b78(void) {
    int32_t v1 = g1; // 0x400b78
    int32_t v2 = frame_dummy(v1); // 0x400b78
    g5 = v2;
    int32_t v3 = g1; // 0x400b80
    function_400b88(v3);
    int32_t v4 = g5; // 0x400b84_1
    return v4;
}

// Address range: 0x400b88 - 0x400b9f
int32_t function_400b88(int32_t a1) {
    // 0x400b88
    __do_global_ctors_aux(a1);
    int32_t v1;
    g8 = v1;
    return 0;
}

// Address range: 0x400ba0 - 0x400c2f
void _PROCEDURE_LINKAGE_TABLE_(void) {
    // 0x400ba0
    return;
}

// Address range: 0x400c30 - 0x400ccf
void __ZNSt8ios_base4InitD1Ev__GLIBCXX_3_4(void) {
    // 0x400c30
    return;
}

// Address range: 0x400cd0 - 0x400d0f
void __ZSt4endlIcSt11char_traitsIcEERSt13basic_ostreamIT_T0_ES6___GLIBCXX_3_4(void) {
    // 0x400cd0
    return;
}

// Address range: 0x400d10 - 0x400d1b
int32_t _ftext(void) {
    // 0x400d10
    function_400d1c();
    int32_t v1 = g5; // 0x400d18_1
    return v1;
}

// Address range: 0x400d1c - 0x400d6f
int32_t function_400d1c(void) {
    // 0x400d1c
    __libc_start_main();
    // branch -> 0x400d68
    while (true) {
        // 0x400d68
        // branch -> 0x400d68
    }
}

// Address range: 0x400d70 - 0x400e0b
void __do_global_dtors_aux(void) {
    char v1 = g15; // 0x400d7c
    if (v1 == 0) {
        // 0x400d90
        g15 = 1;
        // branch -> 0x400df0
    }
}

// Address range: 0x400e0c - 0x400e3b
int32_t frame_dummy(int32_t a1) {
    // 0x400e0c
    g1 = 0x410000;
    int32_t v1 = g13; // 0x400e10
    if (v1 != 0) {
        // 0x400e18
        // branch -> 0x400e34
    }
    int32_t v2 = g5; // 0x400e34
    return v2;
}

// Address range: 0x400e3c - 0x400e3f
int32_t function_400e3c(int32_t a1) {
    int32_t v1 = g5; // 0x400e3c_1
    return v1;
}

// Address range: 0x400e40 - 0x400fc3
// Demangled:     decrypt(std::string, std::string)
int32_t _Z7decryptSsSs(int32_t a1, int32_t a2, int32_t a3, int32_t a4, int32_t a5, int32_t a6, int32_t a7) {

    // interesting! char arrays are good
    // 0x400e40
    char * v1;
    char * v2 = v1; // bp-24

    // then it calls this function 5 times..
    // is that maybe reading a char from args?
    _ZNSsixEj();
    // g11 is a pointer to a char*
    int32_t v3 = g11; // 0x400e70
    // deref
    char v4 = *(char *)v3; // 0x400e70
    // convert to int
    int32_t v5 = (int32_t)v4; // $v0

    _ZNSsixEj();
    int32_t v6 = v5; // 0x400e88
    char v7 = *(char *)v6; // 0x400e88
    v5 = (int32_t)v7;

    _ZNSsixEj();
    int32_t v8 = v5; // 0x400ea0
    char v9 = *(char *)v8; // 0x400ea0
    v5 = (int32_t)v9;

    _ZNSsixEj();
    int32_t v10 = v5; // 0x400eb8
    char v11 = *(char *)v10; // 0x400eb8
    v5 = (int32_t)v11;

    _ZNSsC1ERKSs();
    v2 = NULL;
    int32_t v12 = 0; // $s0
    _ZNKSs4sizeEv();
    uint32_t v13 = v12; // 0x400f6c1
    uint32_t v14 = v5; // 0x400f6c2
    if (v13 < v14) {
        int32_t v15 = (int32_t)&v2; // 0x400f34_0
        char * v16 = NULL; // 0x400edc_0
        // branch -> 0x400edc
        while (true) {
            // 0x400edc
            v5 = (int32_t)v16;
            _ZNSsixEj();
            int32_t v17 = v5; // 0x400ef0
            v12 = v17;
            char * v18 = v2; // 0x400ef4_0
            v5 = (int32_t)v18;
            _ZNSsixEj();
            int32_t v19 = v5; // 0x400f08
            char v20 = *(char *)v19; // 0x400f08
            char * v21 = v2; // 0x400f0c_0
            int32_t v22 = (int32_t)v21; // 0x400f0c
            int32_t v23 = v22 & -0x7ffffffd; // 0x400f18
            int32_t v24 = v23; // 0x400f38
            if (v23 <= 0xffffffff) {
                int32_t v25 = (v22 + 3 | -4) + 1; // 0x400f30
                v24 = v25;
                // branch -> 0x400f34
            }
            char v26 = *(char *)(v24 + v15 + 4); // 0x400f3c
            int32_t v27 = v12; // 0x400f4c
            *(char *)v27 = v26 ^ v20;
            char * v28 = v2; // 0x400f50_0
            int32_t v29 = (int32_t)v28 + 1; // 0x400f54
            v5 = v29;
            v2 = (char *)v29;
            v12 = v29;
            _ZNKSs4sizeEv();
            uint32_t v30 = v12; // 0x400f6c
            uint32_t v31 = v5; // 0x400f6c
            if (v30 < v31) {
                char * v32 = v2;
                v16 = v32;
                // branch -> 0x400edc
                continue;
            }
        }
    }
    // 0x400fa4
    return a1;
}

// Address range: 0x400fc4 - 0x40128f
int main(int argc, char ** argv) {
    // set up some vars, all equal to 0?
    // 0x400fc4
    int32_t v1;
    int32_t v2 = v1; // bp-352
    int32_t v3;
    int32_t v4 = v3; // bp-328
    int32_t v5;
    int32_t v6 = v5; // bp-324
    int32_t v7;
    int32_t v8 = v7; // bp-320
    int32_t v9;
    int32_t v10 = v9; // bp-304

    // get a pointer to v2
    g6 = (int32_t)&v2;
    int32_t v11; // 0x401270

    // if there are at least 2 arguments
    // arg0 is binary name, arg1 is arg
    if (argc >= 2) {
        int32_t v12 = *(int32_t *)((int32_t)argv + 4); // 0x40101c
        g9 = v12;
        _ZNSaIcEC1Ev(); // allocate char?
        _ZNSsC1EPKcRKSaIcE();
        _ZNSaIcED1Ev();
        if (argc < 3) {
            // 0x401068
            _ZNKSs4sizeEv();
            // branch -> 0x4010b4
        }
        // 0x4010b4
        _ZNSt14basic_ifstreamIcSt11char_traitsIcEEC1EPKcSt13_Ios_Openmode();
        _ZNSsC1Ev();
        int32_t v13 = (int32_t)&v10; // 0x40110c_0
        _ZNSsC1Ev();
        int32_t v14 = (int32_t)&v6; // 0x401124_0
        int32_t v15 = (int32_t)&v4; // 0x401140_0
        int32_t v16 = (int32_t)&v8; // 0x401154_0
        // branch -> 0x4011e8
        while (true) {
            int32_t v17 = v13; // $v0
            _ZSt7getlineIcSt11char_traitsIcESaIcEERSt13basic_istreamIT_T0_ES7_RSbIS4_S5_T1_E();
            int32_t v18 = v10; // 0x401200
            int32_t v19 = *(int32_t *)(v18 - 12); // 0x401208
            int32_t v20 = v17; // 0x40120c
            v17 = v20 + v19;
            _ZNKSt9basic_iosIcSt11char_traitsIcEEcvPvEv();
            int32_t v21 = v17; // 0x40121c
            if (v21 == 0) {
                int32_t v22 = 0; // $s0
                _ZNSsD1Ev();
                _ZNSsD1Ev();
                _ZNSt14basic_ifstreamIcSt11char_traitsIcEED1Ev();
                _ZNSsD1Ev();
                int32_t v23 = v22;
                v11 = v23;
                // branch -> 0x401270
                // 0x401270
                return v11;
            }
            // 0x401124
            _ZNSsC1ERKSs();
            _ZNSsC1ERKSs();
            g11 = v15;
            bool v24 = g4;
            int32_t v25 = v24 ? (int32_t)__libc_csu_init : 0; // 0x40116c
            _Z7decryptSsSs(v16, v14, v15, v25, 0, 0, 0);
            _ZNSsaSERKSs();
            _ZNSsD1Ev();
            _ZNSsD1Ev();
            _ZNSsD1Ev();
            _ZStlsIcSt11char_traitsIcESaIcEERSt13basic_ostreamIT_T0_ES7_RKSbIS4_S5_T1_E();
            _ZNSolsEPFRSoS_E();
            // branch -> 0x4011e8
        }
    } else {
        // 0x400fec
        _ZStlsISt11char_traitsIcEERSt13basic_ostreamIcT_ES5_PKc();
        v11 = 0;
        // branch -> 0x401270
    }
    // 0x401270
    return v11;
}

// Address range: 0x401290 - 0x40137f
int32_t function_401290(int32_t a1, int32_t a2, int32_t a3) {
    int32_t v1 = a1; // $s0
    _ZNSaIcED1Ev();
    int32_t v2 = v1; // 0x4012a4
    int32_t v3 = v2; // $a0
    _Unwind_Resume();
    int32_t v4 = v3; // 0x4012b4
    v1 = v4;
    _ZNSsD1Ev();
    _ZNSsD1Ev();
    _ZNSsD1Ev();
    _ZNSsD1Ev();
    _ZNSsD1Ev();
    _ZNSt14basic_ifstreamIcSt11char_traitsIcEED1Ev();
    _ZNSsD1Ev();
    int32_t v5 = v1; // 0x401370
    int32_t v6 = v5; // $v0
    _Unwind_Resume();
    int32_t v7 = v6; // 0x40137c_1
    return v7;
}

// Address range: 0x401380 - 0x401403
// Demangled:     __static_initialization_and_destruction_0(int, int)
int32_t _Z41__static_initialization_and_destruction_0ii(int32_t a1, int32_t a2, int32_t a3, int32_t a4) {
    // 0x401380
    if (a1 == 1) {
        // 0x4013a4
        if (a2 == 0xffff) {
            // 0x4013b4
            g1 = 0x4118f0;
            _ZNSt8ios_base4InitC1Ev();
            __cxa_atexit(__ZNSt8ios_base4InitD1Ev__GLIBCXX_3_4, (char *)&g16, (char *)&g14);
            // branch -> 0x4013ec
        }
    }
    int32_t v1 = g5; // 0x4013fc
    return v1;
}

// Address range: 0x401404 - 0x40143b
int32_t _GLOBAL__sub_I__Z7decryptSsSs(void) {
    // 0x401404
    int32_t v1;
    int32_t v2 = v1; // bp-32
    g6 = (int32_t)&v2;
    int32_t v3 = g3; // 0x40141c
    bool v4 = g4;
    int32_t v5 = v4 ? (int32_t)__libc_csu_init : 0; // 0x40141c
    int32_t v6 = _Z41__static_initialization_and_destruction_0ii(1, 0xffff, v3, v5); // 0x40141c
    return v6;
}

// Address range: 0x40143c - 0x40143f
int32_t function_40143c(int32_t a1) {
    int32_t v1 = g5; // 0x40143c_1
    return v1;
}

// Address range: 0x401440 - 0x401447
int32_t __libc_csu_fini(int32_t a1) {
    int32_t v1 = g5; // 0x401440
    return v1;
}

// Address range: 0x401448 - 0x4014f7
int32_t __libc_csu_init(int32_t a1, int32_t a2, int32_t a3, int32_t a4, int32_t a5) {
    int32_t v1 = 0; // $s5
    int32_t v2 = 0; // $s4
    int32_t v3 = 0; // $s3
    int32_t v4 = 0; // $s2
    int32_t v5 = 0; // $s1
    int32_t v6 = g10; // 0x401450
    g10 = _init;
    v3 = a1;
    v2 = a2;
    v1 = a3;
    int32_t v7 = _init(); // 0x401484
    g7 = v6 + 0x183c8;
    int32_t v8 = *(int32_t *)(v6 + 0x103f0); // 0x401494
    v5 = v8;
    int32_t v9 = *(int32_t *)(v6 + 0x103f4); // 0x401498
    int32_t v10 = v9 - v8; // 0x4014a0
    v4 = v10 / 4;
    int32_t v11; // 0x4014f0_11
    if (v10 < 4) {
        v11 = v7;
        // 0x4014d4
        return v11;
    }
    int32_t v12 = 0; // 0x4014b4
    int32_t v13 = v8; // 0x4014b0
    while (true) {
        int32_t v14 = *(int32_t *)v13; // 0x4014b0
        g10 = v14;
        g9 = v12 + 1;
        int32_t v15 = v3; // 0x4014b8
        g1 = v15;
        int32_t v16 = v2; // 0x4014bc
        g2 = v16;
        g8 = 0x4014c8;
        int32_t v17 = v1; // 0x4014c0
        g3 = v17;
        ((int32_t (*)())v14)();
        int32_t v18 = g9; // 0x4014c8
        uint32_t v19 = v4; // 0x4014c8
        int32_t v20 = v5; // 0x4014cc
        int32_t v21 = v20 + 4; // 0x4014cc
        v5 = v21;
        if (v18 >= v19) {
            v11 = 0;
            // break -> 0x4014d4
            break;
        }
        v12 = v18;
        v13 = v21;
        // continue -> 0x4014b0
    }
    // 0x4014d4
    return v11;
}

// Address range: 0x4014f8 - 0x4014ff
int32_t function_4014f8(int32_t a1) {
    int32_t v1 = g5; // 0x4014fc_1
    return v1;
}

// Address range: 0x401500 - 0x401553
int32_t __do_global_ctors_aux(int32_t a1) {
    int32_t v1 = g12; // 0x401504
    g11 = -1;
    if (v1 == -1) {
        // 0x401540
        return 0;
    }
    int32_t v2 = -1; // $s1
    int32_t v3 = 0x411780; // 0x40152c
    int32_t v4 = v3 - 4; // $s0
    unknown_ffffffff();
    int32_t v5 = v4; // 0x401534
    int32_t v6 = *(int32_t *)v5; // 0x401534
    // branch -> 0x40152c
    while (v6 != v2) {
        int32_t v7 = v2; // 0x401538
        v3 = v5;
        v4 = v3 - 4;
        unknown_ffffffff();
        v5 = v4;
        v6 = *(int32_t *)v5;
        // continue -> 0x40152c
    }
    // 0x401540
    return 0;
}

// Address range: 0x401554 - 0x40155f
int32_t function_401554(int32_t a1) {
    int32_t v1 = g5; // 0x40155c_1
    return v1;
}

// Address range: 0x401560 - 0x40157f
int32_t ___libc_start_main(int32_t a1) {
    int32_t v1 = g7; // 0x401560
    int32_t v2 = *(int32_t *)(v1 - 0x7ff0); // 0x401560
    g10 = v2;
    int32_t v3 = g8; // 0x401564
    int32_t v4 = v3; // $t7
    g8 = 0x401570;
    ((int32_t (*)())v2)();
    int32_t v5 = v4; // 0x40157c_1
    return v5;
}

// Address range: 0x401580 - 0x40159f
int32_t _fini(int32_t a1) {
    int32_t v1 = g2; // 0x401598
    function_4015a0(a1, v1);
    int32_t v2 = g5; // 0x40159c_1
    return v2;
}

// Address range: 0x4015a0 - 0x4015b7
int32_t function_4015a0(int32_t a1, int32_t a2) {
    // 0x4015a0
    __do_global_dtors_aux();
    int32_t v1 = g5; // 0x4015ac
    return v1;
}

// --------------- Dynamically Linked Functions ---------------

// int __cxa_atexit(void(*func)(void *), void * arg, void * dso_handle);
// void __libc_start_main(void);
// void _Unwind_Resume(void);
// void _ZNKSs4sizeEv(void);
// void _ZNKSt9basic_iosIcSt11char_traitsIcEEcvPvEv(void);
// void _ZNSaIcEC1Ev(void);
// void _ZNSaIcED1Ev(void);
// void _ZNSolsEPFRSoS_E(void);
// void _ZNSsaSERKSs(void);
// void _ZNSsC1EPKcRKSaIcE(void);
// void _ZNSsC1ERKSs(void);
// void _ZNSsC1Ev(void);
// void _ZNSsD1Ev(void);
// void _ZNSsixEj(void);
// void _ZNSt14basic_ifstreamIcSt11char_traitsIcEEC1EPKcSt13_Ios_Openmode(void);
// void _ZNSt14basic_ifstreamIcSt11char_traitsIcEED1Ev(void);
// void _ZNSt8ios_base4InitC1Ev(void);
// void _ZSt7getlineIcSt11char_traitsIcESaIcEERSt13basic_istreamIT_T0_ES7_RSbIS4_S5_T1_E(void);
// void _ZStlsIcSt11char_traitsIcESaIcEERSt13basic_ostreamIT_T0_ES7_RKSbIS4_S5_T1_E(void);
// void _ZStlsISt11char_traitsIcEERSt13basic_ostreamIcT_ES5_PKc(void);

// --------------------- Meta-Information ---------------------

// Detected compiler/packer: gcc (4.6.3)
// Detected language: C++
// Detected functions: 25
// Decompiler release: v2.1.2 (2016-01-27)
// Decompilation date: 2016-04-24 15:52:46