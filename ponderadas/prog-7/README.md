# Atividade 7: Integração entre o Kafka cloud e o HiveMQ

## Enunciado

Nessa atividade, deve-se desenvonver a integração entre o cluster do HiveMQ e um provedor de Kafka em nuvem.

## Estrutura de pastas

<pre><code>prog-7/
│
├── go.mod
├── publisher.go
├── pipe_test.go
├── .env
└── main.go</code></pre>

Onde:   
```go.mod```: Módulo do Go.
```publisher.go```: Arquivo que possui o código necessário para criar um publicador e um loop para as mensagens serem publicadas;
```pipe_test.go```: Arquivo que possui o código necessário para testar a pipeline;
```.env```: Arquivo de ambiente para guardar de forma segura alguns valores que podem ser secretos, nesse caso, você precisa completar alguns valores, como explicado na sessão ```Configurando .env```

## Como usar

Primeiro, certifique-se de criar uma conta no [HiveMQ](https://www.hivemq.com) e no [MongoDB](https://cloud.mongodb.com/) com um servidor e um banco de dados configurados e de possuir o [Go](https://go.dev/dl/) e o [Docker](https://www.docker.com) instalados:

Assim, instale as dependências neste diretório:
<pre><code>go mod tidy</code></pre>

### Configurando .env

Agora, assim como dito anteriormente, crie um arquivo ```.env``` e complete ele com os seguintes valores:
<code><pre>BROKER_ADDR="your address"
HIVE_USER="your user"
HIVE_PSWD="your password"
MONGO_URL="mongodb+srv://<usuario>:<senha>@<host>/<banco>"</pre></code>


### Teste
O teste é composto pelo seguinte esquema:

#### Inicialização do mongo

![alt text](media/mongo.png)

#### Publicando um dado

![alt text](media/publishing.png)

#### Acessando o dado da collection do mongo

![alt text](media/filter.png)

#### Comprarando 

![alt text](media/comparando.png)

## Demonstração
[6d97bc78-81a2-4e92-820b-34d063d284f0.webm](https://github.com/Lukovsk/Inteli-Modulo-9/assets/99260684/ec2e5763-e563-407d-a53a-38e0d50536dc)




