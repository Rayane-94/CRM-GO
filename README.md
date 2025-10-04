# CRM-GO

CRM-GO est une application CRM développée en **Go**, permettant de gérer et organiser des données clients de manière simple et efficace.

---

## Description

Ce projet contient les fichiers principaux pour démarrer une application CRM en Go.  
Il permet d’ajouter, modifier et tester des fonctionnalités CRM via des tests unitaires.

---

## Architecture

CRM-GO/
├── cmd/ # Commandes CLI (Cobra)
│ ├── root.go # Commande racine + init config/store
│ ├── add.go # add
│ ├── list.go # list
│ ├── update.go # update --id ...
│ └── delete.go # delete --id ...
├── internal/
│ ├── domain/
│ │ └── contact.go # Entité Contact
│ ├── app/
│ │ └── service.go # Logique métier (ContactService)
│ └── store/ # Interface Storer + implémentations
│ ├── storer.go
│ ├── memory/ # Stockage mémoire (tests)
│ ├── json/ # Stockage fichier JSON
│ └── gorm/ # Stockage GORM/SQLite
├── config.yaml # Configuration (type de stockage)
├── go.mod / go.sum
└── main.go # Entrée du programme

Exemple de sortie lors d'n ajout update et delete à partir des flags
<img width="959" height="385" alt="image" src="https://github.com/user-attachments/assets/2d35aeed-8ae0-4fea-8bd1-f5736ab7da04" />

## Fonctionnement

1. Cloner le projet :
```bash
git clone https://github.com/Rayane-94/CRM-GO.git
cd CRM-GO
```
Lancement go run afin d'ajouter un contact: 
```bash
go run . add --first ... --last ... --email ... --phone ... --company "..." --notes "..."
```

on peut supprimer on contact à partir de l'id: 
```bash
go run . delete --id 1
```

De meme pour l'update

```bash
go run . update --id 1 --first Jean
```




