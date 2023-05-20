import numpy as np
import skfuzzy as fuzz
from skfuzzy import control as ctrl
# New Antecedent/Consequent objects hold universe variables and membership
# functions
def fuzzy_cal(self_hp, work_hp, co_worker_hp):
# input
    points = ctrl.Antecedent(np.arange(0, 101, 1), 'self-happiness')
    workspace_happiness = ctrl.Antecedent(np.arange(0, 101, 1), 'workspace-happiness')
    co_worker_happiness = ctrl.Antecedent(np.arange(0, 101, 1), 'co_worker_happiness')
    # output
    score = ctrl.Consequent(np.arange(0, 101, 1), 'score')
    # Auto-membership function population is possible with .automf(3, 5, or 7)
    points['saddest'] = fuzz.trapmf(points.universe, [0, 0, 20, 50])
    points['normal'] = fuzz.trimf(points.universe, [20, 50, 80])
    points['happiest'] = fuzz.trapmf(points.universe, [50, 90, 100, 100])

    workspace_happiness['saddest'] = fuzz.trapmf(workspace_happiness.universe, [0, 0, 20, 50])
    workspace_happiness['normal'] = fuzz.trimf(workspace_happiness.universe, [20, 50, 80])
    workspace_happiness['happiest'] = fuzz.trapmf(workspace_happiness.universe, [50, 90, 100, 100])

    co_worker_happiness['saddest'] = fuzz.trapmf(co_worker_happiness.universe, [0, 0, 20, 50])
    co_worker_happiness['normal'] = fuzz.trimf(co_worker_happiness.universe, [20, 50, 80])
    co_worker_happiness['happiest'] = fuzz.trapmf(co_worker_happiness.universe, [50, 90, 100, 100])
    # Custom membership functions can be built interactively with a familiar,
    # Pythonic API
    score['low'] = fuzz.trapmf(score.universe, [0, 0, 20, 50])
    score['medium'] = fuzz.trimf(score.universe, [20, 50, 80])
    score['high'] = fuzz.trapmf(score.universe, [50, 90, 100, 100])

    # Rules
    rule1 = ctrl.Rule(points['saddest'] | workspace_happiness['saddest'] | co_worker_happiness['saddest'], score['low'])
    rule2 = ctrl.Rule(workspace_happiness['normal'] | points['normal'] | co_worker_happiness['normal'], score['medium'])
    rule3 = ctrl.Rule(workspace_happiness['happiest'] | points['happiest'] | co_worker_happiness['happiest'], score['high'])

    scoring_ctrl = ctrl.ControlSystem([rule1, rule2, rule3])

    scoring = ctrl.ControlSystemSimulation(scoring_ctrl)
    # Pass inputs to the ControlSystem using Antecedent labels with Pythonic API
    # Note: if you like passing many inputs all at once, use .inputs(dict_of_data)
    scoring.input['self-happiness'] = self_hp
    scoring.input['workspace-happiness'] = work_hp
    scoring.input['co_worker_happiness'] = co_worker_hp

    # Crunch the numbers
    scoring.compute()

    print(scoring.output['score'])

    return round(scoring.output['score'])

def fuzzy_cal_points(values):
    points = ctrl.Antecedent(np.arange(0, 101, 1), 'point')
    # output
    score = ctrl.Consequent(np.arange(0, 101, 1), 'score')
    # Auto-membership function population is possible with .automf(3, 5, or 7)
    points['saddest'] = fuzz.trapmf(points.universe, [0, 0, 20, 50])
    points['normal'] = fuzz.trimf(points.universe, [20, 50, 80])
    points['happiest'] = fuzz.trapmf(points.universe, [50, 90, 100, 100])
    # Custom membership functions can be built interactively with a familiar,
    # Pythonic API
    score['low'] = fuzz.trapmf(score.universe, [0, 0, 20, 50])
    score['medium'] = fuzz.trimf(score.universe, [20, 50, 80])
    score['high'] = fuzz.trapmf(score.universe, [50, 90, 100, 100])
     # Rules
    rule1 = ctrl.Rule(points['saddest'], score['low'])
    rule2 = ctrl.Rule(points['normal'], score['medium'])
    rule3 = ctrl.Rule(points['happiest'], score['high'])

    scoring_ctrl = ctrl.ControlSystem([rule1, rule2, rule3])

    scoring = ctrl.ControlSystemSimulation(scoring_ctrl)
    # Pass inputs to the ControlSystem using Antecedent labels with Pythonic API
    # Note: if you like passing many inputs all at once, use .inputs(dict_of_data)
    for value in values:
        scoring.input['point'] = value
     # Crunch the numbers
    scoring.compute()

    print(scoring.output['score'])

    return round(scoring.output['score'])
 
