# Fiche Notion : Qu'est-ce qu'une API ? 🤔

## 1. Définition d'une API 📖

**API** est l'acronyme de **Application Programming Interface** (Interface de Programmation d'Application en français). En termes simples, une API est un ensemble de règles et de protocoles qui permet à différentes applications ou systèmes de communiquer entre eux.

Imaginez une API comme un **pont** entre deux logiciels qui ne se connaissent pas forcément, mais qui ont besoin de partager des informations. Une API permet à ces logiciels de s'échanger des données ou d'utiliser des fonctionnalités de l'un à l'autre, sans que l'utilisateur ait à se soucier de ce qui se passe en coulisses.

### ⚙️ Le rôle d'une API

Une API permet de :

- **Envoyer des requêtes** d'une application à une autre.
- **Recevoir des réponses** avec des données ou des actions effectuées par l'application qui reçoit la requête.
- **Standardiser** la manière dont ces échanges se font.

> 📌 **Exemple :** Imaginez que vous utilisez une application de météo sur votre téléphone. Cette application utilise une API pour demander les prévisions à un serveur météo. Le serveur répond avec les données demandées, et l'application affiche la météo à l'utilisateur.

---

## 2. Comment fonctionne une API ? 🔄

Une API fonctionne généralement via des **requêtes HTTP** (le protocole utilisé pour naviguer sur le web). Elle repose sur un modèle d'échange où une application "demande" des informations ou des actions à une autre, et cette dernière "répond" avec ce qui est nécessaire.

Voici les étapes clés d'une interaction avec une API :

1. **L'application cliente** (celle qui fait la demande) envoie une **requête** à une API.
   - Cette requête peut demander des données, déclencher une action, ou tout autre type de demande (comme un service ou une opération).
2. **Le serveur** (l'application qui reçoit la demande) traite la requête et **renvoie une réponse**.
   - Cette réponse contient soit les informations demandées, soit un message confirmant que l'action a été réalisée avec succès (ou non).

### 🚀 Les principales méthodes HTTP utilisées par une API :

- **GET** : Pour **demander** des données à une API. (Ex : obtenir la liste des utilisateurs d'un site)
- **POST** : Pour **envoyer** des données à l'API. (Ex : ajouter un nouvel utilisateur)
- **PUT** : Pour **mettre à jour** des données. (Ex : modifier les informations d'un utilisateur)
- **DELETE** : Pour **supprimer** des données. (Ex : supprimer un utilisateur)

> 📌 **Exemple :** Lorsque vous commandez un plat via une appli de livraison, cette application envoie une requête **POST** à une API pour créer votre commande, puis une requête **GET** pour récupérer l'état de la livraison.

---

## 3. Pourquoi les API sont-elles si importantes ? 🌍

Les API sont aujourd'hui partout et essentielles dans le monde numérique. Elles permettent aux entreprises et aux développeurs d'intégrer facilement des services entre différentes applications, sans avoir à tout développer en interne.

### Quelques raisons pour lesquelles les API sont importantes :

- **Modularité** : Une API permet de réutiliser des fonctionnalités sans réécrire le code. Par exemple, au lieu de créer son propre système de carte bancaire, une entreprise peut utiliser une API bancaire existante.
- **Interopérabilité** : Les API permettent à différentes applications de travailler ensemble, même si elles sont développées par des entreprises différentes. Par exemple, vous pouvez utiliser Google Maps dans une application Uber grâce à une API.

- **Économie de temps et de ressources** : En intégrant des API externes, les développeurs gagnent du temps et des ressources, car ils n'ont pas à développer des services qui existent déjà.

> 📌 **Exemple concret** : Imaginez que vous créiez une application de réservation de voyages. Grâce aux API, vous pouvez :
>
> - Intégrer un service de cartographie (comme Google Maps) pour montrer des itinéraires.
> - Connecter un système de paiement (comme PayPal) pour les transactions.
> - Obtenir des prévisions météo pour la destination via une autre API.

---

## 4. Types d'API 🔍

Il existe plusieurs types d'API en fonction de l'usage et de la manière dont elles sont partagées. Voici les principales catégories :

### a. **API publiques (ou ouvertes)** 🔓

Ces API sont ouvertes à tous les développeurs, qui peuvent les utiliser pour intégrer les services de l'entreprise dans leurs propres applications. Exemples d'API publiques :

- API de Google Maps
- API de Twitter
- API de Spotify

### b. **API privées** 🔒

Ces API sont utilisées en interne par une entreprise pour que ses différents systèmes communiquent entre eux. Elles ne sont pas accessibles au public.

### c. **API partenaires** 🤝

Ces API sont partagées avec des partenaires spécifiques pour permettre l'intégration des systèmes entre deux entreprises collaborant ensemble.

---

## 5. Cas concrets d'utilisation des API 🔧

Les API sont utilisées dans presque tous les domaines de la technologie moderne. Voici quelques exemples que vous rencontrez probablement tous les jours :

1. **Utiliser Google pour se connecter à une autre appli**

   - Lorsque vous vous connectez à un site ou une application avec votre compte Google, une **API Google** est utilisée pour récupérer vos informations et valider l'accès.

2. **Applications de paiement** 💳

   - Les API permettent aux applications comme Uber ou Deliveroo de traiter vos paiements via des services comme PayPal ou Stripe, sans que vous deviez saisir vos informations bancaires à chaque utilisation.

3. **Réseaux sociaux** 🌐

   - Lorsque vous partagez un article ou une vidéo directement sur Twitter depuis un autre site, c'est grâce à une **API de Twitter** qui permet d'envoyer ces données à votre compte.

4. **Applications météo** 🌦️
   - Les applications de météo utilisent des **API publiques** pour obtenir les prévisions météorologiques à jour et vous les afficher sur votre écran.

---

## 6. API et sécurité 🔒

L'utilisation d'une API nécessite de respecter des normes de sécurité pour protéger les données échangées. Voici quelques pratiques courantes :

- **Clés API** : Un identifiant unique donné à un utilisateur ou une application pour utiliser l'API. Cela permet de limiter l'accès et d'authentifier les requêtes.
- **OAuth** : Un protocole permettant aux utilisateurs de se connecter à des applications via des services tiers (Google, Facebook, etc.), sans partager leur mot de passe avec l'application.
- **Chiffrement** : Les requêtes et réponses entre les applications sont souvent chiffrées pour protéger les données sensibles.

---

### Résumé clé 🔑

Une API est un outil puissant qui permet aux applications de communiquer entre elles de manière standardisée. Elles sont partout autour de nous et facilitent la création d'applications riches et interconnectées sans avoir à tout développer de zéro. Grâce aux API, les applications peuvent partager des données, accéder à des services, et offrir des fonctionnalités que nous utilisons tous les jours.
