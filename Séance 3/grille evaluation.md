# Grille d'évaluation pour l'ajout de l'authentification avec JWT 🛡️

## **A. Implémentation de l'authentification (`/login`)** (20 points)

| Critère                           | Excellent                                                                          | Bon                                                       | Satisfaisant                                              | Insuffisant                          | Non réalisé              |
| --------------------------------- | ---------------------------------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | ------------------------------------ | ------------------------ |
| **Vérification des identifiants** | Vérification robuste avec hachage des mots de passe et gestion précise des erreurs | Vérification correcte mais gestion des erreurs incomplète | Vérification fonctionnelle mais peu sécurisée ou fragile  | Vérification basique ou incomplète   | Pas de vérification      |
| **Génération du token JWT**       | Token bien généré avec claims pertinents (email, exp) et clé sécurisée             | Token généré avec configuration acceptable                | Token généré mais claims ou sécurité insuffisants         | Génération incorrecte ou incomplète  | Pas de génération        |
| **Gestion des réponses**          | Réponses claires et structurées (200, 401 avec messages explicites)                | Réponses majoritairement claires mais perfectibles        | Réponses suffisantes mais manquent de clarté ou cohérence | Réponses peu adaptées ou incorrectes | Pas de réponses adaptées |

---

## **B. Middleware d'authentification** (20 points)

| Critère                             | Excellent                                                               | Bon                                                    | Satisfaisant                                         | Insuffisant                            | Non réalisé                |
| ----------------------------------- | ----------------------------------------------------------------------- | ------------------------------------------------------ | ---------------------------------------------------- | -------------------------------------- | -------------------------- |
| **Validation des tokens**           | Middleware robuste, vérifie et extrait les données avec précision       | Middleware valide les tokens mais manque de robustesse | Middleware fonctionnel mais partiellement implémenté | Middleware incomplet ou dysfonctionnel | Pas de middleware          |
| **Gestion des erreurs**             | Gestion précise des erreurs avec messages clairs (401 pour les erreurs) | Gestion des erreurs correcte mais améliorable          | Gestion suffisante mais manque de cohérence          | Gestion peu claire ou insuffisante     | Pas de gestion des erreurs |
| **Ajout d'utilisateur au contexte** | Ajout réussi et exploitable dans les routes protégées                   | Ajout fonctionnel mais partiellement exploitable       | Ajout partiellement fonctionnel mais non optimisé    | Ajout mal implémenté ou inutilisable   | Pas d’ajout au contexte    |

---

## **C. Protection des routes** (20 points)

| Critère                       | Excellent                                                                             | Bon                                                | Satisfaisant                                   | Insuffisant                                  | Non réalisé                |
| ----------------------------- | ------------------------------------------------------------------------------------- | -------------------------------------------------- | ---------------------------------------------- | -------------------------------------------- | -------------------------- |
| **Application du middleware** | Middleware appliqué à toutes les routes critiques (chats, consultations, traitements) | Middleware appliqué mais certaines routes manquent | Middleware appliqué mais protection incomplète | Middleware appliqué sur très peu de routes   | Pas de middleware appliqué |
| **Restrictions d'accès**      | Accès bien limité aux utilisateurs authentifiés, cohérence totale                     | Restrictions en place mais quelques failles        | Restrictions fonctionnelles mais basiques      | Restrictions mal implémentées ou inefficaces | Pas de restrictions        |

---

## **D. Gestion des tokens** (20 points)

| Critère                   | Excellent                                            | Bon                                                      | Satisfaisant                            | Insuffisant                            | Non réalisé                |
| ------------------------- | ---------------------------------------------------- | -------------------------------------------------------- | --------------------------------------- | -------------------------------------- | -------------------------- |
| **Expiration des tokens** | Expiration bien gérée et durée appropriée            | Expiration gérée mais durée ou configuration perfectible | Expiration fonctionnelle mais simpliste | Expiration mal configurée ou inadaptée | Pas d’expiration gérée     |
| **Sécurité des tokens**   | Clé sécurisée et bien protégée, configuration solide | Clé sécurisée mais amélioration possible                 | Configuration basique mais acceptable   | Configuration peu sécurisée            | Pas de sécurité des tokens |

---

## **E. Documentation et clarté du code** (20 points)

| Critère                                | Excellent                                                      | Bon                                                | Satisfaisant                                     | Insuffisant                                | Non réalisé           |
| -------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------- | ------------------------------------------------ | ------------------------------------------ | --------------------- |
| **Documentation utilisateur**          | Instructions claires pour `/login` et l'utilisation des tokens | Documentation présente mais incomplète             | Documentation minimale mais compréhensible       | Documentation peu claire ou partielle      | Pas de documentation  |
| **Lisibilité et organisation du code** | Code bien structuré, lisible et commenté                       | Code correct mais quelques améliorations possibles | Code acceptable mais confusion dans la structure | Code difficilement lisible ou mal organisé | Code incompréhensible |

---

## **Challenge Optionnel (Bonus : 10 points)**

| Critère                | Excellent                                                        | Bon                                               | Satisfaisant                              | Insuffisant                                   | Non réalisé          |
| ---------------------- | ---------------------------------------------------------------- | ------------------------------------------------- | ----------------------------------------- | --------------------------------------------- | -------------------- |
| **Rôles utilisateurs** | Implémentation robuste et fonctionnelle                          | Rôles implémentés mais partiellement fonctionnels | Rôles implémentés mais simplistes         | Rôles mal configurés ou inutilisables         | Pas de rôles         |
| **Refresh token**      | Fonctionnalité complète et sécurisée pour prolonger les sessions | Fonctionnalité en place mais perfectible          | Implémentation basique mais fonctionnelle | Implémentation partielle ou non fonctionnelle | Pas de refresh token |

---

### Total : **100 points** (+10 bonus)

| Score global        | Niveau       |
| ------------------- | ------------ |
| **90-100** (+bonus) | Excellent    |
| **75-89**           | Très bon     |
| **60-74**           | Bon          |
| **50-59**           | Satisfaisant |
| **<50**             | Insuffisant  |
