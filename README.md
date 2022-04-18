# quiviv-cli
## Objectif

Développer un exécutable en GO ayant une interface en ligne de commande. Il doit être capable d'exécuter :

- `quiviv-cli` help pour afficher les options possible de ces exécutables
- `quiviv-cli config --path /tmp/quiviv.yaml` pour lire le fichier de configuration au format **yaml** [https://www.redhat.com/fr/topics/automation/what-is-yaml] et afficher sur la sortie standard le contenu au format **json**. L'option **--path** est requise donc un message sera affiché sur la sortie d'erreur pour signaler que la commande est invalide.
```
$ cat /tmp/quiviv.yaml

name: Martin Développeur 

job: Développeur 

skill: Élite 

employed: True 

foods: 

    - Pomme 

    - Orange 

    - Fraise 

    - Mangue 

languages: 

    perl: Élite

    python: Élite

    pascal: Bases

 

$ quiviv-cli config --path /tmp/quiviv.yaml

{

    "name": "Martin Développeur",

    "job": "Développeur", 

    "skill": "Élite",

    "employed": true,

    "foods": ["Pomme", "Orange", "Fraise", "Mangue"],

    "languages": {

        "perl": "Élite" ,

        "python": "Élite", 

        "pascal": "Bases"

    }

}

 
```
- `quiviv-cli list --dir /tmp` pour lister tout le contenu du dossier spécifié en affichant sur la sortie standard un élément par ligne. L'option **--dir** est optionnelle donc si elle n'est pas spécifiée utiliser le dossier où se trouve l'exécutable.
```
$ quiviv-cli list --dir /tmp

quiviv.yaml

cs_debug.log

example.conf

$ quiviv-cli list

Desktop

Documents

Downloads

Music

Pictures

Videos
```

### Les contraintes
- Il faudra utiliser uniquement le language **GO** [https://go.dev].
- Il faudra compiler une version pour **linux** nommé quiviv-cli et une pour **windows** nommé `quiviv-cli.exe`
- Ne pas hésiter à utiliser les librairies les plus adaptées.
