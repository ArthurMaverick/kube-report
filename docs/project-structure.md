## Application structure
```
kube-report/
├── cmd/                         # Entrypoints para o aplicativo (CLI, main files)
│   ├── kube-report/             # Aplicação principal
│   │   └── main.go
│   └── kube-report-cli/         # Interface de linha de comando, se aplicável
│       └── main.go
├── pkg/                         # Código reutilizável (bibliotecas, componentes)
│   ├── client/                  # Conexões e operações com Kubernetes
│   │   ├── client.go
│   │   ├── interface.go
│   │   └── mocks/               # Mocks para testes
│   ├── report/                  # Lógica de criação de relatórios
│   │   ├── generator.go         # Componente para gerar relatórios
│   │   ├── output/              # Tipos de saída para relatórios
│   │   │   ├── pdf.go
│   │   │   ├── csv.go
│   │   │   └── json.go
│   │   ├── formatters/          # Formatadores de dados para saída
│   │   │   ├── format_csv.go
│   │   │   ├── format_pdf.go
│   │   │   └── format_json.go
│   ├── config/                  # Carregamento e parsing das configurações
│   │   └── config.go
│   └── utils/                   # Utilitários e helpers gerais
│       └── file_helpers.go
├── api/                         # Interfaces de API (ex: REST ou gRPC)
│   ├── v1/                      # Versão da API
│   │   ├── handlers/            # Handlers da API
│   │   │   ├── report_handler.go
│   │   │   └── health_handler.go
│   │   └── routes.go            # Configuração de rotas
├── internal/                    # Implementações internas (não exportadas)
│   └── services/                # Serviços internos
│       ├── report_service.go    # Lógica de negócios para relatórios
│       └── health_service.go    # Serviço de monitoramento de saúde
├── tests/                       # Testes unitários e de integração
│   ├── integration/             # Testes de integração
│   └── unit/                    # Testes unitários
├── docs/                        # Documentação (arquitetura, explicação de uso)
│   └── design.md
├── scripts/                     # Scripts para CI/CD, build, etc.
│   ├── build.sh
│   └── test.sh
└── Makefile                     # Configuração de build
```

## Folders
- `cmd/`: Application entry files, such as main.go, that initialize the service and CLI commands.
- `pkg/`: Reusable modules and components that implement specific functionalities.
- `client/`: Configurations and connections with Kubernetes.
- `report/`: Logic for creating and formatting reports, as well as generating files in different formats.
- `config/`: Loading and parsing of configurations.
- `utils/`: General utilities, such as helpers for file handling.
- `api/`: API interfaces to expose reports via HTTP, if the service requires an API.
- `internal/`: Contains business-specific internal implementations.
- `tests/`: Structure for organizing unit and integration tests.
- `docs/`: Documentation files, such as system design and architecture.