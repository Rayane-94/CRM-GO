# CRM-GO

CRM-GO est une application CRM développée en **Go**, permettant de gérer et organiser des données clients de manière simple et efficace.

---

## Description

Ce projet contient les fichiers principaux pour démarrer une application CRM en Go.  
Il permet d’ajouter, modifier et tester des fonctionnalités CRM via des tests unitaires.

---

## Architecture

RM-GO/
│
├─ main.go # Fichier principale
├─ go.mod # Fichier de module Go pour la gestion des dépendances a ne pas toucher 
└─ test.go # Contient tous les tests unitaires

## Fonctionnement

1. Cloner le projet :
```bash
git clone https://github.com/Rayane-94/CRM-GO.git
cd CRM-GO
```
Option 1 lancement normal: 
```bash
go run main.go
```
Cette commande vous permet de vous deplacer manuellement dans l'application

Option 2 lancement avec un flag :
```
go run main.go --ajouter
```
Cette commande permet d’ajouter un utilisateur sans passer par le menu interactif.
