# TP : Construire une API complète de gestion de tâches 📝

## Partie 4 : À vous de jouer !

**Durée estimée : 1h30 - 2h**

### Objectifs

- Créer une API complète avec le routeur de votre choix (Go standard, Gin, Chi)
- Manipuler les opérations CRUD (Create, Read, Update, Delete)
- Travailler avec des structures de données en mémoire
- Gérer les erreurs et les statuts HTTP appropriés

### Description de l'API

Vous allez créer une API pour gérer des **tâches** (TODO list). Cette API doit permettre à un utilisateur de :

1. **Créer une tâche** : L'utilisateur peut créer une nouvelle tâche en envoyant un nom et une description.
2. **Récupérer la liste des tâches** : L'utilisateur peut récupérer toutes les tâches créées.
3. **Récupérer une tâche spécifique** : L'utilisateur peut récupérer les détails d'une tâche en particulier à partir de son ID.
4. **Mettre à jour une tâche** : L'utilisateur peut modifier le nom ou la description d'une tâche.
5. **Supprimer une tâche** : L'utilisateur peut supprimer une tâche par son ID.

#### Routes à implémenter :

1. **POST /tasks** : Crée une nouvelle tâche.

   - **Corps de la requête** (JSON) :
     ```json
     {
       "name": "Titre de la tâche",
       "description": "Description de la tâche"
     }
     ```
   - **Réponse** (JSON) :
     ```json
     {
       "id": 1,
       "name": "Titre de la tâche",
       "description": "Description de la tâche"
     }
     ```

2. **GET /tasks** : Récupère la liste de toutes les tâches.

   - **Réponse** (JSON) :
     ```json
     [
       {
         "id": 1,
         "name": "Titre de la tâche",
         "description": "Description de la tâche"
       },
       ...
     ]
     ```

3. **GET /tasks/{id}** : Récupère les détails d'une tâche spécifique par son ID.

   - **Réponse** (JSON) :
     ```json
     {
       "id": 1,
       "name": "Titre de la tâche",
       "description": "Description de la tâche"
     }
     ```

4. **PUT /tasks/{id}** : Met à jour une tâche spécifique.

   - **Corps de la requête** (JSON) :
     ```json
     {
       "name": "Nouveau titre",
       "description": "Nouvelle description"
     }
     ```
   - **Réponse** (JSON) :
     ```json
     {
       "id": 1,
       "name": "Nouveau titre",
       "description": "Nouvelle description"
     }
     ```

5. **DELETE /tasks/{id}** : Supprime une tâche spécifique.
   - **Réponse** : Statut HTTP `204 No Content`.

---

### Étapes

### 1. Choisissez votre routeur 🔄

Vous avez le choix entre :

- **Go standard** (`net/http`)
- **Gin** (framework minimaliste)
- **Chi** (framework léger et flexible)

En fonction du routeur que vous choisissez, référez-vous aux TP précédents pour savoir comment initialiser un routeur et configurer des routes.

### 2. Créez les structures de données 🗂️

Vous allez travailler avec des données en mémoire. Créez une structure Go pour représenter une **tâche** :

