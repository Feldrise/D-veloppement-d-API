# Projet Final : Développement d'une API RESTful pour une Application Mobile 🚀

## Objectifs 🎯

Votre mission principale est de concevoir une **API RESTful** en Go, qui servira de base pour votre future application mobile. Cette API devra répondre aux standards de qualité, intégrer une base de données, et inclure un système d'authentification sécurisé. Ce projet vise à consolider vos compétences en développement d'API, tout en préparant un socle robuste pour une application mobile moderne.

---

## Étapes du Projet 🛠️

### 1. **Définition des Fonctionnalités**

- Définissez les fonctionnalités principales que votre API devra offrir en fonction des besoins de l'application mobile que vous développerez.
  - Exemple de fonctionnalités possibles : gestion des utilisateurs, gestion d'un contenu spécifique (articles, produits, tâches), gestion de permissions spécifiques.
- Formalisez rapidement les besoins techniques dans un petit document de conception qui détaille les routes principales, les modèles de données, et les relations entre les entités.

---

### 2. **Architecture et Organisation**

- Structurez votre projet Go de manière modulaire pour séparer clairement les responsabilités (ex. : modules pour l'authentification, les utilisateurs, et les données spécifiques).
- Définissez et organisez les routes principales de l'API :
  - Exemple : `/api/v1/users` pour la gestion des utilisateurs, `/api/v1/tasks` pour des tâches.

---

### 3. **Connexion à une Base de Données**

- Configurez une base de données relationnelle (PostgreSQL ou MySQL).
- Utilisez **GORM** pour interagir avec votre base de données et définissez vos modèles en respectant les relations (1-N, N-N).
- Créez et documentez vos migrations pour que la structure de la base soit réplicable.

---

### 4. **Implémentation des Routes et des Fonctions**

- Implémentez les routes nécessaires en respectant les verbes HTTP (GET, POST, PUT, DELETE) et en renvoyant des réponses cohérentes.
- Sécurisez vos routes sensibles en utilisant des middlewares pour l'authentification et l'autorisation.

---

### 5. **Gestion de l'Authentification**

- Implémentez un système d'authentification basé sur **JWT (JSON Web Tokens)** ou un autre mécanisme sécurisé.
- Définissez des règles claires pour les accès utilisateur, par exemple :
  - Utilisateur connecté vs non connecté.
  - Droits spécifiques à certains utilisateurs (administrateur, modérateur, etc.).

---

### 6. **Documentation de l'API**

- Documentez les fonctionnalités de votre API en utilisant un outil tel que **Swagger** ou un autre générateur de documentation.
- Veillez à inclure des descriptions claires pour chaque endpoint (paramètres, réponses attendues, et codes d'erreur).

---

### 7. **Tests et Validation**

- Testez votre API avec des outils comme Postman ou Insomnia.
- Validez les cas d'erreur (400, 404, 500) et assurez-vous que les réponses sont explicites.

---

### 8. **Préparation pour l'Intégration Mobile**

- Assurez-vous que votre API est prête à être consommée par une application Flutter :
  - Données structurées en JSON.
  - Routes et endpoints documentés.
  - Réponses rapides et adaptées aux besoins d'une application mobile.

---

### 9. **Présentation Finale**

Lors de la séance finale, vous devrez présenter votre API sous la forme d'une démonstration interactive et technique. Voici ce qui est attendu pour la présentation :

#### a) **Structure et Fonctionnalités**

- Expliquez l'architecture de votre API (routes principales, organisation des fichiers, connexion à la base de données).
- Présentez les fonctionnalités principales que votre API offre.

#### b) **Démonstration Technique**

- Effectuez des appels à vos routes (via Postman, Bruno ou curl) pour démontrer :
  - La création et l'authentification d'un utilisateur.
  - Les fonctionnalités principales de gestion de données (CRUD).
  - Les aspects de sécurité, comme les routes protégées par authentification.

#### c) **Documentation**

- Présentez votre documentation Swagger (ou autre) en montrant comment elle aide à comprendre et utiliser l'API.
- Mentionnez les cas d'utilisation spécifiques pour une application mobile.

#### d) **Retour sur le Développement**

- Discutez des défis rencontrés et des choix techniques réalisés :
  - Pourquoi tel modèle de données ?
  - Pourquoi JWT pour l'authentification ?
- Mentionnez également les améliorations ou extensions que vous envisagez pour cette API.

#### e) **Code Source et Collaboration**

- Montrez où se trouve votre code source (dépôt GitHub).
- Expliquez comment le projet a été organisé au sein de votre équipe.

---

### Livrables 📦

1. **Code Source** : Déposé sur GitHub avec un `README.md` clair, incluant :
   - Les instructions pour exécuter l'API.
   - La description des principales fonctionnalités.
2. **Base de Données** : Script des migrations ou dump SQL.
3. **Documentation de l'API** : Export Swagger ou équivalent.
4. **Présentation** : Préparez une démonstration interactive pour la session finale.

---

## Ressources 📚

- [Documentation Go](https://go.dev/doc/)
- [GORM Documentation](https://gorm.io/docs/)
- [Swagger](https://swagger.io/)
- [JWT Introduction](https://jwt.io/introduction)

💡 **Conseil : Collaborez efficacement, testez votre code régulièrement, et documentez chaque étape de votre progression. Une API bien conçue est un atout précieux pour n'importe quel projet mobile.**
