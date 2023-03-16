# FeelMe Fuzzy APIs

## This is api service for calculate fuzzy logic using skfuzzy library

We will need a few steps to setup our project environments and requirement

1. simply run this command

```bash
sh init-project.sh
```

or

```bash
python3 -m venv fm-fuzz-venv
pip install -r ./requirements.txt
```

2. Let's start our virtual environment for development by run following command

```bash
source venv.sh
```

and to start server, simply run

```bash
uvicorn main:app --reload --app-dir app
```

and that's it, have fun coding!
