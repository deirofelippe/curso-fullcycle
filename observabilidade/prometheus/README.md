# README

## Token do Telegram

- Procure por `BotFather` no Telegram
- Escreva `/start` para listar comandos
- Escreva `/newbot` para criar
  - Nome do bot: `Teste Automações`
  - Username do bot que termine com bot: `TesteAutomacoes_Bot`
  - Copie o token recebido
- Crie um grupo
  - Adicione o bot no grupo e o adicione como admin
  - Envie uma mensagem no grupo
  - Acesse a seguinte url para buscar o id do chat: `https://api.telegram.org/bot<TOKEN>/getUpdates`
- Com o id do chat e o token do bot, consegue enviar mensagem
  - Acesse a url e receberá a notificação: `https://api.telegram.org/bot<TOKEN>/sendMessage?chat_id=<CHAT_ID>&text=Teste`
