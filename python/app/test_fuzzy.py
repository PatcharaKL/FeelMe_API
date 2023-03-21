from app.fuzzy import fuzzy_cal
def test_fuzzy_cal():
    assert fuzzy_cal(self_hp=80,work_hp=70.1,co_worker_hp=50.564) == 65