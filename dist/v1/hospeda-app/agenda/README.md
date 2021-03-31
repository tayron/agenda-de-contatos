# Agenda de Contato
Sistema para poder gerenciar e consultar contatos interno de uma empresa

## Instalação
1. Configure os dados de conexão do banco de dados no arquivo .env antes da instalação
2. Execute ```sudo install.sh```
Após a instalação você deverá receber a seguinte mensagem: 
```
Agenda de Contatos - versão 1.0

● Removendo instalação antiga
● Removendo serviço antigo
● Arquivos copiados
● Permissões setadas
● Arquivo de serviço copiado
● Atualizando lista de serviços do sistema operacional
● Executando serviço
● agenda-contato.service
   Loaded: loaded (/etc/systemd/system/agenda-contato.service; disabled; vendor 
   Active: active (running) since Wed 2021-03-31 13:03:51 -03; 6ms ago
 Main PID: 25771 (agenda)
    Tasks: 1 (limit: 4915)
   Memory: 2.1M
   CGroup: /system.slice/agenda-contato.service
           └─25771 /usr/local/hospeda-app/agenda/agenda

mar 31 13:03:51 tayron systemd[1]: Started agenda-contato.service.
```

Embora tenha configurado em qual porta a aplicação será executada, pode usar o seguinte comando para ver em qual porta está sendo usada e até mesmo ver as mensagens de erro com o seguinte comando:
```
journalctl -u agenda-contato -f
```

A saída do comando será algo como: 
```
-- Logs begin at Wed 2021-03-31 08:46:54 -03. --
mar 31 13:03:21 tayron systemd[1]: agenda-contato.service: Main process exited, code=exited, status=2/INVALIDARGUMENT
mar 31 13:03:21 tayron systemd[1]: agenda-contato.service: Failed with result 'exit-code'.
mar 31 13:03:39 tayron systemd[1]: /etc/systemd/system/agenda-contato.service:1: Assignment outside of section. Ignoring.
mar 31 13:03:39 tayron systemd[1]: /etc/systemd/system/agenda-contato.service:2: Assignment outside of section. Ignoring.
mar 31 13:03:51 tayron systemd[1]: agenda-contato.service: Service RestartSec=30s expired, scheduling restart.
mar 31 13:03:51 tayron systemd[1]: agenda-contato.service: Scheduled restart job, restart counter is at 10.
mar 31 13:03:51 tayron systemd[1]: Stopped agenda-contato.service.
mar 31 13:03:51 tayron systemd[1]: Started agenda-contato.service.
mar 31 13:03:51 tayron agenda[25771]: ####################### Agenda de Contatos - versão 1.0 #######################
mar 31 13:03:51 tayron agenda[25771]: Servidor executando em: http://127.0.0.1:3007
```

3. Abra o navegador e acesse 127.0.0.1:3007 (3007 - porta da aplicação configurada no arquivo .env)
