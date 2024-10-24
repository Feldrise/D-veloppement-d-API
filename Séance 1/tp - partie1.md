# TP : Créer ta première API avec Go 🐹🚀

## Partie 1 : Mise en place de l'API

**Durée estimée : 45-60 minutes**

### Objectifs

- Découvrir les bases du langage Go
- Mettre en place un serveur HTTP simple
- Apprendre à définir et gérer des routes
- Manipuler des réponses HTTP pour interagir avec le client

### Étapes

### 1. Crée ton projet Go 🎉

- Commence par créer un répertoire pour ton projet, puis initialise un module Go :

```bash
mkdir premiere-api-go
cd premiere-api-go
go mod init premiere-api-go
```

- Ce module va permettre à Go de gérer les dépendances du projet.

### 2. Crée ton serveur HTTP 🚀

- Crée un fichier `main.go` et commence par écrire un programme basique qui va créer un serveur HTTP.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Bienvenue dans mon API Go ! 🚀")
    })

    // Démarre le serveur sur le port 8080
    log.Println("Serveur démarré sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### Explication du code :

- `http.HandleFunc("/", ...)` : Cette ligne définit une route pour la racine `/` qui répondra à toutes les requêtes avec le message `"Bienvenue dans mon API Go ! 🚀"`.
- `http.ListenAndServe(":8080", nil)` : Cette fonction lance le serveur HTTP sur le port `8080`. Le second argument `nil` signifie que nous n'utilisons pas de routeur personnalisé pour l'instant (ça vient après 😉).

- Lance ton serveur en tapant :

```bash
go run main.go
```

- Ouvre ton navigateur et accède à l'URL suivante : **http://localhost:8080**
  - Tu devrais voir apparaître le message **"Bienvenue dans mon API Go ! 🚀"**.

### 3. Ajoute des routes simples 🌐

- Maintenant, on va ajouter d'autres routes pour que notre API réponde à plusieurs types de requêtes. Modifie le fichier `main.go` pour ajouter des routes pour **GET** et **POST** :

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

// Gestionnaire pour la route /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World! 🌍")
}

// Gestionnaire pour la route /greet avec un paramètre de requête
func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "stranger"
    }
    fmt.Fprintf(w, "Hello, %s! 👋", name)
}

func main() {
    // Route racine
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Bienvenue dans mon API Go ! 🚀")
    })

    // Route /hello
    http.HandleFunc("/hello", helloHandler)

    // Route /greet avec un paramètre "name"
    http.HandleFunc("/greet", greetHandler)

    // Démarre le serveur sur le port 8080
    log.Println("Serveur démarré sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### Explication des nouvelles routes :

- `http.HandleFunc("/hello", helloHandler)` : Quand tu vas sur `/hello`, le serveur retourne `"Hello, World! 🌍"`.
- `http.HandleFunc("/greet", greetHandler)` : Cette route prend un paramètre `name` dans l'URL, et si le paramètre est fourni (ex: `/greet?name=Ynov`), l'API te répond avec `"Hello, Ynov! 👋"`. Si aucun nom n'est fourni, il renvoie `"Hello, stranger! 👋"`.

- Relance ton serveur (`Ctrl + C` pour l'arrêter et `go run main.go` pour le redémarrer).
  - Teste la route **/hello** : [http://localhost:8080/hello](http://localhost:8080/hello)
  - Teste la route **/greet** avec un paramètre : [http://localhost:8080/greet?name=Ynov](http://localhost:8080/greet?name=Ynov)

### 4. Ajoute une route POST pour recevoir des données 📩

- Ajoutons maintenant une route qui utilise la méthode **POST**. Nous allons accepter des données envoyées par le client.

Modifie ton fichier `main.go` pour ajouter une route **POST** :

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World! 🌍")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "stranger"
    }
    fmt.Fprintf(w, "Hello, %s! 👋", name)
}

// Gestionnaire pour la route POST /submit
func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    r.ParseForm()
    name := r.FormValue("name")
    if name == "" {
        name = "inconnu"
    }

    fmt.Fprintf(w, "Données reçues : Bonjour, %s! 🎉", name)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Bienvenue dans mon API Go ! 🚀")
    })

    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/greet", greetHandler)
    http.HandleFunc("/submit", submitHandler)

    log.Println("Serveur démarré sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### Explication :

- La route **/submit** accepte uniquement les requêtes **POST**. Si une requête **GET** est envoyée à cette route, le serveur renvoie une erreur `405 Method Not Allowed`.
- La fonction `r.ParseForm()` analyse les données envoyées par le client (dans le corps de la requête), et `r.FormValue("name")` permet de récupérer la valeur du champ `name` envoyé.

Pour tester cette route, utilise un outil comme **Postman** ou **curl** :

- Avec **curl** :

```bash
curl -X POST -d "name=Ynov" http://localhost:8080/submit
```

Cela devrait retourner : **"Données reçues : Bonjour, Ynov! 🎉"**

---

## Conclusion 🎓

Bravo ! 🎉 Vous avez réussi à créer votre première API en Go ! Dans ce TP, vous avez appris :

- Comment créer un serveur HTTP en Go.
- Comment définir des routes et gérer les méthodes HTTP (GET et POST).
- Comment interagir avec des paramètres dans les URLs et les formulaires.
