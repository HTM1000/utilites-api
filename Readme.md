# 📊 INSS Web Scraper API

Este projeto é uma API escrita em Go que realiza **web scraping** para obter informações das tabelas do INSS diretamente do site [debit.com.br](https://www.debit.com.br/tabelas/tabelas-inss). Os dados coletados são disponibilizados por meio de um endpoint REST em formato JSON.

## 📋 Funcionalidades

- Coleta automática das alíquotas do INSS.
- Expõe os dados através de uma API REST no endpoint `/api/inss`.
- Retorna informações no formato JSON, incluindo:
  - Data da tabela
  - Faixas de contribuição
  - Percentual de alíquota correspondente

---

## 📂 Estrutura do Projeto

```
.
├── main.go              # Ponto de entrada da aplicação
├── handlers/            # Contém os handlers da API
│   └── inss_handler.go  # Handler responsável pelo endpoint /api/inss
├── services/            # Lógica de web scraping
│   └── scrape_inss.go   # Função para coletar os dados das tabelas do INSS
├── entity/              # Definições dos modelos de dados
│   └── inss.go          # Modelo para representar as tabelas do INSS
├── go.mod               # Arquivo de configuração do Go
└── go.sum               # Dependências do projeto
```

---

## 🚀 Como Rodar o Projeto

1. Clone o repositório:
   ```sh
   git clone https://github.com/HTM1000/table-inss.git
   ```

2. Navegue até a pasta do projeto:
   ```sh
   cd table-inss
   ```

3. Instale as dependências:
   ```sh
   go mod tidy
   ```

4. Execute o servidor:
   ```sh
   go run main.go
   ```

5. Acesse o endpoint no navegador ou usando ferramentas como curl ou Postman:
   ```sh
   http://localhost:8080/api/inss
   ```

## 🛠️ Tecnologias Utilizadas

- **Go** — Linguagem principal do projeto
- **Colly** — Biblioteca para web scraping
- **net/http** — Módulo nativo para construir a API REST
- **json** — Para serialização das respostas

## 📤 Exemplo de Resposta da API

```json
[
  {
    "data": "01/01/2025",
    "faixa": "Até R$ 1.302,00",
    "aliquota": "7.5%"
  },
  {
    "data": "01/01/2025",
    "faixa": "De R$ 1.302,01 até R$ 2.571,29",
    "aliquota": "9%"
  }
]
```

## 📈 Melhorias Futuras

- Implementar caching dos resultados para evitar múltiplas requisições.
- Adicionar testes unitários para a função de scraping e o handler.
- Criar uma documentação OpenAPI/Swagger para facilitar o consumo da API.

## ⚠️ Observações

Este projeto foi desenvolvido para fins de aprendizado.

