# Hackコンピュータ

## cmd/test

Hack機械語のプログラム（`.hack`）をHackコンピュータで実行し、実行後のレジスタとメモリを表示する。

```sh
go run ./cmd/test [flags] program.hack
```

| flag | 説明 | default |
| --- | --- | --- |
| `-cycles n` | 実行する命令数の上限 | `10000` |
| `-ram list` | 表示するRAMのアドレス（例: `0-2,16`） | `0-15` |
| `-set list` | 実行前に書き込むメモリの値（例: `0=3,1=5`） | なし |
| `-trace` | 1命令ごとにレジスタを表示する | `false` |

`.hack` は1行に16桁の2進数で1命令を書く。空行と `//` 以降のコメントは無視する。

次のような停止ループに入るか、`-cycles` の上限に達すると実行を止める。

```
(END)
    @END
    0;JMP
```

### 例

```sh
$ go run ./cmd/test -set 0=6,1=7 -ram 0-2 cmd/test/testdata/Mult.hack
halted after 80 cycles
PC=14 A=14 D=0
RAM[0] = 0
RAM[1] = 7
RAM[2] = 42
```

`cmd/test/testdata/` にサンプルプログラムがある（`.asm` はアセンブル前のソース）。

| ファイル | 内容 |
| --- | --- |
| `Add.hack` | RAM[0] = 2 + 3 |
| `Max.hack` | RAM[2] = max(RAM[0], RAM[1]) |
| `Mult.hack` | RAM[2] = RAM[0] × RAM[1]（RAM[0] ≥ 0） |
