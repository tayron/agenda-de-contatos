#!/bin/bash
echo && \
echo "---------------------------------------------------------" && \
echo "Subindo aplicação:" && \
echo "---------------------------------------------------------" && \
docker-compose up --build -d && \

echo && \
echo && \
echo "---------------------------------------------------------" && \
echo "Obtendo ip da aplicação:" && \
echo "---------------------------------------------------------" && \
echo && \
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' agenda.contato