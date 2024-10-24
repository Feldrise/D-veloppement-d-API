# TP : Créez une API avec Gin 🐹🍸

## Partie 2 : Passage à Gin

**Durée estimée : 45-60 minutes**

### Objectifs

- Découvrir et utiliser le framework Gin
- Recréer des routes basiques avec Gin
- Apprendre à gérer les requêtes et réponses plus facilement
- Comparer l'utilisation de Gin avec l'implémentation standard de Go

### Étapes

### 1. Installation de Gin 🚀

Avant de commencer, vous devez installer le package Gin dans votre projet. Gin est disponible via Go Modules, donc vous pouvez l’ajouter en utilisant la commande suivante dans votre répertoire de projet :

```bash
go get -u github.com/gin-gonic/gin
```

Gin est maintenant installé, vous pouvez l'importer et l'utiliser dans votre fichier `main.go`.

---

### 2. Créez votre premier serveur Gin 🎉

Modifiez votre fichier `main.go` pour qu'il utilise Gin à la place de `net/http`. Vous allez recréer les mêmes routes que précédemment, mais cette fois avec Gin.

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Crée un nouveau routeur Gin
    router := gin.Default()

    // Route racine
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Bienvenue dans mon API Go avec Gin ! 🚀",
        })
    })

    // Route /hello
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World! 🌍",
        })
    })

    // Route /greet avec un paramètre de requête
    router.GET("/greet", func(c *gin.Context) {
        name := c.DefaultQuery("name", "stranger")
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, " + name + "! 👋",
        })
    })

    // Démarre le serveur sur le port 8080
    router.Run(":8080")
}
```

#### Explication des changements avec Gin :

- **Gin Default Router** : `gin.Default()` crée un routeur Gin avec un ensemble de middlewares prêts à l'emploi (comme le logger).
- **Gestionnaire de route** : Au lieu d'utiliser `http.HandleFunc`, on utilise `router.GET` pour définir les routes GET. La fonction prend un contexte (`c *gin.Context`) pour accéder aux informations de la requête et générer des réponses.
- **c.JSON(...)** : Cette méthode permet de formater facilement les réponses HTTP au format JSON, une opération très fréquente dans les API.

#### Testez les routes comme avant :

- Accédez à [http://localhost:8080/](http://localhost:8080/) pour voir la route racine.
- Testez [http://localhost:8080/hello](http://localhost:8080/hello).
- Essayez la route `greet` avec un paramètre, ex : [http://localhost:8080/greet?name=Ynov](http://localhost:8080/greet?name=Ynov).

---

### 3. Gérer les requêtes POST avec Gin 📩

Maintenant que nous avons recréé nos routes **GET** avec Gin, voyons comment gérer les requêtes **POST**. Gin simplifie l'extraction des données dans le corps d'une requête POST.

Ajoutons une route **POST** pour recevoir des données envoyées par le client.

Modifiez `main.go` pour ajouter cette route :

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    // Route racine
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Bienvenue dans mon API Go avec Gin ! 🚀",
        })
    })

    // Route /hello
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World! 🌍",
        })
    })

    // Route /greet
    router.GET("/greet", func(c *gin.Context) {
        name := c.DefaultQuery("name", "stranger")
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, " + name + "! 👋",
        })
    })

    // Route POST /submit pour recevoir des données
    router.POST("/submit", func(c *gin.Context) {
        var json struct {
            Name string `json:"name" binding:"required"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Le champ 'name' est requis"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Données reçues : Bonjour, " + json.Name + "! 🎉",
        })
    })

    // Démarre le serveur sur le port 8080
    router.Run(":8080")
}
```

#### Explication de la route POST :

- `router.POST("/submit", ...)` : Nous définissons une route **POST** pour recevoir des données.
- `c.ShouldBindJSON(...)` : Gin gère la liaison des données JSON envoyées dans le corps de la requête directement avec une structure Go.
- Le champ `binding:"required"` permet de s'assurer que le champ `name` est bien présent, sinon une erreur `400 Bad Request` est renvoyée.

#### Testez avec **curl** ou **Postman** :

- Avec **curl** :

```bash
curl -X POST http://localhost:8080/submit -d '{"name":"Ynov"}' -H "Content-Type: application/json"
```

Cela devrait retourner : **"Données reçues : Bonjour, Ynov! 🎉"**

---

### 4. Différences clés entre Gin et net/http 🤔

#### Points communs :

- Les deux approches permettent de créer des serveurs HTTP en Go et de définir des routes pour répondre aux requêtes.

#### Ce que Gin simplifie :

- **Gestion des routes** : Gin propose des méthodes plus simples pour définir des routes (`router.GET`, `router.POST`, etc.) par rapport à `http.HandleFunc`.
- **Gestion des réponses JSON** : Gin fournit des méthodes directes pour formater les réponses en JSON (`c.JSON()`), ce qui rend le code plus lisible et plus concis.
- **Parsing des requêtes** : Gin permet de lier automatiquement les données des requêtes (via `ShouldBindJSON`) aux structures Go, avec des vérifications intégrées (comme `binding:"required"`).
- **Middlewares** : Gin inclut des middlewares par défaut, comme le logging des requêtes et la gestion des erreurs, alors que dans `net/http`, il faut souvent les ajouter manuellement.

---

### 5. Conclusion 🎓

Félicitations ! 🎉 Vous avez maintenant une API fonctionnelle utilisant **Gin**, un framework Go léger mais puissant. Vous avez vu :

- Comment créer un serveur avec Gin.
- Comment définir des routes GET et POST.
- Comment formater des réponses JSON facilement.
- Comment valider et extraire des données dans les requêtes.

Gin simplifie énormément la création d'API en Go, tout en vous permettant de maintenir un code clair et organisé. 🎊
