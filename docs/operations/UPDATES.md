# Updates

## MVP
Use Debian packages (`.deb`) and an authenticated package repository/update flow because it minimizes custom infrastructure.

Requirements:
- signed packages/repository;
- no unattended breaking migration without recovery;
- clear reboot requirement;
- version/rollback compatibility documented.

## Acesso remoto no Q4OS

O repositório fornece um bootstrap idempotente para instalar OpenSSH,
configurar uma conta não-root e instalar o atualizador remoto restrito ao
pacote `tv-shell`:

```bash
wget -qO- https://raw.githubusercontent.com/raillen/calmtv/main/scripts/calm-tv-remote-setup | sudo sh
```

Esse primeiro modo mantém senha SSH para permitir o bootstrap. O modo
recomendado instala uma chave pública e desativa senha:

```bash
wget -qO- https://raw.githubusercontent.com/raillen/calmtv/main/scripts/calm-tv-remote-setup \
  | sudo sh -s -- --authorized-key-file "$HOME/.ssh/id_ed25519.pub"
```

Depois, de outro computador na mesma rede:

```bash
ssh usuario@IP_DO_NOTEBOOK sudo -n /usr/local/sbin/calm-tv-remote-update
```

O atualizador consulta a última Release do GitHub, exige versão maior que a
instalada, baixa o `.deb`, valida o SHA-256, instala com APT e reinicia o
display manager somente depois da instalação. O SSH não deve ser exposto
diretamente à Internet; use uma rede confiável ou VPN.

O SHA-256 protege a integridade do arquivo baixado, mas ainda não substitui um
repositório APT assinado. Assinatura/proveniência completa permanece gate de
release e hardening posterior.

## Future
Evaluate transactional/A-B image update or `systemd-sysupdate` only after the product is stable enough to justify the added image/partition complexity.

Update mechanism changes require an ADR.
