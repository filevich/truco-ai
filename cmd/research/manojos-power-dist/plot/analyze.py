import json

with open('/tmp/dist-envido.json', 'r') as f: data_envido = json.loads(f.read())
with open('/tmp/dist-flor.json', 'r') as f: data_flor = json.loads(f.read())
with open('/tmp/dist-power-sum.json', 'r') as f: data_power = json.loads(f.read())

# cast keys as ints
data_envido = {int(k):v for k,v in data_envido.items()}
data_flor = {int(k):v for k,v in data_flor.items()}
data_power = {int(k):v for k,v in data_power.items()}

def analyse_continuity(data):
    min_key, max_key = min(data.keys()), max(data.keys())
    print(f"min_key: {min_key}")
    print(f"max_key: {max_key}")
    not_in_data = [k for k in range(min_key, max_key + 1) if k not in data]
    are_keys_continuous = len(not_in_data) == 0
    print(f"are they continuous? {are_keys_continuous}")
    if not are_keys_continuous: print(f"no, these are missing: {not_in_data}")

def total_sum(data):
    return sum([v for v in data.values()])

def cdf(data, key):
    s = 0
    for i in range(0, key +1):
        s += data[i] if i in data else 0
    t = total_sum(data)
    return s / t

print("analysing data_envido")
analyse_continuity(data_envido)
print(f"data_envido total sum: {total_sum(data_envido)}")

print("analysing data_flor")
analyse_continuity(data_flor)
print(f"data_flor total sum: {total_sum(data_flor)}")

print("analysing data_power")
analyse_continuity(data_power)
print(f"data_power total sum: {total_sum(data_power)}")
print(f"data_power.cdf(84)={cdf(data_power, 84)}")
