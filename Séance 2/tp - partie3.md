# TP : Introduction à GORM pour la Gestion de la Base de Données 🐘

## Partie 3 : Introduction à GORM

**Durée estimée :** 1 heure 30 minutes

### Objectifs

Dans cette partie, vous allez découvrir GORM, un ORM (Object Relational Mapping) pour Go, qui simplifie l’interaction avec les bases de données. Vous allez apprendre à configurer GORM, à définir des modèles, et à effectuer des opérations de base : création, lecture, mise à jour et suppression (CRUD). Nous allons nous appuyer sur la feature `agecalculator` pour stocker les âges calculés dans une base de données SQLite.

### Préparation

1. **Installation de GORM et SQLite**

   - Installez GORM et son pilote pour SQLite avec la commande suivante :
     ```bash
     go get -u gorm.io/gorm
     go get -u gorm.io/driver/sqlite
     ```

2. **Création de la structure de base pour la base de données**
   - Dans le dossier `database/`, créez un fichier `db.go` pour la configuration de la connexion à la base de données.
   - Assurez-vous que le dossier `database/` est structuré comme suit :
     ```
     animal-api/
     ├── database/
     │   ├── db.go           # Configuration de la base de données
     │   └── dbmodel/        # Modèles pour GORM
     ```

---

### Étapes

### 1. Configuration de la Base de Données avec GORM

1. **Configuration de la connexion SQLite**

   - Dans `database/db.go`, configurez une connexion à une base de données SQLite. Si le fichier de base de données n’existe pas, il sera automatiquement créé par SQLite.

     Exemple de code pour `database/db.go` :

     ```go
     package database

     import (
         "gorm.io/driver/sqlite"
         "gorm.io/gorm"
         "log"
     )

     var DB *gorm.DB

     func InitDatabase() {
         var err error
         DB, err = gorm.Open(sqlite.Open("animal_api.db"), &gorm.Config{})
         if err != nil {
             log.Fatal("Failed to connect to database:", err)
         }
         log.Println("Database connected")
     }
     ```

2. **Appel de l'initialisation dans `main.go`**

   - Dans `main.go`, importez le package `database` et appelez `database.InitDatabase()` dans le `main` pour initialiser la connexion à la base de données lorsque le serveur démarre.

     ```go
     package main

     import (
         "log"
         "net/http"
         "github.com/go-chi/chi/v5"
         "github.com/yourusername/animal-api/database"
         "github.com/yourusername/animal-api/pkg/agecalculator"
         "github.com/yourusername/animal-api/pkg/soundidentifier"
     )

     func main() {
         database.InitDatabase()

         r := chi.NewRouter()
         r.Mount("/api/v1/agecalculator", agecalculator.Routes())
         r.Mount("/api/v1/soundidentifier", soundidentifier.Routes())

         log.Println("Serving on :8080")
         http.ListenAndServe(":8080", r)
     }
     ```

---

### 2. Création d'un Modèle GORM : `AgeEntry`

1. **Définition du modèle `AgeEntry`**

   - Dans le dossier `database/dbmodel/`, créez un fichier `age_entry.go` pour définir le modèle `AgeEntry`. Ce modèle contiendra l’âge humain et l’âge en années de chat, ainsi qu’un champ `ID` auto-généré.

     Exemple de code pour `database/dbmodel/age_entry.go` :

     ```go
     package dbmodel

     import "gorm.io/gorm"

     type AgeEntry struct {
         gorm.Model
         HumanAge int `json:"human_age"`
         CatAge   int `json:"cat_age"`
     }
     ```

