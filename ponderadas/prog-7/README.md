# Atividade 6: Integração do simulador com o MongoDB

## Enunciado

Nessa atividade, deve-se desenvolver a integração entre o simulador desenvolvido nas três primeiras atividades e um dashboard desenvolvido usando o Metabase, com persistência de dados em um banco de dados a sua escolha.

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
```subscriber.go```: Arquivo que possui o código necessário para criar um subscriber a fim de visualizar as mensagens publicadas E guardá-las no banco de dados;
```main.go```: Arquivo que possui o código necessário para setar o ambiente do banco de dados e inicializar o sistema;
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


### Backend
Para iniciar o publisher, que publicará, constantemente, dados dos sensores, o subscriber, que mostrar-los-á no terminal e guardar-los-á no banco de dados, basta executar o arquivo ```main.go``` junto com todos os outros arquivos ```.go```:
<pre><code>go run *.go</code></pre>

### Rodando o Metabase
Como uma espécie de frontend, utilizamos o [Metabase](https://www.metabase.com/). Para rodá-lo, com o docker instalado, baixe a imagem do metabase:

<pre><code>docker pull metabase/metabase</code></pre>

Depois, basta rodar o seguinte comando neste diretório, ele criará um novo container com a imagem do metabase e colocará o banco de dados local como volume para o frontend:

<pre><code>docker run -d -p 3000:3000 -v $(pwd)/db/db.db:/db.db --name metabase metabase/metabase</code></pre>

Dessa forma, basta ir no seu [localhost:3000](http://localhost:3000) e configurar seu metabase para criar um dashboard com os dados dos sensores.

## Demonstração
[84fc15b7-e167-453b-b876-1f0e0454c132.webm](https://github.com/Lukovsk/Inteli-Modulo-9/assets/99260684/fb7aafbd-c9c9-4e86-b788-783c148a6991)


