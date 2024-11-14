# Projet Final : API pour une Clinique Vétérinaire 🐾

## Contexte

Dans ce projet final, vous allez développer une API pour une clinique vétérinaire qui s'occupe principalement des chats. Cette API permettra de gérer les informations des patients (chats), les consultations vétérinaires, et les traitements administrés. Ce projet rassemble toutes les compétences que vous avez acquises dans les TPs précédents : structuration de code, utilisation de modèles et de GORM, création de repositories, et configuration centralisée.

## Objectifs

- Structurer l’API avec des routes organisées par fonctionnalité.
- Utiliser GORM pour les opérations CRUD (Create, Read, Update, Delete) sur les patients, les consultations et les traitements.
- Implémenter une architecture modulaire avec des repositories et une configuration centralisée.

---

## Structure du Projet

Votre projet doit être structuré comme suit :

```
vet-clinic-api/
├── config/              # Configuration générale
│   └── config.go        # Fichier principal de configuration
├── database/            # Configuration de la base de données
│   ├── database.go      # Connexion et migration
│   └── dbmodel/         # Modèles pour GORM et Repositories
│       ├── cat.go       # Modèle et Repository pour les patients (chats)
│       ├── visit.go     # Modèle et Repository pour les consultations
│       └── treatment.go # Modèle et Repository pour les traitements
├── pkg/                 # Code source de l'application
│   ├── cat/             # Fonctionnalité de gestion des patients
│   │   ├── routes.go    # Routes pour les patients
│   │   └── controller.go # Logique de gestion des patients
│   ├── visit/           # Fonctionnalité de gestion des consultations
│   │   ├── routes.go    # Routes pour les consultations
│   │   └── controller.go # Logique de gestion des consultations
│   └── treatment/       # Fonctionnalité de gestion des traitements
│   |   ├── routes.go    # Routes pour les traitements
│   |   └── controller.go # Logique de gestion des traitements
|   └── models/           # Les models JSON de l'application
└── main.go              # Point d'entrée de l'application
```

---

## Consignes Détails par Composant

### 1. Configuration de l'Application

Dans le dossier `config/` :

- Créez un fichier `config.go` qui initialise la connexion à la base de données avec GORM et configure les repositories pour chaque modèle.
- La fonction `New` doit retourner une structure `Config` qui centralise :
  - Les repositories pour les chats, les consultations et les traitements.
  - Les paramètres de connexion et de configuration générale.
- Dans `New`, initialisez la connexion à la base de données en utilisant `gorm.Open`, puis configurez les repositories (en appelant chaque repository en passant la session de base de données).

### 2. Gestion de la Base de Données

Dans le dossier `database/` :

- Créez un fichier `database.go` qui contient la fonction `Migrate` pour gérer la migration des modèles.
  - Cette fonction doit appeler `AutoMigrate` sur chaque modèle (chat, consultation, traitement) afin de créer les tables nécessaires lors du démarrage de l’application.

Dans le dossier `database/dbmodel/` :

- Créez un fichier par modèle de données (`cat.go`, `visit.go`, `treatment.go`), chacun contenant :
  - La structure de données du modèle pour GORM (ex : `Cat`, `Visit`, `Treatment`).
  - Une interface `Repository` pour les méthodes CRUD (ex : `CatRepository`, `VisitRepository`, `TreatmentRepository`).
  - Une structure qui implémente cette interface en interagissant avec GORM (par exemple `catRepository`).
  - Pour les méthodes des repositories, utilisez les méthodes de GORM pour implémenter les fonctions de base (Create, FindByID, FindAll, Update, Delete).

---

### 3. Fonctionnalités et Routes

Dans le dossier `pkg/`, créez un dossier par fonctionnalité (`cat/`, `visit/`, `treatment/`) avec la structure suivante :

1. **Fichier `routes.go`**

   - Définit les routes spécifiques à chaque fonctionnalité.
   - Utilise `chi.Router` pour organiser les endpoints REST pour chaque modèle. Par exemple, dans `cat/routes.go`, créez des routes pour :
     - Créer un nouveau patient chat (POST `/cats`)
     - Récupérer tous les chats (GET `/cats`)
     - Récupérer un chat spécifique par ID (GET `/cats/{id}`)
     - Mettre à jour un chat (PUT `/cats/{id}`)
     - Supprimer un chat (DELETE `/cats/{id}`)

2. **Fichier `controller.go`**
   - Implémente la logique de chaque route en appelant les méthodes du repository correspondant.
   - Utilise le package `render` de `chi` pour structurer les réponses JSON.
   - Chaque route doit utiliser la configuration centralisée, en accédant au repository via `Config`.
   - Le contrôleur peut inclure des fonctionnalités avancées, telles que des filtres dans les requêtes pour récupérer uniquement certains types de consultations ou traitements.

### 4. Fichier principal `main.go`

- Configurez `main.go` pour initialiser la configuration (en appelant `config.New`) et pour configurer les routes principales de l’application.
- Utilisez `chi` pour configurer les routes globales :
  - Les routes pour chaque fonctionnalité doivent être montées sous un préfixe d’API, comme `/api/v1/cats`, `/api/v1/visits`, et `/api/v1/treatments`.
  - Activez le middleware CORS pour permettre les requêtes cross-origin, et configurez `render.SetContentType` pour définir le type de contenu par défaut en JSON.
- Si besoin, utilisez un fichier de configuration (`.env` ou autre) pour gérer les constantes de l'application (port de connexion, paramètres de la base de données, etc.).

---

## Exigences Fonctionnelles

### Gestion des Chats (Patients)

1. **Création d’un chat** : permet de créer un nouvel enregistrement de chat avec des informations comme le nom, l’âge, la race et le poids.
2. **Récupération des chats** : permet de récupérer tous les chats ou un chat spécifique par ID.
3. **Mise à jour d’un chat** : permet de mettre à jour les informations d’un chat existant.
4. **Suppression d’un chat** : permet de supprimer un chat de la base de données.

### Gestion des Consultations

1. **Création d’une consultation** : enregistre une nouvelle consultation pour un chat donné, avec des informations comme la date, le motif et le vétérinaire en charge.
2. **Historique des consultations** : permet de récupérer l'historique complet des consultations pour un chat spécifique.

### Gestion des Traitements

1. **Création d’un traitement** : ajoute un traitement prescrit (ex : un médicament) pour une consultation donnée.
2. **Récupération des traitements** : permet de consulter tous les traitements prescrits pour une consultation.

---

### Challenge Avancé

1. **Historique des soins pour un chat** :
   - Ajoutez une route `/cats/{id}/history` qui récupère l’historique complet des soins pour un chat : consultations et traitements associés.
2. **Filtrage des consultations** :
   - Permettez un filtrage des consultations par date, vétérinaire, ou motif.
3. **Gestion des relations avancées avec GORM** :
   - Pour chaque consultation, chargez également les traitements associés.
