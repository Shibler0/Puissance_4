# 🎮 PROJGO – Jeu PUISSANCE 4 (Go)

## 📖 Description
**PUISSANCE** **4** est un petit jeu développé en **Go (Golang)**.  
Le joueur selectionne son nom (facultatif), celuie du deuxieme joueur (facultatif), puis peut jouer au puissance 4.

Tout se joue directement dans le navigateur de votre choix.

---

## 🚀 Fonctionnalités principales
- 👤 **NOM DES JOUEURS** : choix du nom des deux joueurs.  
- ✅ **SAUVEGARDER SA PARTIE** : mettre en pause pour continuer plus tard.
- 🔄 **CONTINUER SA PARTIE** : continuer une partie mise en pause. 
- 📋 **ACCES A L'HISTORIQUE** : accées a toutes les parties jouer précédement.    
- 🏳️ **ABANDONNER SA PARTIE** : refuser de poursuivre une partie en cours, la supprime definitivement.  
- 🚨 **RESET LA PARTIE** : reinitialise la grille du puissance 4, dans la partie.  
  
---

## 🗂 Structure du projet
   ```bash
   PUISSANCE_4/
   │── src/
   │   ├── controller/
   │   │   └── controller.go
   │   ├── grid/
   │   │   └── grid.go
   │   ├── router/
   │   │   └── router.go
   │   ├── static/
   │   │   ├── style.css
   │   │   └── stylehome.css
   │   ├── structure/
   │   │   └── structure.go
   │   ├── template/
   │   │   ├── home.html
   │   │   ├── play.html
   │   │   └── save.html
   │   ├── utils/
   │   │   └── utils.go
   │   ├── gamehistoric.json
   │   ├── gamesave.json
   │   ├── go.mod
   │   └── main.go
   ├── .gitignore
   └── README.md
    
   ```
---

## ▶️ Lancer le projet
1. Cloner le dépôt :  
   ```bash
   git clone https://github.com/ton-compte/PUISSANCE_4.git
   cd PUISSANCE_4

2. Lancer le jeu :
   ```bash
    cd src
    go run main.go

