# TP : Créez une API avec Chi 🐹🍃

## Partie 3 : Découverte de Chi

**Durée estimée : 45-60 minutes**

### Objectifs

- Découvrir et utiliser le framework Chi pour structurer une API en Go
- Apprendre à définir des routes et des groupes de routes
- Utiliser les middlewares avec Chi pour gérer des fonctionnalités globales (logging, etc.)
- Comparer Chi avec Gin et l'implémentation standard de Go

### Étapes

### 1. Installation de Chi 🛠️

Avant de commencer, vous devez installer le package Chi dans votre projet. Chi est disponible via Go Modules, donc vous pouvez l’ajouter en utilisant la commande suivante dans votre répertoire de projet :

```bash
go get github.com/go-chi/chi/v5
```

Chi est maintenant installé et vous pouvez l'utiliser dans votre projet Go.

---

### 2. Créer votre serveur avec Chi 🚀

Modifions le fichier `main.go` pour initialiser un serveur Chi et recréer les routes de base, comme dans les TP précédents avec Go standard et Gin.

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    // Initialise un nouveau routeur Chi
    r := chi.NewRouter()

    // Utilisation d'un middleware de logging fourni par Chi
    r.Use(middleware.Logger)

    // Route racine
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Bienvenue dans mon API Go avec Chi ! 🚀"))
    })

    // Route /hello
    r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World! 🌍"))
    })

    // Route /greet avec un paramètre de requête
    r.Get("/greet", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            name = "stranger"
        }
        w.Write([]byte("Hello, " + name + "! 👋"))
    })

    // Démarre le serveur sur le port 8080
    http.ListenAndServe(":8080", r)
}
```

#### Explication :

- **Chi Router** : Le routeur Chi est initialisé avec `chi.NewRouter()`. On peut ensuite y attacher des routes avec `r.Get` pour gérer les requêtes **GET**.
- **Middleware** : Chi offre plusieurs middlewares intégrés, comme `Logger`, qui log les requêtes HTTP (ce que Gin propose par défaut).
- Les routes sont définies avec `r.Get()` et, comme dans les exemples précédents, on renvoie des réponses simples en utilisant `w.Write()`.

#### Testez les routes :

- Accédez à [http://localhost:8080/](http://localhost:8080/) pour voir la route racine.
- Testez [http://localhost:8080/hello](http://localhost:8080/hello).
- Essayez la route **greet** avec un paramètre, ex : [http://localhost:8080/greet?name=Ynov](http://localhost:8080/greet?name=Ynov).

---

### 3. Ajoutez des routes POST avec Chi 📩

Maintenant que nous avons recréé les routes **GET**, ajoutons une route **POST** pour recevoir des données du client, comme nous l'avons fait dans les TP précédents.

Modifions `main.go` pour ajouter une route **POST** :

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
}

func main() {
    r := chi.NewRouter()

    // Middleware pour logger les requêtes
    r.Use(middleware.Logger)

    // Route racine
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Bienvenue dans mon API Go avec Chi ! 🚀"))
    })

    // Route /hello
    r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World! 🌍"))
    })

    // Route /greet
    r.Get("/greet", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            name = "stranger"
        }
        w.Write([]byte("Hello, " + name + "! 👋"))
    })

    // Route POST /submit pour recevoir des données
    r.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
        var person Person

        // Décodage des données JSON reçues
        err := json.NewDecoder(r.Body).Decode(&person)
        if err != nil || person.Name == "" {
            http.Error(w, "Le champ 'name' est requis", http.StatusBadRequest)
            return
        }

        w.Write([]byte("Données reçues : Bonjour, " + person.Name + "! 🎉"))
    })

    // Démarre le serveur sur le port 8080
    http.ListenAndServe(":8080", r)
}
```

#### Explication :

- **Route POST** : On définit une route POST avec `r.Post("/submit", ...)`.
- **JSON Parsing** : La fonction `json.NewDecoder(r.Body).Decode(&person)` permet de lire et de décoder le corps de la requête JSON vers une structure Go (`Person`).
- Si le nom est vide ou si le JSON est mal formé, on renvoie une erreur `400 Bad Request`.

#### Testez avec **curl** ou **Postman** :

- Avec **curl** :

```bash
curl -X POST http://localhost:8080/submit -d '{"name":"Ynov"}' -H "Content-Type: application/json"
```

Cela devrait retourner : **"Données reçues : Bonjour, Ynov! 🎉"**

---

### 4. Utilisation des groupes de routes avec Chi 🌐

Une des forces de Chi est la possibilité de définir des **groupes de routes**, ce qui permet d'organiser votre API plus facilement, surtout si certaines routes partagent des middlewares ou des préfixes communs.

Modifions le fichier `main.go` pour illustrer l'utilisation des groupes de routes :

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
}

func main() {
    r := chi.NewRouter()

    // Middleware pour logger les requêtes
    r.Use(middleware.Logger)

    // Groupe de routes pour /api
    r.Route("/api", func(api chi.Router) {
        // Route racine /api
        api.Get("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Bienvenue dans l'API version Chi ! 🌐"))
        })

        // Route /api/hello
        api.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello depuis /api/hello ! 👋"))
        })

        // Route POST /api/submit
        api.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
            var person Person
            err := json.NewDecoder(r.Body).Decode(&person)
            if err != nil || person.Name == "" {
                http.Error(w, "Le champ 'name' est requis", http.StatusBadRequest)
                return
            }
            w.Write([]byte("Données reçues : Bonjour, " + person.Name + "! 🎉"))
        })
    })

    // Démarre le serveur sur le port 8080
    http.ListenAndServe(":8080", r)
}
```

#### Explication :

- **Groupes de routes** : `r.Route("/api", func(api chi.Router) { ... })` permet de regrouper les routes sous le préfixe `/api`. Toutes les routes définies dans cette fonction auront `/api` comme préfixe.
- Cela permet de mieux organiser le code, surtout dans des projets plus complexes, où certaines routes peuvent avoir des préfixes ou middlewares communs.

#### Testez les nouvelles routes :

- Route [http://localhost:8080/api/hello](http://localhost:8080/api/hello)
- Route POST `/api/submit` avec **curl** :

```bash
curl -X POST http://localhost:8080/api/submit -d '{"name":"Chi"}' -H "Content-Type: application/json"
```

---

### 5. Différences clés entre Chi et Gin/Go Standard 🤔

#### Similitudes :

- Comme Gin et l'implément

ation standard de Go, Chi permet de créer des serveurs HTTP, définir des routes et gérer des requêtes et réponses.

#### Ce que Chi apporte de plus :

- **Simplicité et légèreté** : Chi est un framework minimaliste qui vous donne juste ce dont vous avez besoin, sans fonctionnalités supplémentaires.
- **Groupes de routes** : Chi propose une manière simple et flexible de regrouper les routes, ce qui est très utile pour organiser des API complexes.
- **Middlewares** : Comme Gin, Chi intègre des middlewares, mais vous avez aussi la possibilité de les définir et de les appliquer uniquement à certains groupes de routes ou à des routes spécifiques.

---

### 6. Conclusion 🎓

Bravo ! 🎉 Vous avez maintenant une API fonctionnelle en utilisant **Chi**, un framework léger et flexible pour construire des API en Go. Vous avez appris à :

- Créer un serveur HTTP avec Chi.
- Définir des routes GET et POST.
- Utiliser les middlewares pour ajouter des fonctionnalités globales.
- Organiser les routes avec des groupes.

Avec Chi, vous avez une approche minimaliste mais puissante pour structurer des API, tout en gardant un code propre et maintenable. 🌟
