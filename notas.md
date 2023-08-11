# Principais recursos do Go

- Linguagem de programação compilada e estaticamente digitada.
- Suporte de simultaneidade e coleta de lixo.
- Biblioteca e conjunto de ferramentas robustos.
- Multiprocessamento e rede de alto desempenho.
- Conhecido pela legibilidade e usabilidade (como Python).

# Instalando no linux
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

