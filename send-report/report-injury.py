import requests
import json
import rospy

# host address and port
HOST = "http://10.100.1.200:8000"

# server endpoint
ENDPOINT = "/api/status"

# creating formatted url
url = "{}{}".format(HOST, ENDPOINT)


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

#  
def post_report(r):
    print(url)
    # data = marshal_report(report)
    requests.post(url, json = report)



# Runs the the following code if this script is run by itself
if __name__ == "__main__":
    post_report(report)
