# Go WebSocket Chat Application

Ce projet est une application de chat simple utilisant WebSocket en Go. Il comprend une interface frontend en HTML pour se connecter à différents salons de discussion, envoyer des messages, et se connecter avec un système d'authentification basique.

Ce projet a pour objectif d'explorer et d'expérimenter avec l'utilisation des WebSockets en Go. Il s'agit d'une initiative d'apprentissage visant à comprendre les bases de la communication bidirectionnelle en temps réel entre un serveur Go et des clients, ainsi qu'à expérimenter diverses fonctionnalités et cas d'utilisation possibles.

## Fonctionnalités

- Se connecter à un salon de chat via WebSocket
- Envoyer et recevoir des messages en temps réel
- Changer de salon de chat
- Authentification basique avec un nom d'utilisateur et un mot de passe
- Communication sécurisée avec TLS

## Technologies utilisées

- Backend : Go avec le package `gorilla/websocket` pour la gestion des WebSocket
- Frontend : HTML et JavaScript pour l'interface utilisateur
- Sécurité : Certificats TLS générés avec OpenSSL

## Installation et configuration

### Prérequis

- [Go](https://golang.org/dl/) installé (version 1.16+ recommandée)
- [OpenSSL](https://www.openssl.org/) pour générer les certificats TLS

### Étapes d'installation

1. **Clone le projet :**
    ```bash
    git clone <url_du_projet>
    cd <nom_du_projet>
    ```

2. **Générer les certificats TLS :**

Ce projet utilise HTTPS et nécessite un certificat et une clé pour démarrer le serveur. Un script bash est inclus pour générer ces fichiers avec OpenSSL.

    ```bash
    bash ./gencert.bash
    ```

3. **Démarrer le serveur :**

    ```bash
    go run *.go
    ```

    > pour le dev vous pouvez utiliser `air` pour recharger automatiquement le serveur
    > ```bash
    > air
    > ```

4. **Accéder à l'application de chat :** Ouvrez votre navigateur et allez à l'adresse https://localhost:8080.

## Fonctionnalités de l'interface utilisateur

- **Chat** : Vous pouvez envoyer des messages dans le salon de discussion et voir les messages des autres utilisateurs en temps réel.
- **Changer de salon** : Vous pouvez choisir un nouveau salon de discussion en saisissant un nom dans le champ Chatroom.
- **Authentification** : Utilisez admin comme nom d'utilisateur et 123 comme mot de passe pour vous connecter. Après la connexion, une session WebSocket sera établie.

## Structure du projet

- main.go : Le point d'entrée du serveur HTTP.
- client.go : Gère la connexion WebSocket pour chaque client.
- manager.go : Gère les clients, les événements, et la logique des salons de chat.
- event.go : Définition des types d'événements utilisés dans l'application.
- frontend/index.html : L'interface utilisateur côté client avec les formulaires de connexion et de chat.
- gencert.bash : Script pour générer les certificats SSL avec OpenSSL.

## Sécurité

Le projet utilise des connexions WebSocket sécurisées avec TLS (wss://). Les certificats auto-signés générés sont utilisés pour cette démonstration.

## Améliorations possibles

- Authentification renforcée : Ajouter un vrai système de gestion d'utilisateurs avec des bases de données.
- Reconnexion automatique : Mettre en place un système de reconnexion en cas de déconnexion du WebSocket.
- Gestion avancée des salons : Ajouter la création dynamique des salons avec une persistance en base de données.

## Crédits

Ce projet est une expérimentation pour explorer les WebSocket en Go avec une interface frontend basique. Il utilise le package `gorilla/websocket` pour la communication WebSocket.

Cette expérientation est inspirée par le tutoriel de [ProgrammingPercy](https://www.youtube.com/watch?v=pKpKv9MKN-E).
