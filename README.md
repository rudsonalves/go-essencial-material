# Go Essencial - Material Público

Este repositório concentra os materiais públicos do projeto Go Essencial.

No momento, o Volume 1 (Fundamentos) está consolidado e o Volume 2
(Desenvolvimento de APIs) está em processo de adição, revisão e evolução.

## Conteúdo Disponível

- Livro em PDF: `vol_1-fundamentos/livro-golang.pdf`
- Exemplos em Go: `vol_1-fundamentos/codigo/`
- Visão do volume: `vol_1-fundamentos/README.md`
- Estrutura inicial do Volume 2: `vol_2-desenvolvimento-de-apis/`
- Códigos já publicados do Volume 2: `vol_2-desenvolvimento-de-apis/codigo/`

## Estrutura

```text
.
├── README.md
├── LICENSE.md
├── vol_1-fundamentos/
	├── livro-golang.pdf
	├── README.md
	└── codigo/
		├── cap04/
		├── cap05/
		├── ...
		└── cap17/
└── vol_2-desenvolvimento-de-apis/
	├── README.md
	└── codigo/
		├── cap01/
		├── cap03/
		├── cap04/
		├── cap05/
		└── cap07/
```

## Como Usar

Leitura sequencial:

1. Comece pelo PDF em `vol_1-fundamentos/livro-golang.pdf`.
2. Ao longo da leitura, rode os exemplos correspondentes em `vol_1-fundamentos/codigo/`.
3. Para acompanhar o avanço do Volume 2, consulte `vol_2-desenvolvimento-de-apis/README.md`.

Consulta rápida:

1. Abra a pasta do capítulo em `vol_1-fundamentos/codigo/capXX/`.
2. Execute o exemplo desejado com `go run`.

Exemplo:

```bash
go run vol_1-fundamentos/codigo/cap05/p01/p01.go
```

Exemplo (Volume 2):

```bash
go run vol_2-desenvolvimento-de-apis/codigo/cap01/s01/main.go
```

## Requisitos

- Go instalado na máquina (versão recente)
- Terminal com suporte a `go run`

## Licença

Este material usa licença dupla:

- Conteúdo editorial (PDF, textos e materiais didáticos): CC BY-NC-SA 4.0
- Código-fonte dos exemplos Go: MIT

Os detalhes completos estão em `LICENSE.md`.

## Projeto de Origem

Este material é gerado a partir do projeto principal `go_essencial`, com sincronização automatizada via alvo `make material`.
