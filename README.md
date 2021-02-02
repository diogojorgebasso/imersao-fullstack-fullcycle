# CodePIX - transação instantânea de capital

> Código realizado com 💖 e tutorado por Wesley Willians, participe gratuitamente em: https://imersao.fullcycle.com.br/

![image](https://user-images.githubusercontent.com/65865529/106500126-37e17700-64a0-11eb-866b-b80fa1223d9c.png)

## Principais desafios do projeto:

- comunicação rápida e eficiente;
- criação e consulta instantânea das chaves (Síncrona);
- garantia de que nenhuma transação seja perdida;

## Stacks:

- gRPC: framework de desenvolvimento que tem por objetivo facilitar o processo de comunicação entre sistemas de uma forma rápida, leve e independente de linguagem, também muito utilizado para microserviços, utilizando protocol buffers e HTTP 2.0;
- Apache Kafka: processamento poderoso de dados, tal como Github implementa;
- Nextjs no FrontEnd e Nodejs no backEnd;
- Docker e microserviços;
- Postgres;
- PgAdmin;
- ZooKeeper.

### Como o software interage com o as stacks:

- será capaz de atuar como um servidor gRPC;
- consumir e publicar mensagens no apache Kafka;
- Ambas operações devem ser realizadas de forma simultânea ao executar o serviço;
- Trabalhar com DDD (Domain Driven Domain)

A aplicação é flexível para a implementação de outros formatos: REST, graphQL e afins, mas utilizaremos como Design Pattern...

### Design de software

Arquitetura Hexagonal / Ports and Adapters.

### Estrutura e camadas das pastas

![image](https://user-images.githubusercontent.com/65865529/106503595-a0caee00-64a4-11eb-9b76-217ce3b8c385.png)

### Licença

Este repo está licenciado sob [MIT](LICENSE)
