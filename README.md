# 🎮 PROJGO – Jeu PUISSANCE 4 (Go)

## 📖 Description
**PUISSANCE** **4** est un petit jeu développé en **Go (Golang)**.  
Le joueur selectionne son nom (facultatif), celuie du deuxieme joueur (facultatif), puis peut jouer au puissance 4.

Tout se joue directement dans le navigateur de votre choix.

---

## 🚀 Fonctionnalités principales
- 👤 **NOM DES JOUEURS** : choix du nom et de la classe (Humain, Elfe, Nain).  
- 📊 **SAUVEGARDER SA PARTIE** : HP, Mana, Argent, Expérience, Niveau.  
- ⚔ **ACCES A L'HISTORIQUE** : affrontements contre des gobelins d’entraînement, avec attaques, compétences et potions.  
- 🧪 **REJOUER SA PARTIE** : possibilité de consommer des potions de vie et de mana.  
- 🔥 **ABANDONNER SA PARTIE** : acquisition et utilisation de sorts comme *Fire Ball*.  
- 🛒 **RESET LA PARTIE** : achat de potions, matériaux, amélioration de l’inventaire, nouvelles compétences.  
  
---

## 🗂 Structure du projet
   ```bash
   PUISSANCE_4/
   │── src/
   │   ├── controller/
   │   │
   │   ├── grid/
   │   │
   │   ├── router/
   │   │
   │   ├── static/
   │   │
   │   │
   │   │
   │   ├── structure/
   │   │
   │   ├── template/
   │   │
   │   │
   │   │
   │   ├── utils/
   │   │
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
   git clone https://github.com/ton-compte/PROJETRED.git
   cd PROJETRED

2. Lancer le jeu :
   ```bash
    cd src
    go run main.go

