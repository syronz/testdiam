let SessionLoad = 1
let s:so_save = &so | let s:siso_save = &siso | set so=0 siso=0
let v:this_session=expand("<sfile>:p")
silent only
cd ~/sandbox/diameter/go-diameter
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
set shortmess=aoO
badd +1 README.md
badd +73 examples/client/DEFAULT.md
badd +262 examples/client/client.go
badd +3 examples/client/README.md
badd +226 examples/client/ccr_sample.go
badd +172 examples/server/ccr_server.go
badd +1 examples/server/server.go
badd +1 diam/dict/doc.go
badd +1 diam/dict/default.go
badd +4 diam/commands.go
badd +261 diam/message_test.go
badd +49 diam/sm/smpeer/metadata.go
badd +18 diam/codes.go
badd +1 diam/avp/codes.go
badd +19 diam/datatype/time.go
badd +49 diam/datatype/time_test.go
badd +202 diam/message.go
badd +1 AVP_TYPES.md
badd +13 diam/datatype/enum.go
badd +4 diam/datatype/octetstring.go
badd +1 diam/datatype/diamident.go
badd +1 diam/datatype/group.go
badd +1 diam/reflect.go
badd +1 diam/reflect_test.go
badd +0 diam/datatype/address.go
argglobal
%argdel
edit examples/client/DEFAULT.md
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe '1resize ' . ((&lines * 45 + 30) / 60)
exe 'vert 1resize ' . ((&columns * 116 + 119) / 238)
exe '2resize ' . ((&lines * 45 + 30) / 60)
exe 'vert 2resize ' . ((&columns * 121 + 119) / 238)
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 57 - ((19 * winheight(0) + 22) / 45)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
57
normal! 050|
wincmd w
argglobal
if bufexists("examples/client/README.md") | buffer examples/client/README.md | else | edit examples/client/README.md | endif
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 18 - ((13 * winheight(0) + 22) / 45)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
18
normal! 053|
wincmd w
exe '1resize ' . ((&lines * 45 + 30) / 60)
exe 'vert 1resize ' . ((&columns * 116 + 119) / 238)
exe '2resize ' . ((&lines * 45 + 30) / 60)
exe 'vert 2resize ' . ((&columns * 121 + 119) / 238)
tabedit diam/dict/default.go
set splitbelow splitright
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=3
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 574 - ((24 * winheight(0) + 29) / 58)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
574
normal! 042|
tabedit diam/reflect_test.go
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 118 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 119 + 119) / 238)
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=3
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 177 - ((37 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
177
normal! 03|
wincmd w
argglobal
if bufexists("diam/reflect.go") | buffer diam/reflect.go | else | edit diam/reflect.go | endif
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=5
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 379 - ((43 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
379
normal! 06|
wincmd w
exe 'vert 1resize ' . ((&columns * 118 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 119 + 119) / 238)
tabedit examples/server/ccr_server.go
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 116 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 121 + 119) / 238)
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=3
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 218 - ((52 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
218
normal! 0107|
wincmd w
argglobal
if bufexists("examples/client/ccr_sample.go") | buffer examples/client/ccr_sample.go | else | edit examples/client/ccr_sample.go | endif
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=5
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 241 - ((43 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
241
normal! 053|
wincmd w
exe 'vert 1resize ' . ((&columns * 116 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 121 + 119) / 238)
tabedit AVP_TYPES.md
set splitbelow splitright
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 193 - ((49 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
193
normal! 0
tabedit diam/avp/codes.go
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 119 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 118 + 119) / 238)
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=1
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 462 - ((26 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
462
normal! 048|
wincmd w
argglobal
if bufexists("diam/datatype/address.go") | buffer diam/datatype/address.go | else | edit diam/datatype/address.go | endif
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=3
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 91 - ((51 * winheight(0) + 28) / 57)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
91
normal! 0
wincmd w
exe 'vert 1resize ' . ((&columns * 119 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 118 + 119) / 238)
tabedit examples/server/ccr_server.go
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
set nosplitbelow
set nosplitright
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 86 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 151 + 119) / 238)
argglobal
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=3
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 62 - ((20 * winheight(0) + 29) / 58)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
62
normal! 03|
wincmd w
argglobal
if bufexists("examples/server/ccr_server.go") | buffer examples/server/ccr_server.go | else | edit examples/server/ccr_server.go | endif
setlocal fdm=syntax
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=3
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 96 - ((3 * winheight(0) + 29) / 58)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
96
normal! 0
wincmd w
exe 'vert 1resize ' . ((&columns * 86 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 151 + 119) / 238)
tabnext 4
if exists('s:wipebuf') && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20 winminheight=1 winminwidth=1 shortmess=filnxtToOF
let s:sx = expand("<sfile>:p:r")."x.vim"
if file_readable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &so = s:so_save | let &siso = s:siso_save
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
