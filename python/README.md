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

and to start server, in /python simply run

```bash
uvicorn app.main:app --reload
```

To run test

```bash
pytest
```

To build fm-fuzzy image, from /python run

```bash
docker build -t fm-fuzz .
```

To run docker container

```bash
docker run -d --name fm-fuzz -p 80:80 fm-fuzz-ms
```

and that's it, have fun coding!
