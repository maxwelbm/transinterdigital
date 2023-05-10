# Imagem base com Go 1.19
FROM golang:1.19

# Crie um diretório de trabalho dentro do container
WORKDIR /go/src/app

# Defina a variável de ambiente PATH com o diretório bin do Go
ENV PATH="/go/bin:${PATH}"

# Copie o conteúdo do diretório local para o diretório de trabalho do container
COPY . .

# Execute o comando make build quando o container for iniciado
CMD ["make", "build"]

