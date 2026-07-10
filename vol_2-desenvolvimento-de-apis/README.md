# Go Essencial — Volume 2: Desenvolvimento de APIs

Este volume reúne a trilha de estudo focada em desenvolvimento de APIs em Go.

O conteúdo está em evolução: capítulos, exemplos e revisões são publicados de
forma incremental.

## Objetivos

- Construir base sólida para desenvolvimento de APIs com Go.
- Evoluir do essencial para padrões de produção.
- Manter um material consultável para estudo e referência.

## Status Atual

Neste repositório, o Volume 2 está em andamento.

- Novos capítulos e seções são adicionados gradualmente.
- Durante a revisão do conteúdo, os códigos são reimplementados e testados.
- A documentação deste volume acompanha o estado publicado em `codigo/`.

## Estrutura Publicada

O conteúdo disponível no momento está em `codigo/`, com os seguintes capítulos
já presentes:

- `cap01/` (seções `s01` a `s06`)
- `cap03/` (práticas `p01` a `p05` e seções `s01` a `s07`)
- `cap04/` (projeto `go-list/`)
- `cap05/` (seção `s01`)
- `cap07/` (seções `s01` a `s03`)

## Organização Planejada do Volume

O roteiro de temas segue esta progressão:

1. HTTP com net/http
2. Roteamento
3. Middleware
4. Serialização com JSON
5. Validação de dados
6. Configuração da aplicação
7. Persistência com banco de dados
8. Autenticação e autorização
9. Upload e download de arquivos
10. Testes de APIs
11. Logging
12. Observabilidade
13. Deploy

## Como Executar Exemplos

Exemplo simples:

```bash
go run codigo/cap01/s01/main.go
```

Exemplo de projeto estruturado:

```bash
go run codigo/cap04/go-list/cmd/*.go
```

## Geração de Material

Use o Makefile da raiz do repositório:

```bash
make pdf VOLUME=2
```

O artefato final é gerado conforme a configuração do projeto principal.