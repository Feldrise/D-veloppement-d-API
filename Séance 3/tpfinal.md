# Projet Final : Ajout de l'Authentification avec JWT à l'API de la Clinique Vétérinaire 🐾🔒

## Contexte

Bravo ! Vous avez déjà développé une API pour une clinique vétérinaire 🐾. Maintenant, il est temps de franchir une nouvelle étape en **sécurisant votre API**. Dans ce projet, vous allez ajouter un système d'authentification basé sur **JWT** pour protéger vos routes et limiter l'accès aux utilisateurs authentifiés.

Ce projet final est moins guidé : vous devrez mobiliser toutes vos compétences pour intégrer cette nouvelle fonctionnalité.

---

## Objectifs

1. Ajouter une **authentification utilisateur** à votre API.
2. Générer et valider des **JWT** pour gérer les sessions utilisateur.
3. Protéger les routes critiques (comme la gestion des chats et des consultations).
4. Gérer les erreurs liées à l’authentification (ex : tokens expirés ou invalides).

---

## Travail attendu

### 1. **Ajout de l'authentification utilisateur**

- **Fonctionnalité** : Implémentez une nouvelle route `/login` qui permet à un utilisateur de se connecter.
- **Détails** :
  - Le client envoie ses identifiants (email et mot de passe).
  - Si l’authentification est réussie, un **token JWT** est généré et renvoyé au client.
  - Si l’authentification échoue, retournez une réponse HTTP appropriée (`401 Unauthorized`).

---

### 2. **Sécurisation des routes**

- **Fonctionnalité** : Ajoutez un middleware pour vérifier les **tokens JWT** sur certaines routes.
- **Détails** :
  - Toutes les routes liées à la gestion des chats, des consultations, et des traitements doivent être sécurisées.
  - Si un token est invalide ou manquant, retournez une réponse HTTP `401 Unauthorized`.

---

### 3. **Gestion des utilisateurs**

- Vous pouvez choisir d’ajouter une table pour les utilisateurs dans votre base de données **OU** utiliser une base d’utilisateurs simulée.
- **Exigences** :
  - Chaque utilisateur doit avoir un email et un mot de passe.
  - Les mots de passe doivent être **hachés** avec `bcrypt`.

---

### 4. **Expirations et sécurité**

- **Token d'accès** :
  - Ajoutez une expiration courte pour les **access tokens** (par exemple, 1 heure).
- **Option avancée** (Challenge) :
  - Implémentez un mécanisme de **refresh token** pour prolonger les sessions utilisateur.

---

## Consignes techniques

### Architecture

- Votre projet doit rester **structuré et modulaire**, comme spécifié dans le projet initial.
- Ajoutez une nouvelle **feature** `authentication` pour gérer l’authentification.

### Points importants

1. **Validation des tokens** :

   - Vérifiez la validité des tokens dans le middleware.
   - Chargez les informations utilisateur depuis le token pour les ajouter au contexte de la requête.

2. **Gestion des erreurs** :
   - Fournissez des réponses claires pour les erreurs d’authentification (ex : token expiré, utilisateur inexistant).

---

## Exigences fonctionnelles

1. **Route de connexion (`/login`)** :

   - Vérifie les identifiants et renvoie un JWT en cas de succès.
   - Gère les erreurs pour les identifiants invalides.

2. **Middleware d'authentification** :

   - Vérifie que chaque requête inclut un JWT valide.
   - Retourne une erreur 401 pour les requêtes non authentifiées.

3. **Protection des routes** :

   - Toutes les routes existantes pour les chats, consultations, et traitements doivent être protégées.

4. **Gestion des expirations** :
   - Les tokens doivent expirer après un certain temps.

---

## Livrables

1. Code source mis à jour avec l’authentification intégrée.
2. Documentation dans un fichier `README.md` expliquant comment utiliser l’authentification :
   - Comment générer un token (via `/login`).
   - Comment utiliser ce token pour accéder aux routes sécurisées.

---

## Évaluation

Une grille d'évaluation sera utilisée pour juger votre travail.

---

## Challenge optionnel 🌟

1. **Ajoutez des rôles utilisateur** :

   - Par exemple : un rôle `admin` qui peut modifier les données, et un rôle `user` qui peut uniquement lire les données.

2. **Ajoutez un refresh token** :
   - Implémentez une route `/refresh` pour générer un nouveau access token à partir d’un refresh token valide.

---

### À vous de jouer !

**Vous avez toutes les clés en main pour sécuriser votre API et montrer vos compétences !** 🚀
