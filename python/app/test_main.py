from fastapi.testclient import TestClient

from .main import app

client = TestClient(app)

def test_read_root():
    response = client.get(f"/v1/fuzzy/?self_hp=80&work_hp=70&co_worker_hp=50")
    assert response.status_code == 200
    assert response.json() == {"value":65}
