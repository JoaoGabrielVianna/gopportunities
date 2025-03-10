🚀 Gopportunities API

<p align="center">
  <img src="./assets/GopportunitiesHeader.svg" alt="GoJob Header">
</p>

🎉 Bem-vindo ao Gopportunities, uma API moderna para oportunidades de emprego construída com Golang!

## 🌟 Funcionalidades

- **Golang**: Aproveite o poder e a eficiência de uma das linguagens de programação mais bem remuneradas do mercado. 💻⚡
- **Go-Gin**: Gerencie rotas de forma eficaz com este roteador rápido e minimalista. 🔥
- **SQLite**: Utilize um banco de dados leve e confiável para armazenar suas informações. 📦
- **GORM**: Simplifique a comunicação com o banco de dados usando este ORM robusto. 🛠️
- **Swagger**: Documente e teste sua API de maneira intuitiva e interativa. 📖✨

## 🛠️ Estrutura do Projeto

🗂️ O projeto segue uma estrutura de pacotes moderna para manter o código organizado e de fácil manutenção.


## 🚀 Como Começar

1. **Clone o repositório**:

    ```bash
    git clone https://github.com/arthur404dev/gopportunities.git

2. **Instale as dependências**:

    Certifique-se de que o Golang esteja instalado em sua máquina. Em seguida, navegue até o diretório do projeto e execute:

    ```bash
    go mod download

1. Execute a aplicação:

    ```bash
    go run main.go

## 🔍 Documentação da API

A documentação interativa da API está disponível através do Swagger. Após iniciar a aplicação, acesse:
    
```bash
http://localhost:8080/swagger/index.html
```    

## 🚧 Dificuldades

- Quando fui gerar os schemas do SQLite no Windows, tive que lidar com C e isso exigiu o MinGW. 😅 No começo, foi complicado porque precisei configurar o PATH e instalar o MSYS2 para conseguir o make. Sem isso, não dava pra compilar o que eu precisava. 😤 
- Acabei tendo que ajustar a arquitetura do MinGW também (32-bit ou 64-bit) e, depois de algumas tentativas, deu certo. Foi um processo meio frustrante, mas no fim, aprendi bastante sobre as ferramentas e como superar esses obstáculos no Windows. 🙌

Se você passar por isso, não desista! Com paciência, tudo se resolve. 💪


## 🔧 Baseado na Aula do [Arthur404dev](https://github.com/arthur404dev): Criando minha Primeira API com Go! 🚀

Agradeço demais pelo conteúdo incrível que ele compartilhou. Foi uma ótima introdução ao desenvolvimento de APIs com Golang, e me ajudou a superar desafios como configuração do MinGW no Windows e a integração com o SQLite. 💻💡

Estou animado com os aprendizados e pronto para o próximo desafio! 💪
