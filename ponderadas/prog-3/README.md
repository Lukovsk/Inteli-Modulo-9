# Atividade 2: Simulador de ataques usando MQTT

## Enunciado

Atividade prática em grupo, de subir um broker remoto e um broker local do MQTT para conduzir cenários de análise de vulnerabilidade (dentro do CIA Triad), identificando situações onde pode ser comprometido cada um dos três pilares: Confiabilidade, Integridade e Disponibilidade.


## O que é CIA Triad?

No contexto de cibersegurança, "CIA Triad" é uma sigla que representa três princípios fundamentais da segurança da informação: Confiabilidade, Integridade e Disponibilidade. Esses são pilares essenciais para garantir a segurança dos dados e sistemas de informação em qualquer ambiente computacional. Eles também formam a base da segurança da informação e são amplamente utilizados para orientar o desenvolvimento e a implementação de estratégias de cibersegurança em organizações e sistemas de informação.

### Confiabilidade

Princípio referente à garantia de que as informações estão acessíveis *apenas* para aqueles que possuem permissão para acessá-las. Isso envolve medidas como:

- Criptografia;
- Controle de acesso;
- Políticas de privacidade.

### Integridade

Princípio referente à garantia de que as informações são precisas, completas e não foram alteradas de forma *não autorizada*. Isso é alcançado através de técnicas como:

- Controle de versão;
- Assinaturas digitais;
- Checksums.

### Disponibilidade

Princípio referente à garantia de que as informações e recursos estão disponíveis quando necessário. Isso envolve medidas como:

- Prevenção de falhas;
- Redundância;
- Backups;
- Planos de continuidade de negócios.

## Ataques

### Confiabilidade

No tocante a um servidor MQTT, a confiabilidade é garantida pela autenticação e criptografia. No contexto deste simulador, o arquivo de configuração do servidor MQTT possui uma grave vulnerabilidade, já que a autenticação está desabilidade, permitindo que qualquer pessoa acesse o servidor e possa se inscrever ou publicar mensagens nos tópicos em questão.
A fim de habilitar a autenticação, é necessário que as configurações não permitam acessos anônimos e necessitem de uma senha para acessar o servidor. Isso pode ser feito adicionado as seguintes configurações ao arquivo:
<pre><code>allow_anonymous false
password_file /mosquitto/config/passwd</code></pre>

### Integridade

No tocante a um servidor MQTT, a integridade é garantida pela constante verificação e criptografia das mensagens recebidas e publicadas. No contexto deste simulador, garantindo a confiabilidade citada acima, com autenticação, já resolvemos a evidente vulnerabilidade de qualquer pessoa anônima publicar mensagens falsas pelo broker. Para além disso, também podemos verificar que, se alguém conseguir se autenticar, precisamos saber de onde essa pessoa está publicando mensagens, colocando sistemas de logs com informações sobre quem está publicando cada mensagem, garantindo que possamos rastrear e interferir ataques que vêm de "dentro" (com uma pessoa que deveria ser confiável autenticada, por exemplo). Inclusive, a fim de garantir a integridade desses logs, sistemas envolvendo bancos de dados, ferramentas de backups e criptografias com certificados de integridade envolvidas também são desejáveis.

### Disponibilidade

No tocante a um servidor MQTT, a disponibilidade é garantidade pela capacidade de processamento, de rede e, também, pela opção de acesso administrativo para que os dados sejam acessíveis àqueles que deveriam poder acessá-los. No contexto deste simulador, há uma vulnerabilidade onde o servidor MQTT estar exposto com pouca memória e processamento alocado, facilitando um ataque de negação de serviço, já que a sobrecarga do sistema é iminente. Podemos evitar isso colocando configurações de memória e processamento coerentes com a demanda do serviço.