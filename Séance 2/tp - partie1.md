# TP : Structurer des Features avec Chi 🐱🐶

## Partie 1 : Structuration et Création de Deux Features

**Durée estimée :** 45 minutes

### Objectifs

Dans ce TP, vous allez structurer votre projet pour intégrer deux fonctionnalités indépendantes en utilisant le framework Chi. L'objectif est de bien comprendre comment organiser les "features" dans des dossiers séparés avec leurs propres routes et contrôleurs. Cela vous permettra de créer des fonctionnalités modulaires et de faciliter la maintenance de votre API. 📂

Les deux fonctionnalités que vous allez implémenter sont :

1. Calculer son âge en âge de chat 🐱
2. Identifier les bruits d’animaux 🐶🐱🦁

### Étapes

1. **Création de la structure de base du projet**

   - Créez un nouveau dossier pour votre projet que vous nommerez `animal-api`.
   - Organisez le projet comme suit :
     ```
     animal-api/
     ├── pkg/                # Code source de l'application
     │   ├── agecalculator/  # Feature pour calculer l'âge en âge de chat
     │   │   ├── routes.go   # Définitions des routes
     │   │   └── controller.go # Logique de calcul
     │   ├── soundidentifier/ # Feature pour identifier les bruits d’animaux
     │   │   ├── routes.go   # Définitions des routes
     │   │   └── controller.go # Logique de traitement
     └── main.go             # Point d'entrée de l'application
     ```

2. **Implémentation de la feature `agecalculator` pour l'âge en âge de chat 🐱**

   - Dans `pkg/agecalculator/`, créez un fichier `routes.go` pour définir les routes et un fichier `controller.go` pour la logique de calcul.
   - Dans `routes.go`, ajoutez une route `/age-in-cat-years` qui permet à l’utilisateur de fournir son âge en années humaines et reçoit l’âge correspondant en âge de chat.

     Exemple de code pour `pkg/agecalculator/routes.go` :

     ```go
     package agecalculator

     import (
         "github.com/go-chi/chi/v5"
         "net/http"
     )

     func Routes() chi.Router {
         router := chi.NewRouter()
         router.Get("/age-in-cat-years", AgeInCatYearsHandler)
         return router
     }
     ```

   - Dans `controller.go`, écrivez la fonction qui calcule l’âge en années de chat. Utilisez une règle simple : les 2 premières années humaines valent 25 ans de chat chacune, puis chaque année humaine supplémentaire vaut 4 ans de chat.

     Exemple de code pour `pkg/agecalculator/controller.go` :

     ```go
     package agecalculator

     import (
         "net/http"
         "strconv"
         "encoding/json"
     )

     func AgeInCatYearsHandler(w http.ResponseWriter, r *http.Request) {
         ageStr := r.URL.Query().Get("age")
         age, err := strconv.Atoi(ageStr)
         if err != nil || age < 0 {
             http.Error(w, "Invalid age parameter", http.StatusBadRequest)
             return
         }

         catAge := calculateCatAge(age)

         w.Header().Set("Content-Type", "application/json")
         json.NewEncoder(w).Encode(map[string]int{"age_in_cat_years": catAge})
     }

     func calculateCatAge(humanAge int) int {
         if humanAge <= 2 {
             return humanAge * 12
         }
         return 24 + (humanAge-2)*4
     }
     ```

3. **Implémentation de la feature `soundidentifier` pour identifier les bruits d'animaux 🐶🐱**

   - Dans `pkg/soundidentifier/`, créez un fichier `routes.go` pour définir les routes et un fichier `controller.go` pour la logique de traitement.
   - Dans `routes.go`, ajoutez une route `/animal-sound` qui prend un nom d’animal en paramètre et retourne le bruit qu’il fait (exemple : "chien" → "woof").

   - Dans `controller.go`, écrivez la logique qui, en fonction du nom de l’animal fourni (`chien`, `chat`, etc.), retourne le bruit correspondant en JSON. Utilisez un switch simple pour gérer plusieurs animaux.

4. **Connexion des routes dans le `main.go`**

   - Dans le fichier `main.go`, importez les packages `agecalculator` et `soundidentifier`, puis ajoutez-les au routeur principal.

     Exemple de code pour `main.go` :

     ```go
     package main

     import (
         "log"
         "net/http"
         "github.com/go-chi/chi/v5"
         "github.com/yourusername/animal-api/pkg/agecalculator"
         "github.com/yourusername/animal-api/pkg/soundidentifier"
     )

     func main() {
         r := chi.NewRouter()
         r.Mount("/api/v1/agecalculator", agecalculator.Routes())
         r.Mount("/api/v1/soundidentifier", soundidentifier.Routes())

         log.Println("Serving on :8080")
         http.ListenAndServe(":8080", r)
     }
     ```

5. **Tester les endpoints**
   - Démarrez votre serveur en exécutant `go run main.go`.
   - Testez vos endpoints :
     - **Age en âge de chat** : `[GET] http://localhost:8080/api/v1/agecalculator/age-in-cat-years?age=5`
     - **Identification des bruits d’animaux** : `[GET] http://localhost:8080/api/v1/soundidentifier/animal-sound?animal=chien`

---

🎉 **Bravo !** Vous avez structuré votre projet en Go avec deux features indépendantes, chacune ayant ses propres routes et contrôleurs. Vous commencez à appréhender une structure modulaire en utilisant Chi !
