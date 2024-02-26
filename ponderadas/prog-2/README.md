# Atividade 2: Teste de um simulador de dispositivos IoT

## Enunciado

Utilizando o simulador de dispositivos IoT desenvolvido na atividade passada e utilizando os conceitos de TDD vistos no decorrer da semana, implemente testes automatizados para validar o simulador. Seus testes obrigatoriamente devem abordar os seguintes aspectos:

Recebimento - garante que os dados enviados pelo simulador são recebidos pelo broker.
Validação dos dados - garante que os dados enviados pelo simulador chegam sem alterações.
Confirmação da taxa de disparo - garante que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.

## Conteúdo

Este é um simulador de sensor MiCS-6814 que gera dados fictícios para simular a leitura de gases como CO, NO2 e NH3, e os envia via MQTT.
Aqui contém um simulador de sensor MiCS-6814 que gera dados fictícios para simular a leitura de gases e comunicação MQTT com testes automáticos.

## Estrutura de pastas
<pre><code>prog-2/
│
├── go.mod
├── publisher.go
├── sub/subscriber.go
├── simulation_test.go
└── mosquito.conf</code></pre>

Onde:
```go.mod```: Módulo do Go.
```publisher.go```: Arquivo que possui o código necessário para criar um publicador e um loop para as mensagens serem publicadas;
```sub/subscriber.go```: Arquivo que possui o código necessário para criar um subscriber a fim de visualizar as mensagens publicadas;
```simulation_test.go```: Arquivo que possui o código necessário com testes automatizados para simular ocasiões em que o publisher possa não funcionar;
```mosquito.conf```: Arquivo necessário para a criação do broker mqtt;

## Como usar
Primeiro, certifique-se de possuir o Go e o Mosquitto MQTT Broker que podem ser instalados, respectivamente, nos seguintes links:

- [Go](https://go.dev/dl/)
- [Mosquitto](https://mosquitto.org/download/)

Instale as dependências neste diretório:
<pre><code>go mod tidy</code></pre>

Agora, com 4 terminais, podemos:

### Broker MQTT
Inicie o broker MQTT (garantindo que que o Mosquitto MQTT Broker esteja instalado). Com isso, podemos iniciar o broker MQTT:
<pre><code>mosquitto -c mosquito.conf</code></pre>

### Publisher
Para iniciar o publisher, basta executar o arquivo ```publisher.go```:
<pre><code>go run publisher.go</code></pre>

### Visualizar as mensagens recebidas
Também podemos visualizar as mensagens recebidas pelo broker, então vamos nos inscrever utilizando o ```sub/subscriber.go```:
<pre><code>go run sub/subsciber.go</code> </pre>

### Testar o ambiente MQTT
Por último, vamos testar todo esse ambiente que criamos. Nesta automação, 3 testes são realizados:
- Conexão;
- Validação das mensagens;
- Teste de publicação;

Para isso, executamos o teste em Go:
<pre><code>go test -v -cover</code> </pre>

Resultado esperado:
<pre><code>=== RUN   TestConection
        simulation_test.go:20: Connection with broker MQTT succeeded
    --- PASS: TestConection (0.00s)
    === RUN   TestDataValidation
        simulation_test.go:34: Data validation successfull
    --- PASS: TestDataValidation (0.00s)
    === RUN   TestPublisher
        simulation_test.go:84: Message received
    --- PASS: TestPublisher (0.00s)
    PASS
    coverage: 10.5% of statements
    ok  	mqtt-go	0.003s</code></pre>

## Demonstração
[go.webm](https://github.com/Lukovsk/Inteli-Modulo-9/assets/99260684/7fef6ce1-09af-444e-87bb-fb80cd93bb80)
