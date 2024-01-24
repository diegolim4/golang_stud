# Principais recursos do Go

- Linguagem de programação compilada e estaticamente digitada.
- Suporte de simultaneidade e coleta de lixo.
- Biblioteca e conjunto de ferramentas robustos.
- Multiprocessamento e rede de alto desempenho.
- Conhecido pela legibilidade e usabilidade (como Python).

# Instalando no Linux

    1 - Baixa o pacote go
        - wget https://golang.org/dl/go1.17.linux-amd64.tar.gz

    2 - Descompacte o arquivo
        - sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz

    3 - Configure as variáveis de ambiente
        - export GOROOT=/usr/local/go
        - export GOPATH=$HOME/go
        - export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

    4 - Atualizando sua var de ambiente
        - source ~/.profile

    5 - Verifique a instalação
        - go version

# Desistalando

    1 - Remova o diretório de instalação
        - sudo rm -rf /usr/local/go

    2 - Remova as entradas das variáveis de ambiente
        - export GOROOT=/usr/local/go
        - export GOPATH=$HOME/go
        - export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

    3 - Atualize as variáveis de ambiente
        - source ~/.profile

#### Link doc

- https://go.dev/doc/

#### Go Workspace

Por padrão essa pasta deve ser chamada go e deve ficar na raiz do nosso usuário.

O Go Workspace deve possuir três pastas. A primeira delas é a pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo de pacotes que vamos importando no nosso código.Também deve ter a pasta src, onde escreveremos o código fonte de cada aplicação, e a pasta bin, onde ficará os compilados do nosso código Go.

#### Executando um programa em Go

- go build name-file.go
    Com isso nosso código sera compilado e será criado um executável

- go run name-file.go
    Já compila e execulta
