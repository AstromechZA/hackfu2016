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




1st test:

```0-6
{-20.908693667561487 s-K&s+ atther          eonemo          ight}
{-20.91996849332944  s-K&s4 atthem          eonemp          ight}
{-20.936705345018755 s-K&* atthis          eonean          ight}
{-20.938158262467894 s-K&s- atthet          eonemi          ight}
{-20.95711943025619  s-K&s= atthed          eonemy          ight}
{-20.97292417847424  s-K&7 atthin          eoneas          ight}
{-20.97419585957021  s-K&w7 atthan          eoneis          ight}
{-20.986267508286957 s-K&w* atthas          eonein          ight}
{-20.997552626566335 s-K&s1 attheh          eonemu          ight}
{-21.011619189768144 s-K&s) atthep          eonemm          ight}
{-21.04273327255421  s-K&s7 atthen          eonems          ight}
{-21.049219102095144 s-K&s! atthex          eoneme          ight}
{-21.05095367807731  s-K&s  atthey          eonemd          ight}
{-21.057237873895893 s-K&w- atthat          eoneii          ight}
{-21.05797478419424  s-K&s0 atthei          eonemt          ight}
{-21.059484449303877 s-K&s* atthes          eonemn          ight}
{-21.061522263287806 s-K&s6 attheo          eonemr          ight}
{-21.07357742748411  s-K&6 atthio          eonear          ight}
{-21.079223711768492 s-K&w0 atthai          eoneit          ight}
{-21.08554684957394  s-K&4 atthim          eoneap          ight}
{-21.10873380403602  w+V q* erings          aiscon          maur}
{-21.113167892247983 s-V q* atings          eoscon          igur}
{-21.115666446524298 s-K&s, attheu          eonemh          ight}
{-21.11599447168216  s-K&w) atthap          eoneim          ight}
{-21.12609400299471  s-K&y+ atthor          eonego          ight}
{-21.126882137118738 s-K&w/ atthav          eoneik          ight}
{-21.13070188021117  w+V q6 eringo          aiscor          maur}
{-21.13513596842313  s-V q6 atingo          eoscor          igur}
{-21.13909622967662  s-K&s/ atthev          eonemk          ight}
{-21.1392701095737   s-K&- atthit          eoneai          ight}
{-21.14090863606735  s-K&s3 atthej          eonemw          ight}
{-21.142163439084513 s-K&w# atthaz          eoneig          ight}
{-21.14429724140745  w+V q) eringp          aiscom          maur}
{-21.145071496711136 s-K&s. atthew          eonemj          ight}
{-21.147384224281772 s-K&s> attheg          eonemz          ight}
{-21.14873132961941  s-V q) atingp          eoscom          igur}
{-21.150053790055892 w+V q1 eringh          aiscou          maur}
{-21.154487878267854 s-V q1 atingh          eoscou          igur}
{-21.15475252720856  w+K&s+ erther          ainemo          maht}
{-21.155420743864273 t-W+{6 fthemo          borher          ngty}
{-21.157973275750162 w+Q+w7 ernean          aithis          mary}
{-21.15807175664758  w+V q0 eringi          aiscot          maur}
{-21.161858652753438 s-K&s2 atthek          eonemv          ight}
{-21.162505844859542 s-V q0 atingi          eoscot          igur}
{-21.163036329836544 s-K&0 atthii          eoneat          ight}
{-21.166027352976513 w+K&s4 erthem          ainemp          maht}
{-21.166896956764617 s+V q* arings          eiscon          iaur}
{-21.16766030214956  t-W+w7 fthean          borhis          ngty}
{-21.167938308065626 s-K&w  atthay          eoneid          ight}
{-21.168246625712136 s-K&w4 attham          eoneip          ight}
{-21.169197707721487 s-K&s( attheq          eoneml          ight}
{-21.169662210749692 s-K&+ atthir          eoneao          ight}
{-21.171998576692673 s+Q+w7 arnean          eithis          iary}
{-21.176995420226046 w+K&* erthis          ainean          maht}
{-21.180867713673578 s-K&w+ atthar          eoneio          ight}
{-21.183895775657586 s-K&s< atthee          eonemx          ight}
{-21.18421712211497  w+K&s- erthet          ainemi          maht}
{-21.18670939947041  t-W+7 fthein          borhas          ngty}
{-21.18886503293976  s+V q6 aringo          eiscor          iaur}
{-21.18978441080747  s-K&# atthiz          eoneag          ight}
{-21.190490567757635 s-K&w( atthaq          eoneil          ight}
{-21.190699102836902 ,Q+{6 munemo          inther          efry}
{-21.190806034516843 s-K&s5 atthel          eonemq          ight}
{-21.192445973160385 {-K&s+ itther          monemo          aght}
{-21.194192109205648 w+Q+w* erneas          aithin          mary}
{-21.195862673209874 w.V q* ewings          alscon          mdur}
{-21.20058949975622  w+Q+* erneis          aithan          mary}
{-21.201906194494867 w+V q- eringt          aiscoi          maur}
{-21.201973928604858 ,Q+{) munemp          inthem          efry}
{-21.20246039413604  s+V q) aringp          eiscom          iaur}
{-21.203178289903263 w+K&s= erthed          ainemy          maht}
{-21.203720798928334 {-K&s4 itthem          monemp          aght}
{-21.206340282706826 s-V q- atingt          eoscoi          igur}
{-21.20722664849028  t-W+* ftheis          borhan          ngty}
{-21.208216942784485 s+V q1 aringh          eiscou          iaur}
{-21.20821741014816  s+Q+w* arneas          eithin          iary}
{-21.210242074324366 p+V q* brings          fiscon          jaur}
{-21.212042894949583 s-K&w, atthau          eoneih          ight}
```

