# Rucksackproblem-Solver

Travis - Develop Branch
![Travis - Develop Branch](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-kobaes.svg?token=uuWEU7CbK6apjnjsgMfR&branch=develop)


## About

Dieses Programm ermöglicht ein Lösen des Rucksacksproblems via Kommandozeile.  
Es wurde ein greedy und ein dynamischer Algorithmus zur Problemlösung verwendet.

## Getting started

Das Repository kann mithilfe des folgenden Befehls geklont werden:

```
$ git clone https://github.com/ob-algdatii-ss19/leistungsnachweis-kobaes
```

Danach muss das noch das Go-Executable erzeugt werden mit:

```
$ go build
```


## Usage

Das Programm kann folgendermaßen aufgerufen werden:

```
$ leistungsnachweis-kobaes
```

Dadurch werden alle möglichen Befehle angezeigt:

```
Rucksackproblem-Solver  
The following flags are available:  
  --all  
        flag which specifies to use both greedy and dynamic algorithm  
  --configfile  
        flag which specifies a path to a config file  
  --dynamic  
        flag which specifies to use dynamic algorithm  
  --greedy  
        flag which specifies to use greedy algorithm  
  --help  
        flag which specifies that help should be shown  
```

Eine Beispiel-Ausgabe des greedy Algorithmus ist:

```
$ leistungsnachweis-kobaes --greedy  

Using the following items for computing knapsack:
Item: Oven, Volume: 3, Worth: 30
Item: Giraffe, Volume: 100, Worth: 50
Item: EmpireStateBuilding, Volume: 1000000, Worth: 100000000
Item: JewingGum, Volume: 1, Worth: 2

Greedy Algorithm:
Oven JewingGum 
Total Worth: 32
Total Volume: 4
```

# Authors
(Sebastian Baumann|https://github.com/baschte83), (Korbinian Karl|https://github.com/korbster), (Ehsan Moslehi|https://github.com/eca852)
