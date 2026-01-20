# Elm Wordle

Un clone du jeu populaire Wordle, écrit entièrement en Elm.
Ce projet a été réalisé dans le cadre du cours ELP.

## Prérequis

Pour exécuter ce projet, vous devez avoir installé :
*   [Elm](https://guide.elm-lang.org/install/elm.html) (Version 0.19.1)
*   [Python 3](https://www.python.org/downloads/) (pour le serveur local)

## Installation et Lancement

1.  **Compiler le code Elm**
    Transforme le code source Elm en JavaScript.
    ```bash
    elm make src/Main.elm --output=main.js
    ```

2.  **Lancer le serveur local**
    Nécessaire pour charger le fichier `words.txt` sans restrictions de navigateur.
    ```bash
    python -m http.server
    ```

3.  **Jouer**
    Ouvrez votre navigateur à l'adresse suivante :
    [http://localhost:8000](http://localhost:8000)

## Comment jouer

*   Le but est de deviner le mot caché en 6 essais.
*   Tapez un mot de 5 lettres et appuyez sur **Entrée**.
*   Codes couleurs :
    *   🟩 **Vert** : La lettre est dans le mot et au bon endroit.
    *   🟨 **Jaune** : La lettre est dans le mot mais au mauvais endroit.
    *   ⬜ **Gris** : La lettre n'est pas dans le mot.

## Structure du projet

*   `src/Main.elm` : Le code source de l'application (logique, modèle, vue).
*   `index.html` : La page web qui charge l'application.
*   `elm.json` : Configuration du projet et dépendances.
*   `words.txt` : Liste des mots valides pour le jeu.