```
&{-21.43389516792415 P s Z'       thedev          usthen        }
&{-21.43389516792415 Q;b,Z?       usthen          thedev        }
&{-21.433908710192032 P s Z0       thedea          usthey        }
&{-21.433908710192032 Q;b,Z(       usthey          thedea        }
&{-21.441595875755013 V;b,Z?       rsthen          shedev        }
&{-21.441595875755013 W s Z'       shedev          rsthen        }
&{-21.4416094180229 W s Z0       shedea          rsthey        }
&{-21.441609418022903 V;b,Z(       rsthey          shedea        }
&{-21.44221793814366 P s ^=       thedal          usthat        }
&{-21.44221793814366 Q;b,^%       usthat          thedal        }
&{-21.449918645974527 W s ^=       shedal          rsthat        }
&{-21.44991864597453 V;b,^%       rsthat          shedal        }
&{-21.453253884396254 P s Z=       thedel          usthet        }
&{-21.453253884396254 Q;b,Z%       usthet          thedel        }
&{-21.460954592227118 V;b,Z%       rsthet          shedel        }
&{-21.460954592227118 W s Z=       shedel          rsthet        }
&{-21.46168032231871 P s Z!       thedep          ustheh        }
&{-21.46168032231871 Q;b,Z9       ustheh          thedep        }
&{-21.4621961390012 P'b,^%       tothat          utedal        }
&{-21.4621961390012 Q<s ^=       utedal          tothat        }
&{-21.46268199381392 F'b,^%       bothat          ctedal        }
&{-21.46268199381392 G<s ^=       ctedal          bothat        }
&{-21.465800658256182 P s Z%       thedet          usthel        }
&{-21.465800658256182 Q;b,Z=       usthel          thedet        }
&{-21.469381030149574 W s Z!       shedep          rstheh        }
&{-21.469381030149574 V;b,Z9       rstheh          shedep        }
&{-21.469644388573826 P'b,Z?       tothen          utedev        }
&{-21.469644388573826 Q<s Z'       utedev          tothen        }
&{-21.469657930841713 Q<s Z0       utedea          tothey        }
&{-21.469657930841716 P'b,Z(       tothey          utedea        }
&{-21.47013024338655 G<s Z'       ctedev          bothen        }
&{-21.47013024338655 F'b,Z?       bothen          ctedev        }
&{-21.470143785654432 G<s Z0       ctedea          bothey        }
&{-21.470143785654436 F'b,Z(       bothey          ctedea        }
&{-21.472772285256795 P s Z<       thedem          ustheu        }
&{-21.472772285256795 Q;b,Z$       ustheu          thedem        }
&{-21.473501366087046 W s Z%       shedet          rsthel        }
&{-21.473501366087046 V;b,Z=       rsthel          shedet        }
&{-21.477965702054107 P b,^%       ththat          usedal        }
&{-21.477965702054107 Q;s ^=       usedal          ththat        }
&{-21.48047299308766 W s Z<       shedem          rstheu        }
&{-21.480472993087663 V;b,Z$       rstheu          shedem        }
&{-21.48389337674382 Q;b!M%       ustert          theirl        }
&{-21.483893376743822 P s-M=       theirl          ustert        }
&{-21.484931006293028 P s Z"       thedes          usthek        }
&{-21.484931006293028 Q;b,Z:       usthek          thedes        }
&{-21.48635292216575 A:b,^%       erthat          diedal        }
&{-21.48635292216575 @!s ^=       diedal          erthat        }
&{-21.486755070489426 P s Z?       theden          usthev        }
&{-21.486755070489426 Q;b,Z'       usthev          theden        }

```

