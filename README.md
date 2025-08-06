# 🎵 Groupie Tracker

## 📌 Description

**Groupie Tracker** est un projet web en Go qui consiste à consommer une API externe contenant des informations sur des groupes et artistes musicaux (noms, membres, dates, lieux de concerts, etc.) et à afficher ces données sur un site web convivial et dynamique.

Le but est de manipuler les données reçues, de les organiser et de les présenter sous forme de visualisations variées (cartes, tableaux, listes...) via une interface utilisateur claire et responsive.

---

## 🧠 Objectifs

- Consommer et manipuler une API RESTful contenant :
  - **Artists** : nom(s), image, date de création, premier album, membres.
  - **Locations** : villes où les concerts ont eu lieu ou auront lieu.
  - **Dates** : dates des concerts passés ou à venir.
  - **Relation** : relie les artistes avec leurs concerts (dates + lieux).

- Créer un **site web en Go** (backend) qui :
  - Affiche les informations des artistes.
  - Permet d'interagir avec le serveur via des événements (client → serveur).
  - Gère proprement les erreurs et ne plante jamais.

---

## 🛠️ Fonctionnalités

- 🔍 Recherche par nom d'artiste ou membre.
- 📅 Visualisation des dates de concerts.
- 🌍 Carte ou liste des lieux de concerts.
- 👥 Détails par artiste (biographie, albums, membres…).
- 🔁 Interactions en temps réel avec le serveur via requêtes HTTP.

---

## 🧪 Tests

- Des fichiers de test unitaires sont recommandés.
- Utilisation des tests Go standards (`go test`).

---

## ✅ Instructions

- Le **backend doit être écrit en Go** uniquement.
- Le site doit être **robuste** (pas de crash).
- Toutes les erreurs doivent être **gérées proprement**.
- Suivre les **bonnes pratiques** de développement.
- **Aucune bibliothèque externe** autorisée — uniquement les packages standards de Go.

---

## 🚀 Lancer le projet

### Prérequis

- [Go](https://golang.org/dl/)
- Un navigateur web

### Commandes

```bash
# Cloner le projet
git clone https://your-repo-url/groupie-tracker.git
cd groupie-tracker

# Lancer le serveur
go run main.go
