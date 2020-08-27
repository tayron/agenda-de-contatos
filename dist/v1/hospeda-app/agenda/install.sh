#!/bin/bash

echo "Agenda de Contatos - versão 1.0"
echo 

if [ -e /usr/local/hospeda-app/ ] && rm -R /usr/local/hospeda-app/; then
    echo "● Removendo instalação antiga"
fi

if [ -e /etc/systemd/system/agenda-contato.service ] && rm -R /etc/systemd/system/agenda-contato.service; then
    echo "● Removendo serviço antigo"
fi

if cp -R ../../hospeda-app/ /usr/local/hospeda-app/; then 
    echo "● Arquivos copiados"
else
    echo "● Erro ao copiar os arquivos para /usr/local/"
fi

if chmod +x /usr/local/hospeda-app/agenda; then 
    echo "● Permissões setadas"
else 
    echo "● Erro ao dar permissão de execução"
fi 

if cp agenda-contato.service /etc/systemd/system/agenda-contato.service; then 
    echo "● Arquivo de serviço copiado"
else 
    echo "● Erro ao copiar arquivo de serviço para /etc/system/"
fi

if systemctl daemon-reload; then 
    echo "● Atualizando lista de serviços do sistema operacional"
else 
    echo "● Erro ao atualizar lista de serviços do sistema operacional"
fi

if service agenda-contato start; then
    echo "● Executando serviço"
else
    echo "● Falha ao iniciar o serviço, execute systemctl -u agenda-contato -f para mais detalhes"
fi

service agenda-contato status

echo 
echo "Fim da instalação"
