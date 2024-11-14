# Cours de Développement d'API avec Go

Bienvenue dans le repository du cours de Développement d'API avec Go !

## Objectifs du Cours

Ce cours a pour objectif d'initier les étudiants aux concepts fondamentaux du développement d'API RESTful en utilisant le langage Go. À la fin de ce cours, les étudiants seront capables de :

- Comprendre les principes d'une API RESTful et les bonnes pratiques de conception.
- Implémenter des routes et gérer les requêtes HTTP (GET, POST, PUT, DELETE) avec Go.
- Manipuler des bases de données (CRUD) via une API avec Go.
- Gérer les erreurs et assurer une validation efficace des données.
- Mettre en place des tests unitaires et d'intégration pour les API.
- Optimiser les performances et la sécurité d'une API.

## Structure du Cours

Le cours est organisé en 10 séances de 3h30 chacune. Voici une vue d'ensemble des principales séances :

1. **Introduction à Go et au développement d'API RESTful**.
2. **Gestion des routes et des requêtes HTTP en Go**.
3. **Introduction à la gestion des bases de données avec Go**.
4. **Création de routes pour des opérations CRUD**.
5. **Validation des données et gestion des erreurs**.
6. **Authentification et autorisation dans les API**.
7. **Mise en place de tests unitaires pour une API en Go**.
8. **Mise en place de tests d'intégration pour une API**.
9. **Travail sur le projet avec supervision**.
10. **Présentation des projets et évaluation finale**.

## Projet Final

Le projet final consistera en la création d'une API RESTful complète, intégrant toutes les compétences acquises durant le cours. Le projet est un élément central de l'évaluation finale et devra inclure des fonctionnalités telles que :

- Gestion d'utilisateurs avec authentification.
- Gestion des rôles et des permissions.
- Accès sécurisé aux données via JWT.
- Tests unitaires et d'intégration pour les principales routes de l'API.

## Comment Utiliser ce Repository

Ce repository contient le code source, les exemples et les exercices pratiques pour chaque séance. Vous y trouverez :

- Des exemples de code pour illustrer chaque concept.
- Des exercices pratiques pour chaque module.
- Les ressources supplémentaires nécessaires au développement des projets.

Clonez ce repository pour accéder au contenu du cours, et n'hésitez pas à ouvrir une issue en cas de questions ou de problèmes.

## Contributeurs

- DENIS Victor

## Ressources Supplémentaires

- [Documentation officielle de Go](https://golang.org/doc/)
- [Introduction à Go](https://tour.golang.org/)

Nous vous souhaitons un excellent apprentissage du développement d'API avec Go ! 🚀

---

Voici un petit helper magique pour la fonction Update :

```go
package helper

import (
	"strings"

	"github.com/mitchellh/mapstructure"
)

func toCamelCase(input string) string {
	isToUpper := false
	var result string
	for i, v := range input {
		if i == 0 {
			result += strings.ToLower(string(v))
		} else if v == '_' {
			isToUpper = true
		} else {
			if isToUpper {
				result += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				result += string(v)
			}
		}
	}
	return result
}

// ApplyChanges function to decode map into struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	camelCaseKeys := make(map[string]interface{})
	for k, v := range changes {
		camelCaseKeys[toCamelCase(k)] = v
	}

	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
	})

	if err != nil {
		return err
	}

	return dec.Decode(camelCaseKeys)
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
```

La fonction s'utilise comme ça :

```go
var data map[string]interface{}

err := json.NewDecoder(r.Body).Decode(&data)
if err != nil {
    render.Render(w, r, errors.ErrInvalidRequest(err))
    return
}

company, err := config.CompanyRepository.FindByID(uint(idUint), true)

if err != nil {
    render.Render(w, r, errors.ErrServerError(err))
    return
}


helper.ApplyChanges(data, company)
```