2. **Création automatique de la table**

   - Modifiez `database.InitDatabase()` pour faire un auto-migration de `AgeEntry`. Cela créera automatiquement la table correspondante si elle n'existe pas.

     Mise à jour de `database/db.go` :

     ```go
     package database

     import (
         "log"
         "gorm.io/driver/sqlite"
         "gorm.io/gorm"
         "github.com/yourusername/animal-api/database/dbmodel"
     )

     var DB *gorm.DB

     func InitDatabase() {
         var err error
         DB, err = gorm.Open(sqlite.Open("animal_api.db"), &gorm.Config{})
         if err != nil {
             log.Fatal("Failed to connect to database:", err)
         }

         DB.AutoMigrate(&dbmodel.AgeEntry{})
         log.Println("Database connected and migrated")
     }
     ```

---

### 3. Modification du Contrôleur `agecalculator` pour Utiliser GORM

1. **Enregistrement des âges dans la base de données**

   - Modifiez le contrôleur `AgeInCatYearsHandler` pour sauvegarder chaque calcul d’âge dans la base de données en utilisant GORM.

     Exemple de code pour `pkg/agecalculator/controller.go` :

     ```go
     package agecalculator

     import (
         "net/http"
         "github.com/go-chi/render"
         "github.com/yourusername/animal-api/pkg/model"
         "github.com/yourusername/animal-api/database"
         "github.com/yourusername/animal-api/database/dbmodel"
     )

     func AgeInCatYearsHandler(w http.ResponseWriter, r *http.Request) {
         req := &model.AgeRequest{}
         if err := render.Bind(r, req); err != nil {
             render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
             return
         }

         catAge := calculateCatAge(req.HumanAge)
         res := &model.AgeResponse{CatAge: catAge}

         // Sauvegarde dans la base de données
         ageEntry := dbmodel.AgeEntry{HumanAge: req.HumanAge, CatAge: catAge}
         database.DB.Create(&ageEntry)

         render.JSON(w, r, res)
     }
     ```

2. **Récupération des enregistrements depuis la base de données**

   - Ajoutez une route pour récupérer tous les calculs enregistrés. Dans le fichier `pkg/agecalculator/routes.go`, ajoutez une nouvelle route `/history`.

     Exemple de code pour `pkg/agecalculator/routes.go` :

     ```go
     package agecalculator

     import (
         "github.com/go-chi/chi/v5"
         "net/http"
     )

     func Routes() chi.Router {
         router := chi.NewRouter()
         router.Post("/age-in-cat-years", AgeInCatYearsHandler)
         router.Get("/history", AgeHistoryHandler)
         return router
     }
     ```

   - Dans le contrôleur `AgeHistoryHandler`, utilisez GORM pour récupérer les données et les retourner en JSON.

     Exemple de code pour `pkg/agecalculator/controller.go` :

     ```go
     package agecalculator

     import (
         "net/http"
         "github.com/go-chi/render"
         "github.com/yourusername/animal-api/database"
         "github.com/yourusername/animal-api/database/dbmodel"
     )

     func AgeHistoryHandler(w http.ResponseWriter, r *http.Request) {
         var ageEntries []dbmodel.AgeEntry
         database.DB.Find(&ageEntries)

         render.JSON(w, r, ageEntries)
     }
     ```

3. **Tester les opérations de CRUD**

   - **Création d’un calcul d’âge** : Testez l’endpoint `/age-in-cat-years` pour vérifier que les calculs sont enregistrés dans la base.

     ```bash
     curl -X POST http://localhost:8080/api/v1/agecalculator/age-in-cat-years -d '{"human_age": 5}' -H "Content-Type: application/json"
     ```

   - **Lecture de l’historique des calculs** : Testez l’endpoint `/history` pour récupérer tous les calculs enregistrés.
     ```bash
     curl http://localhost:8080/api/v1/agecalculator/history
     ```

### 4. Exercice en Autonomie : Enregistrer les Sons d'Animaux Identifiés 🐶🐱🦁

Pour cette étape, vous allez ajouter une fonctionnalité qui enregistre les sons d’animaux identifiés dans la base de données, en suivant les étapes précédentes.

#### Tâches

