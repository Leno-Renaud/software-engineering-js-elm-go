# Documentation Technique - Elm Wordle

Ce document explique le fonctionnement interne et l'architecture du code source `src/Main.elm`.

## Architecture Globale (The Elm Architecture)

Le programme suit l'architecture standard Elm, divisée en 4 parties interconnectées :

1.  **Model** : L'état de l'application (données).
2.  **View** : La représentation visuelle de l'état (HTML).
3.  **Update** : La logique qui modifie l'état en fonction des messages.
4.  **Subscriptions** : L'écoute des événements externes (clavier).

## 1. Le Modèle (`Model`)

Le type `Model` est une structure (record) qui contient tout ce qui peut changer dans le jeu :

*   `guesses` (`List String`) : La liste des mots déjà validés par le joueur.
*   `currentGuess` (`String`) : Les lettres que le joueur est en train de taper (tampon).
*   `targetWord` (`String`) : Le mot mystère à trouver.
*   `gameState` (`type GameState`) : L'état actuel de la partie (`Playing`, `Won`, `Lost`).
*   `validWords` (`List String`) : Le dictionnaire complet chargé depuis `words.txt`.
*   `errorMessage` (`Maybe String`) : Pour afficher une erreur si le chargement échoue.

## 2. Initialisation (`init`)

Au démarrage (`init`), deux choses se produisent :
1.  L'état est initialisé avec des valeurs par défaut (listes vides).
2.  Une commande HTTP (`Http.get`) est lancée pour récupérer le contenu de `/words.txt`.

## 3. Les Messages (`Msg`)

Les changements d'état sont pilotés par des messages précis :

*   `GotWords (Result ...)` : Reçu quand le fichier `words.txt` a fini de charger (ou a échoué).
*   `NewTargetWord String` : Reçu quand le générateur aléatoire a choisi un mot.
*   `KeyPressed String` : Reçu à chaque pression de touche clavier.
*   `Restart` : Reçu au clic sur le bouton "Restart".

## 4. La Logique de Mise à Jour (`update`)

C'est le cœur du programme. La fonction `update` reçoit un message et l'ancien modèle, et retourne un nouveau modèle.

### Gestion du Clavier (`handleKey`)
Gère les entrées utilisateur :
*   **Lettres (A-Z)** : Ajoutées à `currentGuess` (limité à 5 lettres).
*   **Backspace** : Supprime le dernier caractère de `currentGuess`.
*   **Enter** : Déclenche la validation (`submitGuess`).

### Validation (`submitGuess`)
Vérifie les règles du jeu :
1.  Si la longueur n'est pas 5 lettres → Ignore.
2.  Si le mot n'est pas dans `validWords` → Ignore.
3.  Sinon :
    *   Ajoute le mot à la liste `guesses`.
    *   Vérifie la victoire (mot == `targetWord`).
    *   Vérifie la défaite (6 essais atteints).
    *   Met à jour `gameState`.

## 5. La Vue (`view`)

L'interface est construite dynamiquement à chaque modification du modèle.

*   **Grille (`viewGrid`)** : Affiche toujours 6 lignes.
*   **Lignes (`viewRow`)** :
    *   Les lignes passées affichent les mots de `guesses`.
    *   La ligne courante affiche `currentGuess`.
    *   Les lignes futures sont vides.
*   **Couleurs (`getCellColor`)** :
    *   Compare chaque lettre du mot tenté avec le mot cible.
    *   **Vert** : `guessChar == targetChar`.
    *   **Jaune** : La lettre existe dans `targetWord` (`List.member`).
    *   **Gris** : La lettre est absente.

## 6. Générateur Aléatoire

Le module `Random` est utilisé pour choisir un mot au hasard dans la liste chargée.
La fonction `randomWordGenerator` prend la liste des mots et retourne un générateur qui sélectionne un élément uniformément.
