# TP : Authentification avec JWT 🎟️🔒

## Partie 1 : Découverte des JWT avec un Endpoint de Login 🎉

**Durée estimée :** 45 minutes

### Objectifs

Dans ce TP, vous allez apprendre à intégrer une **authentification avec JWT** dans une API REST. Vous mettrez en place une route pour permettre aux utilisateurs de se connecter et de recevoir un **token JWT**, puis vous utiliserez ce token pour sécuriser d’autres routes.

Au programme :

1. Création d’un **endpoint de login**.
2. Génération de **tokens JWT**.
3. Ajout de protections pour les routes privées. 🚀

---

### Étapes

### **1. Préparation du projet**

- Créez un projet appelé `jwt-auth-api` avec la structure suivante :

  ```
  jwt-auth-api/
  ├── authentication/
  │   ├── controller.go   # Gestion des routes d'authentification
  │   ├── jwt.go          # Logique pour générer et valider les tokens
  │   ├── middleware.go   # Middleware pour sécuriser les routes
  ├── main.go             # Point d'entrée de l'application
  ```

- **Initialisez votre module Go :**

  ```bash
  go mod init jwt-auth-api
  ```

- **Installez les dépendances :**
  ```bash
  go get github.com/golang-jwt/jwt/v4
  go get github.com/go-chi/chi/v5
  go get golang.org/x/crypto/bcrypt
  ```

---

### **2. Créez l’Endpoint de Login** 🚪

1. Dans `authentication/controller.go`, implémentez un endpoint `/login` qui permet de vérifier l'email et le mot de passe d’un utilisateur.
2. Utilisez une base d’utilisateurs simulée :

   ```go
   var users = map[string]string{
       "bob@example.com": " $2a$14$saaZ6k5RPBVZMeEySLOiduHkLXs8BioPolfcA4ysbd591hBW48x2.", // password: "password"
   }
   ```

3. Ajoutez une fonction pour vérifier les identifiants, puis générez un JWT si la connexion est réussie.

   Exemple :

   ```go
   func Login(w http.ResponseWriter, r *http.Request) {
       var payload struct {
           Email    string `json:"email"`
           Password string `json:"password"`
       }
       json.NewDecoder(r.Body).Decode(&payload)

       hashedPassword, exists := users[payload.Email]
       if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password)) != nil {
           http.Error(w, "Invalid email or password", http.StatusUnauthorized)
           return
       }

       token, err := GenerateToken("your_secret_key", payload.Email)
       if err != nil {
           http.Error(w, "Failed to generate token", http.StatusInternalServerError)
           return
       }

       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(map[string]string{"token": token})
   }
   ```

---

### **3. Génération des Tokens JWT** 🔑

1. Dans `authentication/jwt.go`, implémentez une fonction pour générer un token JWT contenant des informations sur l’utilisateur :

   ```go
   func GenerateToken(secret, email string) (string, error) {
       token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
           "email": email,
           "exp":   time.Now().Add(time.Hour * 2).Unix(),
       })
       return token.SignedString([]byte(secret))
   }
   ```

2. Ajoutez également une fonction pour valider les tokens :

   ```go
   func ParseToken(secret, tokenString string) (string, error) {
       token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
           return []byte(secret), nil
       })

       if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
           return claims["email"].(string), nil
       }
       return "", err
   }
   ```

---

### **4. Middleware pour Sécuriser les Routes** 🔒

1. Dans `authentication/middleware.go`, implémentez un middleware pour vérifier le JWT et extraire les informations de l’utilisateur :

   ```go
   func AuthMiddleware(secret string) func(http.Handler) http.Handler {
       return func(next http.Handler) http.Handler {
           return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
               authHeader := r.Header.Get("Authorization")
               if authHeader == "" {
                   http.Error(w, "Missing token", http.StatusUnauthorized)
                   return
               }

               email, err := ParseToken(secret, authHeader)
               if err != nil {
                   http.Error(w, "Invalid token", http.StatusUnauthorized)
                   return
               }

               ctx := context.WithValue(r.Context(), "email", email)
               next.ServeHTTP(w, r.WithContext(ctx))
           })
       }
   }
   ```

2. Ajoutez une fonction pour extraire l’utilisateur depuis le contexte :
   ```go
   func GetUserFromContext(ctx context.Context) string {
       email, _ := ctx.Value("email").(string)
       return email
   }
   ```

---

### **5. Ajoutez une Route Protégée** 🛡️

1. Dans `main.go`, ajoutez une route sécurisée `/protected` qui affiche un message de bienvenue si l’utilisateur est authentifié.

   Exemple :

   ```go
   r.Group(func(r chi.Router) {
       r.Use(authentication.AuthMiddleware("your_secret_key"))
       r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
           user := authentication.GetUserFromContext(r.Context())
           w.Write([]byte(fmt.Sprintf("Welcome, %s!", user)))
       })
   })
   ```

2. **Testez la route** avec et sans token.

---

🎉 **Bravo !** Vous venez de mettre en place une API REST sécurisée avec JWT ! Vous avez appris à :

- Générer des tokens JWT.
- Vérifier et valider ces tokens.
- Protéger des routes privées avec un middleware.

**Challenge supplémentaire** : Ajoutez des rôles ou permissions dans les claims JWT pour gérer des accès différenciés (ex : `admin`, `user`). 🚀