```go
type Task struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

Puis, définissez une liste de tâches pour stocker les tâches créées :

```go
var tasks []Task
var nextID = 1
```

- **`tasks`** : Un tableau en mémoire pour stocker les tâches créées.
- **`nextID`** : Un compteur pour assigner un ID unique à chaque nouvelle tâche.

### 3. Implémentez les routes CRUD 🌐

#### a. Route **POST /tasks** (Créer une tâche)

- Récupérez les données envoyées dans le corps de la requête (JSON).
- Assignez un ID unique à la tâche et ajoutez-la à la liste `tasks`.
- Retournez la tâche nouvellement créée avec son ID.

#### b. Route **GET /tasks** (Récupérer toutes les tâches)

- Retournez la liste complète des tâches au format JSON.

#### c. Route **GET /tasks/{id}** (Récupérer une tâche spécifique)

- Récupérez l'ID dans l'URL et cherchez la tâche correspondante dans la liste.
- Si elle existe, renvoyez-la. Sinon, retournez une erreur `404 Not Found`.

#### d. Route **PUT /tasks/{id}** (Mettre à jour une tâche)

- Récupérez l'ID dans l'URL et trouvez la tâche correspondante.
- Modifiez les champs `name` et `description` de la tâche existante.
- Retournez la tâche mise à jour.

#### e. Route **DELETE /tasks/{id}** (Supprimer une tâche)

- Récupérez l'ID dans l'URL et supprimez la tâche correspondante de la liste `tasks`.
- Retournez un statut `204 No Content` si la suppression est réussie.

### 4. Gérer les erreurs et les statuts HTTP 📡

- Assurez-vous que chaque route renvoie le statut HTTP approprié (`200 OK`, `201 Created`, `404 Not Found`, `204 No Content`, etc.).
- En cas de données manquantes ou mal formées dans une requête POST/PUT, renvoyez une réponse `400 Bad Request` avec un message d'erreur approprié.

### 5. Testez l'API avec **Postman** ou **curl** 🛠️

- Testez chaque route une à une pour vous assurer que l'API fonctionne comme prévu.
- Par exemple, pour créer une nouvelle tâche avec **curl** :

```bash
curl -X POST http://localhost:8080/tasks -d '{"name":"Étudier Go", "description":"Apprendre les bases de Go"}' -H "Content-Type: application/json"
```

- Pour récupérer la liste des tâches créées :

```bash
curl http://localhost:8080/tasks
```

---

### 6. Bonus : Améliorez votre API ✨

Si vous terminez rapidement et voulez aller plus loin :

- **Ajoutez une validation** : Assurez-vous que les champs `name` et `description` sont toujours fournis et non vides lors de la création ou la mise à jour d'une tâche.
- **Ajoutez des filtres ou des paramètres de tri** : Permettez à l'utilisateur de trier les tâches par nom ou de filtrer par certains critères (par exemple, récupérer seulement les tâches contenant un certain mot).
- **Utilisez un fichier ou une base de données** : Au lieu de stocker les tâches en mémoire, essayez de les sauvegarder dans un fichier JSON ou une base de données SQLite simplifiée.

---

## Critères d'évaluation

Voici une grille d'évaluation simple pour noter ce TP.

### A. Fonctionnalité de l'API (10 points)

| Critère                                        | Excellent (10)           | Bon (7)       | Satisfaisant (5) | Insuffisant (2) | Non réalisé (0) |
| ---------------------------------------------- | ------------------------ | ------------- | ---------------- | --------------- | --------------- |
| Création de tâche (POST /tasks)                | Fonctionnelle et robuste | Fonctionnelle | Partiellement    | Faible          | Non réalisé     |
| Récupération de toutes les tâches (GET /tasks) | Fonctionnelle et robuste | Fonctionnelle | Partiellement    | Faible          | Non réalisé     |
| Récupération d'une tâche (GET /tasks/{id})     | Fonctionnelle et robuste | Fonctionnelle | Partiellement    | Faible          | Non réalisé     |
| Mise à jour d'une tâche (PUT /tasks/{id})      | Fonctionnelle et robuste | Fonctionnelle | Partiellement    | Faible          | Non réalisé     |
| Suppression d'une tâche (DELETE /tasks/{id})   | Fonctionnelle et robuste | Fonctionnelle | Partiellement    | Faible          | Non réalisé     |

### B. Qualité du code (5 points)

| Critère                         | Excellent (5) | Bon (4) | Satisfaisant (3) | Insuffisant (1) | Non réalisé (0) |
| ------------------------------- | ------------- | ------- | ---------------- | --------------- | --------------- |
| Lisibilité du code              | Très claire   | Claire  | Moyenne          | Faible          | Non réalisé     |
| Utilisation correcte du routeur | Très claire   | Claire  | Moyenne          | Faible          | Non réalisé     |
| Gestion des erreurs             | Très claire   | Claire  | Moyenne          | Faible          | Non réalisée    |

---

### Conclusion 🎓

Bravo, vous êtes maintenant prêts à créer une **API complète** ! Vous avez eu l'opportunité de choisir votre propre framework Go (ou de rester sur la version standard) pour construire un projet qui reflète des scénarios réels d'API CRUD.

Ce TP vous permettra de consolider vos compétences en développement d'API et de vous familiariser avec la gestion des requêtes HTTP et des données en Go.