1. **Création d'un modèle `AnimalSound`**

   - Dans `database/dbmodel/`, créez un fichier `animal_sound.go` pour définir le modèle `AnimalSound`. Ce modèle contiendra le nom de l’animal et le son qu’il produit.

     Indications :

     - Le modèle pourrait avoir les champs `AnimalName` et `Sound`.
     - N'oubliez pas d'inclure `gorm.Model` pour bénéficier des champs automatiques d’ID, de création et de mise à jour.

2. **Configuration de la migration**

   - Modifiez la fonction `InitDatabase` dans `database/db.go` pour inclure `AnimalSound` dans l'auto-migration de GORM.

3. **Modification de `soundidentifier` pour Utiliser GORM**

   - Modifiez le contrôleur `AnimalSoundHandler` pour :
     - Enregistrer chaque son d’animal identifié dans la table `AnimalSound`.
     - Envoyer une réponse JSON avec le son identifié.

4. **Création d'une Nouvelle Route pour Récupérer les Sons Enregistrés**
   - Dans `soundidentifier`, créez une route `/history` pour récupérer l’historique des sons d’animaux identifiés.
   - Ajoutez un nouveau contrôleur `AnimalSoundHistoryHandler` pour récupérer et envoyer en JSON tous les enregistrements de `AnimalSound`.

#### Indications

- Utilisez `database.DB.Create` pour insérer de nouvelles entrées dans la base de données.
- Utilisez `database.DB.Find` pour récupérer tous les enregistrements de sons d’animaux dans `AnimalSoundHistoryHandler`.
- Testez les endpoints pour vérifier que les sons sont bien enregistrés et récupérables.

---

### Challenge : Optimiser les Requêtes avec des Filtres

En autonomie, améliorez la fonctionnalité d’historique pour permettre de filtrer les sons par nom d’animal. Par exemple, la route `/history` pourrait accepter un paramètre `animal`, et ne retourner que les sons de cet animal spécifique.

#### Tâches

1. **Ajoutez un Filtre dans le Contrôleur `AnimalSoundHistoryHandler`**

   - Modifiez `AnimalSoundHistoryHandler` pour accepter un paramètre de requête `animal`. Si ce paramètre est présent, filtrez les résultats pour ne retourner que les sons enregistrés pour cet animal spécifique.
   - Utilisez `database.DB.Where` pour appliquer le filtre lorsque `animal` est fourni dans la requête.

     Exemple de code pour le contrôleur `AnimalSoundHistoryHandler` :

     ```go
     package soundidentifier

     import (
         "net/http"
         "github.com/go-chi/render"
         "github.com/yourusername/animal-api/database"
         "github.com/yourusername/animal-api/database/dbmodel"
     )

     func AnimalSoundHistoryHandler(w http.ResponseWriter, r *http.Request) {
         animalName := r.URL.Query().Get("animal")
         var sounds []dbmodel.AnimalSound

         if animalName != "" {
             database.DB.Where("animal_name = ?", animalName).Find(&sounds)
         } else {
             database.DB.Find(&sounds)
         }

         render.JSON(w, r, sounds)
     }
     ```

2. **Testez le Filtre sur l’Historique des Sons**

   - Démarrez votre serveur et testez l'endpoint `/history` avec et sans le paramètre `animal` pour vérifier le comportement.
   - Exemple de requêtes avec `curl` :

     ```bash
     # Sans filtre, renvoie tous les sons enregistrés
     curl http://localhost:8080/api/v1/soundidentifier/history

     # Avec filtre, renvoie uniquement les sons pour "chien"
     curl http://localhost:8080/api/v1/soundidentifier/history?animal=chien
     ```

---

🎉 **Félicitations !** Vous avez maintenant introduit une fonctionnalité avancée avec GORM en permettant de filtrer les résultats en fonction de paramètres de requête. Vous avez appris à utiliser GORM pour interagir avec la base de données et optimisé votre API pour répondre à des requêtes spécifiques.
