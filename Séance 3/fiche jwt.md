# 🌟 Fiche Théorique : Les JSON Web Tokens (JWT) 🚀

Bienvenue dans le monde des **JWT**, les super-héros de l’authentification moderne ! 🦸‍♂️🦸‍♀️ Aujourd’hui, on va découvrir ensemble **ce qu’est un JWT**, **comment il fonctionne** et pourquoi on l’utilise. Prêt·e à explorer ? C’est parti ! 🎉

---

## 🧐 C’est quoi un JWT ?

Un **JSON Web Token (JWT)**, c’est une petite chaîne de texte (un **token**) qui transporte de manière **sécurisée** des informations entre deux parties (souvent un **client** et un **serveur**). 🛡️

Il est souvent utilisé pour :

- **Authentifier un utilisateur** (qui tu es).
- **Autoriser l’accès** à des ressources protégées (ce que tu peux faire).

---

## 🧩 Comment est structuré un JWT ?

Un JWT est composé de **trois parties**, séparées par des points (`.`). 🟢🔵🟣

1. **Header (En-tête)** 📰  
   Il contient des informations sur le **type de token** (ici, JWT) et l’**algorithme** utilisé pour sécuriser le tout (ex : SHA-256).

   > Exemple : `{"alg": "HS256", "typ": "JWT"}`

2. **Payload (Charge utile)** 📦  
   C’est ici que résident les informations (appelées **claims**). Elles peuvent indiquer :

   - Qui est l’utilisateur (`sub`).
   - Quand le token a été émis (`iat`).
   - Quand il expire (`exp`).
     > **⚠️ Attention** : Ces infos ne sont pas cryptées, donc ne mettez jamais de données sensibles ici !

3. **Signature** ✍️  
   C’est la clé magique 🔑 qui garantit que personne n’a trafiqué le token. Elle est créée avec :
   - Le **Header**.
   - Le **Payload**.
   - Une **clé secrète** connue uniquement du serveur.

Un JWT ressemble à ça :

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

---

## 🔑 Pourquoi utiliser des JWT ?

### ✅ **Avantages** :

1. **Stateless** 🌐  
   Pas besoin de garder des sessions en mémoire côté serveur. Tout est dans le token !

2. **Sécurisé** 🛡️  
   La **signature** garantit que les données n’ont pas été modifiées.

3. **Portable** 🧳  
   Les JWT fonctionnent partout : sites web, applis mobiles, microservices… C’est comme un passeport international ! 🌍

---

### ⚠️ **Attention aux inconvénients** :

1. **Taille du token** 🐘  
   Les JWT peuvent être volumineux (surtout si le payload est chargé).

2. **Expiration** ⏰  
   Une fois délivré, un JWT ne peut pas être invalidé avant son expiration (à moins d’avoir un système de blacklist).

3. **Pas d’infos sensibles dans le Payload** 🚨  
   Les JWT sont **encodés** (Base64) mais pas **cryptés**. Quelqu’un peut les lire !

---

## 🏗️ Comment ça marche dans une API ?

1. **Étape 1 : Authentification initiale** 👤  
   L’utilisateur se connecte en fournissant ses identifiants (ex : email + mot de passe).

2. **Étape 2 : Génération du token** 🪄  
   Si les identifiants sont valides, le serveur génère un JWT contenant des infos sur l’utilisateur (ex : `id`, `roles`, etc.) et le renvoie au client.

3. **Étape 3 : Accès aux ressources** 📂  
   À chaque requête, le client inclut le JWT dans le **header Authorization** comme ceci :

   ```
   Authorization: Bearer <token>
   ```

   Le serveur vérifie le token et accorde (ou non) l’accès.

4. **Étape 4 : Expiration et rafraîchissement** 🔄  
   Si le token expire, un mécanisme de **refresh token** peut être utilisé pour prolonger la session.

---

## 🛠️ Les "claims" dans le Payload

Les **claims** sont comme les cartes d’identité de ton token. Ils contiennent des infos sur :

- **L’utilisateur** (ex : `sub` = identifiant de l’utilisateur).
- **Le contexte** (ex : `role`, `permissions`).
- **Les règles de validité** (ex : `exp` = expiration).

### Les claims les plus courants :

- `iss` : L’émetteur du token (qui l’a créé).
- `sub` : Le sujet (à qui appartient le token).
- `aud` : L’audience (qui peut utiliser ce token).
- `exp` : La date d’expiration (quand le token expire).
- `iat` : La date d’émission (quand le token a été créé).
- `nbf` : "Not Before" (date avant laquelle le token est invalide).

---

## 🎯 Bonnes pratiques avec les JWT

1. **Toujours utiliser HTTPS** 🌐  
   Pour éviter que quelqu’un vole ton token en plein vol ! 🕵️‍♂️

2. **Mettre une date d’expiration courte** ⏳  
   Les tokens doivent avoir une durée de vie limitée. Pas d’éternité ici !

3. **Ne jamais stocker de données sensibles dans le token** ❌  
   Ex : Pas de mots de passe, d’infos bancaires, etc.

4. **Stocker le JWT dans un endroit sécurisé côté client** 🔒

   - Préfère les **cookies httpOnly** pour plus de sécurité.
   - Si tu utilises le local storage, fais attention aux attaques XSS.

5. **Mettre en place un mécanisme de Refresh Token** 🔄  
   Permet de prolonger la session sans redemander à l’utilisateur de se connecter.

---

## 🌈 En résumé

Les **JWT**, c’est comme un **passeport numérique**. Ils te permettent de prouver qui tu es et d'accéder à des ressources sans avoir à te réauthentifier à chaque fois. Ils sont sécurisés, efficaces et ultra-pratiques dans une architecture API moderne. 💪

Avec ces bases, tu es prêt·e à plonger dans l’univers des JWT et à les utiliser dans tes propres projets ! 🚀
