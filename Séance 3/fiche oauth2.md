# 🌟 Fiche Théorique : OAuth 2.0 🚀

Bienvenue dans le monde de **OAuth 2.0**, un standard puissant et indispensable pour l’authentification et l’autorisation dans les applications modernes. 🌐 Aujourd’hui, on va découvrir ce qu’est **OAuth 2.0**, comment il fonctionne et pourquoi il est devenu incontournable. Prêt·e à te transformer en expert·e ? C’est parti ! 🎉

---

## 🧐 C’est quoi OAuth 2.0 ?

**OAuth 2.0**, c’est un protocole qui permet à une application (le **client**) de demander l’accès à des ressources protégées au nom d’un utilisateur, **sans jamais avoir accès à son mot de passe**. 🛡️

Imagine que tu veux utiliser une app pour planifier tes vacances, et cette app a besoin d'accéder à tes photos sur Google Photos. Avec **OAuth 2.0**, tu autorises cette app à accéder uniquement à tes photos, sans lui donner ton mot de passe Google. ✨

---

## 🤓 Les rôles principaux dans OAuth 2.0

1. **Resource Owner** 🧑  
   La personne qui possède les ressources protégées (toi, l’utilisateur·rice).

2. **Resource Server** 🗄️  
   Le service qui héberge les ressources (ex : l’API de Google Photos).

3. **Client** 📱  
   L’application qui demande l’accès aux ressources (ex : ton application de planning).

4. **Authorization Server** 🛡️  
   Le serveur qui gère les autorisations et délivre les **tokens** (ex : les serveurs OAuth de Google).

---

## 🔑 Les tokens dans OAuth 2.0

OAuth repose sur deux types de **tokens** pour autoriser l'accès :

1. **Access Token** 🎟️

   - Permet au client d'accéder à une ressource protégée.
   - A une durée de vie limitée (quelques minutes ou heures).

2. **Refresh Token** 🔄
   - Permet de demander un nouvel access token quand celui-ci expire, sans re-demander l'autorisation à l'utilisateur.

---

## 🧩 Les différents **flows** (flux) d’OAuth 2.0

OAuth 2.0 propose plusieurs **flux** adaptés à différents scénarios. Les plus courants sont :

### 1. **Authorization Code Flow** 🗝️

👉 Utilisé pour les applications web ou backend.  
 **Étapes** :

- L’utilisateur donne son autorisation via une interface sécurisée.
- Le client reçoit un **code d’autorisation**.
- Ce code est échangé contre un **access token** auprès du serveur d’autorisation.

### 2. **Implicit Flow** 🚀 _(à éviter aujourd’hui)_

👉 Utilisé pour les applications frontend (sans serveur backend).  
 ⚠️ Ce flux est maintenant considéré comme moins sécurisé car les tokens sont exposés côté client.

### 3. **Client Credentials Flow** 🤝

👉 Utilisé quand une application accède directement à des ressources **sans utilisateur**.  
 Exemple : un microservice accédant à une API interne.

### 4. **Password Grant** 🔐 _(déprécié)_

👉 Utilisé pour les applications qui demandent directement les identifiants de l’utilisateur.  
 ⚠️ Fortement déconseillé car le client a accès au mot de passe.

---

## 🛠️ Comment ça fonctionne ?

### 🚦 Exemple simple d’Authorization Code Flow :

1. **Étape 1 : L’utilisateur s’authentifie** 🔑

   - Le client redirige l’utilisateur vers l’**Authorization Server** (ex : Google).
   - L’utilisateur entre ses identifiants et autorise l’accès.

2. **Étape 2 : Le client reçoit un code d’autorisation** 📜

   - Une fois l’utilisateur authentifié, l’Authorization Server renvoie un **code temporaire** au client.

3. **Étape 3 : Le client échange le code contre un token** 🎟️

   - Le client envoie le **code** et ses **identifiants client** (client ID et secret) au serveur pour obtenir un **Access Token** (et parfois un Refresh Token).

4. **Étape 4 : L’accès est autorisé** ✅
   - Le client utilise l’Access Token pour accéder aux ressources protégées sur le **Resource Server**.

---

## 🎯 Pourquoi utiliser OAuth 2.0 ?

### ✅ **Avantages** :

1. **Sécurité** 🔒

   - Pas besoin de partager ton mot de passe avec l’application cliente.

2. **Granularité** 🧐

   - Tu peux accorder des accès limités (ex : lire tes photos mais pas les supprimer).

3. **Expiration des permissions** ⏳

   - Les tokens ont une durée de vie courte, limitant l’impact en cas de vol.

4. **Interopérabilité** 🌍
   - OAuth est un standard largement adopté (Google, Facebook, GitHub, etc.).

---

### ⚠️ **Inconvénients** :

1. **Complexité** 🌀

   - L’implémentation est plus compliquée qu’un simple login/mot de passe.

2. **Dépendance** 🤝
   - Tu dépends du bon fonctionnement de l’Authorization Server.

---

## 🚨 Bonnes pratiques avec OAuth 2.0

1. **Toujours utiliser HTTPS** 🌐  
   Pour protéger les tokens en transit.

2. **Protéger les secrets clients** 🔑  
   Ne partage jamais ton client ID et ton secret.

3. **Configurer les scopes intelligemment** 🧐  
   Donne à l’application uniquement les accès nécessaires.

4. **Utiliser le PKCE (Proof Key for Code Exchange)** 🔒  
   PKCE renforce la sécurité du flux d’Authorization Code, surtout pour les applis frontend.

5. **Mettre en place une gestion efficace des tokens expirés** ⏰  
   Implémente le Refresh Token Flow pour prolonger les sessions.

---

## 🌈 En résumé

OAuth 2.0, c’est **la solution moderne pour déléguer l’accès** en toute sécurité, sans jamais exposer tes identifiants. C’est un standard incontournable pour les applications qui s’interconnectent avec des services tiers. 🌐

Avec ses **flux variés** et ses mécanismes robustes, OAuth s’impose comme un pilier de l’écosystème API moderne. Alors, prêt·e à sécuriser tes applications comme un·e pro ? 💪 🚀
