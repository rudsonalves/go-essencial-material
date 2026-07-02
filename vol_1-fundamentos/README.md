# Go Essencial — Volume 1: Fundamentos

Este projeto reúne uma revisão dos fundamentos da linguagem Go organizada em formato de livro.

A proposta principal não é seguir a estrutura tradicional de um curso, mas construir um material de estudo progressivo, coerente e consultável. O livro parte do contexto histórico e da filosofia da linguagem para, aos poucos, revisar os principais recursos de Go com foco em entendimento sólido, escrita idiomática e aplicação prática.

No futuro, este material poderá servir como base para cursos, aulas ou trilhas de estudo. Neste momento, porém, o objetivo central é consolidar o conteúdo em forma de livro.

## Objetivos

- Revisar os fundamentos de Go de forma conceitual e prática.
- Explicar as decisões da linguagem antes de apresentar apenas a sintaxe.
- Construir um material que funcione tanto como leitura sequencial quanto como referência.
- Manter uma base organizada para possíveis cursos derivados no futuro.

## Estrutura do Livro

O conteúdo principal está no diretório `book/`.

Capítulos planejados:

1. Introdução
2. O Ambiente de Desenvolvimento
3. Variáveis e Tipos
4. Controle de Fluxo
5. Coleções
6. Funções
7. Ponteiros
8. Structs
9. Métodos
10. Interfaces
11. Erros
12. Packages
13. Módulos
14. Biblioteca Padrão
15. Concorrência
16. Testes
17. Interfaces Avançadas

Cada capítulo deve combinar explicação conceitual, exemplos de código, observações sobre uso idiomático e pontos de atenção comuns.

### Convenção de Arquivos

Cada pasta dentro de `book/` representa um capítulo. Os arquivos de seção usam o formato `XXxx-nome.md`, em que `XX` representa o número do capítulo e `xx` representa o número da seção dentro do capítulo.

Exemplos:

```text
0100-intro.md
1004-interfaces-na0biblioteca-padrao.md
1410-desafios.md
```

Essa convenção evita ambiguidades na ordenação do PDF quando o livro passa de 9 capítulos ou quando um capítulo passa de 9 seções.

## Organização do Projeto

```text
.
├── book/                 # Conteúdo do livro em Markdown
├── dist/                 # Arquivos gerados
├── tools/pandoc/         # Filtros e ajustes para geração do PDF
└── README.md             # Visão geral do projeto
```

## Geração do PDF

O projeto utiliza Pandoc e XeLaTeX para gerar o livro em PDF. Os comandos de build ficam no `Makefile` da raiz do repositório.

Para gerar o PDF:

```bash
make pdf
```

O arquivo final será criado em:

```text
dist/vol_1-fundamentos.pdf
```

Quando este volume é sincronizado para o repositório público de material, o PDF é publicado como:

```text
vol_1-fundamentos/vol_1-fundamentos.pdf
```

Para limpar os arquivos gerados:

```bash
make clean
```

## Licença

Este volume usa licença dupla:

- O conteúdo textual do livro, imagens, diagramas, arquivos Markdown, PDFs e demais materiais editoriais são licenciados sob Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International (CC BY-NC-SA 4.0).
- Os exemplos de código em Go localizados em `codigo/` são licenciados sob a licença MIT.

Veja os detalhes no arquivo `../LICENSE.md`.

## Programação dos Volumes

Este volume funciona como base conceitual para uma série maior. A programação inicial dos volumes é:

1. **Go Essencial — Volume 1: Fundamentos**: sintaxe, tipos, controle de fluxo, coleções, funções, ponteiros, structs, métodos, interfaces, erros, packages, módulos, biblioteca padrão, concorrência e testes.
2. **Go Essencial — Volume 2: Desenvolvimento de APIs**: HTTP, JSON, context, middleware, banco de dados, autenticação, arquitetura, logging, observabilidade e deploy.
3. **Go Essencial — Volume 3: Concorrência e Sistemas**: concorrência avançada, performance, profiling, redes, sistemas distribuídos, gRPC, filas e processamento paralelo.
4. **Go Essencial — Volume 4: Engenharia de Software em Go**: Clean Architecture, DDD, testes avançados, segurança, ferramentas e projetos completos.

A escolha do nome **Go Essencial** busca preservar longevidade para a obra: ele continua fazendo sentido tanto como estudo inicial quanto como uma coleção madura e composta por vários volumes.

## Ideia Central

A pergunta que orienta este projeto é:

> Como compreender Go a partir das ideias que moldaram a linguagem?

Por isso, a revisão não se limita a listar recursos. O objetivo é entender por que Go favorece simplicidade, composição, ferramentas integradas, tratamento explícito de erros e concorrência como parte do modelo da linguagem.