```
&{-23.539497571172465 j-B-             plei            time    }
&{-23.539497571172465 n(J!             time            plei    }
&{-23.56645336878846 n4C!             tude            ppli    }
&{-23.56645336878846 j1K-             ppli            tude    }
&{-23.571025845977726 (I'             einc            alfo    }
&{-23.571025845977726 {-A+             alfo            einc    }
&{-23.571830674471297 {%C!             adde            eali    }
&{-23.5718306744713  K-             eali            adde    }
&{-23.57457941469794  K!             eale            addi    }
&{-23.57457941469794 {%C-             addi            eale    }
&{-23.575929784155452 (C!             eide            alli    }
&{-23.575929784155456 {-K-             alli            eide    }
&{-23.57828548296643 /F&             enab            akin    }
&{-23.578285482966432 {*N*             akin            enab    }
&{-23.58186938371228 s-C-             ildi            mile    }
&{-23.58186938371228 w(K!             mile            ildi    }
&{-23.585167744429484 -K!             elle            aidi    }
&{-23.585167744429484 {(C-             aidi            elle    }
&{-23.58936164596346 {(I!             aine            elfi    }
&{-23.589361645963464 -A-             elfi            aine    }
&{-23.58988851116658 7N+             evio            asac    }
&{-23.58988851116658 {2F'             asac            evio    }
&{-23.5910966772861 {7N*             avin            esab    }
&{-23.591096677286103 2F&             esab            avin    }
&{-23.591454074156236 {(C!             aide            elli    }
&{-23.59145407415624 -K-             elli            aide    }
&{-23.592795096093994 s-K,             illh            midd    }
&{-23.592795096093997 w(C              midd            illh    }
&{-23.595762274530493 2J%             esma            avem    }
&{-23.595762274530497 {7B)             avem            esma    }
&{-23.596813055944445 6N*             ewin            arab    }
&{-23.596813055944445 {3F&             arab            ewin    }
&{-23.59807041745404 3N!             erie            awai    }
&{-23.59807041745404 {6F-             awai            erie    }
&{-23.600024054293154 n(J%             tima            plem    }
&{-23.600024054293154 j-B)             plem            tima    }
&{-23.600318722159837 v-C-             lldi            hile    }
&{-23.60031872215984 r(K!             hile            lldi    }
&{-23.600483577791415 {%K-             adli            eade    }
&{-23.60048357779142  C!             eade            adli    }
&{-23.600770346126485  J%             eama            adem    }
&{-23.600770346126485 {%B)             adem            eama    }
&{-23.60082449856918 v-B              lled            himh    }
&{-23.60082449856918 r(J,             himh            lled    }
&{-23.601189617234024 2F-             esai            avie    }
&{-23.601189617234024 {7N!             avie            esai    }
&{-23.60203768526417 3J-             ermi            awee    }
&{-23.60203768526417 {6B!             awee            ermi    }
&{-23.60559193883078 -K%             ella            aidm    }
&{-23.60559193883078 {(C)             aidm            ella    }
```

