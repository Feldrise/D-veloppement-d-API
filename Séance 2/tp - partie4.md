# TP : Structurer l’Accès aux Données avec GORM et des Repositories

## Partie 4 : Structuration avancée avec GORM et des Repositories

**Durée estimée :** 2 heures

### Objectifs

Dans cette partie, vous allez :

1. Configurer l’accès aux données avec GORM dans un fichier de configuration centralisé.
2. Utiliser des repositories pour structurer les interactions avec la base de données.
3. Refactoriser les routes et contrôleurs pour exploiter cette nouvelle structure.

Nous allons structurer notre API pour qu’elle soit modulaire, en s’appuyant sur une architecture similaire à celle fournie dans votre exemple.

### Étapes

---

### 1. Configuration de l’Application avec GORM

1. **Création du fichier de configuration `config.go`**

   - Dans un nouveau dossier `config/`, créez un fichier `config.go` qui initialisera les constantes et les repositories, et maintiendra une instance de configuration centrale pour l’application.
   - Ajoutez une structure `Config` pour contenir les informations de connexion et les repositories.

     Exemple de code pour `config/config.go` :

     ```go
     package config

     import (
         "log"
         "gorm.io/driver/sqlite"
         "gorm.io/gorm"
         "yourproject/database"
         "yourproject/database/dbmodel"
     )

     type Config struct {
         // Connexion aux repositories
         AgeEntryRepository     dbmodel.AgeEntryRepository
         AnimalSoundRepository  dbmodel.AnimalSoundRepository
     }

     func New() (*Config, error) {
         config := Config{}

         // Initialisation de la connexion à la base de données
         databaseSession, err := gorm.Open(sqlite.Open("animal_api.db"), &gorm.Config{})
         if err != nil {
             return &config, err
         }

         // Migration des modèles
         database.Migrate(databaseSession)

         // Initialisation des repositories
         config.AgeEntryRepository = dbmodel.NewAgeEntryRepository(databaseSession)
         config.AnimalSoundRepository = dbmodel.NewAnimalSoundRepository(databaseSession)

         return &config, nil
     }
     ```

2. **Configuration de la Migration**

   - Dans `database/database.go`, configurez la migration automatique pour vos modèles `AgeEntry` et `AnimalSound`.

     Exemple de code pour `database/database.go` :

     ```go
     package database

     import (
         "log"
         "yourproject/database/dbmodel"
         "gorm.io/gorm"
     )

     func Migrate(db *gorm.DB) {
         db.AutoMigrate(
             &dbmodel.AgeEntry{},
             &dbmodel.AnimalSound{},
         )
         log.Println("Database migrated successfully")
     }
     ```

---

### 2. Création des Repositories pour les Modèles

Les repositories permettent de centraliser les interactions avec la base de données, rendant le code plus modulaire et testable.

1. **Création du repository pour `AgeEntry`**

   - Dans `database/dbmodel/age_entry.go`, définissez une interface `AgeEntryRepository` avec les méthodes CRUD.
   - Ajoutez une structure `ageEntryRepository` qui implémente ces méthodes.

     Exemple de code pour `database/dbmodel/age_entry.go` :

     ```go
     package dbmodel

     import (
         "context"
         "gorm.io/gorm"
     )

     type AgeEntry struct {
         gorm.Model
         HumanAge int `json:"human_age"`
         CatAge   int `json:"cat_age"`
     }

     type AgeEntryRepository interface {
         Create(entry *AgeEntry) (*AgeEntry, error)
         FindAll() ([]*AgeEntry, error)
     }

     type ageEntryRepository struct {
         db *gorm.DB
     }

     func NewAgeEntryRepository(db *gorm.DB) AgeEntryRepository {
         return &ageEntryRepository{db: db}
     }

     func (r *ageEntryRepository) Create(entry *AgeEntry) (*AgeEntry, error) {
         if err := r.db.Create(entry).Error; err != nil {
             return nil, err
         }
         return entry, nil
     }

     func (r *ageEntryRepository) FindAll() ([]*AgeEntry, error) {
         var entries []*AgeEntry
         if err := r.db.Find(&entries).Error; err != nil {
             return nil, err
         }
         return entries, nil
     }
     ```

