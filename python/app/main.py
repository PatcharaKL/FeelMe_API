# FeelMe Fuzzy APIs

from typing import Union
from pydantic import BaseModel
from fastapi import FastAPI
from fuzzy import fuzzy_cal

class Happiness_Points(BaseModel):
    self_hp: int
    work_hp: int
    co_worker_hp: int

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}


@app.get("/v1/fuzzy/")
def test_fuzzy(hp: Happiness_Points):
    result = fuzzy_cal(hp.self_hp, hp.work_hp, hp.co_worker_hp)
    return {"value": result}