# Instalação no Q4OS

O pacote é para Debian/Q4OS amd64. Ele adiciona `TV Shell` como uma sessão
opcional e não remove o desktop atual.

## Instalar

```bash
sudo apt update
sudo apt install ./tv-shell_0.1.1_amd64.deb
```

O pacote aceita os nomes modernos `pkexec`/`polkitd` e mantém fallback para o
pacote transitório `policykit-1` das versões Debian anteriores.

Se o APT solicitar dependências, confirme a instalação. Depois encerre a
sessão atual e escolha `TV Shell` no menu de sessões da tela de login.

Para voltar ao Q4OS, encerre a sessão e escolha o desktop original no mesmo
menu.

## Verificar

```bash
test -x /usr/bin/tv-shell-session
test -f /usr/share/xsessions/tv-shell.desktop
command -v matchbox-window-manager
```

Com a sessão TV Shell ativa, a partir de uma cópia do projeto execute:

```bash
./scripts/target-preflight build/target-preflight
./scripts/target-session-check build/target-session-check
```

O pacote não promete aceleração VA-API, streaming comercial ou RetroArch
quando os respectivos componentes não estiverem instalados no Q4OS.