```
8 - 14
-21.5195 '        edanag          thaveb      ' 's ^?{&'
-21.5250 '        edanad          thavea      ' 's ^?{%'
-21.5299 '        thathl          edalli      ' 'b,^%r-'
-21.5299 '        edalli          thathl      ' 's ^=v('
-21.5341 '        reathl          cialli      ' 'd!^%r-'
-21.5341 '        cialli          reathl      ' 'u-^=v('
-21.5346 '        thenaw          edever      ' 'b,Z?{6'
-21.5470 '        edalld          thatha      ' 's ^=v%'
-21.5470 '        thatha          edalld      ' 'b,^%r '
-21.5502 '        cialaw          reater      ' 'u-^={6'
-21.5512 '        reatha          cialld      ' 'd!^%r '
-21.5512 '        cialld          reatha      ' 'u-^=v%'
-21.5598 '        cialar          reatew      ' 'u-^={3'
-21.5626 '        thathi          edalll      ' 'b,^%r('
-21.5626 '        edalll          thathi      ' 's ^=v-'
-21.5668 '        reathi          cialll      ' 'd!^%r('
-21.5668 '        cialll          reathi      ' 'u-^=v-'
-21.5687 '        thenai          edevel      ' 'b,Z?{('
-21.5693 '        edanal          thavei      ' 's ^?{-'
-21.5694 '        cialad          reatea      ' 'u-^={%'
-21.5721 '        edallw          thathr      ' 's ^=v6'
-21.5721 '        thathr          edallw      ' 'b,^%r3'
-21.5748 '        cialav          reates      ' 'u-^={7'
-21.5763 '        reathr          ciallw      ' 'd!^%r3'
-21.5763 '        ciallw          reathr      ' 'u-^=v6'
-21.5778 '        cialai          reatel      ' 'u-^={('
-21.5801 '        cialac          reatef      ' 'u-^={"'
-21.5811 '        cialas          reatev      ' 'u-^={2'
-21.5816 '        edanam          thaveh      ' 's ^?{,'
-21.5831 '        cialaa          reated      ' 'u-^={ '
-21.5852 '        edanav          thaves      ' 's ^?{7'
-21.5856 '        thenak          edeven      ' 'b,Z?{*'
-21.5872 '        edanaw          thaver      ' 's ^?{6'
```

