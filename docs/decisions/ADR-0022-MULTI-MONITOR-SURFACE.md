# ADR-0022 — Superfície única para múltiplos monitores

**Status:** Accepted
**Date:** 2026-08-13

## Contexto

Quando o X11 possui dois monitores, `gtk_window_fullscreen()` pode tratar o
espaço virtual inteiro como uma única área. A Home do Calm TV então fica
dividida entre a tela do notebook e o monitor externo, o que é inadequado para
uma interface controlada por D-pad.

## Decisão

O Calm TV terá uma única superfície de interface ativa por vez:

- a sessão inicia em tela cheia no monitor primário;
- a Shell usa fullscreen específico do monitor, nunca fullscreen do espaço
  virtual inteiro;
- durante desenvolvimento, `TV_SHELL_MONITOR=<índice>` permite selecionar um
  monitor; valor inválido retorna ao monitor primário;
- se o monitor escolhido desaparecer, a próxima inicialização usa o monitor
  primário disponível;
- o segundo monitor permanece livre e não cria um segundo FocusManager.

O caminho futuro de Display poderá oferecer modos explícitos:

1. **Uma tela** — padrão e recomendado para TV;
2. **Espelhar** — mesma interface nas duas saídas, somente quando o driver
   suportar isso de forma confiável;
3. **Tela auxiliar** — informações secundárias, sem navegação independente,
   somente depois que o MVP principal estiver funcional.

## Alternativas consideradas

- **Estender a Shell entre monitores:** rejeitado; prejudica legibilidade,
  foco e previsibilidade do controle remoto.
- **Duas interfaces interativas independentes:** rejeitado; duplica estado,
  input e ciclo de vida sem benefício para o appliance de 2 GB.
- **Desligar sempre o segundo monitor:** não adotado; o desktop existente e o
  uso de desenvolvimento continuam disponíveis.

## Consequências

A correção inicial é pequena, reversível e compatível com Xorg/Matchbox. A
seleção persistente por nome de conector e os modos espelhado/auxiliar ficam
como extensão das Quick Settings, sem atrasar o MVP.
