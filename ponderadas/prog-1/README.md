# Atividade 1: Teste de um simulador de dispositivos IoT

## Enunciado

Utilizando o simulador de dispositivos IoT desenvolvido na atividade passada e utilizando os conceitos de TDD vistos no decorrer da semana, implemente testes automatizados para validar o simulador. Seus testes obrigatoriamente devem abordar os seguintes aspectos:

- Recebimento - garante que os dados enviados pelo simulador são recebidos pelo broker.
- Validação dos dados - garante que os dados enviados pelo simulador chegam sem alterações.
- Confirmação da taxa de disparo - garante que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.

## Conteúdo

Este é um simulador de sensor MiCS-6814 que gera dados fictícios para simular a leitura de gases como CO, NO2 e NH3, e os envia via MQTT.

## Estrutura de pastas
<pre><code>prog-1/
│
├── main.py
├── mosquito.conf
└── requirements.txt</code></pre>

Onde:
```main.py```: Arquivo que possui o código necessário para criar um publicador e um loop para as mensagens serem publicadas;
```mosquito.conf```: Arquivo necessário para a criação do broker mqtt;
```requirements.txt```: Arquivo de texto com as dependências necessárias para executar a aplicação.

## Como usar
Primeiro, certifique-se de possuir o python e o Mosquitto MQTT Broker que podem ser instalados, respectivamente, nos seguintes links:

- [Python](https://www.python.org/)
- [Mosquitto](https://mosquitto.org/download/)

Agora, abra um ambiente virtual python executando o comando:
<pre><code>python -m venv venv</code></pre>

Instale as dependências neste diretório:
<pre><code>python -m pip install -r requirements.txt</code></pre>

Agora, com 3 terminais, podemos:

### Broker MQTT
Inicie o broker MQTT (garantindo que que o Mosquitto MQTT Broker esteja instalado). Com isso, podemos iniciar o broker MQTT:
<pre><code>mosquitto -c mosquito.conf</code></pre>

### Publisher
Para iniciar o publisher, basta executar o arquivo ```main.py``` com a venv ligada e as dependências instaladas:
<pre><code>python3 main.py</code></pre>

### Visualizar as mensagens recebidas
Também podemos visualizar as mensagens recebidas pelo broker, então vamos nos inscrever utilizando o ```mosquit_sub```:
<pre><code>mosquitto_sub -h localhost -p 1891 -v -t "sensor/mics6814"</code> </pre>

## Video
[mosquito-p1.webm](https://github.com/Lukovsk/Inteli-Modulo-9/assets/99260684/405f93c1-5486-4218-99aa-c567a2d504c0)