```
10 - 16
-21.6144 '          tparab          thewin    ' 'K!{3F&'
-21.6306 '          outime          omplei    ' 'P$n(J!'
-21.6306 '          omplei          outime    ' 'P<j-B-'
-21.6315 '          tparin          thewab    ' 'K!{3N*'
-21.6350 '          anadde          aveali    ' '^?{%C!'
-21.6378 '          anaddi          aveale    ' '^?{%C-'
-21.6378 '          sparab          shewin    ' 'L!{3F&'
-21.6391 '          tparie          thewai    ' 'K!{3N!'
-21.6438 '          inalfo          iveinc    ' 'V?{-A+'
-21.6470 '          eparab          ehewin    ' 'Z!{3F&'
-21.6503 '          tpariu          theway    ' 'K!{3N1'
-21.6537 '          tparib          thewan    ' 'K!{3N&'
-21.6549 '          sparin          shewab    ' 'L!{3N*'
-21.6584 '          inadde          iveali    ' 'V?{%C!'
-21.6607 '          onalfo          oveinc    ' 'P?{-A+'
-21.6611 '          inaddi          iveale    ' 'V?{%C-'
-21.6624 '          sparie          shewai    ' 'L!{3N!'
-21.6636 '          nlawai          nterie    ' 'Q={6F-'
-21.6641 '          eparin          ehewab    ' 'Z!{3N*'
-21.6684 '          tpario          thewac    ' 'K!{3N+'
-21.6716 '          eparie          ehewai    ' 'Z!{3N!'
-21.6737 '          spariu          sheway    ' 'L!{3N1'
-21.6747 '          autime          amplei    ' '^$n(J!'
-21.6747 '          amplei          autime    ' '^<j-B-'
-21.6757 '          elawai          eterie    ' 'Z={6F-'
-21.6771 '          sparib          shewan    ' 'L!{3N&'
-21.6796 '          alawai          aterie    ' '^={6F-'
-21.6800 '          flawai          fterie    ' 'Y={6F-'
-21.6821 '          onawai          overie    ' 'P?{6F-'
-21.6829 '          epariu          eheway    ' 'Z!{3N1'
-21.6838 '          inawai          iverie    ' 'V?{6F-'
-21.6846 '          slawai          sterie    ' 'L={6F-'
-21.6846 '          nparab          nhewin    ' 'Q!{3F&'
-21.6850 '          enawai          everie    ' 'Z?{6F-'
-21.6862 '          eparib          ehewan    ' 'Z!{3N&'
-21.6867 '          tparia          thewam    ' 'K!{3N%'
-21.6899 '          tparah          thewid    ' 'K!{3F,'
-21.6901 '          tpariz          thewav    ' 'K!{3N>'
-21.6905 '          aladem          ateama    ' '^={%B)'
-21.6911 '          outima          omplem    ' 'P$n(J%'
-21.6911 '          omplem          outima    ' 'P<j-B)'
-21.6918 '          spario          shewac    ' 'L!{3N+'
-21.6920 '          analfo          aveinc    ' '^?{-A+'
-21.6931 '          isaidi          ikelle    ' 'V"{(C-'
-21.6953 '          aladei          ateame    ' '^={%B-'
-21.6957 '          onawin          overab    ' 'P?{6N*'
-21.6964 '          nlawee          ntermi    ' 'Q={6B!'
-21.6974 '          inawin          iverab    ' 'V?{6N*'
-21.6984 '          nlawin          nterab    ' 'Q={6N*'
-21.6986 '          enawin          everab    ' 'Z?{6N*'
-21.6986 '          tlawai          tterie    ' 'K={6F-'
-21.6994 '          isaide          ikelli    ' 'V"{(C!'
-21.6998 '          dparab          dhewin    ' '[!{3F&'
-21.7000 '          aladec          ateamo    ' '^={%B''
-21.7010 '          epario          ehewac    ' 'Z!{3N+'
-21.7010 '          nlawab          nterin    ' 'Q={6F&'
-21.7018 '          nparin          nhewab    ' 'Q!{3N*'
-21.7027 '          inalli          iveide    ' 'V?{-K-'
-21.7032 '          tpaidi          thelle    ' 'K!{(C-'
-21.7034 '          nwillh          nomidd    ' 'Q&s-K,'
-21.7043 '          inalfa          iveinm    ' 'V?{-A%'
-21.7053 '          italfo          ileinc    ' 'V%{-A+'
-21.7084 '          aladea          ateamm    ' '^={%B%'
-21.7085 '          elawee          etermi    ' 'Z={6B!'
-21.7086 '          twillh          tomidd    ' 'K&s-K,'
-21.7087 '          byhimh          balled    ' '](r(J,'
-21.7093 '          nparie          nhewai    ' 'Q!{3N!'
-21.7095 '          tpaide          thelli    ' 'K!{(C!'
-21.7096 '          esaidi          ekelle    ' 'Z"{(C-'
-21.7101 '          sparia          shewam    ' 'L!{3N%'
-21.7105 '          elawin          eterab    ' 'Z={6N*'
-21.7108 '          inalfi          iveine    ' 'V?{-A-'
-21.7120 '          emplei          eutime    ' 'Z<j-B-'
-21.7120 '          eutime          emplei    ' 'Z$n(J!'
-21.7123 '          twasac          toevio    ' 'K&{2F''
-21.7124 '          alawee          atermi    ' '^={6B!'
-21.7128 '          flawee          ftermi    ' 'Y={6B!'
-21.7131 '          elawab          eterin    ' 'Z={6F&'
-21.7132 '          sparah          shewid    ' 'L!{3F,'
-21.7135 '          spariz          shewav    ' 'L!{3N>'
-21.7136 '          isaidm          ikella    ' 'V"{(C)'
-21.7138 '          tparax          thewit    ' 'K!{3F<'
-21.7144 '          alawin          aterab    ' '^={6N*'
-21.7147 '          rparab          rhewin    ' 'M!{3F&'
-21.7148 '          flawin          fterab    ' 'Y={6N*'
-21.7159 '          esaide          ekelli    ' 'Z"{(C!'
-21.7163 '          ssaidi          skelle    ' 'L"{(C-'
-21.7170 '          dparin          dhewab    ' '[!{3N*'
-21.7170 '          alawab          aterin    ' '^={6F&'
-21.7174 '          flawab          fterin    ' 'Y={6F&'
-21.7174 '          slawee          stermi    ' 'L={6B!'
-21.7178 '          eladem          eteama    ' 'Z={%B)'
-21.7182 '          ewtime          eoplei    ' 'Z&n(J!'
-21.7182 '          eoplei          ewtime    ' 'Z>j-B-'
-21.7193 '          eparia          ehewam    ' 'Z!{3N%'
-21.7194 '          slawin          sterab    ' 'L={6N*'
-21.7195 '          onawab          overin    ' 'P?{6F&'
-21.7195 '          shadde          speali    ' 'L9{%C!'
-21.7196 '          onalli          oveide    ' 'P?{-K-'
-21.7198 '          rwillh          romidd    ' 'M&s-K,'
```

do these look like keys:

attherustinlawabeonemotheenterinight

