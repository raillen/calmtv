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

## Atualizações automáticas — MVP

**Status: Planejado.** O caminho automático ainda não está conectado ao menu
de Configurações nem possui um `systemd.timer` instalado. O bootstrap remoto e
o comando manual acima são a base operacional existente.

O comportamento alvo é:

1. um serviço leve verifica periodicamente a origem de updates, sem manter
   processo pesado residente;
2. a Shell mostra a versão instalada e a existência de nova versão em
   **Configurações → Sistema → Atualizações**;
3. o usuário pode escolher **Instalar e reiniciar**;
4. a instalação ocorre por APT/PolicyKit, fora do processo GTK e sem `sudo`
   chamado pela UI;
5. o pacote é baixado em arquivo temporário e validado por repositório
   assinado, checksum, nome, versão e arquitetura;
6. a Shell mostra progresso, erro compreensível e resultado;
7. o display manager só é reiniciado depois da instalação bem-sucedida;
8. após o retorno, um health check confirma que a nova Shell iniciou;
9. se a inicialização falhar, o recovery mantém o desktop anterior acessível e
   oferece rollback para a versão anterior quando o mecanismo de pacotes
   permitir.

Durante o MVP, a política padrão deve ser **verificação automática com
confirmação humana para instalar**. Atualização totalmente silenciosa fica
condicionada a metadados assinados, janela configurável, rollback testado e
recuperação independente da Shell.

Contrato de configuração proposto, ainda não implementado:

```text
updates.enabled
updates.channel
updates.check_interval
updates.install_policy = confirm | unattended
updates.restart_policy = ask | after_install
```

Critérios de aceite do Goal de atualização automática:

- nenhum update é instalado quando a assinatura/checksum é inválido;
- falha de rede não altera o sistema nem bloqueia a Home;
- versões iguais ou antigas não são instaladas;
- a UI não congela durante consulta/download/instalação;
- reinício exige confirmação, exceto em política explicitamente configurada;
- atualização interrompida deixa diagnóstico e caminho de recovery;
- o fluxo funciona por D-pad, teclado e sem terminal.

## Future
Evaluate transactional/A-B image update or `systemd-sysupdate` only after the product is stable enough to justify the added image/partition complexity.

Update mechanism changes require an ADR.
