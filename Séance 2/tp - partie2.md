# TP : Utiliser des Modèles avec `chi/render` pour Structurer les Features 🐱🐶

## Partie 2 : Introduction aux Modèles avec `chi/render`

**Durée estimée :** 1 heure

### Objectifs

Dans cette deuxième partie, vous allez introduire des modèles en utilisant le package `render` de Chi pour structurer et gérer plus facilement les entrées et les sorties de vos fonctionnalités. Vous allez implémenter un modèle pour la fonctionnalité de calcul de l’âge en âge de chat, puis vous serez guidés pour créer un modèle pour la fonctionnalité de reconnaissance des bruits d’animaux en autonomie.

### Préparation

Avant de commencer, assurez-vous que le package `chi/render` est installé. Si ce n’est pas le cas, installez-le avec la commande :

```bash
go get github.com/go-chi/render
```

### Étapes

---

### 1. Exemple Guidé : Utilisation de `render` pour l'Âge en Âge de Chat 🐱

Pour cette première étape, nous allons structurer notre code en utilisant `render` pour gérer les requêtes et réponses avec un modèle. Cela simplifiera le code et centralisera la gestion des erreurs et de la sérialisation.

#### Étapes

1. **Création d’un modèle `AgeRequest`**

   - Dans le dossier `pkg/model/`, créez un fichier `age.go` qui définira une structure `AgeRequest`. Ce modèle représentera les données d’entrée pour le calcul de l’âge en âge de chat.

     Exemple de code pour `pkg/model/age.go` :

     ```go
     package model

     type AgeRequest struct {
         HumanAge int `json:"human_age"`
     }

     type AgeResponse struct {
         CatAge int `json:"cat_age"`
     }
     ```

2. **Utilisation de `render` pour structurer les réponses**

   - Dans le contrôleur `AgeInCatYearsHandler` de `agecalculator`, utilisez `render.Bind` pour décoder la requête dans le modèle `AgeRequest`, puis utilisez `render.JSON` pour structurer la réponse dans un modèle `AgeResponse`.

     Exemple de code pour `pkg/agecalculator/controller.go` :

     ```go
     package agecalculator

     import (
         "net/http"
         "github.com/go-chi/render"
         "github.com/yourusername/animal-api/pkg/model"
     )

     func AgeInCatYearsHandler(w http.ResponseWriter, r *http.Request) {
         req := &model.AgeRequest{}
         if err := render.Bind(r, req); err != nil {
             render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
             return
         }

         catAge := calculateCatAge(req.HumanAge)
         res := &model.AgeResponse{CatAge: catAge}
         render.JSON(w, r, res)
     }

     func calculateCatAge(humanAge int) int {
         if humanAge <= 2 {
             return humanAge * 12
         }
         return 24 + (humanAge-2)*4
     }
     ```

3. **Mise à jour du modèle pour la compatibilité avec `render`**

   - Mettez à jour `AgeRequest` pour qu’il implémente l’interface `render.Binder`. Cela vous permet de gérer la validation des données d’entrée.

     Exemple de code pour `pkg/model/age.go` :

     ```go
     package model

     import "errors"

     type AgeRequest struct {
         HumanAge int `json:"human_age"`
     }

     func (a *AgeRequest) Bind(r *http.Request) error {
         if a.HumanAge < 0 {
             return errors.New("human_age must be a positive integer")
         }
         return nil
     }

     type AgeResponse struct {
         CatAge int `json:"cat_age"`
     }
     ```

4. **Tester la nouvelle structure**

   - Démarrez votre serveur et testez l’endpoint `age-in-cat-years` en envoyant une requête POST avec un corps JSON pour vérifier que le modèle est bien utilisé.

     Exemple de requête avec `curl` :

     ```bash
     curl -X POST http://localhost:8080/api/v1/agecalculator/age-in-cat-years -d '{"human_age": 5}' -H "Content-Type: application/json"
     ```

---

### 2. Exercice en Autonomie : Modèle avec `render` pour Identifier les Bruits d'Animaux 🐶🐱🦁

Pour cette seconde étape, vous allez créer un modèle de données pour la fonctionnalité `soundidentifier`, en suivant l'exemple précédent. L’objectif est de structurer la requête et la réponse en utilisant `render` pour plus de clarté et de maintenabilité.

#### Tâches

1. **Création du modèle pour l’identification de son d'animal**

   - Dans `pkg/model/`, créez un fichier `sound.go` avec deux structures :
     - `SoundRequest` pour représenter le nom de l’animal reçu.
     - `SoundResponse` pour représenter le son renvoyé en fonction de l’animal.

2. **Implémentation de `render.Binder` pour SoundRequest**

   - Assurez-vous que `SoundRequest` implémente l’interface `render.Binder` pour valider les entrées. Par exemple, vérifiez que le champ `AnimalName` n'est pas vide.

3. **Utilisation de `render` dans le contrôleur `AnimalSoundHandler`**
   - Modifiez le contrôleur pour utiliser `render.Bind` pour décoder `SoundRequest` et `render.JSON` pour structurer la réponse avec `SoundResponse`.

#### Indications

- `SoundRequest` peut contenir un champ `AnimalName` pour le nom de l’animal, tandis que `SoundResponse` contiendra un champ `Sound` pour le bruit de l’animal.
- Pour tester, utilisez une requête POST similaire à celle de la première feature.

---

### 3. Challenge : Identification Multiple des Bruits d'Animaux 🎉

En autonomie, ajoutez un modèle et une route permettant de demander plusieurs animaux dans une seule requête et d'obtenir un dictionnaire `{animal: bruit}` en réponse.

#### Tâches

1. **Modifiez le modèle `SoundRequest` pour prendre une liste d'animaux.**
2. **Créez une nouvelle route** dans `soundidentifier` pour traiter les requêtes où plusieurs animaux sont demandés.
3. **Utilisez un modèle `MultipleSoundsResponse`** pour structurer la réponse en utilisant `render.JSON` pour obtenir un dictionnaire `{animal: bruit}`.

Testez chaque étape en vous assurant que la structure des données et les réponses sont correctement formatées en JSON.

---

🎉 **Bravo !** Vous avez appris à structurer vos fonctionnalités avec des modèles et le package `render` pour simplifier la gestion des entrées et sorties JSON de votre API Go avec Chi.
