# 🌟 Fiche Théorique : Bien Structurer son API en Go 🚀

Salut ! Voici une fiche de rappel sur **comment organiser proprement le code de ton API** en Go pour que ton projet reste clair, maintenable et scalable. 🎯 On va revoir **les bonnes pratiques**, **une structure typique de projet**, et des conseils pour bien démarrer. C'est parti ! 🎉

---

## 🗂️ Pourquoi structurer son API ?

Une bonne structure de projet permet de :

1. **Faciliter la collaboration** 👥 : Tout le monde sait où chercher le code.
2. **Simplifier la maintenance** 🔧 : Corriger un bug ou ajouter une fonctionnalité devient facile.
3. **Encourager les bonnes pratiques** ✅ : Séparer les responsabilités et éviter les dépendances inutiles.

---

## 🏗️ Structure typique d'un projet Go API

Voici une structure **recommandée** pour ton projet API. Elle est pensée pour les API RESTful et s'appuie sur des principes comme la **séparation des responsabilités** et la **modularité**.

### Arborescence

```
mon-api/
├── config/              # Configuration générale
│   └── config.go        # Fichier principal de configuration
├── database/            # Gestion de la base de données
│   ├── database.go      # Connexion et migrations
│   └── dbmodel/         # Modèles pour GORM et Repositories
│       ├── model1.go    # Modèle et Repository pour une table
│       ├── model2.go    # Modèle et Repository pour une autre table
├── pkg/                 # Code source de l'application
│   ├── feature1/        # Fonctionnalité 1
│   │   ├── routes.go    # Définition des routes
│   │   └── controller.go # Logique métier
│   ├── feature2/        # Fonctionnalité 2
│   │   ├── routes.go    # Définition des routes
│   │   └── controller.go # Logique métier
│   └── authentication/  # Gestion de l'authentification
│       ├── routes.go    # Routes pour l'authentification
│       ├── controller.go # Logique d'authentification
│       ├── jwt.go       # Gestion des tokens JWT
│       └── middleware.go # Middleware pour protéger les routes
│   └── models/          # Modèles JSON échangés par l'API
└── main.go              # Point d'entrée de l'application
```

---

## 📦 Description des dossiers

### `config/` 📄

- **Objectif** : Centraliser la configuration de l'application (ex : variables d'environnement, ports, configurations externes).
- **Contenu** :
  - `config.go` : Fichier principal qui charge les configurations.

### `database/` 🛢️

- **Objectif** : Gérer les connexions, migrations, et interactions avec la base de données.
- **Contenu** :
  - `database.go` : Initialise la connexion avec la DB.
  - `dbmodel/` : Contient les modèles et repositories pour interagir avec la DB.

### `pkg/` 📂

- **Objectif** : Contenir toutes les fonctionnalités de l'application sous forme de modules indépendants.

#### Sous-dossiers :

1. **`featureX/`** : Chaque fonctionnalité a son propre module avec :
   - `routes.go` : Définit les endpoints de l'API.
   - `controller.go` : Contient la logique métier.
2. **`authentication/`** : Gestion de l'authentification avec :

   - `jwt.go` : Gestion des JWT.
   - `middleware.go` : Middleware pour valider les tokens et sécuriser les routes.

3. **`models/`** : Définit les modèles JSON échangés entre le client et le serveur.

### `main.go` 🚀

- **Objectif** : Point d'entrée de l'application.
- **Contenu** : Initialise le serveur, charge les configurations, et démarre les routes.

---

## 🛠️ Bonnes pratiques pour structurer ton API

1. **Séparer les responsabilités** 🧹 : Chaque fichier ou dossier a une fonction bien définie.
2. **Modularité** 🧩 : Les fonctionnalités doivent être indépendantes les unes des autres.
3. **Nommage clair** ✍️ : Les noms des fichiers et des dossiers doivent refléter leur contenu.
4. **Respecter le principe DRY (Don't Repeat Yourself)** 🚫 : Évite de dupliquer du code.
5. **Documenter le code** 📘 : Ajoute des commentaires pour expliquer les parties complexes.

---

## 🔑 Exemple de modèle JSON et base de données

### Modèle JSON

```go
package model

import (
	"errors"
	"net/http"
)

type Company struct {
	ID                uint     `json:"id"`
	Name              string   `json:"name"`
	Description       *string  `json:"description"`
	Website           string   `json:"website"`
	Email             *string  `json:"email"`
	CyberSolutions    []string `json:"cyber_solutions"`
	CyberExpertises   []string `json:"cyber_expertises"`
}
```

### Modèle Base de Données

```go
package dbmodel

import (
	"strings"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name              string
	Description       *string
	Website           string
	Email             *string
	CyberSolutions    string
	CyberExpertises   string
}

func (company *Company) ToModel() *model.Company {
	return &model.Company{
		ID:             company.ID,
		Name:           company.Name,
		Description:    company.Description,
		Website:        company.Website,
		Email:          company.Email,
		CyberSolutions: strings.Split(company.CyberSolutions, ","),
		CyberExpertises: strings.Split(company.CyberExpertises, ","),
	}
}
```

---

## 🎯 En résumé

Structurer ton API, c'est comme organiser ton bureau avant de commencer à travailler. Une bonne organisation rend tout plus simple ! En suivant cette structure, tu seras sûr·e de créer des applications **scalables**, **claires**, et faciles à maintenir. Prêt·e à coder ? 🚀
