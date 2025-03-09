.PHONY: default run build test docs clean

# Variáveis
APP_NAME = gopportunities

# Tarefas

# Tarefa padrão, chamada quando não é especificada nenhuma tarefa
default: run

# Rodar o código Go
run:
	@go run main.go

# Compilar o código Go e gerar o binário
build:
	@go build -o $(APP_NAME) main.go

# Rodar os testes
test:
	@go test ./...

# Gerar a documentação usando swag
docs:
	@swag init

# Limpar os arquivos binários e a documentação gerada
clean:
	@rm -rf $(APP_NAME)
	@rm -rf docs
