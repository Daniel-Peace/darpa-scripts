import requests
import json

id = 0
team = "CSUCI"
system = "SYSTEM"
type = "TYPE"
value = 0

# Dictionary for holding report information
report = {
    "casualty": id,
    "team": team,
    "system": system,
    "type": type,
    "value": value
}

# Takes a dictionary and returns that dictionary as a formatted json string
def marshal_report(r):
    json_report = json.dumps(r)
    print(json_report)
    return json_report

# Runs the the following code if this script is run by itself
if __name__ == "__main__":
    marshal_report(report)
