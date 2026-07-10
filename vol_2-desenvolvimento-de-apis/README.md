# Go Essencial — Volume 2: Desenvolvimento de APIs

Este volume reúne uma trilha progressiva para construção de APIs HTTP em Go,
partindo de conceitos fundamentais de `net/http` até tópicos de deploy,
observabilidade e operação.

## Objetivos

- Construir base sólida para desenvolvimento de APIs com Go.
- Evoluir do essencial para padrões de produção.
- Manter um material consultável para estudo e referência.

## Estrutura do Volume

O conteúdo principal está em `book/`, organizado por capítulos:

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

## Geração do PDF

Use o Makefile da raiz do repositório:

```bash
make pdf VOLUME=2
```

O PDF final é gerado em `dist/` na raiz do projeto.