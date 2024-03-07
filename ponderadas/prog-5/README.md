# Atividade 5: Integração do simulador com Metabase

## Enunciado

Nessa atividade, deve-se desenvolver a integração entre o simulador desenvolvido nas três primeiras atividades e um dashboard desenvolvido usando o Metabase, com persistência de dados em um banco de dados a sua escolha.

## Estrutura de pastas
<pre><code>prog-5/
│
├── db/db.db
├── go.mod
├── publisher.go
├── subscriber.go
├── main.go
└── mosquito.conf</code></pre>

Onde:
```db/db.db```: Banco de dados sqlite com os dados simulados do sensor.
```go.mod```: Módulo do Go.
```publisher.go```: Arquivo que possui o código necessário para criar um publicador e um loop para as mensagens serem publicadas;
```subscriber.go```: Arquivo que possui o código necessário para criar um subscriber a fim de visualizar as mensagens publicadas E guardá-las no banco de dados;
```main.go```: Arquivo que possui o código necessário para setar o ambiente do banco de dados e inicializar o sistema;
```.env```: Arquivo de ambiente para guardar de forma segura alguns valores que podem ser secretos, nesse caso, você precisa completar alguns valores, como explicado na sessão ```Configurando .env```

## Como usar

Primeiro, certifique-se de criar uma conta no [HiveMQ](https://www.hivemq.com) com um servidor configurado e de possuir o [Go](https://go.dev/dl/) e o [Docker](https://www.docker.com) instalados:

Assim, instale as dependências neste diretório:
<pre><code>go mod tidy</code></pre>

### Configurando .env

Agora, assim como dito anteriormente, crie um arquivo ```.env``` e complete ele com os seguintes valores:
<code><pre>BROKER_ADDR="your address"
HIVE_USER="your user"
HIVE_PSWD="your password"</pre></code>


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
[metabase.webm](https://github.com/Lukovsk/Inteli-Modulo-9/assets/99260684/fd04b2af-36b2-4b7b-8bb7-c18b92d6b15f)
