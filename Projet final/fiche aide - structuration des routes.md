# 🌟 Fiche Théorique : Structurer ses Routes d'API RESTful 🚦

L’organisation des **routes d’API** est essentielle pour rendre ton application **compréhensible**, **scalable** et **facile à maintenir**. Dans cette fiche, nous allons voir **les bonnes pratiques**, **les conventions à suivre** et un exemple concret pour bien structurer tes routes d’API. Let's go ! 🚀

---

## 🛤️ Qu’est-ce qu’une route d’API ?

Une **route d’API** permet de mapper une requête HTTP (GET, POST, PUT, DELETE, etc.) à une fonctionnalité de ton application.  
Par exemple :

- `GET /api/v1/companies` : Récupérer toutes les entreprises.
- `POST /api/v1/companies` : Créer une nouvelle entreprise.
- `PUT /api/v1/companies/:id` : Mettre à jour une entreprise existante.
- `DELETE /api/v1/companies/:id` : Supprimer une entreprise.

---

## 🧩 Bonnes pratiques pour structurer ses routes

### 1. **Respecter les conventions RESTful** 🌍

REST est une architecture basée sur des **ressources**. Chaque ressource doit avoir une route **claire** et **cohérente**.

| Méthode HTTP | Action    | Exemple             |
| ------------ | --------- | ------------------- |
| `GET`        | Lire      | `GET /users`        |
| `POST`       | Créer     | `POST /users`       |
| `PUT`        | Modifier  | `PUT /users/:id`    |
| `DELETE`     | Supprimer | `DELETE /users/:id` |

**Conseil** : Utilise des **noms de ressources au pluriel** (ex : `users`, `companies`) et évite les verbes dans les noms des routes (ex : pas de `GET /getUser`).

---

### 2. **Versionner tes routes** 📦

Pour éviter les conflits ou problèmes lors des mises à jour, ajoute un **versionnage** dans tes URLs.

Exemple :

- `GET /api/v1/companies` (version 1)
- `GET /api/v2/companies` (version 2)

---

### 3. **Utiliser des identifiants clairs** 🆔

Lorsque tu fais référence à une ressource spécifique, utilise des **identifiants uniques** dans l’URL.

Exemple :

- `GET /companies/:id` : Récupérer une entreprise par son ID.
- `PUT /companies/:id` : Mettre à jour une entreprise spécifique.

---

### 4. **Gérer les relations entre ressources** 🤝

Pour représenter les relations, structure tes URLs avec des **niveaux hiérarchiques**.

Exemple :

- `GET /companies/:id/employees` : Récupérer les employés d’une entreprise.
- `POST /companies/:id/employees` : Ajouter un employé à une entreprise.

---

### 5. **Eviter les routes trop profondes** 🌊

Limite la profondeur de tes routes à **2 ou 3 niveaux maximum** pour éviter la complexité.

🚫 Mauvais :

```
GET /companies/:id/departments/:depId/employees/:empId
```

✅ Bon :

```
GET /employees/:id
```

---

### 6. **Séparer les routes publiques et privées** 🔒

- **Routes publiques** : Pas besoin d’authentification (ex : `GET /public/products`).
- **Routes protégées** : Nécessitent un token ou une session valide (ex : `POST /orders`).

Utilise des middlewares pour appliquer cette séparation !

---

## 🏗️ Structurer des routes relationnelles

Parfois, tu dois gérer des relations : voici comment le faire correctement.

### Exemple avec des employés dans des entreprises

Routes :

- `GET /companies/:id/employees` : Liste des employés d’une entreprise.
- `POST /companies/:id/employees` : Ajouter un employé à une entreprise.

---

## 🎯 En résumé

1. **Organise tes routes par fonctionnalité** : Grouper les endpoints similaires.
2. **Respecte les conventions RESTful** : Utilise les bonnes méthodes HTTP.
3. **Versionne ton API** : Facilite l’évolution de tes services.
4. **Gère les relations hiérarchiquement** : Evite les routes trop profondes.
5. **Sépare routes publiques et protégées** : Utilise des middlewares.

Avec cette structure, ton API sera **claire**, **efficace**, et **facile à maintenir**. Prêt·e à construire des routes comme un·e pro ? 🚀
