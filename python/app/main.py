# FeelMe Fuzzy APIs

from typing import Union
from pydantic import BaseModel
from fastapi import FastAPI, HTTPException
from .fuzzy import fuzzy_cal
from .fuzzy import fuzzy_cal_points

class Happiness_Points(BaseModel):
    self_hp: int
    work_hp: int
    co_worker_hp: int
app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.get("/health-check")
def read_health_check():
    return {"status": "Healthy"}

@app.get("/v1/fuzzy/")
def test_fuzzy(self_hp: int, work_hp: int, co_worker_hp: int):
    result = fuzzy_cal(self_hp, work_hp, co_worker_hp)
    return {"value": result}
@app.get("/v1/fuzzy/cal/")
def cal_fuzzy_self_hp(point: int):
    result = fuzzy_cal_points(point)
    return {"value": result}