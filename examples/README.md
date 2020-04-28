## Running main.go

**Download data**

```bash
$ # Download darpa dataset in MIDAS format
$ curl -O https://www.comp.nus.edu.sg/~sbhatia/assets/datasets/darpa_midas.csv
$ # Download original darpa dataset
$ curl -O https://www.comp.nus.edu.sg/~sbhatia/assets/datasets/darpa_original.csv
```

```bash
$ go run main.go > scores.txt
$ python auc.py # checking the AUC
AUC:  0.9171735033552131

$ go run main.go --norelations > scores.txt # use midas instead of midasR
$ python auc.py # checking the AUC
AUC:  0.9467124341220924
```
