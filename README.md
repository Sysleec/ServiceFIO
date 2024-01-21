# ServiceFIO

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/6295/6295417.png" width="100" />
</p>
<p align="center">
    <h1 align="center">SERVICEFIO</h1>
</p>
<p align="center">
    
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/Sysleec/ServiceFIO?style=flat&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/Sysleec/ServiceFIO?style=flat&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/Sysleec/ServiceFIO?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/Sysleec/ServiceFIO?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
		<em>Developed with the software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/GNU%20Bash-4EAA25.svg?style=flat&logo=GNU-Bash&logoColor=white" alt="GNU%20Bash">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
</p>
<hr>

## 📍 Overview

# HTTP Handlers for People Management

## Create People
- **Method:** POST
- **URL:** `http://45.8.97.234:8080/v1/people`
- **Body:**
  ```json
  {
      "Name": "Nike",
      "Surname": "Balatov",
      "Patronymic": "Jerokovich" // optional
  }
  ```

## Get People
- **Method:** GET
- **URL:** `http://45.8.97.234:8080/v1/people`
- **Optional Filters:** page, limit, filter
- **Example:** `http://45.8.97.234:8080/v1/people?page=1&limit=10&filter=Maks`

## Delete People
- **Method:** DELETE
- **URL:** `http://45.8.97.234:8080/v1/people/{id}`

## Update People
- **Method:** PATCH
- **URL:** `http://45.8.97.234:8080/v1/people/{id}`
- **Body:**
  ```json
  {
      "Name": "Nike", // optional
      "Surname": "Balatov", // optional
      "Patronymic": "Jerokovich" // optional
  }
  ```

---

## 📂 Repository Structure

```sh
└── ServiceFIO/
    ├── .env
    ├── .golangci.pipeline.yaml
    ├── Makefile
    ├── cmd
    │   └── main.go
    ├── docker-compose.yaml
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── api
    │   │   └── people
    │   ├── converter
    │   │   └── people.go
    │   ├── model
    │   │   └── people.go
    │   ├── repository
    │   │   ├── people
    │   │   └── repository.go
    │   ├── service
    │   │   ├── enrich
    │   │   ├── people
    │   │   └── service.go
    │   └── utils
    │       └── json.go
    ├── migration.Dockerfile
    ├── migration.sh
    ├── sql
    │   ├── queries
    │   │   └── people.sql
    │   └── schema
    │       └── 001_people.sql
    └── sqlc.yaml
```

---

## 🚀 Getting Started

***Requirements***

Ensure you have the following dependencies installed on your system:

* **Go**: `version 1.21.6`

### ⚙️ Installation

1. Clone the ServiceFIO repository:

```sh
git clone https://github.com/Sysleec/ServiceFIO
```

2. Change to the project directory:

```sh
cd ServiceFIO
```

3. Install the dependencies:

```sh
go build -o myapp
```

### 🤖 Running ServiceFIO

Use the following command to run ServiceFIO:

```sh
./myapp
```

### 🧪 Tests

To execute tests, run:

```sh
go test
```

---

## 🛠 Project Roadmap

- [X] `► INSERT-TASK-1`
- [ ] `► INSERT-TASK-2`
- [ ] `► ...`

---

## 🤝 Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Submit Pull Requests](https://github/Sysleec/ServiceFIO/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github/Sysleec/ServiceFIO/discussions)**: Share your insights, provide feedback, or ask questions.
- **[Report Issues](https://github/Sysleec/ServiceFIO/issues)**: Submit bugs found or log feature requests for Servicefio.

<details closed>
    <summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your GitHub account.
2. **Clone Locally**: Clone the forked repository to your local machine using a Git client.
   ```sh
   git clone https://github.com/Sysleec/ServiceFIO
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to GitHub**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.

Once your PR is reviewed and approved, it will be merged into the main branch.

</details>

---

## 👏 Acknowledgments

- List any resources, contributors, inspiration, etc. here.

[**Return**](#-quick-links)

---