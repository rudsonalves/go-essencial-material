# Go Essencial — Volume 2: Desenvolvimento de APIs

Este volume reúne uma trilha progressiva para construção de APIs HTTP em Go,
partindo de conceitos fundamentais de `net/http` até tópicos de deploy,
observabilidade e operação.

## Objetivos

- Construir base sólida para desenvolvimento de APIs com Go.
- Evoluir do essencial para padrões de produção.
- Manter um material consultável para estudo e referência.

## Estrutura do Volume

O conteúdo principal está em `book/`, organizado por capítulos. Os códigos executáveis que acompanham o texto ficam em `codigo/`.

Há dois tipos de código de acompanhamento:

- `codigo/capXX/sNN/`: aplicação pequena e autocontida que acompanha uma seção e pode ser usada como semente para experimentos;
- `codigo/capXX/go-list/`: estado acumulado da GoList depois das funcionalidades desenvolvidas naquele capítulo.

Os diretórios `s01`, `s02` e seguintes também ajudam a verificar se o código apresentado no texto compila e se comporta como descrito. Já os diretórios `go-list` preservam a evolução do projeto por capítulo, permitindo comparar cada etapa sem perder as versões anteriores.

Exemplos:

```text
codigo/cap01/s01/
codigo/cap01/s02/
codigo/cap04/go-list/
```

No material publicado, consulte como referência o [código de `cap01/s01`](https://github.com/rudsonalves/go-essencial-material/tree/main/vol_2-desenvolvimento-de-apis/codigo/cap01/s01) e a [GoList do capítulo 4](https://github.com/rudsonalves/go-essencial-material/tree/main/vol_2-desenvolvimento-de-apis/codigo/cap04/go-list).

Os capítulos abordam:

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
