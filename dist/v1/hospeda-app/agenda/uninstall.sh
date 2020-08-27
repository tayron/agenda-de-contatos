#!/bin/bash

echo "Agenda de Contatos - versão 1.0"
echo 

if service agenda-contato stop; then
    echo "● Parando serviço"
else
    echo "● Falha ao para o serviço"
fi

if [ -e /usr/local/hospeda-app/ ] && rm -R /usr/local/hospeda-app/; then
    echo "● Removendo instalação antiga"
fi

if [ -e /etc/systemd/system/agenda-contato.service ] && rm /etc/systemd/system/agenda-contato.service; then
    echo "● Removendo serviço antigo"
fi

if systemctl daemon-reload; then 
    echo "● Atualizando lista de serviços do sistema operacional"
else 
    echo "● Erro ao atualizar lista de serviços do sistema operacional"
fi

echo 
echo "Fim da instalação"
