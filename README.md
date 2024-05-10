# Funcionamentos avançados do Kafka

## Primeiros passos

- Importante baixar o TDM-GCC-64 para Windows clique [aqui](https://jmeubank.github.io/tdm-gcc/)

![img.png](img.png)
- Logo após executar o .exe, pegue o diretório em que foi baixado (Ex: `C:\TDM-GCC-64\bin`)
- Abra seu gerenciador de variáveis
![img_1.png](img_1.png)
![img_2.png](img_2.png)

---

- Suba o docker-compose no terminal: ```docker-compose up -d```

- Após baixar as imagens do kafka, acesse o bash: `docker exec -it gokafka bash`

- Utilize o comando: `go mod init github.com/gui-meireles/fc2-kafka-avanced`

- Baixe as dependências do GoLang em sua IDE (Eu utilizei o Intellij)

- Rode o arquivo `main.go` que está dentro de `cmd/producer` com o comando: `go run cmd/producer/main.go` 
dentro do bash do `gokafka`

> Caso apareça a mensagem "Hello Go" significa que está funcionando

---
### Testar o envio de mensagens para o kafka:

- Abra um outro terminal
- Abra o bash do kafka: `docker exec -it fc2-kafka-advanced-kafka-1 bash`
- Vamos criar um tópico: `kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3`
- Logo após, abra o console do consumer no mesmo bash:
`kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste`

> Com o console do consumer aberto, vamos rodar nossa aplicação com o comando `go run cmd/producer/main.go` dentro
do bash `docker exec -it gokafka bash` (No outro terminal)

**Com isso, você verá que no console do consumer chegará uma "Mensagem"**

---

