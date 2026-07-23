# golang-api

> GitHub template for golang-api projects.

## Usage

Click **"Use this template"** to create a new repository.

Then clone and start coding:

```bash
git clone https://github.com/guilhermelinosp/golang-api-template.git my-project
cd my-project
```

## CI/CD

| Workflow | Trigger | Description |
|---|---|---|
| `pipeline` | Push to `main` | Release → build → push |
| `pr-check` | Pull request | Lint, merge check, secrets scan, label |

Uses reusable workflows from [ci-templates](https://github.com/guilhermelinosp/ci-templates).

## License

[Apache 2.0](LICENSE)
