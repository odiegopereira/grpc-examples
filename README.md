# gRPC Examples

Este repositório contém um exemplo de implementação de um servidor e cliente gRPC. Siga as instruções abaixo para clonar o repositório, configurar o ambiente e executar o servidor e cliente.

## Como testar

### Pré-requisitos

- [Go](https://golang.org/doc/install) instalado na máquina.
- [Node.js](https://nodejs.org/en/download/) instalado na máquina.

### Clonar o repositório

```bash
git clone https://github.com/odiegopereira/grpc-examples
cd grpc-examples
```

### Executar o servidor

1. Navegue até o diretório do servidor gRPC:

    ```bash
    cd grpc-server
    ```

2. Execute o servidor:

    ```bash
    go run .
    ```

### Executar o cliente

1. Abra um novo terminal e navegue até o diretório do cliente gRPC:

    ```bash
    cd grpc-client
    ```

2. Instale as dependências do cliente:

    ```bash
    npm install
    ```

3. Execute o cliente para pegar um conselho:

    ```bash
    node grpc-client.js
    ```

### Estrutura do Projeto

O repositório está organizado da seguinte forma:

```
grpc-tests/
├── grpc-server/
│   ├── main.go
│   ├── advice/
│   │   └── advice.proto
│   └── ...
└── grpc-client/
    ├── grpc-client.js
    ├── package.json
    └── ...
```

- **grpc-server/**: Contém o código-fonte do servidor gRPC implementado em Go.
- **grpc-client/**: Contém o código-fonte do cliente gRPC implementado em Node.js.

### Solução de Problemas

- **Erro ao conectar ao servidor**: Verifique se o servidor está em execução e se você está no diretório correto ao executar os comandos.
- **Dependências não instaladas**: Certifique-se de ter executado `npm install` no diretório do cliente.

### Contribuição

Sinta-se à vontade para abrir issues ou pull requests se encontrar problemas ou tiver sugestões de melhorias. 

### Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

Com estas instruções, você deve ser capaz de configurar e executar o projeto sem problemas. Se precisar de mais informações, não hesite em consultar a documentação oficial das ferramentas utilizadas ou abrir uma issue no repositório.