2. **Création du repository pour `AnimalSound`**
   - Répétez l’étape précédente pour créer un repository `AnimalSoundRepository` dans `database/dbmodel/animal_sound.go`.

---

### 3. Mise à Jour des Features pour Utiliser les Repositories

1. **Modification de `agecalculator` pour utiliser `AgeEntryRepository`**

   - Dans le contrôleur `AgeInCatYearsHandler`, utilisez le repository pour créer un nouvel enregistrement dans la base de données et pour récupérer les données d’historique.

     Exemple de code pour `pkg/agecalculator/controller.go` :

     ```go
     package agecalculator

     import (
         "net/http"
         "github.com/go-chi/render"
         "yourproject/config"
         "yourproject/pkg/model"
         "yourproject/database/dbmodel"
     )

     type AgeCalculatorConfig struct {
         *config.Config
     }

     func New(configuration *config.Config) *AgeCalculatorConfig {
         return &AgeCalculatorConfig{configuration}
     }

     func (config *AgeCalculatorConfig) AgeInCatYearsHandler(w http.ResponseWriter, r *http.Request) {
         req := &model.AgeRequest{}
         if err := render.Bind(r, req); err != nil {
             render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
             return
         }

         catAge := calculateCatAge(req.HumanAge)
         ageEntry := &dbmodel.AgeEntry{HumanAge: req.HumanAge, CatAge: catAge}
         config.AgeEntryRepository.Create(ageEntry)

         res := &model.AgeResponse{CatAge: catAge}
         render.JSON(w, r, res)
     }

     func (config *AgeCalculatorConfig) AgeHistoryHandler(w http.ResponseWriter, r *http.Request) {
         entries, err := config.AgeEntryRepository.FindAll()
         if err != nil {
             render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
             return
         }
         render.JSON(w, r, entries)
     }
     ```

2. **Mise en place des Routes**

   - Dans `pkg/agecalculator/routes.go`, ajustez le code pour que les routes utilisent l’instance configurée.

     ```go
     package agecalculator

     import (
         "github.com/go-chi/chi/v5"
         "yourproject/config"
     )

     func Routes(configuration *config.Config) *chi.Mux {
         ageCalculatorConfig := New(configuration)
         router := chi.NewRouter()

         router.Post("/age-in-cat-years", ageCalculatorConfig.AgeInCatYearsHandler)
         router.Get("/history", ageCalculatorConfig.AgeHistoryHandler)

         return router
     }
     ```

---

### 4. Configuration du `main.go`

1. **Initialisation de la Configuration et des Routes**

   - Dans `main.go`, initialisez la configuration, puis passez-la aux routes de chaque feature.

     Exemple de code pour `main.go` :

     ```go
     package main

     import (
         "log"
         "net/http"
         "yourproject/config"
         "yourproject/pkg/agecalculator"
         "github.com/go-chi/chi/v5"
     )

     func Routes(configuration *config.Config) *chi.Mux {
         router := chi.NewRouter()
         router.Mount("/api/v1/agecalculator", agecalculator.Routes(configuration))
         return router
     }

     func main() {
         // Initialisation de la configuration
         configuration, err := config.New()
         if err != nil {
             log.Panicln("Configuration error:", err)
         }

         // Initialisation des routes
         router := Routes(configuration)

         log.Println("Serving on :8080")
         log.Fatal(http.ListenAndServe(":8080", router))
     }
     ```

---

🎉 **Félicitations !** Vous avez structuré votre projet en utilisant une configuration centralisée et des repositories avec GORM. Cette approche rend votre API modulaire, testable et facile à maintenir.
