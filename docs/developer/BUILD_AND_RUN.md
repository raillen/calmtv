# Build and Run

## Development build

The shell and pure services are built from the repository root:

```text
go test ./...
go build -trimpath -o build/tv-shell ./cmd/tv-shell
```

GTK3/system libraries remain system dependencies; Go modules manage Go dependencies.

For the Debian image build, enable SQLite FTS5 and use the project Makefile:

```text
make test
make build
make image
```

The image consumes a `.deb` from `build/apt/` so the shell is installed as a
controlled package rather than copied into the image ad hoc. Build that
package in a Debian build environment before `make image`. `make image`
requires `live-build`; missing tools and a missing package fail explicitly.
For a windowed local run, use `TV_SHELL_WINDOWED=1` with an X11 display:

```text
TV_SHELL_WINDOWED=1 build/tv-shell
```

## Testar como sessão adicional no Debian

O pacote instala uma sessão selecionável sem remover o desktop existente:

```text
/usr/share/xsessions/tv-shell.desktop
/usr/bin/tv-shell-session
```

Em uma máquina Debian amd64 de teste:

```bash
sudo apt install ./tv-shell_*.deb
```

The package declares `pkexec` and `polkitd` with a fallback to Debian's older
`policykit-1` transitional package.

Depois, encerre a sessão atual, selecione `Calm TV` no menu de sessões da
tela de login e entre normalmente. Para voltar, selecione o desktop anterior
no mesmo menu. A instalação não troca o display manager nem altera a sessão
padrão.

Se a sessão não aparecer, valide:

```bash
test -x /usr/bin/tv-shell-session
test -f /usr/share/xsessions/tv-shell.desktop
command -v matchbox-window-manager
```

O launcher encerra Matchbox ao sair da Shell. Em caso de falhas repetidas, o
launcher da Shell entra no modo de recuperação e preserva o desktop original.

Antes do primeiro teste no notebook, execute no próprio Debian:

```bash
./scripts/target-preflight build/target-preflight
cat build/target-preflight/environment.txt
cat build/target-preflight/boot.txt
```

Em uma configuração com dois monitores, a sessão usa o monitor primário e não
atravessa as duas telas. Para testar explicitamente o segundo monitor:

```bash
TV_SHELL_MONITOR=1 /usr/bin/tv-shell
```

`TV_SHELL_MONITOR` é um índice técnico para desenvolvimento; a seleção por
nome do monitor será exposta nas Quick Settings quando o adapter de display
estiver conectado à UI.

Para manutenção remota, consulte [`docs/operations/UPDATES.md`](../operations/UPDATES.md).
Para abrir um terminal local avançado, pressione `Menu` na Home, entre em
Diagnóstico e selecione **Terminal avançado**. O Q4OS precisa ter um terminal
gráfico instalado; `qterminal` é uma opção leve.

Esse relatório registra o hardware e as dependências sem instalar ou alterar
serviços.

Depois de selecionar a sessão, execute `./scripts/target-session-check` para
registrar Xorg, Matchbox, Shell, displays conectados, unidades de usuário
falhas, memória e o PSS inicial da Shell.

## Target image

Debian image generation lives under `image/live-build/`. Package metadata lives under `packaging/debian/`.

Expected CI layers:
1. format/lint/unit tests;
2. integration tests;
3. Debian package build;
4. image build;
5. VM smoke;
6. self-hosted hardware gates.

Do not make a manually configured developer machine the only reproducible build path.
